# Keygent

<a href="https://www.producthunt.com/products/keygent?utm_source=badge-follow&utm_medium=badge&utm_source=badge-keygent" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/follow.svg?product_id=1252692&theme=light" alt="Keygent - Run&#0032;terminal&#0032;commands&#0032;with&#0032;safely&#0032;injected&#0032;env&#0032;vars | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

Keygent is a lightweight CLI tool for running commands with injected environment variables in a safe and predictable way.

It lets you temporarily override or add environment variables for a single command execution without affecting your system environment.

## Features

- Inject environment variables per command
- Load environment variables from .env file
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

- ### Basic syntax

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

- ### Multiple environment variables

```bash
keygent set NAME=Heshan AGE=22 -- node -e "console.log(process.env.NAME, process.env.AGE)"
```

Output:

```
Heshan 22
```

- ### Override system environment

```bash
keygent set PATH=custom -- node -e "console.log(process.env.PATH)"
```

- ### Using quoted values

```bash
keygent set MESSAGE="hello world" -- node -e "console.log(process.env.MESSAGE)"
```

Output:

```
hello world
```

- ### Load from `.env` file

```bash
keygent env .env -- node -e "console.log(process.env.MESSAGE)"
```

## How it works

Keygent:

1. Reads system environment variables
2. Applies user-defined overrides
3. Builds a temporary environment
4. Runs the target command with that environment
5. Does NOT modify your system environment
