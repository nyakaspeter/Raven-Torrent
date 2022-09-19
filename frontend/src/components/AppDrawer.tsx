import { useHookstate } from "@hookstate/core";
import { CloudDownloadOutlined, HomeOutlined, MovieOutlined, SettingsOutlined, TvOutlined } from "@mui/icons-material";
import { Box, Drawer, List, ListItem, ListItemIcon, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";
import { uiStore } from "../stores/ui";

type Props = {};

const AppDrawer = (props: Props) => {
  const { menuOpen } = useHookstate(uiStore);

  return (
    <Drawer className="no-drag" anchor="left" open={menuOpen.value} onClose={() => menuOpen.set(false)}>
      <List
        sx={{
          height: "100%",
          marginY: 3,
          display: "flex",
          flexDirection: "column",
        }}
      >
        <ListItem
          button
          component={Link}
          to="/"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <HomeOutlined />
          </ListItemIcon>
          <ListItemText primary="Kezdőlap" />
        </ListItem>

        <ListItem
          button
          component={Link}
          to="/movies"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <MovieOutlined />
          </ListItemIcon>
          <ListItemText primary="Filmek" />
        </ListItem>

        <ListItem
          button
          component={Link}
          to="/shows"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <TvOutlined />
          </ListItemIcon>
          <ListItemText primary="Sorozatok" />
        </ListItem>

        <ListItem
          button
          component={Link}
          to="/torrents"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <CloudDownloadOutlined />
          </ListItemIcon>
          <ListItemText primary="Torrentek" />
        </ListItem>

        <Box flex="1" />

        <ListItem
          button
          component={Link}
          to="/debug"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <SettingsOutlined />
          </ListItemIcon>
          <ListItemText primary="Debug" />
        </ListItem>

        <ListItem
          button
          component={Link}
          to="/settings"
          onClick={() => menuOpen.set(false)}
          style={{ paddingLeft: 48, paddingRight: 72 }}
        >
          <ListItemIcon>
            <SettingsOutlined />
          </ListItemIcon>
          <ListItemText primary="Beállítások" />
        </ListItem>
      </List>
    </Drawer>
  );
};

export default AppDrawer;
