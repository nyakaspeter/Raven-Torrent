import { useHookstate } from "@hookstate/core";
import { Button } from "@mui/material";
import { useEffect } from "react";
import { DiscoverMovies, GetMovieTorrents, StartMediaPlayer } from "../../../wailsjs/go/main/App";
import { settingsStore } from "../../stores/settings";

type Props = {};

const Debug = (props: Props) => {
  const { externalPlayer, jackettApiAddress, jackettApiKey } = useHookstate(settingsStore);

  useEffect(() => {
    DiscoverMovies(
      {
        genreIds: [],
        maxReleaseDate: "",
        minReleaseDate: "",
        sortBy: "",
      },
      "hu",
      1
    ).then((res) => console.log(res));

    GetMovieTorrents(
      { searchText: "" },
      { jackett: { enabled: true, apiAddress: jackettApiAddress.value, apiKey: jackettApiKey.value } }
    ).then((res) => console.log(res));
  }, []);

  return (
    <div>
      <Button
        onClick={() =>
          StartMediaPlayer({ path: externalPlayer.executablePath.value!!, args: externalPlayer.args.value!! }).then(
            (res) => console.log(res)
          )
        }
      >
        Play
      </Button>
    </div>
  );
};

export default Debug;
