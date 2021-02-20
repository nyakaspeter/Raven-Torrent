package memorystorage

import (
	"fmt"
	"sync"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
)

// Just to set pieces UnComplete
type memoryClient struct {
	pc storage.PieceCompletion
}

type memoryTorrent struct {
	cl *memoryClient
	pl int64
	ih metainfo.Hash
	np int // Just to set pieces UnComplete

	storageMutex *sync.Mutex
}

type memoryPiece struct {
	trt *memoryTorrent

	p metainfo.Piece
}

func NewMemoryStorage() storage.ClientImpl {
	ret := &memoryClient{
		pc: storage.NewMapPieceCompletion(),
	}

	return ret
}

func (me *memoryClient) Close() error {
	//return me.pc.Close()
	return nil
}

func (me *memoryClient) OpenTorrent(info *metainfo.Info, infoHash metainfo.Hash) (storage.TorrentImpl, error) {
	return &memoryTorrent{
		cl: me,
		pl: info.PieceLength,
		ih: infoHash,
		np: info.NumPieces(),

		storageMutex: &sync.Mutex{},
	}, nil
}

func (me *memoryTorrent) Piece(p metainfo.Piece) storage.PieceImpl {
	return &memoryPiece{
		trt: me,
		p:   p,
	}
}

func (me *memoryTorrent) Close() error {
	// Set all pieces UnComplete
	for key := 0; key < me.np; key++ {
		me.cl.pc.Set(metainfo.PieceKey{me.ih, key}, false)
	}

	storageDelete(me.storageMutex)

	return nil
}

func (sp *memoryPiece) pieceKey() metainfo.PieceKey {
	return metainfo.PieceKey{sp.trt.ih, sp.p.Index()}
}

func (sp *memoryPiece) chunkKey(index int) string {
	return fmt.Sprintf("%d", index)
}

func (sp *memoryPiece) Completion() storage.Completion {
	ret, _ := sp.trt.cl.pc.Get(sp.pieceKey())
	return ret
}

func (sp *memoryPiece) MarkComplete() error {
	sp.trt.cl.pc.Set(sp.pieceKey(), true)
	return nil
}

func (sp *memoryPiece) MarkNotComplete() error {
	sp.trt.cl.pc.Set(sp.pieceKey(), false)
	return nil
}

func (sp *memoryPiece) ReadAt(b []byte, off int64) (n int, err error) {
	ci := sp.p.Index()
	bToRead := sp.trt.pl - off
	//log.Printf("Got read for chunk (%d) offset (%d).", ci, off)
	for len(b) != 0 {
		//ck := sp.chunkKey(int(ci))
		var rLen int
		if len(b) < int(bToRead) {
			rLen = len(b)
		} else {
			rLen = int(bToRead)
		}
		_b := make([]byte, rLen)
		i, rerr := storageReadAt(sp.trt.storageMutex, int(ci), _b, off)
		//log.Printf("Doing read for chunk (%d) offset (%d) for (%d) bytes and got (%d) bytes.", ci, off, rLen, i)
		n1 := copy(b, _b[:i])
		off = 0
		ci++
		b = b[n1:]
		n += n1
		if rerr != nil {
			//log.Printf("Error Reading During Read: %s", rerr)
			err = rerr
			return
		}
	}
	return
}

func (sp *memoryPiece) WriteAt(b []byte, off int64) (n int, err error) {
	ci := sp.p.Index()
	//log.Printf("At chunk (%d), got bytes (%d) and offset (%d).", ci, len(b), off)
	bToWrite := sp.trt.pl - off
	var btw int
	for len(b) != 0 {
		if len(b) > int(bToWrite) {
			btw = int(bToWrite)
		} else {
			btw = len(b)
		}
		//ck := sp.chunkKey(int(ci))
		n1, werr := storageWriteAt(sp.trt, int(ci), b[:btw], off)
		//log.Printf("Writing (%d) bytes [confirm %d] to chunk (%d) - total bytes (%d) - offset (%d) - written (%d) bytes.", btw, len(b[:btw]), ci, len(b), off, n1)
		if werr != nil {
			//log.Printf("Error Writing During Write: %s", werr)
			err = werr
			return
		}
		if n1 > len(b) {
			break
		}
		b = b[n1:]
		off = 0
		bToWrite = sp.trt.pl - off
		ci++
		n += n1
	}
	return
}
