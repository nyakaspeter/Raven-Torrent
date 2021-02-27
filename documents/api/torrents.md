# Get All Running Torrents

Get all active torrents.

**URL** : `/api/torrents`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type          | Description                                 |
| --------- | ------------- | ------------------------------------------- |
| `success` | bool          | Indicates whether the query was successful. |
| `results` | array[object] | Array of objects.                           |

**Object [ results ]** :

| Name     | Type   | Description                  |
| -------- | ------ | ---------------------------- |
| `name`   | string | Filename with extension.     |
| `hash`   | string | 40 characters long infohash. |
| `length` | string | File length in bytes.        |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/torrents`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "name": "Sintel",
      "hash": "08ada5a7a6183aae1e09d831df6748d566095a10",
      "length": "129302391"
    }
  ]
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "No active torrents found."
}
```
