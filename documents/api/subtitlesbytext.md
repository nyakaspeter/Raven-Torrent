# Search subtitles by query text

Search subtitles for movies or tv shows by query text.

**URL** : `/api/subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type    | Description                                                     |
| --------- | ------- | --------------------------------------------------------------- |
| `text`    | string  | Query text string.                                              |
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

| Name           | Type   | Description                                                                          |
| -------------- | ------ | ------------------------------------------------------------------------------------ |
| `lang`         | string | ISO 639-1 two-letter language code.                                                  |
| `subtitlename` | string | Subtitle filename.                                                                   |
| `releasename`  | string | Subtitle release name.                                                               |
| `subformat`    | string | Subtitle file type. **It is always "srt", because only SubRip files are supported.** |
| `subencoding`  | string | Character encoding of the subtitle file.                                             |
| `subdata`      | string | HTTP url to get the raw content of the subtitle file.                                |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/subtitlesbytext/stranger%20things/lang/ger,eng/season/3/episode/6`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "lang": "de",
      "subtitlename": "Stranger.Things.S03E06.INTERNAL.1080p.WEB.x264-STRiFE.srt",
      "releasename": " Stranger.Things.S03E06.INTERNAL.1080p.WEB.x264-STRiFE",
      "subformat": "srt",
      "subencoding": "UTF-8",
      "subdata": "http://127.0.0.1:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1NjUwYmMyL3NpZC1CbzZ6emV5T1pTZlVtLU9BM2E5NUsxYndzTzQvc3ViYWQvNzgyMjY3OQ==/encode/UTF-8/subtitle.srt"
    },
    {
      "lang": "en",
      "subtitlename": "Stranger.Things.S03E06.720p.WEBRip.X264-METCON.srt",
      "releasename": "Stranger.Things.S03E06.720p.WEBRip.X264-METCON",
      "subformat": "srt",
      "subencoding": "UTF-8",
      "subdata": "http://127.0.0.1:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1NTEwYmI4L3NpZC1CbzZ6emV5T1pTZlVtLU9BM2E5NUsxYndzTzQvc3ViYWQvNzgyMjQxNw==/encode/UTF-8/subtitle.srt"
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
