# Ministats – Monitor System Stats from Your Terminal 🖥️📊

**Ministats** is a lightweight CLI tool written in Go to monitor system usage (CPU, memory, processes, and more).
```
A blazing-fast, beautiful, and hackable CLI tool written in Go to monitor system-level resource usage like CPU, memory, top processes, and users. Inspired by Linux power tools — powered by gopsutil, cobra, and lipgloss.
```

## 🚀 Features
- ✅  Total CPU and memory usage
- ✅  Top 5 processes by CPU or memory
- ✅  Logged-in users, system uptime, OS info
- ✅  cobra CLI with flags like --sort=cpu or --sort=mem
- ✅  Styled terminal output with `lipgloss`
- ✅  Manual page integration via `man ministats`

## Tech Stack
<p> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" title="Go"/> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/bash/bash-original.svg" width="40" title="Bash"/> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" width="40" title="Linux"/> </p>

## 📦 Installation

### Option 1: Shell Script (Recommended)

```bash
git clone https://github.com/krrishmahar/ministats.git
cd ministats
./script.sh
```

### Option 2: Debian Package (Coming Soon)

```bash
wget https://github.com/krrishmahar/ministats/releases/latest/download/ministats.deb
sudo apt install ./ministats.deb
```

### Usage

```bash
ministats cpu           # CPU usage
ministats mem           # Memory usage
ministats ps --sort=cpu # Top CPU processes
ministats ps --sort=mem # Top memory processes
ministats users         # Logged-in users
man ministats           # Read manual
```

### Example Architecture

```lua
Go (cobra + gopsutil)
        ↓
    CLI binary
        ↓
Optional .deb package
        ↓
/usr/local/bin + man page

```

## 📃 License

```
MIT License — see LICENSE file for details.
```

### 👥 Contributing

```
Pull requests, issues, and suggestions are welcome!
If this CLI helps you, consider ⭐ starring the repo.
```