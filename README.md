# TaskMaster

TaskMaster is a small terminal-based task logger written in Go. It supports **two modes**:

- **Command mode**: run a single command (e.g. `log`, `add`, `done`) and exit.
- **REPL mode**: run with no args to enter an interactive prompt and execute commands repeatedly.

## Features

- Create tasks with optional **description**, **deadline**, and **status**
- View tasks in compact or detailed form
- Mark tasks done/undo, edit tasks, remove tasks, swap ordering
- Purge completed tasks or clear everything
- Tasks are automatically sorted by deadline after adds/edits

## Requirements

- Go (per `go.mod`: `go 1.25.5`)

## Install / Build

Clone and build:

```bash
go build -o taskmaster ./cmd/app
```

Run:

```bash
./taskmaster
```

## Usage

### REPL mode (interactive)

Run **with no arguments**:

```bash
./taskmaster
```

You’ll get a prompt like:

```text
> 
```

Type commands the same way as command mode (examples below). To exit:

- `exit`
- or press Ctrl-C

### Command mode (single command)

Run with a command name as the first argument:

```bash
./taskmaster help
./taskmaster log
```

## Commands

### `log`

Print tasks.

- `-d`: detailed table (includes description, status, deadline, time period)
- `-l`: long-term only (filters out tasks due within the next 24 hours)
- `-dl`: detailed + long-term

Examples:

```bash
./taskmaster log
./taskmaster log -d
./taskmaster log -l
./taskmaster log -dl
```

### `add`

Add one or more tasks by title, with optional flags **per task**:

- `-desc`: description
- `-time`: deadline offset (see “Time format”)
- `-status`: `done`, `pending`, `overdue`, `paused`

Examples:

```bash
./taskmaster add "Buy milk"
./taskmaster add "File taxes" -desc "collect receipts" -time 2w
./taskmaster add "Gym" -time 1d -status pending "Book tickets" -time 3m
```

### `rm`

Remove task(s) by id (as shown in `log`):

```bash
./taskmaster rm 2
./taskmaster rm 1 3 4
```

### `done` / `undo`

Mark task(s) done / mark them back to pending:

```bash
./taskmaster done 1
./taskmaster undo 1 2
```

### `edit`

Edit a task by id:

- `-title`: new title
- `-desc`: new description
- `-time`: new deadline offset
- `-status`: `done`, `pending`, `overdue`, `paused`

Examples:

```bash
./taskmaster edit 3 -title "Read book" -time 1w
./taskmaster edit 2 -status done
```

### `swap`

Swap positions of two tasks by id:

```bash
./taskmaster swap 1 2
```

### `purge`

Remove completed tasks:

```bash
./taskmaster purge
```

### `clear`

Remove all tasks:

```bash
./taskmaster clear
```

### `help`

Print built-in help:

```bash
./taskmaster help
```

## Time format

When using `-time`, TaskMaster expects an **offset** with a numeric value and a suffix:

- `d` = days (e.g. `2d`)
- `w` = weeks (e.g. `1w`)
- `m` = months (30 days, e.g. `3m`)
- `y` = years (365 days, e.g. `1y`)

Values may be fractional (it parses a float), e.g. `0.5w`.

## Data storage

Tasks are stored as CSV at:

- `~/.tasks.csv`

Each run loads tasks from that file and saves back on exit.

## Notes

- In REPL mode, commands are parsed shell-style (quoted strings work).
- Task ids are **1-based** (the first task is `1`).
