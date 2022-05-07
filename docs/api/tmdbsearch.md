# Search movies or tv shows by query text

Search movies or tv shows by query text.

**URL** : `/api/tmdbsearch/type/{type}/lang/{lang}/page/{page}/text/{text}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type    | Description                                                                                                                                |
| --------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `type`    | string  | Switch between searching movies or tv shows.                                                                                               |
| `lang`    | string  | ISO 639-1 two-letter language code.                                                                                                        |
| `page`    | integer | Specify the page of results to query.                                                                                                      |
| `text`    | string  | Text query to search. **Space characters must be replaced with minus or non-breaking space characters. This value should be URI encoded.** |

**Accepted Values [ type ]** :

| Value   | Description        |
| ------- | ------------------ |
| `movie` | Discover movies.   |
| `tv`    | Discover tv shows. |

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

`GET http://localhost:9000/api/tmdbsearch/type/tv/lang/en/page/1/text/raised-by-wolves`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "page": 1,
      "total_results": 2,
      "total_pages": 1,
      "results": [
        {
          "original_name": "Raised by Wolves",
          "genre_ids": [10765, 18],
          "name": "Raised by Wolves",
          "popularity": 168.868,
          "origin_country": ["US"],
          "vote_count": 433,
          "first_air_date": "2020-09-03",
          "backdrop_path": "/na2xUduK8HviOFT97TiFG2MkJmY.jpg",
          "original_language": "en",
          "id": 85723,
          "vote_average": 7.7,
          "overview": "After Earth is ravaged by a great religious war, an atheistic android architect sends two of his creations, Mother and Father, to start a peaceful, godless colony on the planet Kepler-22b. Their treacherous task is jeopardized by the arrival of the Mithraic, a deeply devout religious order of surviving humans.",
          "poster_path": "/mTvSVKMn2Npf6zvYNbGMJnYLtvp.jpg"
        },
        {
          "original_name": "Raised by Wolves",
          "genre_ids": [35],
          "name": "Raised by Wolves",
          "popularity": 5.674,
          "origin_country": ["GB"],
          "vote_count": 8,
          "first_air_date": "2015-03-16",
          "backdrop_path": "/cSA5pPdkEgFDWFcjdr4S2IohLXd.jpg",
          "original_language": "en",
          "id": 62160,
          "vote_average": 7.5,
          "overview": "Set on a Wolverhampton council estate, Raised By Wolves is modern day reimagining of the childhood of Caitlin Moran and her brothers and sisters.\n\nSingle-mum Della lives in a three bedroom council house with Germaine, Aretha, Yoko, Mariah, Wyatt and baby Cher. She is attempting to raise the children by herself, but does have visits from Grampy, who likes to come around to dispense his wisdom to his grandchildren.",
          "poster_path": "/54ONCnoHQZdMnPZU90QN8ezjhS2.jpg"
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
