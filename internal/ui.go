package internal

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

const (
	lavender = "\x1b[38;5;147m"
	cyan     = "\x1b[38;5;123m"
	softGray = "\x1b[38;5;250m"
	red      = "\x1b[31m"
	green    = "\x1b[32m"
	yellow   = "\x1b[33m"
	reset    = "\x1b[0m"
	bold     = "\x1b[1m"
)

func PrintMoon() string {
	return fmt.Sprintf(`
%s%s
    ╔════════════════════════════════════════════════╗
    ║                                                ║
    ║            WELCOME TO MOON CLI                 ║
    ║                                                ║
    ║      Your elegant command-line companion       ║
    ║                                                ║
    ╚════════════════════════════════════════════════╝
%s

%s
        ███╗   ███╗  ██████╗   ██████╗  ███╗   ██╗
        ████╗ ████║ ██╔═══██╗ ██╔═══██╗ ████╗  ██║
        ██╔████╔██║ ██║   ██║ ██║   ██║ ██╔██╗ ██║
        ██║╚██╔╝██║ ██║   ██║ ██║   ██║ ██║╚██╗██║
        ██║ ╚═╝ ██║ ╚██████╔╝ ╚██████╔╝ ██║ ╚████║
        ╚═╝     ╚═╝  ╚═════╝   ╚═════╝  ╚═╝  ╚═══╝
%s

%s           Illuminate your workflow with elegance
%s
`,
		cyan, bold, reset,
		lavender,
		reset,
		softGray,
		reset,
	)
}

type MessageType int

const (
	Normal  MessageType = iota + 1 // 1
	Success                        // 2
	Good                           // 3
	Warning                        // 4
	Error                          // 5
)

func PrintMoonMessage(msgType MessageType, message string) string {
	var prefix string
	switch msgType {
	case Success:
		prefix = bold + green // Bold Green
	case Good:
		prefix = bold + cyan // Bold Cyan
	case Warning:
		prefix = bold + yellow // Bold Yellow
	case Error:
		prefix = bold + red // Bold Red
	case Normal:
		fallthrough // 'Normal' and 'default' are treated the same
	default:
		prefix = bold // Just bold, as in your original function
	}
	return prefix + message + reset
}

func SelectCommand(cmds []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select a command to paste",
		Items: cmds,
	}
	_, result, err := prompt.Run()
	return result, err
}
