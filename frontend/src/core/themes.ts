import { ThemeOptions } from "@mui/material";

interface SiteTheme extends ThemeOptions {
  name: string;
  logo: string;
}

export const THEMES = new Map<string, SiteTheme>([
  [
    "black",
    {
      name: "Fekete",
      logo: "logo-red.png",
      palette: {
        mode: "dark",
        primary: { main: "#DC1A28" },
        secondary: { main: "#DC1A28" },
        background: {
          default: "#000000",
          paper: "#212121",
        },
      },
    },
  ],
  [
    "dark",
    {
      name: "Sötét",
      logo: "logo-white.png",
      palette: {
        mode: "dark",
        primary: { main: "#DC1A28" },
        secondary: { main: "#DC1A28" },
        background: {
          default: "#212121",
        },
      },
    },
  ],
  [
    "light",
    {
      name: "Világos",
      logo: "logo-grey.png",
      palette: {
        mode: "light",
        primary: { main: "#DC1A28" },
        secondary: { main: "#DC1A28" },
        background: {
          default: "#D4D4D4",
          paper: "#F4F4F4",
        },
      },
    },
  ],
]);

export const DEFAULT_THEME = Array.from(THEMES.keys())[0];
