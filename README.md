# Keygent

Keygent is a lightweight CLI tool for running commands with injected environment variables in a safe and predictable way.

It lets you temporarily override or add environment variables for a single command execution without affecting your system environment.

## Features

- Inject environment variables per command
- Override system environment safely
- Supports quoted values
- Cross platform (Windows, Linux, macOS)
- Simple command execution wrapper
- No shell dependency (direct execution model)

## Installation

### Option 1: Go install

```bash
go install github.com/heshanthenura/keygent/cmd/keygent@latest
```

Make sure Go is installed and `$GOPATH/bin` is in your PATH

### Option 2: Build from source

```bash
git clone https://github.com/heshanthenura/keygent.git
cd keygent
go build -o keygent ./cmd/keygent
```

## Usage

### Basic syntax

```bash
keygent set KEY=value -- command [args...]
```

#### Example: inject environment variable

```bash
keygent set TEST=123 -- node -e "console.log(process.env.TEST)"
```

Output:

```
123
```

### Multiple environment variables

```bash
keygent set NAME=Heshan AGE=22 -- node -e "console.log(process.env.NAME, process.env.AGE)"
```

Output:

```
Heshan 22
```

### Override system environment

```bash
keygent set PATH=custom -- node -e "console.log(process.env.PATH)"
```

### Using quoted values

```bash
keygent set MESSAGE="hello world" -- node -e "console.log(process.env.MESSAGE)"
```

Output:

```
hello world
```

## How it works

Keygent:

1. Reads system environment variables
2. Applies user-defined overrides
3. Builds a temporary environment
4. Runs the target command with that environment
5. Does NOT modify your system environment

## Future improvements

- `.env file support`
- `--env-file` flag
- shell mode (`--shell`)
