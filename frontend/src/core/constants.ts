import { HdTwoTone, MovieTwoTone, TvTwoTone, VideoLibraryTwoTone } from "@mui/icons-material";

export const QUALITY_TAGS = [
  "2160p",
  "1080p",
  "720p",
  "480p",
  "360p",
  "webrip",
  "webdl",
  "web-dl",
  "bluray",
  "bdrip",
  "brrip",
  "dvdrip",
  "dvd",
  "hdtv",
  "tvrip",
];

export const TORRENT_CATEGORIES = {
  all: {
    key: "all",
    icon: VideoLibraryTwoTone,
  },
  movies: {
    key: "movies",
    icon: MovieTwoTone,
  },
  shows: {
    key: "shows",
    icon: TvTwoTone,
  },
};

export const TORRENT_QUALITIES = {
  hd: {
    key: "hd",
    icon: HdTwoTone,
    color: "active",
  },
  sd: {
    key: "sd",
    icon: HdTwoTone,
    color: "disabled",
  },
};

export const EXTERNAL_PLAYER_PRESETS = [
  {
    key: "vlc",
    name: "VLC",
    executablePath: "C:\\Program Files\\VideoLAN\\VLC\\vlc.exe",
    args: "<VIDEO_URL>",
  },
  {
    key: "mpchc",
    name: "MPC-HC",
    executablePath: "C:\\Program Files\\MPC-HC\\mpc-hc64.exe",
    args: "<VIDEO_URL> /sub <SUBTITLE_URL>",
  },
  {
    key: "custom",
    name: "Egyéni",
  },
  {
    key: "disabled",
    name: "Letiltva",
    executablePath: "",
    args: "",
  },
];
