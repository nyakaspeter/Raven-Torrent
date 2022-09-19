import { useHookstate } from "@hookstate/core";
import { Menu } from "@mui/icons-material";
import { Box, IconButton, useTheme } from "@mui/material";
import { uiStore } from "../stores/ui";
import Img from "./Img";

type Props = {};

const AppHeader = (props: Props) => {
  const theme = useTheme();
  const { menuOpen } = useHookstate(uiStore);

  const imageUrl = new URL("../assets/images/logo-red.png", import.meta.url).href;

  return (
    <Box sx={{ width: "100%", height: 56, display: "flex", justifyContent: "center", alignItems: "center" }}>
      <Box sx={{ position: "fixed", left: 0, top: 0, margin: 1 }}>
        <IconButton onClick={() => menuOpen.set(true)}>
          <Menu />
        </IconButton>
      </Box>
      <Img
        alt="Black Raven"
        src={imageUrl}
        sx={{
          width: 160,
        }}
      />
    </Box>
  );
};

export default AppHeader;
