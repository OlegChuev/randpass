package generator

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/OlegChuev/randpass/internal/config"
)

const (
	// Character sets as defined in requirements
	LowerChars  = "abcdefghijklmnopqrstuvwxyz"
	UpperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DigitChars  = "0123456789"
	SymbolChars = "!@#$%^&*()-_=+[]{}<>?/~"
)

// Generator handles password generation
type Generator struct {
	config *config.Config
}

// New creates a new Generator with the given configuration
func New(cfg *config.Config) *Generator {
	return &Generator{config: cfg}
}

// Generate creates a random password based on the configuration
func (g *Generator) Generate() (string, error) {
	// Validate configuration
	if err := g.config.Validate(); err != nil {
		return "", err
	}

	// Build character set based on configuration
	charSet := g.buildCharacterSet()
	if len(charSet) == 0 {
		return "", fmt.Errorf("no character sets available for password generation")
	}

	// Generate password
	password, err := g.generateSecurePassword(charSet, g.config.Length)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %w", err)
	}

	return password, nil
}

// buildCharacterSet constructs the character set based on enabled options
func (g *Generator) buildCharacterSet() string {
	var builder strings.Builder

	if !g.config.NoLower {
		builder.WriteString(LowerChars)
	}
	if !g.config.NoUpper {
		builder.WriteString(UpperChars)
	}
	if !g.config.NoDigits {
		builder.WriteString(DigitChars)
	}
	if !g.config.NoSymbols {
		builder.WriteString(SymbolChars)
	}

	return builder.String()
}

// generateSecurePassword generates a cryptographically secure random password
func (g *Generator) generateSecurePassword(charSet string, length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be positive")
	}
	if len(charSet) == 0 {
		return "", fmt.Errorf("character set cannot be empty")
	}

	password := make([]byte, length)
	charSetLen := len(charSet)

	for i := range length {
		// Generate a random byte
		randomBytes := make([]byte, 1)
		if _, err := rand.Read(randomBytes); err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %w", err)
		}

		// Map the random byte to a character in our set
		randomIndex := int(randomBytes[0]) % charSetLen
		password[i] = charSet[randomIndex]
	}

	return string(password), nil
}
