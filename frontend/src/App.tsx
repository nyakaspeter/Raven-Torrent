import { useHookstate } from "@hookstate/core";
import { Container, createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { useMemo } from "react";
import { QueryClientProvider } from "react-query";
import { HashRouter, Navigate, Route, Routes } from "react-router-dom";
import "./App.css";
import AppDrawer from "./components/AppDrawer";
import AppHeader from "./components/AppHeader";
import { queryClient } from "./core/query";
import { THEMES } from "./core/themes";
import Debug from "./pages/debug/Debug";
import Home from "./pages/home/Home";
import Movies from "./pages/movies/Movies";
import Settings from "./pages/settings/Settings";
import Shows from "./pages/shows/Shows";
import Torrents from "./pages/torrents/Torrents";
import { settingsStore } from "./stores/settings";

const App = () => {
  const { selectedTheme } = useHookstate(settingsStore);

  const theme = useMemo(() => {
    const selectedThemeOptions = THEMES.get(selectedTheme.value);
    return createTheme(selectedThemeOptions);
  }, [selectedTheme.value]);

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider theme={theme}>
        <HashRouter>
          <AppDrawer />
          <AppHeader />
          <Container>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/movies" element={<Movies />} />
              <Route path="/shows" element={<Shows />} />
              <Route path="/torrents" element={<Torrents />} />
              <Route path="/settings" element={<Settings />} />
              <Route path="/debug" element={<Debug />} />
              <Route path="*" element={<Navigate replace to="/" />} />
            </Routes>
          </Container>
        </HashRouter>
        <CssBaseline />
      </ThemeProvider>
    </QueryClientProvider>
  );
};

export default App;
