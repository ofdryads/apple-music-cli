package cmd

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

// add the currently playing song to a given playlist
func addToPlaylist(playlist string) error {
	script := fmt.Sprintf(`
tell application "Music"
    set thePlaylist to (first playlist whose name is "%s")
    set theTrack to current track
    duplicate theTrack to thePlaylist
end tell
`, playlist)

	_, err := mack.Tell("Music", script)
	if err != nil {
		return fmt.Errorf("error adding song to playlist: %w", err)
	}
	return nil
}

// user can have a "default playlist" that songs will be added to if no option selected
// take the playlist name as a part of a command
// like ./music add nameofplaylist
var addToPlaylistCmd = &cobra.Command{
    Use:   "add [playlist]",
    Short: "Add the song now playing to a playlist",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        playlist := args[0]

        rawName, err := mack.Tell("Music", `tell application "Music" to get name of current track`)
        if err != nil {
            fmt.Println("Error getting this song's name:", err)
            return
        }
        songName := string(rawName)

        err = addToPlaylist(playlist)
        if err != nil {
            fmt.Println("Error adding to playlist:", err)
        } else {
            fmt.Printf("'%s' has been added to %s\n", songName, playlist)
        }
    },
}