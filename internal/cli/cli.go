package cli

import (
	"flag"
	"fmt"

	"github.com/OlegChuev/randpass/internal/config"
	"github.com/OlegChuev/randpass/internal/generator"
)

const (
	appName       = "randpass"
	appVersion    = "1.0.0"
	appDesc       = "A simple, fast, secure command-line password generator"
	defaultLength = 16
)

// Execute runs the CLI application
func Execute() error {
	cfg := config.NewConfig()

	// Define flags
	flag.IntVar(&cfg.Length, "length", defaultLength, "Password length")
	flag.IntVar(&cfg.Length, "l", defaultLength, "Password length (shorthand)")

	flag.BoolVar(&cfg.NoLower, "no-lower", false, "Exclude lowercase letters")
	flag.BoolVar(&cfg.NoLower, "nl", false, "Exclude lowercase letters (shorthand)")

	flag.BoolVar(&cfg.NoUpper, "no-upper", false, "Exclude uppercase letters")
	flag.BoolVar(&cfg.NoUpper, "nu", false, "Exclude uppercase letters (shorthand)")

	flag.BoolVar(&cfg.NoDigits, "no-digits", false, "Exclude numbers")
	flag.BoolVar(&cfg.NoDigits, "nd", false, "Exclude numbers (shorthand)")

	flag.BoolVar(&cfg.NoSymbols, "no-symbols", false, "Exclude symbols")
	flag.BoolVar(&cfg.NoSymbols, "ns", false, "Exclude symbols (shorthand)")

	flag.BoolVar(&cfg.Copy, "copy", false, "Copy password to clipboard")
	flag.BoolVar(&cfg.Copy, "c", false, "Copy password to clipboard (shorthand)")

	flag.BoolVar(&cfg.Help, "help", false, "Show help message")
	flag.BoolVar(&cfg.Help, "h", false, "Show help message (shorthand)")

	// Custom usage function
	flag.Usage = func() {
		showHelp()
	}

	// Parse flags
	flag.Parse()

	// Show help if requested
	if cfg.Help {
		showHelp()
		return nil
	}

	// Generate password
	gen := generator.New(cfg)
	password, err := gen.Generate()
	if err != nil {
		return fmt.Errorf("password generation failed: %w", err)
	}

	// Handle clipboard if requested
	if cfg.Copy {
		if err := copyToClipboard(password); err != nil {
			fmt.Println(password)
			return fmt.Errorf("failed to copy to clipboard: %w", err)
		}
		fmt.Println("Password copied to clipboard!")
	} else {
		fmt.Println(password)
	}

	return nil
}

// showHelp displays the help message
func showHelp() {
	fmt.Printf(`%s v%s
%s

Usage:
  %s [flags]

Flags:
  -l, --length int      Password length (default: %d)
  -nl, --no-lower       Exclude lowercase letters
  -nu, --no-upper       Exclude uppercase letters
  -nd, --no-digits      Exclude numbers
  -ns, --no-symbols     Exclude symbols
  -c, --copy            Copy password to clipboard
  -h, --help            Show this help message

Character Sets:
  Lower:   abcdefghijklmnopqrstuvwxyz
  Upper:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
  Digits:  0123456789
  Symbols: !@#$%%^&*()-_=+[]{}<>?/

Examples:
  %s                                    # Default %d-char password
  %s -l 24 --no-symbols                 # 24-char without symbols
  %s --length 12 --no-lower --no-symbols # 12-char uppercase and digits only
  %s -l 20 -c                           # 20-char password copied to clipboard

`, appName, appVersion, appDesc, appName, defaultLength, appName, defaultLength, appName, appName, appName)
}
