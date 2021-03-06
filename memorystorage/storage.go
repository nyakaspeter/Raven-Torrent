package memorystorage

import (
	//"fmt"
	"io"
	"log"
	"math"
	"runtime"
	"sync"

	"github.com/anacrolix/torrent/metainfo"
	lru "github.com/hashicorp/golang-lru"
)

const megaByte = 1024 * 1024

var maxMemorySize int64 // Maximum memory size in MByte

var maxPieceLength int64 // Maximum piece length in MByte

var maxCount int // Number of pieces that LRU cache can hold

var lruStorage *lru.Cache

var needToDeleteKey = -1

var memStats runtime.MemStats

var setMaxCount = true

func SetMemorySize(memorySize int64, pieceLength int64) {
	maxMemorySize = memorySize
	maxPieceLength = pieceLength
	maxCount = int(maxMemorySize / pieceLength)
	lruStorage, _ = lru.NewWithEvict(maxCount, onEvicted)
}

func onEvicted(key interface{}, value interface{}) {
	needToDeleteKey = key.(int)
	runtime.GC()
	//log.Printf("Removed piece from LRU: %d, LRU space: %d/%d", needToDeleteKey, lruStorage.Len(), maxCount)
}

// Restricting all I/O through a single mutex, which would stop simultanious read/writes.
func storageWriteAt(mt *memoryTorrent, key int, b []byte, off int64) (int, error) {
	mt.storageMutex.Lock()
	defer mt.storageMutex.Unlock()

	if setMaxCount == true {
		// 75% of max memory size for LRU cache will keep memory allocation approximately in the right bounds
		elementCount := int(math.Floor(float64(maxMemorySize*megaByte) / float64(mt.pl) * 75 / 100))

		if maxCount != elementCount {
			lruStorage.Resize(elementCount)
			maxCount = elementCount
		}

		log.Printf("LRU cache size: %d", maxCount)

		setMaxCount = false
	}

	newPiece := false

	dataInterface, present := lruStorage.Get(key)
	if present == false {
		newPiece = true
		dataInterface = []byte{}
	}

	ioff := int(off)
	iend := ioff + len(b)
	if len(dataInterface.([]byte)) < iend {
		if len(dataInterface.([]byte)) == ioff {
			if lruStorage.Add(key, append(dataInterface.([]byte), b...)) == true {
				if needToDeleteKey > -1 {
					mt.cl.pc.Set(metainfo.PieceKey{mt.ih, needToDeleteKey}, false)
				}
			}
			return len(b), nil
		}
		// Add zero bytes to the end of data
		if lruStorage.Add(key, append(dataInterface.([]byte), make([]byte, iend-len(dataInterface.([]byte)))...)) == true {
			if needToDeleteKey > -1 {
				mt.cl.pc.Set(metainfo.PieceKey{mt.ih, needToDeleteKey}, false)
			}
		}
	}

	dataInterface, present = lruStorage.Get(key)
	if present == false {
		dataInterface = []byte{}
	}

	copy(dataInterface.([]byte)[ioff:], b)
	if lruStorage.Add(key, dataInterface.([]byte)) == true {
		if needToDeleteKey > -1 {
			mt.cl.pc.Set(metainfo.PieceKey{mt.ih, needToDeleteKey}, false)
		}
	}

	if newPiece {
		//log.Printf("Added new piece to LRU: %d, LRU space: %d/%d", key, lruStorage.Len(), maxCount)
		//runtime.GC()
		logMemStats()
	}

	// Before return check if need to free up some memory
	//FreeMemoryPercent(mt, uint64(maxMemorySize), 15)

	return len(b), nil
}

func storageReadAt(mu *sync.Mutex, key int, b []byte, off int64) (int, error) {
	dataInterface, present := lruStorage.Get(key)
	if present == false {
		dataInterface = []byte{}
	}

	ioff := int(off)
	if len(dataInterface.([]byte)) <= ioff {
		return 0, io.EOF
	}

	n := copy(b, dataInterface.([]byte)[ioff:])
	if n != len(b) {
		return n, io.EOF
	}

	return len(b), nil
}

func FreeMemoryPercent(mt *memoryTorrent, threshold uint64, percent int) {
	runtime.ReadMemStats(&memStats)

	if (memStats.Alloc / megaByte) > threshold {
		var deleteCount = (maxCount * percent) / 100

		if deleteCount == 0 {
			deleteCount++
		}

		log.Printf("Freeing up memory, currently allocated: %v MB\n", (memStats.Alloc / megaByte))

		for i := 0; i < deleteCount; i++ {
			key, _, ok := lruStorage.RemoveOldest()
			if ok == true {
				if needToDeleteKey > -1 {
					mt.cl.pc.Set(metainfo.PieceKey{mt.ih, key.(int)}, false)
				}
			}
		}

		needToDeleteKey = -1

		runtime.GC()
	}
}

func storageDelete(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	setMaxCount = true

	lruStorage.Purge()

	needToDeleteKey = -1

	runtime.GC()
}

func logMemStats() {
	runtime.ReadMemStats(&memStats)
	log.Printf("Memory storage allocated memory: %v MB, NumGC: %v, LRU: %v/%v", (memStats.Alloc / megaByte), memStats.NumGC, lruStorage.Len(), maxCount)
}
