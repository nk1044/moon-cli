# Moon CLI — Lightweight Alias & Command Manager

<span style="color:lavender;">

        ███╗   ███╗  ██████╗   ██████╗  ███╗   ██╗
        ████╗ ████║ ██╔═══██╗ ██╔═══██╗ ████╗  ██║
        ██╔████╔██║ ██║   ██║ ██║   ██║ ██╔██╗ ██║
        ██║╚██╔╝██║ ██║   ██║ ██║   ██║ ██║╚██╗██║
        ██║ ╚═╝ ██║ ╚██████╔╝ ╚██████╔╝ ██║ ╚████║
        ╚═╝     ╚═╝  ╚═════╝   ╚═════╝  ╚═╝  ╚═══╝

</span>

Moon CLI is a lightweight command-line tool for managing and executing aliases, making repetitive terminal tasks faster and consistent.

---

## Usage

```bash
moon [command] [arguments]
````

### Core Commands

| Command                | Description                                |
| ---------------------- | ------------------------------------------ |
| `create <alias>`       | Create a new alias for a command group     |
| `add <alias> <cmd...>` | Add one or more commands under an alias    |
| `<alias>`              | Select and paste a command from alias list |
| `delete <alias>`       | Remove an alias and its commands           |
| `list`                 | List all existing aliases                  |
| `edit <alias>`         | Edit or rename existing alias commands     |
| `rename <old> <new>`   | Rename an alias                            |

### System Commands

| Command                  | Description                             |
| ------------------------ | --------------------------------------- |
| `help, -h, --help`       | Show this help message                  |
| `version, -v, --version` | Show version info                       |
| `init`                   | Initialize or repair Moon data store    |
| `config`                 | Show Moon CLI configuration paths       |
| `clear`                  | Clear screen and temporary cache        |
| `reset`                  | Reset Moon data (requires confirmation) |
| `about`                  | Display info about Moon CLI             |

### Examples

```bash
moon create deploy
moon add deploy "git push" "vercel --prod"
moon deploy
moon delete deploy
```

**Tips:**

* Commands are saved inside `data/store.json`
* When you run `moon <alias>`, it shows a list to copy/paste quickly.
* Aliases make repetitive terminal tasks faster and consistent.

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/nk1044/moon-cli.git
cd moon-cli
```

### 2. Build the project

Make sure you have Go installed. Then run:

```bash
go build -o moon main.go
```

This will generate a `moon` binary in the current directory.

### 3. Add Moon CLI to your PATH

#### macOS / Linux

```bash
# Move binary to a global path
sudo mv moon /usr/local/bin/moon

# Make executable
sudo chmod +x /usr/local/bin/moon

# Verify installation
moon --version
```

#### Windows (PowerShell / CMD)

```powershell
# Move moon.exe to a folder included in your PATH, e.g., C:\Windows\System32
move .\moon.exe C:\Windows\System32\moon.exe

# Verify installation
moon --version
```

Now you can run `moon` from any terminal.

---

## Data Storage

Moon CLI stores alias data in:

```text
data/store.json
```

The CLI automatically creates and updates this file when you add, edit, or remove aliases.

---

## Contributing

Contributions are welcome! Open issues or pull requests for bug fixes, features, or improvements.

---

## License

[MIT License](LICENSE)

---

## Tags

`go` `cli` `command-line` `alias-manager` `productivity` `tool`
