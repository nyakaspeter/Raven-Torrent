# Get more info about movie or tv show by TMDB id

Get more info about movie or tv show by TMDB id.

**URL** : `/api/tmdbinfo/type/{type}/tmdbid/{tmdbid}/lang/{lang}`

**Method** : `GET`

**Query Parameters** :

| Parameter | Type    | Description                                    |
| --------- | ------- | ---------------------------------------------- |
| `type`    | string  | Switch between discovering movies or tv shows. |
| `tmdbid`  | integer | TMDB movie or tv show id.                      |
| `lang`    | string  | ISO 639-1 two-letter language code.            |

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

| Movie results                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | TV Show results                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table> <tr><th>Name</th><th>Type</th></tr><tr><td>`adult`</td><td>bool</td></tr><tr><td>`backdrop_path`</td><td>string or null</td></tr><tr><td>`belongs_to_collection`</td><td>object or null</td></tr><tr><td>`budget`</td><td>integer</td></tr><tr><td>`genres`</td><td>array[object]</td></tr><tr><td>`homepage`</td><td>string or null</td></tr><tr><td>`id`</td><td>integer</td></tr><tr><td>`imdb_id`</td><td>string or null</td></tr><tr><td>`original_language`</td><td>string</td></tr><tr><td>`original_title`</td><td>string</td></tr><tr><td>`overview`</td><td>string or null</td></tr><tr><td>`popularity`</td><td>number</td></tr><tr><td>`poster_path`</td><td>string or null</td></tr><tr><td>`production_companies`</td><td>array[object]</td></tr><tr><td>`production_countries`</td><td>array[object]</td></tr><tr><td>`release_date`</td><td>string</td></tr><tr><td>`revenue`</td><td>integer</td></tr><tr><td>`runtime`</td><td>integer or null</td></tr><tr><td>`spoken_languages`</td><td>array[object]</td></tr><tr><td>`status`</td><td>string</td></tr><tr><td>`tagline`</td><td>string or null</td></tr><tr><td>`title`</td><td>string</td></tr><tr><td>`video`</td><td>bool</td></tr><tr><td>`vote_average`</td><td>number</td></tr><tr><td>`vote_count`</td><td>integer</td></tr><tr><td> </td><td> </td></tr><tr><td> </td><td> </td></tr><tr><td> </td><td> </td></tr> </table> | <table> <tr><th>Name</th><th>Type</th></tr><tr><td>`backdrop_path`</td><td>string or null</td></tr><tr><td>`created_by`</td><td>array[object]</td></tr><tr><td>`episode_run_time`</td><td>array[integer]</td></tr><tr><td>`first_air_date`</td><td>string</td></tr><tr><td>`genres`</td><td>array[object]</td></tr><tr><td>`homepage`</td><td>string</td></tr><tr><td>`id`</td><td>integer</td></tr><tr><td>`in_production`</td><td>bool</td></tr><tr><td>`languages`</td><td>array[string]</td></tr><tr><td>`last_air_date`</td><td>string</td></tr><tr><td>`name`</td><td>string</td></tr><tr><td>`next_episode_to_air`</td><td>object or null</td></tr><tr><td>`networks`</td><td>array[object]</td></tr><tr><td>`number_of_episodes`</td><td>integer</td></tr><tr><td>`number_of_seasons`</td><td>integer</td></tr><tr><td>`origin_country`</td><td>array[string]</td></tr><tr><td>`original_language`</td><td>string</td></tr><tr><td>`original_name`</td><td>string</td></tr><tr><td>`overview`</td><td>string</td></tr><tr><td>`popularity`</td><td>number</td></tr><tr><td>`poster_path`</td><td>string or null</td></tr><tr><td>`production_companies`</td><td>array[object]</td></tr><tr><td>`seasons`</td><td>array[object]</td></tr><tr><td>`status`</td><td>string</td></tr><tr><td>`type`</td><td>string</td></tr><tr><td>`vote_average`</td><td>number</td></tr><tr><td>`vote_count`</td><td>integer</td></tr><tr><td>`external_ids`</td><td>object</td></tr> </table> |

**Object [ belongs_to_collection ]** :

| Name            | Type    |
| --------------- | ------- |
| `id`            | integer |
| `name`          | string  |
| `poster_path`   | string  |
| `backdrop_path` | string  |

**Object [ genres ]** :

| Name   | Type    |
| ------ | ------- |
| `id`   | integer |
| `name` | string  |

**Object [ production_companies ] and [ networks ]** :

| Name             | Type           |
| ---------------- | -------------- |
| `name`           | string         |
| `id`             | integer        |
| `logo_path`      | string or null |
| `origin_country` | string         |

**Object [ production_countries ]** :

| Name         | Type   |
| ------------ | ------ |
| `iso_3166_1` | string |
| `name`       | string |

**Object [ spoken_languages ]** :

| Name        | Type   |
| ----------- | ------ |
| `iso_639_1` | string |
| `name`      | string |

**Object [ created_by ]** :

| Name           | Type    |
| -------------- | ------- |
| `id`           | integer |
| `credit_id`    | string  |
| `name`         | string  |
| `gender`       | integer |
| `profile_path` | string  |

**Object [ seasons ]** :

| Name            | Type    |
| --------------- | ------- |
| `air_date`      | string  |
| `episode_count` | integer |
| `id`            | integer |
| `name`          | string  |
| `overview`      | string  |
| `poster_path`   | string  |
| `season_number` | string  |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/tmdbinfo/type/movie/tmdbid/590223/lang/en`

**Success Response** :

```json
{
  "success": true,
  "results": [
    {
      "adult": false,
      "backdrop_path": "/lA5fOBqTOQBQ1s9lEYYPmNXoYLi.jpg",
      "belongs_to_collection": null,
      "budget": 30000000,
      "genres": [
        {
          "id": 28,
          "name": "Action"
        },
        {
          "id": 12,
          "name": "Adventure"
        },
        {
          "id": 35,
          "name": "Comedy"
        },
        {
          "id": 878,
          "name": "Science Fiction"
        }
      ],
      "homepage": "https://www.paramountmovies.com/movies/love-and-monsters",
      "id": 590223,
      "imdb_id": "tt2222042",
      "original_language": "en",
      "original_title": "Love and Monsters",
      "overview": "Seven years after the Monsterpocalypse, Joel Dawson, along with the rest of humanity, has been living underground ever since giant creatures took control of the land. After reconnecting over radio with his high school girlfriend Aimee, who is now 80 miles away at a coastal colony, Joel begins to fall for her again. As Joel realizes that there’s nothing left for him underground, he decides against all logic to venture out to Aimee, despite all the dangerous monsters that stand in his way.",
      "popularity": 509.995,
      "poster_path": "/r4Lm1XKP0VsTgHX4LG4syAwYA2I.jpg",
      "production_companies": [
        {
          "id": 2575,
          "logo_path": "/9YJrHYlcfHtwtulkFMAies3aFEl.png",
          "name": "21 Laps Entertainment",
          "origin_country": "US"
        },
        {
          "id": 96540,
          "logo_path": "/AgYjTNeIKOh0yvegPLSjq8EOsif.png",
          "name": "Paramount Players",
          "origin_country": "US"
        },
        {
          "id": 746,
          "logo_path": "/kc7bdIVTBkJYy9aDK1QDDTAL463.png",
          "name": "MTV Films",
          "origin_country": "US"
        },
        {
          "id": 13785,
          "logo_path": null,
          "name": "Aurum Producciones",
          "origin_country": "ES"
        },
        {
          "id": 4,
          "logo_path": "/fycMZt242LVjagMByZOLUGbCvv3.png",
          "name": "Paramount",
          "origin_country": "US"
        }
      ],
      "production_countries": [
        {
          "iso_3166_1": "US",
          "name": "United States of America"
        }
      ],
      "release_date": "2020-10-16",
      "revenue": 0,
      "runtime": 109,
      "spoken_languages": [
        {
          "iso_639_1": "en",
          "name": "English"
        }
      ],
      "status": "Released",
      "tagline": "",
      "title": "Love and Monsters",
      "video": false,
      "vote_average": 7.7,
      "vote_count": 286
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
