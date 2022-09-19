import { useHookstate } from "@hookstate/core";
import { Box, FormControl, FormHelperText, InputLabel, MenuItem, Select, TextField } from "@mui/material";
import ChipInput from "../../components/ChipInput";
import { EXTERNAL_PLAYER_PRESETS } from "../../core/constants";
import { THEMES } from "../../core/themes";
import { settingsStore } from "../../stores/settings";

type Props = {};

const Settings = (props: Props) => {
  const { jackettApiAddress, jackettApiKey, externalPlayer, subtitleLanguages, hdQualityTags, selectedTheme } =
    useHookstate(settingsStore);

  return (
    <Box width="100%" height="100%" display="flex" flexDirection="column" alignItems="center" overflow="auto">
      <TextField
        fullWidth
        margin="normal"
        label="Jackett API elérési útvonala"
        helperText="Protokollal és portszámmal együtt, a Raven szerverhez viszonyítva, pl. http://127.0.0.1:9117"
        variant="filled"
        value={jackettApiAddress.value}
        onChange={(event) => jackettApiAddress.set(event.target.value)}
      />

      <TextField
        fullWidth
        margin="normal"
        label="Jackett API kulcs"
        helperText="A Jackett webes felületéről leolvasható API kulcs"
        variant="filled"
        value={jackettApiKey.value}
        onChange={(event) => jackettApiKey.set(event.target.value)}
      />

      <>
        <Box height="16px" />
        <FormControl fullWidth margin="normal" variant="filled" style={{ marginTop: 0 }}>
          <InputLabel>Külső médialejátszó program</InputLabel>
          <Select
            value={externalPlayer.key.value}
            onChange={(event) => {
              const preset = EXTERNAL_PLAYER_PRESETS.find((p) => p.key === event.target.value);
              if (preset) {
                externalPlayer.merge(preset);
              }
            }}
          >
            {EXTERNAL_PLAYER_PRESETS.map((preset) => (
              <MenuItem key={preset.key} value={preset.key}>
                {preset.name}
              </MenuItem>
            ))}
          </Select>
          <FormHelperText>Engedélyezd, ha külső médialejátszót szeretnél használni</FormHelperText>
        </FormControl>

        {externalPlayer.key.value !== "disabled" && (
          <>
            <TextField
              fullWidth
              margin="normal"
              label="Külső lejátszó elérési útvonala"
              helperText="Lejátszó elérési útja, pl. C:\Program Files\MPC-HC\mpc-hc64.exe"
              variant="filled"
              value={externalPlayer.executablePath.value}
              onChange={(event) => {
                externalPlayer.set((p) => ({
                  key: "custom",
                  name: "Egyéni",
                  executablePath: event.target.value,
                  args: p.args,
                }));
              }}
            />

            <TextField
              fullWidth
              margin="normal"
              label="Külső lejátszó indítási argumentumai"
              helperText="Használható a <VIDEO_URL> és a <SUBTITLE_URL> változó"
              variant="filled"
              value={externalPlayer.args.value}
              onChange={(event) => {
                externalPlayer.set((p) => ({
                  key: "custom",
                  name: "Egyéni",
                  executablePath: p.executablePath,
                  args: event.target.value,
                }));
              }}
            />
          </>
        )}
      </>

      <FormControl fullWidth margin="normal">
        <ChipInput
          value={subtitleLanguages.value}
          onChange={(event, value) => subtitleLanguages.set(value)}
          textFieldProps={{ label: "Felirat keresés nyelvei", variant: "filled" }}
        />
        <FormHelperText>ISO 639-2 formátumú nyelvkódok listája, pl. hun, eng, ger</FormHelperText>
      </FormControl>

      <FormControl fullWidth margin="normal">
        <ChipInput
          value={hdQualityTags.value}
          onChange={(event, value) => hdQualityTags.set(value)}
          textFieldProps={{ label: "HD minőség kulcsszavak", variant: "filled" }}
        />
        <FormHelperText>Az ezeket tartalmazó torrentek HD-nek lesznek tekintve</FormHelperText>
      </FormControl>

      <FormControl fullWidth margin="normal" variant="filled">
        <InputLabel>Nyelv</InputLabel>
        <Select value={"hun"} onChange={(event) => {}}>
          <MenuItem value={"hun"}>Magyar</MenuItem>
        </Select>
        <FormHelperText>A felhasználói felület és a filmadatbázis nyelve</FormHelperText>
      </FormControl>

      <FormControl fullWidth margin="normal" variant="filled">
        <InputLabel>Téma</InputLabel>
        <Select
          value={selectedTheme.value}
          onChange={(event) => {
            selectedTheme.set(event.target.value);
          }}
        >
          {Array.from(THEMES.keys()).map((key) => (
            <MenuItem key={key} value={key}>
              {THEMES.get(key)!!.name}
            </MenuItem>
          ))}
        </Select>
        <FormHelperText>Az alkalmazás megjelenítési stílusa</FormHelperText>
      </FormControl>
    </Box>
  );
};

export default Settings;
