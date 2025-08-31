package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

func showPlaylist() ([]string, error) {
	script := `
    set playlistNames to name of playlists
    return playlistNames
  `
	result, err := mack.Tell("Music", script)
	if err != nil {
		return nil, fmt.Errorf("error fetching playlists: %w", err)
	}
	playlistNames := strings.Split(result, ", ")
	return playlistNames, nil
}

// TODO more explicitly handle error where the playlist is empty
func playPlaylist(playlist string) error {
	fmt.Printf("Attempting to play playlist: %s \n", playlist)
	script := fmt.Sprintf(`play playlist named "%s"`, playlist)
	_, err := mack.Tell("Music", script)
	if err != nil {
		return fmt.Errorf("error playing the playlist: %w", err)
	}
	return nil
}

var playlistsCmd = &cobra.Command{
	Use:   "playlists",
	Short: "Display all playlists and select one to play",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		isOpen, err := isMusicOpen()
		if err != nil {
			fmt.Println("Error checking if Apple Music is open:", err)
			return
		}

		if !isOpen {
			fmt.Println("Apple Music is not open!")
			return
		}

		if len(args) == 0 {
			playlistNames, err := showPlaylist()
			if err != nil {
				fmt.Println("Error gettings playlust names:", err)
				return
			}

			if len(playlistNames) == 0 {
				fmt.Println("No playlists found.")
				return
			}

			fzfCmd := exec.Command("fzf")
			fzfCmd.Stdin = strings.NewReader(strings.Join(playlistNames, "\n"))
			fzfCmd.Stderr = os.Stderr

			playlistBytes, err := fzfCmd.Output()
			if err != nil {
				fmt.Println("Error running fzf:", err)
				return
			}

			selectedPlaylist := strings.TrimSpace(string(playlistBytes))
			errr := playPlaylist(selectedPlaylist)
			if errr != nil {
				fmt.Println("Error playing the chosen playlist:", err)
				return
			}

			fmt.Printf("%s has been played\n", selectedPlaylist)
			info, err := getCurrentSongInfo()
			if err != nil {
				fmt.Printf("Error getting current song info: %v", err)
				return
			}

			if info.trackName == "" {
				fmt.Println("Song Skipped")
				return
			}
			fmt.Printf("Now Playing: %s\nBy: %s\n", info.trackName, info.artistName)

		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(playlistsCmd)
}
