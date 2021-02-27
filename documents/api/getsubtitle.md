# Download Subtitle File

Download and serve the raw content of the subtitle file from memory.

**URL** : `/api/getsubtitle/{base64path}/encode/{encode}/subtitle.srt`

**Method** : `GET`

**Query Parameters** :

| Parameter    | Type   | Description                                            |
| ------------ | ------ | ------------------------------------------------------ |
| `base64path` | string | Base64 encoded path of a zip compressed subtitle file. |
| `encode`     | string | Character encoding of the subtitle file.               |

## Success Response

**Code** : `200 OK`

**Response Data** : `The Raw Content Of The Subtitle File.`

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
GET http://localhost:9000/api/getsubtitle/aHR0cDovL2RsLm9wZW5zdWJ0aXRsZXMub3JnL2VuL2Rvd25sb2FkL3NyYy1hcGkvdnJmLWY1NjIwYmMyL3NpZC1sNHdMMUYtOGxTRktLZXluS3VHVWtyT2RLZzkvc3ViYWQvMzg3MTc5Ng==/encode/ASCII/subtitle.srt
```

**Success Response** :

```srt
1
00:00:06,000 --> 00:00:12,074
Advertise your product or brand here
contact www.OpenSubtitles.org today

2
00:01:47,250 --> 00:01:50,500
This blade has a dark past.

3
00:01:51,800 --> 00:01:55,800
It has shed much innocent blood.

4
00:01:58,000 --> 00:02:01,450
You're a fool for traveling alone,
so completely unprepared.

5
00:02:01,750 --> 00:02:04,800
You're lucky your blood's still flowing.

6
00:02:05,250 --> 00:02:06,300
Thank you.
```

**Error Response** :

```json
{
  "success": false,
  "message": "Failed to load the subtitle."
}
```
