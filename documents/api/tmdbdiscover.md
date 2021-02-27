# Discover Movies Or Tv Shows

Discover movies or tv shows.

**URL** : `/api/tmdbdiscover/type/{type}/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}`

**Method** : `GET`

**Query Parameters** :

| Parameter   | Type    | Description                                                                                                                                                     |
| ----------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `type`      | string  | Switch between discovering movies or tv shows.                                                                                                                  |
| `genretype` | string  | Filter by genre IDs.                                                                                                                                            |
| `sort`      | string  | Sort by various order.                                                                                                                                          |
| `date`      | string  | Filter and only include movies or tv shows that have a release or air date that is less than or equal to the specified value. Standard date format: YYYY-MM-DD. |
| `lang`      | string  | ISO 639-1 two-letter language code.                                                                                                                             |
| `page`      | integer | Specify the page of results to query.                                                                                                                           |

**Accepted Values [ type ]** :

| Value   | Description        |
| ------- | ------------------ |
| `movie` | Discover movies.   |
| `tv`    | Discover tv shows. |

**Accepted Values [ genretype ]** :

| Movie genretype                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | TV Show genretype                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table> <tr><th>Value</th><th>Genre</th></tr><tr><td>`all`</td><td>All Genre</td></tr><tr><td>`28`</td><td>Action</td></tr><tr><td>`12`</td><td>Adventure</td></tr><tr><td>`16`</td><td>Animation</td></tr><tr><td>`35`</td><td>Comedy</td></tr><tr><td>`80`</td><td>Crime</td></tr><tr><td>`99`</td><td>Documentary</td></tr><tr><td>`18`</td><td>Drama</td></tr><tr><td>`10751`</td><td>Family</td></tr><tr><td>`14`</td><td>Fantasy</td></tr><tr><td>`36`</td><td>History</td></tr><tr><td>`27`</td><td>Horror</td></tr><tr><td>`10402`</td><td>Music</td></tr><tr><td>`9648`</td><td>Mystery</td></tr><tr><td>`10749`</td><td>Romance</td></tr><tr><td>`878`</td><td>Sci-fi</td></tr><tr><td>`53`</td><td>Thriller</td></tr><tr><td>`10752`</td><td>War</td></tr><tr><td>`37`</td><td>Western</td></tr> </table> | <table> <tr><th>Value</th><th>Genre</th></tr><tr><td>`all`</td><td>All Genre</td></tr><tr><td>`10759`</td><td>Action & Adventure</td></tr><tr><td>`16`</td><td>Animation</td></tr><tr><td>`35`</td><td>Comedy</td></tr><tr><td>`80`</td><td>Crime</td></tr><tr><td>`99`</td><td>Documentary</td></tr><tr><td>`18`</td><td>Drama</td></tr><tr><td>`10751`</td><td>Family</td></tr><tr><td>`10762`</td><td>Kids</td></tr><tr><td>`9648`</td><td>Mystery</td></tr><tr><td>`10763`</td><td>News</td></tr><tr><td>`10764`</td><td>Reality</td></tr><tr><td>`10765`</td><td>Sci-fi & Fantasy</td></tr><tr><td>`10766`</td><td>Soap</td></tr><tr><td>`10767`</td><td>Talk</td></tr><tr><td>`10768`</td><td>War & Politics</td></tr><tr><td>`37`</td><td>Western</td></tr><tr><td> </td><td> </td></tr><tr><td> </td><td> </td></tr> </table> |

**Accepted Values [ sort ]** :

| Value                       | Description                 |
| --------------------------- | --------------------------- |
| `popularity.asc`            | Popularity ascending.       |
| `popularity.desc`           | Popularity descending.      |
| `release_date.asc`          | Release date ascending.     |
| `release_date.desc`         | Release date descending.    |
| `revenue.asc`               | Revenue ascending.          |
| `revenue.desc`              | Revenue descending.         |
| `primary_release_date.asc`  | Primary release ascending.  |
| `primary_release_date.desc` | Primary release descending. |
| `original_title.asc`        | Title ascending.            |
| `original_title.desc`       | Title descending.           |
| `vote_average.asc`          | Vote average ascending.     |
| `vote_average.desc`         | Vote average descending.    |
| `vote_count.asc`            | Vote count ascending.       |
| `vote_count.desc`           | Vote count descending.      |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type          | Description                                 |
| --------- | ------------- | ------------------------------------------- |
| `success` | bool          | Indicates whether the query was successful. |
| `results` | array[object] | Contains the complete TMDB API response.    |

**Object [ results ]** :

| Name            | Type          | Description                         |
| --------------- | ------------- | ----------------------------------- |
| `page`          | integer       | Current page number.                |
| `total_results` | integer       | Total number of movies or tv shows. |
| `total_pages`   | integer       | Total number of ages.               |
| `results`       | array[object] | Array of movie or tv show objects.  |

**Object [ results ]** :

| Movie results                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | TV Show results                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table> <tr><th>Name</th><th>Type</th></tr><tr><td>`poster_path`</td><td>string or null</td></tr><tr><td>`adult`</td><td>bool</td></tr><tr><td>`overview`</td><td>string</td></tr><tr><td>`release_date`</td><td>string</td></tr><tr><td>`genre_ids`</td><td>array[integer]</td></tr><tr><td>`id`</td><td>integer</td></tr><tr><td>`original_title`</td><td>string</td></tr><tr><td>`original_language`</td><td>string</td></tr><tr><td>`title`</td><td>string</td></tr><tr><td>`backdrop_path`</td><td>string or null</td></tr><tr><td>`popularity`</td><td>number</td></tr><tr><td>`vote_count`</td><td>integer</td></tr><tr><td>`video`</td><td>bool</td></tr><tr><td>`vote_average`</td><td>number</td></tr> </table> | <table> <tr><th>Name</th><th>Type</th></tr><tr><td>`poster_path`</td><td>string or null</td></tr><tr><td>`popularity`</td><td>number</td></tr><tr><td>`id`</td><td>integer</td></tr><tr><td>`backdrop_path`</td><td>string or null</td></tr><tr><td>`vote_average`</td><td>number</td></tr><tr><td>`overview`</td><td>string</td></tr><tr><td>`first_air_date`</td><td>string</td></tr><tr><td>`origin_country`</td><td>array[string]</td></tr><tr><td>`genre_ids`</td><td>array[integer]</td></tr><tr><td>`original_language`</td><td>string</td></tr><tr><td>`vote_count`</td><td>integer</td></tr><tr><td>`name`</td><td>string</td></tr><tr><td>`original_name`</td><td>string</td></tr><tr><td> </td><td> </td></tr> </table> |

**Decode Table [ genre_ids ]** :

| Movie genre_ids                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | TV Show genre_ids                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table> <tr><th>Value</th><th>Genre</th></tr><tr><td>`28`</td><td>Action</td></tr><tr><td>`12`</td><td>Adventure</td></tr><tr><td>`16`</td><td>Animation</td></tr><tr><td>`35`</td><td>Comedy</td></tr><tr><td>`80`</td><td>Crime</td></tr><tr><td>`99`</td><td>Documentary</td></tr><tr><td>`18`</td><td>Drama</td></tr><tr><td>`10751`</td><td>Family</td></tr><tr><td>`14`</td><td>Fantasy</td></tr><tr><td>`36`</td><td>History</td></tr><tr><td>`27`</td><td>Horror</td></tr><tr><td>`10402`</td><td>Music</td></tr><tr><td>`9648`</td><td>Mystery</td></tr><tr><td>`10749`</td><td>Romance</td></tr><tr><td>`878`</td><td>Sci-fi</td></tr><tr><td>`53`</td><td>Thriller</td></tr><tr><td>`10752`</td><td>War</td></tr><tr><td>`37`</td><td>Western</td></tr> </table> | <table> <tr><th>Value</th><th>Genre</th></tr><tr><td>`10759`</td><td>Action & Adventure</td></tr><tr><td>`16`</td><td>Animation</td></tr><tr><td>`35`</td><td>Comedy</td></tr><tr><td>`80`</td><td>Crime</td></tr><tr><td>`99`</td><td>Documentary</td></tr><tr><td>`18`</td><td>Drama</td></tr><tr><td>`10751`</td><td>Family</td></tr><tr><td>`10762`</td><td>Kids</td></tr><tr><td>`9648`</td><td>Mystery</td></tr><tr><td>`10763`</td><td>News</td></tr><tr><td>`10764`</td><td>Reality</td></tr><tr><td>`10765`</td><td>Sci-fi & Fantasy</td></tr><tr><td>`10766`</td><td>Soap</td></tr><tr><td>`10767`</td><td>Talk</td></tr><tr><td>`10768`</td><td>War & Politics</td></tr><tr><td>`37`</td><td>Western</td></tr><tr><td> </td><td> </td></tr><tr><td> </td><td> </td></tr> </table> |

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
GET http://localhost:9000/api/tmdbdiscover/type/movie/genretype/all/sort/popularity.desc/date/2020-11-06/lang/en/page/1
```

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "page": 1,
      "total_results": 8810,
      "total_pages": 441,
      "results": [
        {
          "popularity": 647.673,
          "vote_count": 6,
          "video": false,
          "poster_path": "/j8D3jXHtA9cNb4epzvmA6hymKQ4.jpg",
          "id": 499338,
          "adult": false,
          "backdrop_path": "/fUxq6ilPW01roOpqB5g9SOS3zZv.jpg",
          "original_language": "en",
          "original_title": "I Believe",
          "genre_ids": [10751],
          "title": "I Believe",
          "vote_average": 3.2,
          "overview": "A 9 year old boy experiences God's power in a supernatural way.",
          "release_date": "2017-11-07"
        },
        {
          "popularity": 522.452,
          "vote_count": 129,
          "video": false,
          "poster_path": "/xqvX5A24dbIWaeYsMTxxKX5qOfz.jpg",
          "id": 660982,
          "adult": false,
          "backdrop_path": "/75ooojtgiKYm5LcCczbCexioZze.jpg",
          "original_language": "en",
          "original_title": "American Pie Presents: Girls' Rules",
          "genre_ids": [35],
          "title": "American Pie Presents: Girls Rules",
          "vote_average": 6.2,
          "overview": "It's Senior year at East Great Falls. Annie, Kayla, Michelle, and Stephanie decide to harness their girl power and band together to get what they want their last year of high school.",
          "release_date": "2020-10-06"
        },
        {
          "popularity": 374.408,
          "vote_count": 12652,
          "video": false,
          "poster_path": "/gGEsBPAijhVUFoiNpgZXqRVWJt2.jpg",
          "id": 354912,
          "adult": false,
          "backdrop_path": "/askg3SMvhqEl4OL52YuvdtY40Yb.jpg",
          "original_language": "en",
          "original_title": "Coco",
          "genre_ids": [16, 10751, 35, 12, 14, 10402],
          "title": "Coco",
          "vote_average": 8.2,
          "overview": "Despite his family’s baffling generations-old ban on music, Miguel dreams of becoming an accomplished musician like his idol, Ernesto de la Cruz. Desperate to prove his talent, Miguel finds himself in the stunning and colorful Land of the Dead following a mysterious chain of events. Along the way, he meets charming trickster Hector, and together, they set off on an extraordinary journey to unlock the real story behind Miguel's family history.",
          "release_date": "2018-02-27"
        }
      ]
    }
  ]
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "No TMDB data found."
}
```
