package moonuser

import (
	"errors"
	"fmt"
	"moon/internal"
)

// UserCommandHandler routes the user commands like:
//   moon create <alias>
//   moon add <alias> <commands...>
//   moon list
//   moon <alias>
//   moon delete <alias>
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
		return internal.SaveStore(store)

	case "list":
		if len(store.Aliases) == 0 {
			fmt.Println(internal.PrintMoonMessage(internal.Warning, "No aliases found. Use 'moon create <alias>'"))
			return nil
		}
		fmt.Println(internal.PrintMoonMessage(internal.Normal, "Available aliases:\n"))
		for alias, data := range store.Aliases {
			fmt.Printf("• %s → %v\n", alias, data.Commands)
		}
		return nil

	default:
		// Possibly: moon <alias>
		alias := action
		if aliasData, ok := store.Aliases[alias]; ok {
			if len(aliasData.Commands) == 0 {
				fmt.Println(internal.PrintMoonMessage(internal.Warning, "No commands under alias: "+alias))
				return nil
			}
			choice, err := internal.SelectCommand(aliasData.Commands)
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
