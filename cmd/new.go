// create a new playlist

package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

func makePlaylist(playlist string) error {
	script := fmt.Sprintf(`
tell application "Music"
	make new user playlist with properties {name:"%s"}
end tell
    `, playlist)
	_, err := mack.Tell("Music", script)
	if err != nil {
		return fmt.Errorf("error making new playlist: %w", err)
	}
	return nil
}

var newCmd = &cobra.Command{}
