# Search subtitles by inner file hash

Search subtitles for movies or tv shows by torrent's inner file hash.

**URL** : `/api/subtitlesbyfile/{hash}/{base64path}/lang/{lang}`

**Method** : `GET`

**Query Parameters** :

| Parameter    | Type    | Description                                                   |
| ------------ | ------- | ------------------------------------------------------------- |
| `hash`       | string  | 40 characters long infohash.                                  |
| `base64path` | string  | Base64 encoded path with filename.                            |
| `lang`       | integer | ISO 639-2 three-letter language codes separated with a comma. |

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

`GET http://localhost:9000/api/subtitlesbyfile/08ada5a7a6183aae1e09d831df6748d566095a10/U2ludGVsLm1wNA==/lang/eng`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "lang": "en",
      "subtitlename": "sintel_en.srt",
      "releasename": "",
      "subformat": "srt",
      "subencoding": "ASCII",
      "subdata": "http://localhost:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1NjIwYmMyL3NpZC1WMDRMZkRQak5YdkItRnlVbGxiZ3RtcjJpcGYvc3ViYWQvMzg3MTc5Ng==/encode/ASCII/subtitle.srt"
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
