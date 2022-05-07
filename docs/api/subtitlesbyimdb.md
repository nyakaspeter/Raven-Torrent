# Search subtitles by IMDB id

Search subtitles for movies or tv shows by IMDB id.

**URL** : `/api/subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type    | Description                                                     |
| --------- | ------- | --------------------------------------------------------------- |
| `imdb`    | string  | Internet Movie Database identifier.                             |
| `lang`    | string  | ISO 639-2 three-letter language codes separated with a comma.   |
| `season`  | integer | Season number. **Must be set to 0 for movie subtitle search.**  |
| `episode` | integer | Episode number. **Must be set to 0 for movie subtitle search.** |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type          | Description                                 |
| --------- | ------------- | ------------------------------------------- |
| `success` | bool          | Indicates whether the query was successful. |
| `results` | array[object] | Array of objects.                           |

**Object [ results ]** :

| Name           | Type   | Description                                                    |
| -------------- | ------ | -------------------------------------------------------------- |
| `lang`         | string | ISO 639-1 two-letter language code.                            |
| `subtitlename` | string | Subtitle filename.                                             |
| `releasename`  | string | Subtitle release name.                                         |
| `subformat`    | string | Subtitle file type. **Only SubRip (srt) files are supported.** |
| `subencoding`  | string | Character encoding of the subtitle file.                       |
| `subdata`      | string | HTTP url to get the raw content of the subtitle file.          |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/subtitlesbyimdb/tt0460681/lang/eng,hun/season/15/episode/17`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "lang": "en",
      "subtitlename": "Supernatural.s15e17.720p.Unity.WEBRip.x264-4227C - m.srt",
      "releasename": "Supernatural.S15E17.720p.WEB.H264-CAKES",
      "subformat": "srt",
      "subencoding": "UTF-8",
      "subdata": "http://127.0.0.1:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1MzYwYmFlL3NpZC1zSEJrR2tJOVl3MUZWZHhySnRKQ3Z3Y3lMcTcvc3ViYWQvODQwNDQwMQ==/encode/UTF-8/subtitle.srt"
    },
    {
      "lang": "hu",
      "subtitlename": "supernatural.s15e17.1080p.web.h264-cakes.srt",
      "releasename": "Supernatural.S15E17.1080p.WEB.H264-CAKES",
      "subformat": "srt",
      "subencoding": "CP1250",
      "subdata": "http://127.0.0.1:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1M2IwYmI0L3NpZC1zSEJrR2tJOVl3MUZWZHhySnRKQ3Z3Y3lMcTcvc3ViYWQvODQwNTEyNw==/encode/CP1250/subtitle.srt"
    }
  ]
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "No subtitles found."
}
```
