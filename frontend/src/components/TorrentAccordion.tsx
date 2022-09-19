import { useHookstate } from "@hookstate/core";
import { ArrowDownward, ArrowUpward, ExpandMore, GetApp, PlayArrow } from "@mui/icons-material";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Chip,
  CircularProgress,
  IconButton,
  Typography,
} from "@mui/material";
import React, { useRef } from "react";
import { useQuery } from "react-query";
import { AddTorrent } from "../../wailsjs/go/main/App";
import { types } from "../../wailsjs/go/models";
import { formatBytesToString } from "../core/utils";
import TorrentFileList from "./TorrentFileList";

type Props = { torrent: types.MovieTorrent | types.ShowTorrent };

const TorrentAccordion = ({ torrent }: Props) => {
  const torrentLink = torrent.magnet || torrent.torrent;

  const { loadingDefaultFile } = useHookstate({
    loadingDefaultFile: false,
  });

  const files = useQuery(
    ["files", torrentLink],
    async () => {
      const torrentInfo = await AddTorrent(torrentLink);

      torrentInfo.files.sort((a, b) => b.name.split("/").length - a.name.split("/").length);

      let paths = torrentInfo.files.map((f) => {
        return { path: f.name, size: f.length, url: f.url };
      });

      let fileList = [];
      let level = { result: fileList };
      paths.forEach((path) => {
        path.path.split("/").reduce((r, name, i, a) => {
          if (!r[name]) {
            r[name] = { result: [] };
            r.result.push({
              name,
              children: r[name].result,
              size: path.size,
              url: path.url,
            });
          }
          return r[name];
        }, level);
      });

      return fileList;
    },
    {
      enabled: false,
    }
  );

  const filesRef = useRef([]);
  const selectedFileRef = useRef();

  async function listTorrentFiles() {
    //const files = await AddTorrent(torrent.magnet || torrent.torrent);
    // loadingFiles.set(true);
    // filesRef.current = await RavenAPI.addTorrent(torrent.magnet || torrent.torrent);
    // filesRef.current.sort((a, b) => b.name.split("/").length - a.name.split("/").length);
    // let paths = filesRef.current.map((f) => {
    //   return { path: f.name, size: f.length, url: f.url };
    // });
    // let fileList = [];
    // let level = { result: fileList };
    // paths.forEach((path) => {
    //   path.path.split("/").reduce((r, name, i, a) => {
    //     if (!r[name]) {
    //       r[name] = { result: [] };
    //       r.result.push({
    //         name,
    //         children: r[name].result,
    //         size: path.size,
    //         url: path.url,
    //       });
    //     }
    //     return r[name];
    //   }, level);
    // });
    // files.set(fileList);
    // loadingFiles.set(false);
  }

  async function playDefaultFile(season, episode) {
    // loadingDefaultFile.set(true);
    // if (filesRef.current.length === 0) await listTorrentFiles();
    // let results = filesRef.current;
    // if (results.length > 0) {
    //   if (typeof torrent.season !== "undefined") {
    //     let s = season || "1";
    //     let e = episode || "1";
    //     let seasonEpisodeString1 = "s" + s.padStart(2, "0") + "e" + e.padStart(2, "0");
    //     let seasonEpisodeString2 = s + "x" + e.padStart(2, "0");
    //     let seasonEpisodeString3 = s + e.padStart(2, "0");
    //     let filteredResults1 = results.filter((result) => result.name.toLowerCase().includes(seasonEpisodeString1));
    //     let filteredResults2 = results.filter((result) => result.name.toLowerCase().includes(seasonEpisodeString2));
    //     let filteredResults3 = results.filter((result) => result.name.toLowerCase().includes(seasonEpisodeString3));
    //     if (filteredResults1.length > 0) {
    //       results = filteredResults1;
    //     } else if (filteredResults2.length > 0) {
    //       results = filteredResults2;
    //     } else if (filteredResults3.length > 0) {
    //       results = filteredResults3;
    //     }
    //   }
    //   results.sort((a, b) => b.length - a.length);
    //   selectedFileRef.current = results[0];
    //   let path = selectedFileRef.current.name.split("/");
    //   selectedFileRef.current.name = path[path.length - 1];
    //   openPlayDialog();
    // }
    // loadingDefaultFile.set(false);
  }

  function openPlayDialog() {
    // uiState.merge({ playDialogOpen: true, currentTorrent: torrent, currentFile: selectedFileRef.current });
  }

  return (
    <Accordion
      TransitionProps={{ unmountOnExit: true }}
      style={{
        borderRadius: 8,
        width: "100%",
      }}
      onChange={async (e, expanded) => {
        if (expanded && !files.isFetched) await files.refetch();
      }}
    >
      <AccordionSummary expandIcon={<ExpandMore />}>
        <Box display="flex" alignItems="center" width="100%">
          <Box display="flex" alignItems="center" flexWrap="wrap">
            <Typography
              variant="subtitle2"
              style={{
                wordBreak: "break-word",
                marginRight: 8,
                marginTop: 2,
                marginBottom: 2,
              }}
            >
              {torrent.title}
            </Typography>

            <Box display="flex" alignItems="center" flexWrap="wrap">
              {/* {torrent.quality && (
                  <Chip size="small" label={torrent.quality} />
                )}

                {torrent.lang && (
                  <Chip
                    size="small"
                    label={capitalizeString(languageNames.of(torrent.lang))}
                  />
                )} */}

              <Chip size="small" label={torrent.provider} style={{ marginRight: 4 }} />

              <Chip size="small" label={formatBytesToString(torrent.size)} style={{ marginLeft: 4, marginRight: 4 }} />

              <Chip
                icon={<ArrowUpward />}
                size="small"
                label={torrent.seeds}
                style={{ marginLeft: 4, marginRight: 4 }}
              />

              <Chip
                icon={<ArrowDownward />}
                size="small"
                label={torrent.peers}
                style={{ marginLeft: 4, marginRight: 4 }}
              />
            </Box>
          </Box>

          <Box flex="auto" />

          <Box display="flex" alignItems="center">
            {loadingDefaultFile.value ? (
              <CircularProgress
                color="primary"
                style={{
                  width: 24,
                  height: 24,
                  marginRight: 12,
                }}
              />
            ) : (
              <IconButton
                style={{ marginRight: 12 }}
                size="small"
                onClick={async (e) => {
                  e.stopPropagation();
                  await playDefaultFile(torrent.season, torrent.episode);
                }}
              >
                <PlayArrow fontSize="small" />
              </IconButton>
            )}

            <a href={torrent.magnet || torrent.torrent}>
              <IconButton
                size="small"
                onClick={(e) => {
                  e.stopPropagation();
                }}
              >
                <GetApp fontSize="small" />
              </IconButton>
            </a>
          </Box>
        </Box>
      </AccordionSummary>

      <AccordionDetails>
        <Box display="flex" justifyContent="center" width="100%" marginRight="20px" marginBottom="12px">
          {files.isFetching && <CircularProgress color="primary" />}

          {files.isFetched && files.data.length > 0 && (
            <TorrentFileList
              files={files.data}
              onShowPlayDialog={() => openPlayDialog()}
              onFileSelect={(file) => (selectedFileRef.current = file)}
            />
          )}

          {files.isFetched && files.data.length === 0 && <Box margin="10px">Nem sikerült a torrent betöltése</Box>}
        </Box>
      </AccordionDetails>
    </Accordion>
  );
};

export default React.memo(TorrentAccordion);
