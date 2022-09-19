import {
  ChevronRight,
  DescriptionOutlined,
  ExpandMore,
  Folder,
  ImageOutlined,
  InsertDriveFileOutlined,
  MovieOutlined,
  MusicNoteOutlined,
  PlayArrow,
} from "@mui/icons-material";
import { TreeItem, TreeView } from "@mui/lab";
import { Box, Chip, IconButton, Typography } from "@mui/material";
import mime from "mime";
import { formatBytesToString } from "../core/utils";

type Props = {};

const TorrentFileList = (props) => {
  const { files, onShowPlayDialog, onFileSelect } = props;

  let nodeId = 0;

  function fileNameToIcon(fileName) {
    let mimeType = mime.getType(fileName);

    if (mimeType) {
      if (mimeType.startsWith("image")) return <ImageOutlined />;
      if (mimeType.startsWith("video")) return <MovieOutlined />;
      if (mimeType.startsWith("audio")) return <MusicNoteOutlined />;
      if (mimeType.startsWith("text")) return <DescriptionOutlined />;
    }

    return <InsertDriveFileOutlined />;
  }

  const TreeItems = (file) => {
    if (file.children.length !== 0) {
      return (
        <TreeItem
          key={file.name}
          nodeId={`${nodeId++}`}
          label={
            <Box display="flex" alignItems="center">
              <IconButton>
                <Folder />
              </IconButton>
              <Typography variant="body2">{file.name}</Typography>
            </Box>
          }
        >
          {file.children.map((child) => TreeItems(child))}
        </TreeItem>
      );
    } else
      return (
        <TreeItem
          key={file.name}
          nodeId={`${nodeId++}`}
          label={
            <Box display="flex" alignItems="center">
              <IconButton>{fileNameToIcon(file.name)}</IconButton>
              <Typography
                variant="body2"
                style={{
                  wordBreak: "break-word",
                  marginRight: 8,
                  marginTop: 2,
                  marginBottom: 2,
                }}
              >
                {file.name}
              </Typography>
              <Chip size="small" label={formatBytesToString(file.size)} />

              <Box flex="auto" />

              <Box display="flex" alignItems="center" marginRight="16px">
                {mime.getType(file.name) && mime.getType(file.name).startsWith("video") && (
                  <IconButton
                    style={{ marginRight: 0 /*12*/ }}
                    size="small"
                    onClick={(e) => {
                      e.stopPropagation();
                      onFileSelect(file);
                      onShowPlayDialog();
                    }}
                  >
                    <PlayArrow fontSize="small" />
                  </IconButton>
                )}

                {/* <a download={file.name} href={file.url}>
                    <IconButton
                      size="small"
                      onClick={(e) => {
                        e.stopPropagation();
                      }}
                    >
                      <GetAppIcon fontSize="small" />
                    </IconButton>
                  </a> */}
              </Box>
            </Box>
          }
        />
      );
  };

  return (
    <TreeView
      disableSelection
      style={{ flex: "auto" }}
      defaultCollapseIcon={<ExpandMore />}
      defaultExpandIcon={<ChevronRight />}
    >
      {files.length === 1 && files[0].children.length > 0
        ? files[0].children.map((file) => TreeItems(file))
        : files.map((file) => TreeItems(file))}
    </TreeView>
  );
};

export default TorrentFileList;
