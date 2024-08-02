# Taskbook.go
**Tasks, boards & notes for the terminal** 🚀🚀🚀

https://github.com/user-attachments/assets/1588a437-2680-42b0-ab61-9ddbfaf13bca

## Overview

* Create tasks & notes in different boards
* View timeline of tasks & notes
* Filter by boards
* Respects XDG env
* Configurable colors (TOML)
* Archive and restore tasks & notes
* Tab-complete arguments

## Table of Contents
* [Installation](#Installation)
    * [Requirements](#Requirements)
    * [AUR](#AUR)
    * [Git](#Git)
    * [Binary releases](#Binary-releases)
* [Usage](#Usage)
    * [Quickstart](#Quickstart)
    * [Shell Completion](#Shell-completion)
    * [Config](#Config)
        * [Example](#Example)
    * [Commmands](#Commands)
* [Comparison](#Comparison)
* [Thanks to](#Thanks-to)

## Installation
### Requirements
Needs [Nerd Fonts](https://github.com/ryanoasis/nerd-fonts) to render icons correctly.
### AUR
Use your favorite AUR helper or `makepkg` to install.
```sh
paru -S tb.go
```
or
```sh
yay -S tb.go
```
or
```sh
#Manual
git clone https://aur.archlinux.org/tb.git
cd tb
makepkg -si
```
### Git
```sh
git clone https://github.com/araaha/tb.go
cd tb.go
make
make sys-install
```
### Binary releases
You can download binaries from here:
* https://github.com/araaha/tb.go/releases

## Usage
### Quickstart
Create a task via `tb task`. You can begin tasks, or complete them. Alternatively `tb note` creates a note. All items are stored in a `taskbook.json` file stored in
* `$XDG_DATA_HOME`
* `$HOME/.local/share/taskbook`

in that order.

Deleting tasks archives them. To permanently remove them, run `tb archive --remove`. `tb clear` clears every completed task and places them in the archive. For further information, run `tb --help` to see all commands. For each subcommand, you can run `--help` for that flag to see examples.

### Shell Completion
* bash
  ```sh
  #Set up completion for tb
  eval "$(tb completion bash)"
  ```
* zsh
  ```sh
  #Set up completion for tb
  source <(tb completion bash)```
* fish
  ```sh
  #Set up completion for tb
  tb completion fish | source
  ```
* powershell
  ```sh
  #Set up completion for tb
  tb completion powershell | Out-String | Invoke-Expression```

### Config
The config file is stored in TOML. You can specify one with the `--config` flag. Alternatively, it searches in
* `$XDG_CONFIG_HOME/taskbook/taskbook.toml`
* `$HOME/.config/taskbook/taskbook.toml`
in that order.

#### Example

```toml
#Gruvbox
[colors]
white = "#ebdbb2"
red = "#fb4934"
yellow = "#fabd2f"
gray = "#928374"
green = "#b8bb26"
blue = "#83a598"
magenta = "#d3869b"
cyan = "#8ec07c"
```

You can specify the colors in hex format.

### Commands
```sh
tb help

A Taskbook

Usage:
  tb [flags]
  tb [command]

Available Commands:
  archive     Display archived items
  begin       Begin/pause task
  check       Check/uncheck task
  clear       Delete all checked items
  completion  Generate the autocompletion script for the specified shell
  delete      Delete item
  edit        Edit item description
  help        Help about any command
  list        List items by board
  move        Move item between boards
  note        Create note
  priority    Update priority of task
  restore     Restore items from archive
  star        Star/unstar item
  task        Create task
  timeline    Display timeline view

Flags:
      --config string   config file (default is $XDG_CONFIG_HOME/taskbook/taskbook.toml)
  -h, --help            help for tb
  -v, --version         Display current version

Use "tb [command] --help" for more information about a command.
```


## Comparison
Compared to [taskbook](https://github.com/klaudiosinani/taskbook):

- **tb.go** respects `$XDG` env
- Shell-completion (via [cobra](https://github.com/spf13/cobra))
- You can specify which colors to use
- Much more performant due to being compiled
- Uses [Nerd Fonts](https://github.com/ryanoasis/nerd-fonts) icons

## Thanks to
* [taskbook](https://github.com/klaudiosinani/taskbook) - this project was heavily inspired by this one
