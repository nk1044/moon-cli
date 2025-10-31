package moonuser

import (
	"errors"
	"fmt"
	"moon/internal"
	"strings"
)

// UserCommandHandler routes the user commands like:
//
//	moon create <alias>
//	moon add <alias> <commands...>
//	moon list
//	moon <alias>
//	moon delete <alias>
func UserCommandHandler(args []string) error {
	if len(args) < 2 {
		return errors.New("no command provided. Try: moon help")
	}

	action := args[1]
	store, err := internal.LoadStore()
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	switch action {
	case "create":
		if len(args) < 3 {
			return errors.New("usage: moon create <alias>")
		}
		alias := args[2]
		if _, exists := store.Aliases[alias]; exists {
			return fmt.Errorf("alias '%s' already exists", alias)
		}
		store.Aliases[alias] = internal.AliasData{Commands: []string{}}
		msg := internal.PrintMoonMessage(internal.Success, "Created Successfully!")
		fmt.Println(msg)
		return internal.SaveStore(store)

	case "add":
		if len(args) < 4 {
			return errors.New("usage: moon add <alias> <commands...>")
		}
		alias := args[2]
		cmds := args[3:]
		a := store.Aliases[alias]
		a.Commands = append(a.Commands, cmds...)
		store.Aliases[alias] = a
		msg := internal.PrintMoonMessage(internal.Success, "Added Successfully!")
		fmt.Println(msg)
		return internal.SaveStore(store)

	case "delete":
		if len(args) < 3 {
			return errors.New("usage: moon delete <alias>")
		}
		alias := args[2]
		if _, exists := store.Aliases[alias]; !exists {
			return fmt.Errorf("alias '%s' not found", alias)
		}
		delete(store.Aliases, alias)
		msg := internal.PrintMoonMessage(internal.Success, "Deleted Successfully!")
		fmt.Println(msg)
		return internal.SaveStore(store)

	case "edit":
		if len(args) < 3 {
			return errors.New("usage: moon edit <alias>")
		}
		alias := args[2]

		aliasData, exists := store.Aliases[alias]
		if !exists {
			return fmt.Errorf("alias '%s' not found", alias)
		}
		if len(aliasData.Commands) == 0 {
			return fmt.Errorf("alias '%s' has no commands to edit", alias)
		}

		// Step 1: select command to edit/delete
		selectedCmd, err := internal.SelectCommand(aliasData.Commands, "Select a command to edit/delete")
		if err != nil {
			return fmt.Errorf("selection cancelled: %w", err)
		}

		// Step 2: choose action
		actionChoice, err := internal.SelectCommand([]string{"Edit this command", "Delete this command", "Cancel"}, "Choose an action")
		if err != nil || actionChoice == "Cancel" {
			fmt.Println(internal.PrintMoonMessage(internal.Warning, "Cancelled."))
			return nil
		}

		switch actionChoice {
		case "Delete this command":
			// remove selectedCmd from slice
			newCmds := []string{}
			for _, cmd := range aliasData.Commands {
				if cmd != selectedCmd {
					newCmds = append(newCmds, cmd)
				}
			}
			aliasData.Commands = newCmds
			store.Aliases[alias] = aliasData
			if err := internal.SaveStore(store); err != nil {
				return fmt.Errorf("failed to save changes: %w", err)
			}
			fmt.Println(internal.PrintMoonMessage(internal.Success, "Command deleted successfully."))

		case "Edit this command":
			// Step 3: ask user for replacement command text
			fmt.Print(internal.PrintMoonMessage(internal.Normal, "Enter new command: "))
			var newCmd string
			fmt.Scanln(&newCmd)

			// update in slice
			for i, cmd := range aliasData.Commands {
				if cmd == selectedCmd {
					aliasData.Commands[i] = newCmd
					break
				}
			}
			store.Aliases[alias] = aliasData
			if err := internal.SaveStore(store); err != nil {
				return fmt.Errorf("failed to save changes: %w", err)
			}
			fmt.Println(internal.PrintMoonMessage(internal.Success, "Command updated successfully."))
		}

		return nil

	case "list":
    if len(store.Aliases) == 0 {
        fmt.Println(internal.PrintMoonMessage(internal.Warning,
            "No aliases found. Use 'moon create <alias>'"))
        return nil
    }

    fmt.Println(internal.PrintMoonMessage(internal.Normal, "Your saved aliases:\n"))

    maxLen := 0
    for alias := range store.Aliases {
        if len(alias) > maxLen {
            maxLen = len(alias)
        }
    }

    for alias, data := range store.Aliases {
        padding := strings.Repeat(" ", maxLen-len(alias))
        fmt.Printf("  🌙 %s%s │ %d command(s)\n", alias, padding, len(data.Commands))

        for _, cmd := range data.Commands {
            fmt.Printf("     ↳  %s\n", cmd)
        }
        fmt.Println()
    }

    fmt.Println(internal.PrintMoonMessage(internal.Normal,
        "Tip: Run 'moon <alias>' to view and paste commands for any alias.\n"))
    return nil


	default:
		// Possibly: moon <alias>
		alias := action
		if aliasData, ok := store.Aliases[alias]; ok {
			if len(aliasData.Commands) == 0 {
				fmt.Println(internal.PrintMoonMessage(internal.Warning, "No commands under alias: "+alias))
				return nil
			}
			choice, err := internal.SelectCommand(aliasData.Commands, "Select a command to paste")
			if err != nil {
				fmt.Println(internal.PrintMoonMessage(internal.Warning, "Cancelled"))
				return nil
			}
			// fmt.Println("\nPaste this command to run it:")
			// fmt.Println(choice)
			fmt.Println("✔", choice)
			if err := internal.CopyToClipboard(choice); err == nil {
				fmt.Println("Command copied to clipboard! Press ⌘+V or ↑ to recall it.")
			} else {
				fmt.Println("⚠️ Could not copy automatically:", err)
			}

			return nil
		}
		return fmt.Errorf("unknown command or alias: '%s'", alias)
	}
}
