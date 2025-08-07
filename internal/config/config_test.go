package config

import (
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid default config",
			config: &Config{
				Length:    16,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			wantErr: false,
		},
		{
			name: "zero length should error",
			config: &Config{
				Length:    0,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			wantErr: true,
		},
		{
			name: "negative length should error",
			config: &Config{
				Length:    -5,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			wantErr: true,
		},
		{
			name: "all character types disabled should error",
			config: &Config{
				Length:    16,
				NoLower:   true,
				NoUpper:   true,
				NoDigits:  true,
				NoSymbols: true,
			},
			wantErr: true,
		},
		{
			name: "only one character type enabled is valid",
			config: &Config{
				Length:    16,
				NoLower:   false,
				NoUpper:   true,
				NoDigits:  true,
				NoSymbols: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()

	if cfg.Length != 16 {
		t.Errorf("NewConfig() Length = %d, want 16", cfg.Length)
	}

	if cfg.NoLower || cfg.NoUpper || cfg.NoDigits || cfg.NoSymbols || cfg.Copy || cfg.Help {
		t.Errorf("NewConfig() should have all boolean flags as false by default")
	}
}

func TestConfig_GetEnabledCharSets(t *testing.T) {
	cfg := &Config{
		NoLower:   false,
		NoUpper:   true,
		NoDigits:  false,
		NoSymbols: true,
	}

	enabled := cfg.GetEnabledCharSets()

	if !enabled["lower"] {
		t.Error("Expected lower to be enabled")
	}
	if enabled["upper"] {
		t.Error("Expected upper to be disabled")
	}
	if !enabled["digits"] {
		t.Error("Expected digits to be enabled")
	}
	if enabled["symbols"] {
		t.Error("Expected symbols to be disabled")
	}
}
