import { useHookstate } from "@hookstate/core";
import { Search } from "@mui/icons-material";
import { Box, CircularProgress, Divider, Fade, IconButton, InputBase, Paper, useTheme } from "@mui/material";
import React, { useEffect, useMemo, useRef } from "react";
import { useQueries, useQuery } from "react-query";
import { GetMovieTorrents, GetShowTorrents } from "../../../wailsjs/go/main/App";
import { types } from "../../../wailsjs/go/models";
import TorrentAccordion from "../../components/TorrentAccordion";
import TorrentList from "../../components/TorrentList";
import { TORRENT_CATEGORIES, TORRENT_QUALITIES } from "../../core/constants";
import { includesAny } from "../../core/utils";
import { settingsStore } from "../../stores/settings";

type Props = {};

const Torrents = (props: Props) => {
  const theme = useTheme();
  const { hdQualityTags, jackettApiAddress, jackettApiKey } = useHookstate(settingsStore);
  const { torrentSearchQuery, alreadySearched, selectedCategory, selectedQuality } = useHookstate({
    torrentSearchQuery: "",
    alreadySearched: false,
    selectedCategory: TORRENT_CATEGORIES.all.key,
    selectedQuality: TORRENT_QUALITIES.hd.key,
  });

  const torrents = useQuery("torrents", () =>
    Promise.all([
      GetMovieTorrents(
        { searchText: torrentSearchQuery.value },
        { jackett: { enabled: true, apiAddress: jackettApiAddress.value, apiKey: jackettApiKey.value } }
      ),
      GetShowTorrents(
        { searchText: torrentSearchQuery.value },
        { jackett: { enabled: true, apiAddress: jackettApiAddress.value, apiKey: jackettApiKey.value } }
      ),
    ])
  );

  const filteredTorrents = useMemo(() => {
    if (!torrents.data) return [];

    const [movieTorrents, showTorrents] = torrents.data;
    let filtered: (types.MovieTorrent | types.ShowTorrent)[] = [];

    if (selectedCategory.value === "all") {
      filtered = movieTorrents.concat(showTorrents);
    } else if (selectedCategory.value === "movies") {
      filtered = movieTorrents;
    } else if (selectedCategory.value === "shows") {
      filtered = showTorrents;
    }

    if (selectedQuality.value === "hd" && hdQualityTags.value.length > 0) {
      filtered = filtered.filter((torrent) => includesAny(torrent.title, hdQualityTags.value));
    } else if (selectedQuality.value === "sd") {
      filtered = filtered.filter((torrent) => !includesAny(torrent.title, hdQualityTags.value));
    }

    filtered.sort((a, b) => Number(b.seeds) - Number(a.seeds));

    return filtered;
  }, [selectedCategory.value, selectedQuality.value, torrents.data]);

  function switchCategory() {
    const keys = Object.keys(TORRENT_CATEGORIES);
    const currentIndex = keys.indexOf(selectedCategory.value);
    const newIndex = (currentIndex + 1) % keys.length;

    selectedCategory.set(keys[newIndex]);
  }

  function switchQuality() {
    const keys = Object.keys(TORRENT_QUALITIES);
    const currentIndex = keys.indexOf(selectedQuality.value);
    const newIndex = (currentIndex + 1) % keys.length;

    selectedQuality.set(keys[newIndex]);
  }

  return (
    <Box display="flex" flexDirection="column" height="100%" width="100%" alignItems="center">
      <Paper
        component="form"
        onSubmit={async (event) => {
          event.preventDefault();
          await torrents.refetch();
          if (!alreadySearched.value) alreadySearched.set(true);
        }}
        sx={{ marginBottom: 2 }}
        style={{
          padding: "2px 8px",
          display: "flex",
          alignItems: "center",
          maxWidth: "500px",
          width: "100%",
          borderRadius: "50px",
        }}
      >
        <InputBase
          style={{ marginLeft: "8px", flex: 1 }}
          placeholder="Keresett kifejezés"
          value={torrentSearchQuery.value}
          onChange={(e) => torrentSearchQuery.set(e.target.value)}
        />
        <IconButton
          style={{
            padding: "10px",
            color: theme.palette.action[TORRENT_QUALITIES[selectedQuality.value].color],
          }}
          onClick={switchQuality}
        >
          {/* {TORRENT_QUALITIES[selectedQuality.value].icon} */}
        </IconButton>
        <IconButton
          style={{
            padding: "10px",
          }}
          onClick={switchCategory}
        >
          {/* {TORRENT_CATEGORIES[selectedCategory.value].icon} */}
        </IconButton>
        <Divider style={{ height: "28px", margin: "4px" }} orientation="vertical" />
        <IconButton
          type="submit"
          style={{
            padding: "10px",
          }}
        >
          <Search />
        </IconButton>
      </Paper>

      {torrents.isFetching && (
        <CircularProgress
          color="primary"
          style={{
            position: "relative",
            top: "24px",
          }}
        />
      )}

      {!torrents.isFetching && alreadySearched.value && filteredTorrents.length === 0 && (
        <Box marginTop="36px">Nincs találat</Box>
      )}

      <Fade in={!torrents.isFetching}>
        <Box>
          <TorrentList torrents={filteredTorrents} />
        </Box>
      </Fade>

      {/* {!loadingTorrents.value &&
        filteredTorrents.value.map((torrent) => (
          <Box paddingBottom="16px" width="100%" key={torrent.magnet || torrent.torrent}>
            <TorrentAccordion torrent={torrent} />
          </Box>
        ))} */}
    </Box>
  );
};

export default Torrents;
