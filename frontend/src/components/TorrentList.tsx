import { Box } from "@mui/material";
import React from "react";
import { types } from "../../wailsjs/go/models";
import TorrentAccordion from "./TorrentAccordion";

type Props = {
  torrents: (types.MovieTorrent | types.ShowTorrent)[];
};

const TorrentList: React.FC<Props> = ({ torrents }) => (
  <>
    {torrents.map((torrent) => (
      <Box paddingBottom="16px" width="100%" key={torrent.magnet || torrent.torrent}>
        <TorrentAccordion torrent={torrent} />
      </Box>
    ))}
  </>
);

export default React.memo(TorrentList);
