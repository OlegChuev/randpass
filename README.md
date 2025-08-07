# 🔑 randpass

[![CI/CD](https://github.com/OlegChuev/randpass/workflows/CI%2FCD/badge.svg)](https://github.com/OlegChuev/randpass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/OlegChuev/randpass)](https://goreportcard.com/report/github.com/OlegChuev/randpass)
[![codecov](https://codecov.io/gh/OlegChuev/randpass/branch/main/graph/badge.svg)](https://codecov.io/gh/OlegChuev/randpass)
[![Release](https://img.shields.io/github/release/OlegChuev/randpass.svg)](https://github.com/OlegChuev/randpass/releases/latest)
[![License](https://img.shields.io/github/license/OlegChuev/randpass.svg)](LICENSE)

```txt

                      _
                     | |
  _ __ __ _ _ __   __| |_ __   __ _ ___ ___
 | '__/ _` | '_ \ / _` | '_ \ / _` / __/ __|
 | | | (_| | | | | (_| | |_) | (_| \__ \__ \
 |_|  \__,_|_| |_|\__,_| .__/ \__,_|___/___/
                       | |
                       |_|


$ > A simple, fast, secure command-line password generator
```


## Features

- Generate cryptographically secure passwords
- Customizable password length
- Configurable character sets (lowercase, uppercase, digits, symbols)
- Clipboard support
- Clean, simple CLI interface

## Installation

### Download Pre-built Binaries (Recommended)

Download the latest release for your platform from the [Releases page](https://github.com/OlegChuev/randpass/releases):

```bash
# Linux amd64
wget https://github.com/OlegChuev/randpass/releases/latest/download/randpass-linux-amd64.tar.gz
tar -xzf randpass-linux-amd64.tar.gz
chmod +x randpass-linux-amd64
sudo mv randpass-linux-amd64 /usr/local/bin/randpass

# macOS amd64
wget https://github.com/OlegChuev/randpass/releases/latest/download/randpass-macos-amd64.tar.gz
tar -xzf randpass-macos-amd64.tar.gz
chmod +x randpass-macos-amd64
sudo mv randpass-macos-amd64 /usr/local/bin/randpass

# macOS arm64 (Apple Silicon)
wget https://github.com/OlegChuev/randpass/releases/latest/download/randpass-macos-arm64.tar.gz
tar -xzf randpass-macos-arm64.tar.gz
chmod +x randpass-macos-arm64
sudo mv randpass-macos-arm64 /usr/local/bin/randpass
```

### Build from Source

```bash
# Clone and build
git clone <repository-url>
cd randpass
make build

# Or install to $GOPATH/bin
make install
```

### Download Dependencies

```bash
make deps
```

## Usage

### Basic Examples

```bash
# Default 16-character password with all character types
randpass

# 24-character password without symbols
randpass -l 24 --no-symbols

# 12-character password with only uppercase and digits
randpass --length 12 --no-lower --no-symbols

# 20-character password copied to clipboard
randpass -l 20 -c

# Show help
randpass --help
```

### Command Line Options

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--length`, `-l` | int | 16 | Password length |
| `--no-lower`, `-nl` | bool | false | Exclude lowercase letters |
| `--no-upper`, `-nu` | bool | false | Exclude uppercase letters |
| `--no-digits`, `-nd` | bool | false | Exclude numbers |
| `--no-symbols`, `-ns` | bool | false | Exclude symbols |
| `--copy`, `-c` | bool | false | Copy password to clipboard |
| `--help`, `-h` | bool | false | Show help message |

### Character Sets

| Type | Characters |
|------|------------|
| Lower | `abcdefghijklmnopqrstuvwxyz` |
| Upper | `ABCDEFGHIJKLMNOPQRSTUVWXYZ` |
| Digits | `0123456789` |
| Symbols | `!@#$%^&*()-_=+[]{}<>?/` |

## Development

### Quick Start

```bash
# Set up development environment
make dev-setup

# Run tests
make test

# Run benchmarks
make bench

# Run performance regression test
make performance-test

# Build and run
make run

# Format and lint code
make lint
```

### Available Make Targets

- `make build` - Build the binary
- `make run` - Build and run with default settings
- `make test` - Run all tests
- `make lint` - Format code and run go vet
- `make help` - Show all available targets

## Security

- Uses `crypto/rand` for cryptographically secure random number generation
- No password storage or logging
- Secure random byte generation for each character selection

## Dependencies

- `github.com/atotto/clipboard` - Cross-platform clipboard access

## License

See LICENSE file for details.
