package types

type MovieDiscoverParams struct {
	GenreIds       []int  `json:"genreIds"`
	MaxReleaseDate string `json:"maxReleaseDate"`
	MinReleaseDate string `json:"minReleaseDate"`
	SortBy         string `json:"sortBy"`
}

type ShowDiscoverParams struct {
	GenreIds   []int  `json:"genreIds"`
	MaxAirDate string `json:"maxAirDate"`
	MinAirDate string `json:"minAirDate"`
	SortBy     string `json:"sortBy"`
}

type MovieResults struct {
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Results      []Movie `json:"results"`
}

type Movie struct {
	Id               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	ReleaseDate      string  `json:"release_date"`
	Description      string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	GenreIds         []int   `json:"genre_ids"`
	Video            bool    `json:"video"`
	Adult            bool    `json:"adult"`
}

type ShowResults struct {
	Page         int    `json:"page"`
	TotalPages   int    `json:"total_pages"`
	TotalResults int    `json:"total_results"`
	Results      []Show `json:"results"`
}

type Show struct {
	Id               int      `json:"id"`
	Title            string   `json:"name"`
	OriginalTitle    string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	FirstAirDate     string   `json:"first_air_date"`
	Description      string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	BackdropPath     string   `json:"backdrop_path"`
	Popularity       float64  `json:"popularity"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	GenreIds         []int    `json:"genre_ids"`
}

type MovieInfo struct {
	Id                  int          `json:"id"`
	ImdbId              string       `json:"imdb_id"`
	Title               string       `json:"title"`
	OriginalTitle       string       `json:"original_title"`
	OriginalLanguage    string       `json:"original_language"`
	ReleaseDate         string       `json:"release_date"`
	Description         string       `json:"overview"`
	Tagline             string       `json:"tagline"`
	PosterPath          string       `json:"poster_path"`
	BackdropPath        string       `json:"backdrop_path"`
	Popularity          float64      `json:"popularity"`
	VoteAverage         float64      `json:"vote_average"`
	VoteCount           int          `json:"vote_count"`
	Genres              []Genre      `json:"genres"`
	Collections         []Collection `json:"belongs_to_collection"`
	ProductionCompanies []Company    `json:"production_companies"`
	ProductionCountries []Country    `json:"production_countries"`
	SpokenLanguages     []Language   `json:"spoken_languages"`
	RuntimeMinutes      int          `json:"runtime"`
	Budget              int          `json:"budget"`
	Revenue             int          `json:"revenue"`
	Status              string       `json:"status"`
	Homepage            string       `json:"homepage"`
	Video               bool         `json:"video"`
	Adult               bool         `json:"adult"`
}

type ShowInfo struct {
	Id                  int         `json:"id"`
	ExternalIds         ExternalIds `json:"external_ids"`
	Title               string      `json:"name"`
	OriginalTitle       string      `json:"original_name"`
	OriginCountry       []string    `json:"origin_country"`
	Languages           []string    `json:"languages"`
	OriginalLanguage    string      `json:"original_language"`
	FirstAirDate        string      `json:"first_air_date"`
	LastAirDate         string      `json:"last_air_date"`
	LastEpisode         Episode     `json:"last_episode_to_air"`
	NextEpisode         Episode     `json:"next_episode_to_air"`
	Seasons             []Season    `json:"seasons"`
	Description         string      `json:"overview"`
	Tagline             string      `json:"tagline"`
	PosterPath          string      `json:"poster_path"`
	BackdropPath        string      `json:"backdrop_path"`
	Popularity          float64     `json:"popularity"`
	VoteAverage         float64     `json:"vote_average"`
	VoteCount           int         `json:"vote_count"`
	SeasonCount         int         `json:"number_of_seasons"`
	EpisodeCount        int         `json:"number_of_episodes"`
	Genres              []Genre     `json:"genres"`
	CreatedBy           []Creator   `json:"created_by"`
	Networks            []Company   `json:"networks"`
	ProductionCompanies []Company   `json:"production_companies"`
	ProductionCountries []Country   `json:"production_countries"`
	SpokenLanguages     []Language  `json:"spoken_languages"`
	RuntimeMinutes      []int       `json:"episode_run_time"`
	Homepage            string      `json:"homepage"`
	InProduction        bool        `json:"in_production"`
	Status              string      `json:"status"`
	Type                string      `json:"type"`
	Adult               bool        `json:"adult"`
}

type ExternalIds struct {
	ImdbId      string `json:"imdb_id"`
	FreebaseMid string `json:"freebase_mid"`
	FreebaseId  string `json:"freebase_id"`
	TvdbId      int    `json:"tvdb_id"`
	TvrageId    int    `json:"tvrage_id"`
	FacebookId  string `json:"facebook_id"`
	InstagramId string `json:"instagram_id"`
	TwitterId   string `json:"twitter_id"`
}

type Episode struct {
	Id             int     `json:"id"`
	Title          string  `json:"name"`
	SeasonNumber   int     `json:"season_number"`
	EpisodeNumber  int     `json:"episode_number"`
	AirDate        string  `json:"air_date"`
	Description    string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	RuntimeMinutes int     `json:"runtime"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type Season struct {
	Id           int    `json:"id"`
	SeasonNumber int    `json:"season_number"`
	EpisodeCount int    `json:"episode_count"`
	Title        string `json:"name"`
	AirDate      string `json:"air_date"`
	Description  string `json:"overview"`
	PosterPath   string `json:"poster_path"`
}

type SeasonInfo struct {
	Id           int       `json:"id"`
	SeasonNumber int       `json:"season_number"`
	Episodes     []Episode `json:"episodes"`
	Title        string    `json:"name"`
	AirDate      string    `json:"air_date"`
	Description  string    `json:"overview"`
	PosterPath   string    `json:"poster_path"`
}

type Creator struct {
	Id          int    `json:"id"`
	CreditId    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Collection struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type Company struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type Country struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type Language struct {
	Iso6391 string `json:"iso_639_1"`
	Name    string `json:"name"`
}

type ShowIds struct {
	ImdbId   string `json:"imdbId"`
	TvdbId   string `json:"tvdbId"`
	TvMazeId string `json:"tvMazeId"`
}

type TvMazeEpisode struct {
	TvMazeId        int           `json:"id"`
	TvMazeUrl       string        `json:"url"`
	Title           string        `json:"name"`
	SeasonNumber    int           `json:"season"`
	EpisodeNumber   int           `json:"inumberd"`
	Type            string        `json:"type"`
	FirstAirDate    string        `json:"airdate"`
	FirstAirTime    string        `json:"airtime"`
	FirstAirDateUtc string        `json:"airstamp"`
	RuntimeMinutes  int           `json:"runtime"`
	Description     string        `json:"summary"`
	Images          EpisodeImages `json:"image"`
	Links           EpisodeLinks  `json:"_links"`
}

type EpisodeImages struct {
	MediumImageUrl   string `json:"medium"`
	OriginalImageUrl string `json:"original"`
}

type EpisodeLinks struct {
	Self EpisodeLink `json:"self"`
}

type EpisodeLink struct {
	Href string `json:"href"`
}
