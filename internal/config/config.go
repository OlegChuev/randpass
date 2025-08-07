package config

import (
	"fmt"
)

// Config holds the configuration for password generation
type Config struct {
	Length    int
	NoLower   bool
	NoUpper   bool
	NoDigits  bool
	NoSymbols bool
	Copy      bool
	Help      bool
}

// NewConfig creates a new Config with default values
func NewConfig() *Config {
	return &Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
		Copy:      false,
		Help:      false,
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Length <= 0 {
		return fmt.Errorf("password length must be greater than 0")
	}

	if c.NoLower && c.NoUpper && c.NoDigits && c.NoSymbols {
		return fmt.Errorf("at least one character type must be enabled")
	}

	return nil
}

// GetEnabledCharSets returns which character sets are enabled
func (c *Config) GetEnabledCharSets() map[string]bool {
	return map[string]bool{
		"lower":   !c.NoLower,
		"upper":   !c.NoUpper,
		"digits":  !c.NoDigits,
		"symbols": !c.NoSymbols,
	}
}
