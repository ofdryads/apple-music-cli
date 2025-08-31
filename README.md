<h1 align="center">apple-music-cli</h1>
<h2 align="center">(appleMusic-cli-player Fork)</h2>

<div align="center">
  <img width="724" alt="image" src="https://github.com/user-attachments/assets/e41dee7d-fc40-4063-b301-c8d79be85234">
</div>

<h3 align="center">Control your music directly from the terminal</h3>

## Overview
This is a fork of the `appleMusic-cli-player` tool by [talz-a](https://github.com/talz-a). It was made to extend the original tool's features to include these new abilities:
- add songs to playlists
- make new playlists
- go back to the previous song
- (Possibly later) gaming controller support

All credit for creating the project, the original commands, and the album art features goes to talz-a.

**[FROM OG REPO]**

`appleMusic-cli-player` is a command-line interface (CLI) tool designed to control music playback on Apple Music. Built using [Go](https://golang.org/) and the [Cobra](https://github.com/spf13/cobra) CLI library, it allows you to manage your music without leaving the terminal. With commands for playing, pausing, skipping tracks, and more, this tool offers seamless control over your Apple Music experience.

## Usage

- **List of Commands:**

  ```bash
  music [command]
  ```

  - `completion` - Generate the autocompletion script for the specified shell
  - `current` - Display currently playing song information
  - `help` - Get help about any command
  - `next` - Skip to the next track in Apple Music
  - `open` - Open Apple Music
  - `pause` - Pause music playback
  - `play` - Play the current song in Apple Music
  - `playlists` - Choose a playlist to play
  - `shuffle` - Toggle shuffle mode in Apple Music
  - `volume` - Set the volume for Apple Music
  **New commands in this fork:**
  - `back` - Go back to the previous song
  - `add` - Add the song that is currently playing to a playlist
  - `new` - Create a new playlist

- **Flags:**

  ```bash
  -h, --help     Help for music
  -t, --toggle   Help message for toggle
  ```

Use `music [command] --help` for more information about a specific command.

## Run

- Clone the (fork) repository

  ```bash
  git clone https://github.com/ofdryads/apple-music-cli.git
  ```

- Navigate to the project directory

  ```bash
  cd apple-music-cli
  ```

- Build the project

  ```bash
  go build
  ```

- Run the CLI

  ```bash
  ./music
  ```