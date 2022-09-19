import { createState } from "@hookstate/core";

export const uiStore = createState({
  menuOpen: false,
  playDialogOpen: false,
  videoPlayerOpen: false,
  currentTorrent: null,
  currentFile: null,
  currentSubtitles: null,
  currentSubtitle: null,
});
