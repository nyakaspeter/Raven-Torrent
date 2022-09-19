// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';

export function AddTorrent(arg1:string):Promise<types.TorrentInfo>;

export function CastMediaToDevice(arg1:types.MediaParams,arg2:string):Promise<Error>;

export function DiscoverMovies(arg1:types.MovieDiscoverParams,arg2:string,arg3:number):Promise<types.MovieResults>;

export function DiscoverShows(arg1:types.ShowDiscoverParams,arg2:string,arg3:number):Promise<types.ShowResults>;

export function GetActiveTorrents():Promise<Array<types.TorrentInfo>>;

export function GetMediaDevices():Promise<Array<types.MediaDevice>>;

export function GetMovieInfo(arg1:number,arg2:string):Promise<types.MovieInfo>;

export function GetMovieTorrents(arg1:types.MovieParams,arg2:types.SourceParams):Promise<Array<types.MovieTorrent>>;

export function GetShowInfo(arg1:number,arg2:string):Promise<types.ShowInfo>;

export function GetShowSeason(arg1:number,arg2:number,arg3:string):Promise<types.SeasonInfo>;

export function GetShowTorrents(arg1:types.ShowParams,arg2:types.SourceParams):Promise<Array<types.ShowTorrent>>;

export function GetSubtitles(arg1:types.MediaParams,arg2:Array<string>):Promise<Array<types.SubtitleFile>>;

export function GetSubtitlesForEpisode(arg1:types.MediaParams,arg2:types.EpisodeParams,arg3:Array<string>):Promise<Array<types.SubtitleFile>>;

export function RemoveTorrent(arg1:string):Promise<Error>;

export function SearchMovies(arg1:string,arg2:string,arg3:number):Promise<types.MovieResults>;

export function SearchShows(arg1:string,arg2:string,arg3:number):Promise<types.ShowResults>;

export function StartMediaPlayer(arg1:types.MediaPlayerParams):Promise<Error>;
