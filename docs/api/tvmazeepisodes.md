# Get tv show episodes by IMDB id or TVDB id or both

Get all available episodes for a tv show.

**URL** 

- Search by IMDB id: `/api/tvmazeepisodes/imdb/{imdb}`
- Search by TVDB id: `/api/tvmazeepisodes/tvdb/{tvdb}`
- Search by both at once: `/api/tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type | Description |
| --- | --- | --- |
| `tvdb` | string | Internet Movie Database identifier.|
| `imdb` | string | TheTVDB identifier.|

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name | Type | Description |
| --- | --- | --- |
| `success` | bool | Indicates whether the query was successful.|
| `results` | array[object] | Array of objects.|

**Object [ results ]** :

| Name | Type | Description |
| --- | --- | --- |
| `id` | integer | TVMAZE identifier.|
| `url` | string | Episode url.|
| `name` | string | Episode name.|
| `season` | integer | Season number.|
| `number` | integer | Episode number.|
| `type` | string | Type.|
| `airdate` | string | First air date.|
| `airtime` | string | First air time.|
| `airstamp` | string | First air stamp.|
| `runtime` | integer | Runtime.|
| `image` | object | Object of images.|
| `summary` | string | Summary.|
| `_links` | object | Object of known urls that related to the episode.|

**Object [ image ]** :

| Name | Type | Description |
| --- | --- | --- |
| `medium` | string | Medium size image url.|
| `original` | string | Original size image url.|

**Object [ _links ]** :

| Name | Type | Description |
| --- | --- | --- |
| `self` | object | Object of internal TVMAZE episode API urls.|

**Object [ self ]** :

| Name | Type | Description |
| --- | --- | --- |
| `href` | string | TVMAZE episode API url.|

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name | Type | Description |
| --- | --- | --- |
| `success` | bool | Indicates whether the query was successful.|
| `message` | string | Text message that describes the response.|

## Examples

**Request** :

```
GET http://localhost:9000/api/tvmazeepisodes/tvdb/368166/imdb/tt8690918
```

**Success Response** :

```json
{
    "success": true,
    "results": [{
        "id": 1974915,
        "url": "https://www.tvmaze.com/episodes/1974915/resident-alien-1x01-pilot",
        "name": "Pilot",
        "season": 1,
        "number": 1,
        "type": "regular",
        "airdate": "2021-01-27",
        "airtime": "22:00",
        "airstamp": "2021-01-28T03:00:00+00:00",
        "runtime": 55,
        "image": {
            "medium": "https://static.tvmaze.com/uploads/images/medium_landscape/290/725605.jpg",
            "original": "https://static.tvmaze.com/uploads/images/original_untouched/290/725605.jpg"
        },
        "summary": "<p>An alien crashes on Earth and hides in a remote Colorado mountain town; after assuming the identity of the town doctor, his nefarious mission is threatened when he realizes one of the townspeople, a 9-year-old boy, can see his true alien form.</p>",
        "_links": {
            "self": {
                "href": "https://api.tvmaze.com/episodes/1974915"
            }
        }
    }, {
        "id": 2013099,
        "url": "https://www.tvmaze.com/episodes/2013099/resident-alien-1x02-homesick",
        "name": "Homesick",
        "season": 1,
        "number": 2,
        "type": "regular",
        "airdate": "2021-02-03",
        "airtime": "22:00",
        "airstamp": "2021-02-04T03:00:00+00:00",
        "runtime": 60,
        "image": {
            "medium": "https://static.tvmaze.com/uploads/images/medium_landscape/295/738798.jpg",
            "original": "https://static.tvmaze.com/uploads/images/original_untouched/295/738798.jpg"
        },
        "summary": "<p>In his first week at the clinic, Harry struggles to diagnose a strange feeling.</p>",
        "_links": {
            "self": {
                "href": "https://api.tvmaze.com/episodes/2013099"
            }
        }
    }]
}
```

**Error Response** :

```json
{
    "success": false,
    "message": "No TVMaze data found."
}
```