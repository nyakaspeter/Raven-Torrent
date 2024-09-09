# Raven Torrent API
## Version: 0.7.0


### /about

#### GET
##### Summary:

Get application details

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |

### /add/{base64uri}

#### GET
##### Summary:

Get torrent info and streaming URLs

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| base64uri | path | Link to torrent file / magnet link (base64 encoded) | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TorrentFilesResultsResponse](#v0.TorrentFilesResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /cast/{base64location}/{base64query}

#### GET
##### Summary:

Cast media file to TV or media player

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| base64location | path | Base64 encoded location of the device to cast to | Yes | string |
| base64query | path | Base64 encoded URI encoded query string. Supported parameters: video, subtitle, title | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /delete/{hash}

#### GET
##### Summary:

Delete torrent from torrent client

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Infohash of torrent to delete | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /deleteall

#### GET
##### Summary:

Delete all torrents from torrent client

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getmoviemagnet/imdb/{imdb}/providers/{providers}

#### GET
##### Summary:

Get movie torrents by IMDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the movie | Yes | string |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MovieMagnetLinksResponse](#v0.MovieMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}

#### GET
##### Summary:

Get movie torrents by IMDB id and query string

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the movie | Yes | string |
| query | path | URI encoded query string. Supported parameters: title, releaseyear | Yes | string |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MovieMagnetLinksResponse](#v0.MovieMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getmoviemagnet/query/{query}/providers/{providers}

#### GET
##### Summary:

Get movie torrents by query string

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| query | path | URI encoded query string. Supported parameters: title, releaseyear | Yes | string |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MovieMagnetLinksResponse](#v0.MovieMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getshowmagnet/imdb/{imdb}/query/{query}/season/{season}/episode/{episode}/providers/{providers}

#### GET
##### Summary:

Get show torrents by IMDB id and query string

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the show | Yes | string |
| query | path | URI encoded query string. Supported parameters: title | Yes | string |
| season | path | Season number. Use 0 to search for all seasons | Yes | integer |
| episode | path | Episode number. Use 0 to search for all episodes | Yes | integer |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, eztv, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowMagnetLinksResponse](#v0.ShowMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getshowmagnet/imdb/{imdb}/season/{season}/episode/{episode}/providers/{providers}

#### GET
##### Summary:

Get show torrents by IMDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the show | Yes | string |
| season | path | Season number. Use 0 to search for all seasons | Yes | integer |
| episode | path | Episode number. Use 0 to search for all episodes | Yes | integer |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, eztv, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowMagnetLinksResponse](#v0.ShowMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /getshowmagnet/query/{query}/season/{season}/episode/{episode}/providers/{providers}

#### GET
##### Summary:

Get show torrents by query string

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| query | path | URI encoded query string. Supported parameters: title | Yes | string |
| season | path | Season number. Use 0 to search for all seasons | Yes | integer |
| episode | path | Episode number. Use 0 to search for all episodes | Yes | integer |
| providers | path | Torrent providers to use, separated by comma. Possible values: jackett, eztv, 1337x, itorrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowMagnetLinksResponse](#v0.ShowMagnetLinksResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /mediarenderers

#### GET
##### Summary:

Get list of available casting targets

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MediaRenderersResponse](#v0.MediaRenderersResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /restart/downrate/{downrate}/uprate/{uprate}

#### GET
##### Summary:

Restart torrent client with new bandwith limits

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| downrate | path | Maximum download speed in Kbps. Use 0 for unlimited | Yes | integer |
| uprate | path | Maximum upload speed in Kbps. Use 0 to disable uploading | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /startplayer/{base64path}/{base64args}

#### GET
##### Summary:

Launch media player application

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| base64path | path | Base64 encoded path to the media player executable | Yes | string |
| base64args | path | Base64 encoded launch arguments to pass to the media player | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /stats/{hash}

#### GET
##### Summary:

Get torrent download stats

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Infohash of the torrent | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TorrentStatsResponse](#v0.TorrentStatsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /stop

#### GET
##### Summary:

Shut down the application

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.MessageResponse](#v0.MessageResponse) |

### /subtitlesbyfile/{hash}/{base64path}/lang/{lang}

#### GET
##### Summary:

Get subtitles by torrent's inner file hash

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| hash | path | Infohash of the torrent | Yes | string |
| base64path | path | Base64 encoded path with filename (for example: Season.1/Stranger.Things.S01E01.1080p.mkv, encoded to base64) | Yes | string |
| lang | path | ISO 639-2 three-letter language codes, separated by comma | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.SubtitleFilesResultsResponse](#v0.SubtitleFilesResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}

#### GET
##### Summary:

Get subtitles by IMDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the movie or show | Yes | string |
| lang | path | ISO 639-2 three-letter language codes, separated by comma | Yes | string |
| season | path | Season number. Must be set to 0 for movie subtitle search. | Yes | integer |
| episode | path | Episode number. Must be set to 0 for movie subtitle search. | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.SubtitleFilesResultsResponse](#v0.SubtitleFilesResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}

#### GET
##### Summary:

Get subtitles by text

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| text | path | Title of the movie or show | Yes | string |
| lang | path | ISO 639-2 three-letter language codes, separated by comma | Yes | string |
| season | path | Season number. Must be set to 0 for movie subtitle search. | Yes | integer |
| episode | path | Episode number. Must be set to 0 for movie subtitle search. | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.SubtitleFilesResultsResponse](#v0.SubtitleFilesResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbdiscover/type/movie/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}

#### GET
##### Summary:

Discover movies by genre

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| genretype | path | Genre ids separated by comma, or 'all' to search for all genres. Possible values: 28 (Action), 12	(Adventure), 16	(Animation), 35	(Comedy), 80 (Crime), 99 (Documentary), 18 (Drama), 10751 (Family), 14 (Fantasy), 36 (History), 27 (Horror), 10402 (Music), 9648 (Mystery), 10749 (Romance), 878 (Sci-fi), 53 (Thriller), 10752 (War), 37 (Western) | Yes | string |
| sort | path | Sort order. Possible values: popularity.asc, popularity.desc, release_date.asc, release_date.desc, revenue.asc, revenue.desc, primary_release_date.asc, primary_release_date.desc, original_title.asc, original_title.desc, vote_average.asc, vote_average.desc, vote_count.asc, vote_count.desc | Yes | string |
| date | path | Filter and only include movies or tv shows that have a release or air date that is less than or equal to the specified value. Standard date format: YYYY-MM-DD | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |
| page | path | Specify the page of results to query | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbMovieResultsResponse](#v0.TmdbMovieResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbdiscover/type/tv/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}

#### GET
##### Summary:

Discover shows by genre

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| genretype | path | Genre ids separated by comma, or 'all' to search for all genres. Possible values: 10759 (Action & Adventure), 16 (Animation), 35 (Comedy), 80 (Crime), 99 (Documentary), 18 (Drama), 10751 (Family), 10762 (Kids), 9648 (Mystery), 10763 (News), 10764 (Reality), 10765 (Sci-fi & Fantasy), 10766 (Soap), 10767 (Talk), 10768 (War & Politics), 37 (Western) | Yes | string |
| sort | path | Sort order. Possible values: popularity.asc, popularity.desc, release_date.asc, release_date.desc, revenue.asc, revenue.desc, primary_release_date.asc, primary_release_date.desc, original_title.asc, original_title.desc, vote_average.asc, vote_average.desc, vote_count.asc, vote_count.desc | Yes | string |
| date | path | Filter and only include movies or tv shows that have a release or air date that is less than or equal to the specified value. Standard date format: YYYY-MM-DD | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |
| page | path | Specify the page of results to query | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbShowResultsResponse](#v0.TmdbShowResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbinfo/type/movie/tmdbid/{tmdbid}/lang/{lang}

#### GET
##### Summary:

Get movie details

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tmdbid | path | TMDB id of the movie | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbMovieInfoResponse](#v0.TmdbMovieInfoResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbinfo/type/tv/tmdbid/{tmdbid}/lang/{lang}

#### GET
##### Summary:

Get show details

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tmdbid | path | TMDB id of the show | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbShowInfoResponse](#v0.TmdbShowInfoResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbsearch/type/movie/lang/{lang}/page/{page}/text/{text}

#### GET
##### Summary:

Search movies

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| text | path | Text query to search. Space characters must be replaced with minus or non-breaking space characters. This value should be URI encoded | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |
| page | path | Specify the page of results to query | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbMovieResultsResponse](#v0.TmdbMovieResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tmdbsearch/type/tv/lang/{lang}/page/{page}/text/{text}

#### GET
##### Summary:

Search shows

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| text | path | Text query to search. Space characters must be replaced with minus or non-breaking space characters. This value should be URI encoded | Yes | string |
| lang | path | ISO 639-1 two-letter language code | Yes | string |
| page | path | Specify the page of results to query | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TmdbShowResultsResponse](#v0.TmdbShowResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /torrents

#### GET
##### Summary:

Get list of added torrents

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.TorrentListResultsResponse](#v0.TorrentListResultsResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tvmazeepisodes/imdb/{imdb}

#### GET
##### Summary:

Get show episodes by IMDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| imdb | path | IMDB id of the show | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowEpisodesResponse](#v0.ShowEpisodesResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tvmazeepisodes/tvdb/{tvdb}

#### GET
##### Summary:

Get show episodes by TVDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tvdb | path | TVDB id of the show | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowEpisodesResponse](#v0.ShowEpisodesResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### /tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}

#### GET
##### Summary:

Get show episodes by TVDB id and IMDB id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tvdb | path | TVDB id of the show | Yes | string |
| imdb | path | IMDB id of the show | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [v0.ShowEpisodesResponse](#v0.ShowEpisodesResponse) |
| 404 | Not Found | [v0.MessageResponse](#v0.MessageResponse) |

### Models


#### types.Company

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | No |
| logo_path | string |  | No |
| name | string |  | No |
| origin_country | string |  | No |

#### types.Country

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| iso_3166_1 | string |  | No |
| name | string |  | No |

#### types.Creator

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| credit_id | string |  | No |
| gender | integer |  | No |
| id | integer |  | No |
| name | string |  | No |
| profile_path | string |  | No |

#### types.Episode

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| air_date | string |  | No |
| episode_number | integer |  | No |
| id | integer |  | No |
| name | string |  | No |
| overview | string |  | No |
| production_code | string |  | No |
| runtime | integer |  | No |
| season_number | integer |  | No |
| still_path | string |  | No |
| vote_average | number |  | No |
| vote_count | integer |  | No |

#### types.EpisodeImages

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| medium | string |  | No |
| original | string |  | No |

#### types.EpisodeLink

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| href | string |  | No |

#### types.EpisodeLinks

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| self | [types.EpisodeLink](#types.EpisodeLink) |  | No |

#### types.ExternalIds

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| facebook_id | string |  | No |
| freebase_id | string |  | No |
| freebase_mid | string |  | No |
| imdb_id | string |  | No |
| instagram_id | string |  | No |
| tvdb_id | integer |  | No |
| tvrage_id | integer |  | No |
| twitter_id | string |  | No |

#### types.Genre

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | No |
| name | string |  | No |

#### types.Language

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| iso_639_1 | string |  | No |
| name | string |  | No |

#### types.MediaDevice

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| location | string |  | No |
| name | string |  | No |

#### types.Movie

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| adult | boolean |  | No |
| backdrop_path | string |  | No |
| genre_ids | [ integer ] |  | No |
| id | integer |  | No |
| original_language | string |  | No |
| original_title | string |  | No |
| overview | string |  | No |
| popularity | number |  | No |
| poster_path | string |  | No |
| release_date | string |  | No |
| title | string |  | No |
| video | boolean |  | No |
| vote_average | number |  | No |
| vote_count | integer |  | No |

#### types.MovieInfo

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| adult | boolean |  | No |
| backdrop_path | string |  | No |
| budget | integer |  | No |
| genres | [ [types.Genre](#types.Genre) ] |  | No |
| homepage | string |  | No |
| id | integer |  | No |
| imdb_id | string |  | No |
| original_language | string |  | No |
| original_title | string |  | No |
| overview | string |  | No |
| popularity | number |  | No |
| poster_path | string |  | No |
| production_companies | [ [types.Company](#types.Company) ] |  | No |
| production_countries | [ [types.Country](#types.Country) ] |  | No |
| release_date | string |  | No |
| revenue | integer |  | No |
| runtime | integer |  | No |
| spoken_languages | [ [types.Language](#types.Language) ] |  | No |
| status | string |  | No |
| tagline | string |  | No |
| title | string |  | No |
| video | boolean |  | No |
| vote_average | number |  | No |
| vote_count | integer |  | No |

#### types.MovieResults

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| page | integer |  | No |
| results | [ [types.Movie](#types.Movie) ] |  | No |
| total_pages | integer |  | No |
| total_results | integer |  | No |

#### types.MovieTorrent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| hash | string |  | No |
| lang | string |  | No |
| magnet | string |  | No |
| peers | string |  | No |
| provider | string |  | No |
| quality | string |  | No |
| seeds | string |  | No |
| size | string |  | No |
| title | string |  | No |
| torrent | string |  | No |

#### types.Season

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| air_date | string |  | No |
| episode_count | integer |  | No |
| id | integer |  | No |
| name | string |  | No |
| overview | string |  | No |
| poster_path | string |  | No |
| season_number | integer |  | No |

#### types.Show

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| backdrop_path | string |  | No |
| first_air_date | string |  | No |
| genre_ids | [ integer ] |  | No |
| id | integer |  | No |
| name | string |  | No |
| origin_country | [ string ] |  | No |
| original_language | string |  | No |
| original_name | string |  | No |
| overview | string |  | No |
| popularity | number |  | No |
| poster_path | string |  | No |
| vote_average | number |  | No |
| vote_count | integer |  | No |

#### types.ShowInfo

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| adult | boolean |  | No |
| backdrop_path | string |  | No |
| created_by | [ [types.Creator](#types.Creator) ] |  | No |
| episode_run_time | [ integer ] |  | No |
| external_ids | [types.ExternalIds](#types.ExternalIds) |  | No |
| first_air_date | string |  | No |
| genres | [ [types.Genre](#types.Genre) ] |  | No |
| homepage | string |  | No |
| id | integer |  | No |
| in_production | boolean |  | No |
| languages | [ string ] |  | No |
| last_air_date | string |  | No |
| last_episode_to_air | [types.Episode](#types.Episode) |  | No |
| name | string |  | No |
| networks | [ [types.Company](#types.Company) ] |  | No |
| next_episode_to_air | [types.Episode](#types.Episode) |  | No |
| number_of_episodes | integer |  | No |
| number_of_seasons | integer |  | No |
| origin_country | [ string ] |  | No |
| original_language | string |  | No |
| original_name | string |  | No |
| overview | string |  | No |
| popularity | number |  | No |
| poster_path | string |  | No |
| production_companies | [ [types.Company](#types.Company) ] |  | No |
| production_countries | [ [types.Country](#types.Country) ] |  | No |
| seasons | [ [types.Season](#types.Season) ] |  | No |
| spoken_languages | [ [types.Language](#types.Language) ] |  | No |
| status | string |  | No |
| tagline | string |  | No |
| type | string |  | No |
| vote_average | number |  | No |
| vote_count | integer |  | No |

#### types.ShowResults

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| page | integer |  | No |
| results | [ [types.Show](#types.Show) ] |  | No |
| total_pages | integer |  | No |
| total_results | integer |  | No |

#### types.ShowTorrent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| episode | string |  | No |
| hash | string |  | No |
| lang | string |  | No |
| magnet | string |  | No |
| peers | string |  | No |
| provider | string |  | No |
| quality | string |  | No |
| season | string |  | No |
| seeds | string |  | No |
| size | string |  | No |
| title | string |  | No |
| torrent | string |  | No |

#### types.SubtitleFile

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| lang | string |  | No |
| releasename | string |  | No |
| subdata | string |  | No |
| subencoding | string |  | No |
| subformat | string |  | No |
| subtitlename | string |  | No |
| vttdata | string |  | No |

#### types.TorrentFile

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| length | string |  | No |
| name | string |  | No |
| url | string |  | No |

#### types.TvMazeEpisode

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _links | [types.EpisodeLinks](#types.EpisodeLinks) |  | No |
| airdate | string |  | No |
| airstamp | string |  | No |
| airtime | string |  | No |
| id | integer |  | No |
| image | [types.EpisodeImages](#types.EpisodeImages) |  | No |
| name | string |  | No |
| number | integer |  | No |
| runtime | integer |  | No |
| season | integer |  | No |
| summary | string |  | No |
| type | string |  | No |
| url | string |  | No |

#### v0.MediaRenderersResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [types.MediaDevice](#types.MediaDevice) ] |  | No |
| success | boolean |  | No |

#### v0.MessageResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| success | boolean |  | No |

#### v0.MovieMagnetLinksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [types.MovieTorrent](#types.MovieTorrent) ] |  | No |
| success | boolean |  | No |

#### v0.ShowEpisodesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [types.TvMazeEpisode](#types.TvMazeEpisode) ] |  | No |
| success | boolean |  | No |

#### v0.ShowMagnetLinksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [types.ShowTorrent](#types.ShowTorrent) ] |  | No |
| success | boolean |  | No |

#### v0.SubtitleFilesResultsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [types.SubtitleFile](#types.SubtitleFile) ] |  | No |
| success | boolean |  | No |

#### v0.TmdbMovieInfoResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| result | [types.MovieInfo](#types.MovieInfo) |  | No |
| success | boolean |  | No |

#### v0.TmdbMovieResultsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [types.MovieResults](#types.MovieResults) |  | No |
| success | boolean |  | No |

#### v0.TmdbShowInfoResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| result | [types.ShowInfo](#types.ShowInfo) |  | No |
| success | boolean |  | No |

#### v0.TmdbShowResultsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [types.ShowResults](#types.ShowResults) |  | No |
| success | boolean |  | No |

#### v0.TorrentFilesResultsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| hash | string |  | No |
| results | [ [types.TorrentFile](#types.TorrentFile) ] |  | No |
| success | boolean |  | No |

#### v0.TorrentListResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| hash | string |  | No |
| length | string |  | No |
| name | string |  | No |

#### v0.TorrentListResultsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| results | [ [v0.TorrentListResponse](#v0.TorrentListResponse) ] |  | No |
| success | boolean |  | No |

#### v0.TorrentStatsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| downdata | string |  | No |
| downpercent | string |  | No |
| downspeed | string |  | No |
| fulldata | string |  | No |
| peers | string |  | No |
| success | boolean |  | No |