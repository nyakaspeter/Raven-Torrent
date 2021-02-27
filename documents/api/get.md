# Stream Or Download The Selected File

Stream or download the selected file over http connection.

**URL** : `/api/get/{hash}/{base64path}`

**Method** : `GET`

**Query Parameters** :

| Parameter    | Type   | Description                        |
| ------------ | ------ | ---------------------------------- |
| `hash`       | string | 40 characters long infohash.       |
| `base64path` | string | Base64 encoded path with filename. |

## Success Response

**Code** : `200 OK`

**Response Data** : `The Raw Content Of The File.`

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/get/08ada5a7a6183aae1e09d831df6748d566095a10/U2ludGVsLmVuLnNydA==`

**Success Response** :

```srt
1
00:01:47,250 --> 00:01:50,500
This blade has a dark past.

2
00:01:51,800 --> 00:01:55,800
It has shed much innocent blood.

3
00:01:58,000 --> 00:02:01,450
You're a fool for traveling alone,
so completely unprepared.

4
00:02:01,750 --> 00:02:04,800
You're lucky your blood's still flowing.

5
00:02:05,250 --> 00:02:06,300
Thank you.
```

**Error Response** :

```json
{
  "success": false,
  "message": "Invalid base64 path."
}
```
