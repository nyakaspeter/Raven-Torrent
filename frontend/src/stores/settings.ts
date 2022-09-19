import { createState } from "@hookstate/core";
import { Persistence } from "@hookstate/persistence";
import { EXTERNAL_PLAYER_PRESETS } from "../core/constants";
import { DEFAULT_THEME } from "../core/themes";

export const settingsStore = createState({
  jackettApiAddress: "http://localhost:9117",
  jackettApiKey: "",
  subtitleLanguages: ["hun", "eng"],
  hdQualityTags: ["2160p", "1080p", "720p"],
  selectedTheme: DEFAULT_THEME,
  externalPlayer: EXTERNAL_PLAYER_PRESETS[0],
});

settingsStore.attach(Persistence("settings"));
