package moonsystem

import (
	"fmt"
	"moon/internal"
)

func getVersion() string {
	const version = "1.2.0"
	return version
}

var ReservedWords = map[string]bool{
	// Common help/version
	"help":      true,
	"-h":        true,
	"--help":    true,
	"version":   true,
	"-v":        true,
	"--version": true,

	// System-level keywords
	"init":   true,
	"config": true,
	"clear":  true,
	"reset":  true,
	"about":  true,
}

func IsReserved(name string) bool {
	_, isReserved := ReservedWords[name]
	return isReserved
}

func SystemCommandHandler(command string) {
	switch command {

	case "help", "--help", "-h":
		fmt.Println(internal.PrintMoonMessage(internal.Normal, `
🌙 MOON CLI — Lightweight Alias & Command Manager

Usage:
  moon [command] [arguments]

Core Commands:
  create <alias>                Create a new alias for a command group
  add <alias> <cmd...>          Add one or more commands under an alias
  <alias>                       Select and paste a command from alias list
  delete <alias>                Remove an alias and its commands
  list                          List all existing aliases
  edit <alias>                  Edit or rename existing alias commands
  rename <old> <new>            Rename an alias

System Commands:
  help, -h, --help              Show this help message
  version, -v, --version        Show version info
  init                          Initialize or repair Moon data store
  config                        Show Moon CLI configuration paths
  clear                         Clear screen and temporary cache
  reset                         Reset Moon data (requires confirmation)
  about                         Display info about Moon CLI

Examples:
  moon create deploy
  moon add deploy "git push" "vercel --prod"
  moon deploy
  moon delete deploy

Tips:
  - Commands are saved inside data/store.json
  - When you run "moon <alias>", it shows a list to copy/paste quickly.
  - Aliases make repetitive terminal tasks faster and consistent.

For more info, visit: https://github.com/nk1044/moon-cli
`))

	case "-v", "--version", "version":
		fmt.Println(internal.PrintMoonMessage(internal.Normal,
			"🌙 Moon CLI version "+getVersion()+"\nMade with ❤️ by Neeraj Kumar"),
		)

	case "init":
		fmt.Println(internal.PrintMoonMessage(internal.Success,
			"Moon data store initialized successfully (data/store.json)."),
		)

	case "config":
		fmt.Println(internal.PrintMoonMessage(internal.Normal, `
Moon CLI Configuration:
  Data file:   ./data/store.json
  Platform:    auto-detected (macOS/Linux/Windows)
  Clipboard:   pbcopy / xclip / clip (auto)
`))

	case "clear":
		fmt.Print("\033[H\033[2J") // ANSI escape for clearing terminal
		fmt.Println(internal.PrintMoonMessage(internal.Success, "Screen cleared."))

	case "reset":
		fmt.Println(internal.PrintMoonMessage(internal.Warning,
			"⚠️  This will delete all alias data. Run manually if you’re sure:\n  rm -f data/store.json"),
		)

	case "about":
		fmt.Println(internal.PrintMoonMessage(internal.Normal, `
🌙 Moon CLI
A lightweight command alias manager for your terminal.

Features:
  - Create reusable command groups (aliases)
  - Quickly copy commands to clipboard
  - Cross-platform clipboard support
  - Simple JSON-based storage

Author: Neeraj Kumar
Version: `+getVersion()+`
`))

	default:
		fmt.Println(internal.PrintMoonMessage(internal.Warning,
			"Unknown system command: "+command+"\nUse 'moon help' to see available commands."))
	}
}
