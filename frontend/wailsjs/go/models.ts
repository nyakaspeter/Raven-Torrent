export namespace types {
	
	export class EztvParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new EztvParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class EpisodeParams {
	    season: number;
	    episode: number;
	
	    static createFrom(source: any = {}) {
	        return new EpisodeParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.season = source["season"];
	        this.episode = source["episode"];
	    }
	}
	export class MovieParams {
	    imdbId: string;
	    searchText: string;
	
	    static createFrom(source: any = {}) {
	        return new MovieParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.imdbId = source["imdbId"];
	        this.searchText = source["searchText"];
	    }
	}
	export class YtsParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new YtsParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class ItorrentParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ItorrentParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class X1337xParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new X1337xParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class MovieTorrent {
	    hash: string;
	    quality: string;
	    size: string;
	    provider: string;
	    lang: string;
	    title: string;
	    seeds: string;
	    peers: string;
	    magnet: string;
	    torrent: string;
	
	    static createFrom(source: any = {}) {
	        return new MovieTorrent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hash = source["hash"];
	        this.quality = source["quality"];
	        this.size = source["size"];
	        this.provider = source["provider"];
	        this.lang = source["lang"];
	        this.title = source["title"];
	        this.seeds = source["seeds"];
	        this.peers = source["peers"];
	        this.magnet = source["magnet"];
	        this.torrent = source["torrent"];
	    }
	}
	export class ShowDiscoverParams {
	    genreIds: number[];
	    maxAirDate: string;
	    minAirDate: string;
	    sortBy: string;
	
	    static createFrom(source: any = {}) {
	        return new ShowDiscoverParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.genreIds = source["genreIds"];
	        this.maxAirDate = source["maxAirDate"];
	        this.minAirDate = source["minAirDate"];
	        this.sortBy = source["sortBy"];
	    }
	}
	export class Language {
	    iso_639_1: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Language(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.iso_639_1 = source["iso_639_1"];
	        this.name = source["name"];
	    }
	}
	export class Country {
	    iso_3166_1: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Country(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.iso_3166_1 = source["iso_3166_1"];
	        this.name = source["name"];
	    }
	}
	export class Company {
	    id: number;
	    name: string;
	    logo_path: string;
	    origin_country: string;
	
	    static createFrom(source: any = {}) {
	        return new Company(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.logo_path = source["logo_path"];
	        this.origin_country = source["origin_country"];
	    }
	}
	export class Collection {
	    id: number;
	    name: string;
	    poster_path: string;
	    backdrop_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Collection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	    }
	}
	export class Genre {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Genre(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class MovieInfo {
	    id: number;
	    imdb_id: string;
	    title: string;
	    original_title: string;
	    original_language: string;
	    release_date: string;
	    overview: string;
	    tagline: string;
	    poster_path: string;
	    backdrop_path: string;
	    popularity: number;
	    vote_average: number;
	    vote_count: number;
	    genres: Genre[];
	    belongs_to_collection: Collection[];
	    production_companies: Company[];
	    production_countries: Country[];
	    spoken_languages: Language[];
	    runtime: number;
	    budget: number;
	    revenue: number;
	    status: string;
	    homepage: string;
	    video: boolean;
	    adult: boolean;
	
	    static createFrom(source: any = {}) {
	        return new MovieInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.imdb_id = source["imdb_id"];
	        this.title = source["title"];
	        this.original_title = source["original_title"];
	        this.original_language = source["original_language"];
	        this.release_date = source["release_date"];
	        this.overview = source["overview"];
	        this.tagline = source["tagline"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.popularity = source["popularity"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.genres = this.convertValues(source["genres"], Genre);
	        this.belongs_to_collection = this.convertValues(source["belongs_to_collection"], Collection);
	        this.production_companies = this.convertValues(source["production_companies"], Company);
	        this.production_countries = this.convertValues(source["production_countries"], Country);
	        this.spoken_languages = this.convertValues(source["spoken_languages"], Language);
	        this.runtime = source["runtime"];
	        this.budget = source["budget"];
	        this.revenue = source["revenue"];
	        this.status = source["status"];
	        this.homepage = source["homepage"];
	        this.video = source["video"];
	        this.adult = source["adult"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Show {
	    id: number;
	    name: string;
	    original_name: string;
	    origin_country: string[];
	    original_language: string;
	    first_air_date: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    popularity: number;
	    vote_average: number;
	    vote_count: number;
	    genre_ids: number[];
	
	    static createFrom(source: any = {}) {
	        return new Show(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.original_name = source["original_name"];
	        this.origin_country = source["origin_country"];
	        this.original_language = source["original_language"];
	        this.first_air_date = source["first_air_date"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.popularity = source["popularity"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.genre_ids = source["genre_ids"];
	    }
	}
	export class ShowResults {
	    page: number;
	    total_pages: number;
	    total_results: number;
	    results: Show[];
	
	    static createFrom(source: any = {}) {
	        return new ShowResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.total_pages = source["total_pages"];
	        this.total_results = source["total_results"];
	        this.results = this.convertValues(source["results"], Show);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MediaDevice {
	    name: string;
	    location: string;
	
	    static createFrom(source: any = {}) {
	        return new MediaDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.location = source["location"];
	    }
	}
	export class RarbgParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RarbgParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class PopcornTimeParams {
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PopcornTimeParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	    }
	}
	export class JackettParams {
	    enabled: boolean;
	    apiAddress: string;
	    apiKey: string;
	
	    static createFrom(source: any = {}) {
	        return new JackettParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.apiAddress = source["apiAddress"];
	        this.apiKey = source["apiKey"];
	    }
	}
	export class SourceParams {
	    jackett: JackettParams;
	    pt: PopcornTimeParams;
	    yts: YtsParams;
	    rarbg: RarbgParams;
	    itorrent: ItorrentParams;
	    x1337x: X1337xParams;
	    eztv: EztvParams;
	
	    static createFrom(source: any = {}) {
	        return new SourceParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.jackett = this.convertValues(source["jackett"], JackettParams);
	        this.pt = this.convertValues(source["pt"], PopcornTimeParams);
	        this.yts = this.convertValues(source["yts"], YtsParams);
	        this.rarbg = this.convertValues(source["rarbg"], RarbgParams);
	        this.itorrent = this.convertValues(source["itorrent"], ItorrentParams);
	        this.x1337x = this.convertValues(source["x1337x"], X1337xParams);
	        this.eztv = this.convertValues(source["eztv"], EztvParams);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class ShowTorrent {
	    hash: string;
	    quality: string;
	    season: string;
	    episode: string;
	    size: string;
	    provider: string;
	    lang: string;
	    title: string;
	    seeds: string;
	    peers: string;
	    magnet: string;
	    torrent: string;
	
	    static createFrom(source: any = {}) {
	        return new ShowTorrent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hash = source["hash"];
	        this.quality = source["quality"];
	        this.season = source["season"];
	        this.episode = source["episode"];
	        this.size = source["size"];
	        this.provider = source["provider"];
	        this.lang = source["lang"];
	        this.title = source["title"];
	        this.seeds = source["seeds"];
	        this.peers = source["peers"];
	        this.magnet = source["magnet"];
	        this.torrent = source["torrent"];
	    }
	}
	
	
	export class Episode {
	    id: number;
	    name: string;
	    season_number: number;
	    episode_number: number;
	    air_date: string;
	    overview: string;
	    production_code: string;
	    runtime: number;
	    still_path: string;
	    vote_average: number;
	    vote_count: number;
	
	    static createFrom(source: any = {}) {
	        return new Episode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.season_number = source["season_number"];
	        this.episode_number = source["episode_number"];
	        this.air_date = source["air_date"];
	        this.overview = source["overview"];
	        this.production_code = source["production_code"];
	        this.runtime = source["runtime"];
	        this.still_path = source["still_path"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	    }
	}
	export class ShowParams {
	    imdbId: string;
	    searchText: string;
	    season: string;
	    episode: string;
	
	    static createFrom(source: any = {}) {
	        return new ShowParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.imdbId = source["imdbId"];
	        this.searchText = source["searchText"];
	        this.season = source["season"];
	        this.episode = source["episode"];
	    }
	}
	export class ExternalIds {
	    imdb_id: string;
	    freebase_mid: string;
	    freebase_id: string;
	    tvdb_id: number;
	    tvrage_id: number;
	    facebook_id: string;
	    instagram_id: string;
	    twitter_id: string;
	
	    static createFrom(source: any = {}) {
	        return new ExternalIds(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.imdb_id = source["imdb_id"];
	        this.freebase_mid = source["freebase_mid"];
	        this.freebase_id = source["freebase_id"];
	        this.tvdb_id = source["tvdb_id"];
	        this.tvrage_id = source["tvrage_id"];
	        this.facebook_id = source["facebook_id"];
	        this.instagram_id = source["instagram_id"];
	        this.twitter_id = source["twitter_id"];
	    }
	}
	export class SeasonInfo {
	    id: number;
	    season_number: number;
	    episodes: Episode[];
	    name: string;
	    air_date: string;
	    overview: string;
	    poster_path: string;
	
	    static createFrom(source: any = {}) {
	        return new SeasonInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.season_number = source["season_number"];
	        this.episodes = this.convertValues(source["episodes"], Episode);
	        this.name = source["name"];
	        this.air_date = source["air_date"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TorrentFile {
	    name: string;
	    url: string;
	    length: string;
	
	    static createFrom(source: any = {}) {
	        return new TorrentFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.url = source["url"];
	        this.length = source["length"];
	    }
	}
	export class TorrentInfo {
	    name: string;
	    hash: string;
	    length: string;
	    files: TorrentFile[];
	
	    static createFrom(source: any = {}) {
	        return new TorrentInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.hash = source["hash"];
	        this.length = source["length"];
	        this.files = this.convertValues(source["files"], TorrentFile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Creator {
	    id: number;
	    credit_id: string;
	    name: string;
	    gender: number;
	    profile_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Creator(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.credit_id = source["credit_id"];
	        this.name = source["name"];
	        this.gender = source["gender"];
	        this.profile_path = source["profile_path"];
	    }
	}
	export class Season {
	    id: number;
	    season_number: number;
	    episode_count: number;
	    name: string;
	    air_date: string;
	    overview: string;
	    poster_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Season(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.season_number = source["season_number"];
	        this.episode_count = source["episode_count"];
	        this.name = source["name"];
	        this.air_date = source["air_date"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	    }
	}
	export class ShowInfo {
	    id: number;
	    external_ids: ExternalIds;
	    name: string;
	    original_name: string;
	    origin_country: string[];
	    languages: string[];
	    original_language: string;
	    first_air_date: string;
	    last_air_date: string;
	    last_episode_to_air: Episode;
	    next_episode_to_air: Episode;
	    seasons: Season[];
	    overview: string;
	    tagline: string;
	    poster_path: string;
	    backdrop_path: string;
	    popularity: number;
	    vote_average: number;
	    vote_count: number;
	    number_of_seasons: number;
	    number_of_episodes: number;
	    genres: Genre[];
	    created_by: Creator[];
	    networks: Company[];
	    production_companies: Company[];
	    production_countries: Country[];
	    spoken_languages: Language[];
	    episode_run_time: number[];
	    homepage: string;
	    in_production: boolean;
	    status: string;
	    type: string;
	    adult: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ShowInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.external_ids = this.convertValues(source["external_ids"], ExternalIds);
	        this.name = source["name"];
	        this.original_name = source["original_name"];
	        this.origin_country = source["origin_country"];
	        this.languages = source["languages"];
	        this.original_language = source["original_language"];
	        this.first_air_date = source["first_air_date"];
	        this.last_air_date = source["last_air_date"];
	        this.last_episode_to_air = this.convertValues(source["last_episode_to_air"], Episode);
	        this.next_episode_to_air = this.convertValues(source["next_episode_to_air"], Episode);
	        this.seasons = this.convertValues(source["seasons"], Season);
	        this.overview = source["overview"];
	        this.tagline = source["tagline"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.popularity = source["popularity"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.number_of_seasons = source["number_of_seasons"];
	        this.number_of_episodes = source["number_of_episodes"];
	        this.genres = this.convertValues(source["genres"], Genre);
	        this.created_by = this.convertValues(source["created_by"], Creator);
	        this.networks = this.convertValues(source["networks"], Company);
	        this.production_companies = this.convertValues(source["production_companies"], Company);
	        this.production_countries = this.convertValues(source["production_countries"], Country);
	        this.spoken_languages = this.convertValues(source["spoken_languages"], Language);
	        this.episode_run_time = source["episode_run_time"];
	        this.homepage = source["homepage"];
	        this.in_production = source["in_production"];
	        this.status = source["status"];
	        this.type = source["type"];
	        this.adult = source["adult"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MediaParams {
	    videoUrl: string;
	    subtitleUrl: string;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new MediaParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.videoUrl = source["videoUrl"];
	        this.subtitleUrl = source["subtitleUrl"];
	        this.title = source["title"];
	    }
	}
	export class Movie {
	    id: number;
	    title: string;
	    original_title: string;
	    original_language: string;
	    release_date: string;
	    overview: string;
	    poster_path: string;
	    backdrop_path: string;
	    popularity: number;
	    vote_average: number;
	    vote_count: number;
	    genre_ids: number[];
	    video: boolean;
	    adult: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Movie(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.original_title = source["original_title"];
	        this.original_language = source["original_language"];
	        this.release_date = source["release_date"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.backdrop_path = source["backdrop_path"];
	        this.popularity = source["popularity"];
	        this.vote_average = source["vote_average"];
	        this.vote_count = source["vote_count"];
	        this.genre_ids = source["genre_ids"];
	        this.video = source["video"];
	        this.adult = source["adult"];
	    }
	}
	export class MovieResults {
	    page: number;
	    total_pages: number;
	    total_results: number;
	    results: Movie[];
	
	    static createFrom(source: any = {}) {
	        return new MovieResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.total_pages = source["total_pages"];
	        this.total_results = source["total_results"];
	        this.results = this.convertValues(source["results"], Movie);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MovieDiscoverParams {
	    genreIds: number[];
	    maxReleaseDate: string;
	    minReleaseDate: string;
	    sortBy: string;
	
	    static createFrom(source: any = {}) {
	        return new MovieDiscoverParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.genreIds = source["genreIds"];
	        this.maxReleaseDate = source["maxReleaseDate"];
	        this.minReleaseDate = source["minReleaseDate"];
	        this.sortBy = source["sortBy"];
	    }
	}
	export class SubtitleFile {
	    lang: string;
	    subtitlename: string;
	    releasename: string;
	    subformat: string;
	    subencoding: string;
	    subdata: string;
	    vttdata: string;
	
	    static createFrom(source: any = {}) {
	        return new SubtitleFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lang = source["lang"];
	        this.subtitlename = source["subtitlename"];
	        this.releasename = source["releasename"];
	        this.subformat = source["subformat"];
	        this.subencoding = source["subencoding"];
	        this.subdata = source["subdata"];
	        this.vttdata = source["vttdata"];
	    }
	}
	export class MediaPlayerParams {
	    path: string;
	    args: string;
	
	    static createFrom(source: any = {}) {
	        return new MediaPlayerParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.args = source["args"];
	    }
	}

}

