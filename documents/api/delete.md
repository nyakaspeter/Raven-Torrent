# Delete torrent by hash

Delete torrent by 40 characters long infohash.

**URL** : `/api/delete/{hash}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type   | Description                  |
| --------- | ------ | ---------------------------- |
| `hash`    | string | 40 characters long infohash. |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/delete/08ada5a7a6183aae1e09d831df6748d566095a10`

**Success Response** :

```json
{
  "success": true,
  "message": "Torrent deleted."
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "Torrent not found."
}
```
