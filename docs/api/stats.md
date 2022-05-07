# Get running torrent statistics by hash

Get active torrent statistics by 40 characters long infohash.

**URL** : `/api/stats/{hash}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type   | Description                  |
| --------- | ------ | ---------------------------- |
| `hash`    | string | 40 characters long infohash. |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name          | Type   | Description                                 |
| ------------- | ------ | ------------------------------------------- |
| `success`     | bool   | Indicates whether the query was successful. |
| `downspeed`   | string | Download speed.                             |
| `downdata`    | string | Downloaded data size.                       |
| `downpercent` | string | Downloaded data percent.                    |
| `fulldata`    | string | Torrent data size.                          |
| `peers`       | string | Peer and seed counters.                     |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/stats/08ada5a7a6183aae1e09d831df6748d566095a10`

**Success Response** :

```json
{
  "success": true,
  "downspeed": "560 kB/s",
  "downdata": "15 MB",
  "downpercent": "11",
  "fulldata": "129 MB",
  "peers": "8/24"
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "Torrent not found."
}
```
