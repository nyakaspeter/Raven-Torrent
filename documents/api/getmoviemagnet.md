# Get Movie Torrents By IMDB Id And/Or Query Text

Search torrent links for movie.

**URL**

- Search by IMDB id: `/api/getmoviemagnet/imdb/{imdb}/providers/{providers}`
- Search by query text: `/api/getmoviemagnet/query/{query}/providers/{providers}`
- Search by both at once: `/api/getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}`

**Method** : `GET`

**Query Parameters** :

| Parameter   | Type   | Description                                                 |
| ----------- | ------ | ----------------------------------------------------------- |
| `imdb`      | string | Internet Movie Database identifier.                         |
| `query`     | string | Text query to search. **This value should be URI encoded.** |
| `providers` | string | Providers separated by a comma.                             |

**Text Query Parameters** :

| Parameter     | Type   | Description   |
| ------------- | ------ | ------------- |
| `title`       | string | Movie title.  |
| `releaseyear` | string | Release year. |

**Supported Providers For IMDB id** :

| Provider   | Type   | Website             |
| ---------- | ------ | ------------------- |
| `jackett`  | string | JACKETT API         |
| `pt`       | string | POPCORN TIME        |
| `yts`      | string | YTS                 |
| `rarbg`    | string | RARBG               |
| `pto`      | string | POPCORN TIME ONLINE |
| `itorrent` | string | ITORRENT            |

**Supported Providers For Query Text** :

| Provider  | Type   | Website     |
| --------- | ------ | ----------- |
| `jackett` | string | JACKETT API |
| `1337x`   | string | 1337X       |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type          | Description                                 |
| --------- | ------------- | ------------------------------------------- |
| `success` | bool          | Indicates whether the query was successful. |
| `results` | array[object] | Array of objects.                           |

**Object [ results ]** :

| Name       | Type   | Description                         |
| ---------- | ------ | ----------------------------------- |
| `hash`     | string | 40 characters long infohash.        |
| `quality`  | string | Video quality.                      |
| `size`     | string | Torrent data size in bytes.         |
| `provider` | string | Source of the magnet link.          |
| `lang`     | string | ISO 639-1 two-letter language code. |
| `title`    | string | Movie title.                        |
| `seeds`    | string | Currently available seeds.          |
| `peers`    | string | Currently available peers.          |
| `magnet`   | string | Magnet link.                        |
| `torrent`  | string | Torrent file link.                  |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

```
GET http://localhost:9000/api/getmoviemagnet/imdb/tt0093773/query/title%3DPredator%26releaseyear%3D1987/providers/rarbg,yts,1337x
```

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "hash": "EA2AC2D1F0F51D6F0D1284474D25616108F3E59A",
      "quality": "720p",
      "size": "1037356237",
      "provider": "1337X",
      "lang": "en",
      "title": "Predator.1987.REMASTERED.720p.BluRay.999MB.HQ.x265.10bit-GalaxyRG",
      "seeds": "110",
      "peers": "15"
    },
    {
      "hash": "FB132F15FD902A715403C7F57025C3EEC86F77E5",
      "quality": "1080p",
      "size": "1760936591",
      "provider": "YTS",
      "lang": "en",
      "title": "Predator",
      "seeds": "68",
      "peers": "4"
    },
    {
      "hash": "63feb6094d1378928bede0ec715b73aec680606f",
      "quality": "1080p",
      "size": "2181126473",
      "provider": "RARBG",
      "lang": "en",
      "title": "Predator.1987.NEW.REMASTERED.1080p.BluRay.H264.AAC-RARBG",
      "seeds": "16",
      "peers": "1"
    }
  ]
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "No magnet links found."
}
```
