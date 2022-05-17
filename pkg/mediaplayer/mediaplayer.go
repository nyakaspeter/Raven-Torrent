package mediaplayer

import (
	"os/exec"
	"strings"

	"github.com/nyakaspeter/raven-torrent/pkg/mediaplayer/types"
)

func StartMediaPlayer(params types.MediaPlayerParams) error {
	splitArgs := strings.Split(params.ExecutableArgs, " ")

	cmd := exec.Command(params.ExecutablePath, splitArgs...)
	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
