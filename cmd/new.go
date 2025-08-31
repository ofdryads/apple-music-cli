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

var newCmd = &cobra.Command{
	Use:   "new [playlist]",
	Short: "Make a new playlist in Apple Music",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		playlist := args[0]

		err := makePlaylist(playlist)
		if err != nil {
			fmt.Printf("Playlist could not be created: %v\n", err)
			return
		}
		fmt.Printf("Playlist '%s' has now been created \n", playlist)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
