# Get TV Show Episodes By IMDB Id Or TVDB Id

Get episode list for TV show.

**URL** : `/api/getshowepisodes/{query}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type   | Description                                            |
| --------- | ------ | ------------------------------------------------------ |
| `query`   | string | Query to search. **This value should be URI encoded.** |

**Query Parameters** :

| Parameter | Type   | Description      |
| --------- | ------ | ---------------- |
| `imdb`    | string | IMDB id of show. |
| `tvdb`    | string | TVDB id of show. |

## Success Response

**Code** : `200 OK`

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/getshowepisodes/imdb=tt9140560&tvdb=362392`

**Success Response** :

```json
[
  {
    "id": 1969172,
    "url": "https://www.tvmaze.com/episodes/1969172/wandavision-1x01-filmed-before-a-live-studio-audience",
    "name": "Filmed Before a Live Studio Audience",
    "season": 1,
    "number": 1,
    "type": "regular",
    "airdate": "2021-01-15",
    "airtime": "",
    "airstamp": "2021-01-15T12:00:00+00:00",
    "runtime": 27,
    "image": {
      "medium": "https://static.tvmaze.com/uploads/images/medium_landscape/293/732523.jpg",
      "original": "https://static.tvmaze.com/uploads/images/original_untouched/293/732523.jpg"
    },
    "summary": "<p>Wanda and Vision struggle to conceal their powers during dinner with Vision's boss and his wife.</p>",
    "_links": { "self": { "href": "https://api.tvmaze.com/episodes/1969172" } }
  },
  {
    "id": 2001015,
    "url": "https://www.tvmaze.com/episodes/2001015/wandavision-1x02-dont-touch-that-dial",
    "name": "Don't Touch That Dial",
    "season": 1,
    "number": 2,
    "type": "regular",
    "airdate": "2021-01-15",
    "airtime": "",
    "airstamp": "2021-01-15T12:00:00+00:00",
    "runtime": 34,
    "image": {
      "medium": "https://static.tvmaze.com/uploads/images/medium_landscape/293/732526.jpg",
      "original": "https://static.tvmaze.com/uploads/images/original_untouched/293/732526.jpg"
    },
    "summary": "<p>Wanda Maximoff and Vision bring illusion and glamour to Westview's talent show fundraiser.</p>",
    "_links": { "self": { "href": "https://api.tvmaze.com/episodes/2001015" } }
  },
  ...
]
```

**Error Response** :

```json
{
  "success": false,
  "message": "No TVmaze data found."
}
```
