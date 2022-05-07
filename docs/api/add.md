# Add torrent by magnet link

Add torrent by 40 characters long infohash.

**URL** : `/api/add/{base64uri}`

**Method** : `GET`

**Query Parameters** :

| Parameter   | Type   | Description                                     |
| ----------- | ------ | ----------------------------------------------- |
| `base64uri` | string | Base64 encoded magnet link or torrent file name |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type          | Description                                 |
| --------- | ------------- | ------------------------------------------- |
| `success` | bool          | Indicates whether the query was successful. |
| `results` | array[object] | Array of objects.                           |

**Object [ results ]** :

| Name     | Type   | Description                                  |
| -------- | ------ | -------------------------------------------- |
| `name`   | string | Filename with extension.                     |
| `url`    | string | HTTP url to get the raw content of the file. |
| `length` | string | File length in bytes.                        |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/add/08ada5a7a6183aae1e09d831df6748d566095a10`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "name": "Sintel.en.srt",
      "url": "http://localhost:9000/api/get/08ada5a7a6183aae1e09d831df6748d566095a10/U2ludGVsLmVuLnNydA==",
      "length": "1514"
    },
    {
      "name": "Sintel.mp4",
      "url": "http://localhost:9000/api/get/08ada5a7a6183aae1e09d831df6748d566095a10/U2ludGVsLm1wNA==",
      "length": "129241752"
    },
    {
      "name": "poster.jpg",
      "url": "http://localhost:9000/api/get/08ada5a7a6183aae1e09d831df6748d566095a10/cG9zdGVyLmpwZw==",
      "length": "46115"
    }
  ]
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "Failed to add torrent."
}
```
