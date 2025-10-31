package internal

import (
    "fmt"
    "os/exec"
    "runtime"
)

func CopyToClipboard(text string) error {
    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "darwin": // macOS
        cmd = exec.Command("pbcopy")

    case "linux": // Linux
        if isCommandAvailable("xclip") {
            cmd = exec.Command("xclip", "-selection", "clipboard")
        } else if isCommandAvailable("wl-copy") {
            cmd = exec.Command("wl-copy")
        } else {
            return fmt.Errorf("no clipboard utility found (install xclip or wl-clipboard)")
        }

    case "windows": // Windows
        cmd = exec.Command("cmd", "/c", "clip")

    default:
        return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
    }

    in, err := cmd.StdinPipe()
    if err != nil {
        return err
    }

    if err := cmd.Start(); err != nil {
        return err
    }

    _, _ = in.Write([]byte(text))
    _ = in.Close()

    return cmd.Wait()
}

// helper to check if a command exists in PATH
func isCommandAvailable(name string) bool {
    _, err := exec.LookPath(name)
    return err == nil
}
