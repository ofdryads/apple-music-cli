package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

// Note: back behavior is inconsistent - maybe try more manual way of tracking recently played
var backCmd = &cobra.Command{
	Use:   "back",
	Short: "Go back to the previous track in Apple Music",
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

		if _, err := mack.Tell("Music", "previous track"); err != nil {
			fmt.Println("Error skipping the song:", err)
		}

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
	},
}

func init() {
	rootCmd.AddCommand(backCmd)
}
