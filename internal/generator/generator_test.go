package generator

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/OlegChuev/randpass/internal/config"
)

func TestGenerator_Generate(t *testing.T) {
	tests := []struct {
		name    string
		config  *config.Config
		wantErr bool
	}{
		{
			name: "default config",
			config: &config.Config{
				Length:    16,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			wantErr: false,
		},
		{
			name: "only lowercase",
			config: &config.Config{
				Length:    12,
				NoLower:   false,
				NoUpper:   true,
				NoDigits:  true,
				NoSymbols: true,
			},
			wantErr: false,
		},
		{
			name: "all disabled should error",
			config: &config.Config{
				Length:    16,
				NoLower:   true,
				NoUpper:   true,
				NoDigits:  true,
				NoSymbols: true,
			},
			wantErr: true,
		},
		{
			name: "zero length should error",
			config: &config.Config{
				Length:    0,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.config)
			password, err := g.Generate()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Generator.Generate() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Generator.Generate() unexpected error: %v", err)
				return
			}

			if len(password) != tt.config.Length {
				t.Errorf("Generator.Generate() password length = %d, want %d", len(password), tt.config.Length)
			}

			// Verify password contains only allowed characters
			charSet := g.buildCharacterSet()
			for _, char := range password {
				found := false
				for _, allowedChar := range charSet {
					if char == allowedChar {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Generator.Generate() password contains disallowed character: %c", char)
				}
			}
		})
	}
}

func TestGenerator_buildCharacterSet(t *testing.T) {
	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name: "all enabled",
			config: &config.Config{
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			},
			expected: LowerChars + UpperChars + DigitChars + SymbolChars,
		},
		{
			name: "only lowercase",
			config: &config.Config{
				NoLower:   false,
				NoUpper:   true,
				NoDigits:  true,
				NoSymbols: true,
			},
			expected: LowerChars,
		},
		{
			name: "no symbols",
			config: &config.Config{
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: true,
			},
			expected: LowerChars + UpperChars + DigitChars,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.config)
			result := g.buildCharacterSet()
			if result != tt.expected {
				t.Errorf("Generator.buildCharacterSet() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// Benchmark tests for password generation performance

func BenchmarkGenerator_Generate_Default(b *testing.B) {
	cfg := &config.Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_Generate_ShortPassword(b *testing.B) {
	cfg := &config.Config{
		Length:    8,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_Generate_LongPassword(b *testing.B) {
	cfg := &config.Config{
		Length:    64,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_Generate_OnlyLowercase(b *testing.B) {
	cfg := &config.Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   true,
		NoDigits:  true,
		NoSymbols: true,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_Generate_OnlyDigits(b *testing.B) {
	cfg := &config.Config{
		Length:    16,
		NoLower:   true,
		NoUpper:   true,
		NoDigits:  false,
		NoSymbols: true,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_Generate_NoSymbols(b *testing.B) {
	cfg := &config.Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: true,
	}
	g := New(cfg)

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_buildCharacterSet(b *testing.B) {
	cfg := &config.Config{
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	for b.Loop() {
		_ = g.buildCharacterSet()
	}
}

func BenchmarkGenerator_generateSecurePassword(b *testing.B) {
	cfg := &config.Config{Length: 16}
	g := New(cfg)
	charSet := g.buildCharacterSet()

	for b.Loop() {
		_, err := g.generateSecurePassword(charSet, 16)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark different password lengths to understand scaling
func BenchmarkGenerator_PasswordLength(b *testing.B) {
	lengths := []int{4, 8, 16, 32, 64, 128}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("length_%d", length), func(b *testing.B) {
			cfg := &config.Config{
				Length:    length,
				NoLower:   false,
				NoUpper:   false,
				NoDigits:  false,
				NoSymbols: false,
			}
			g := New(cfg)

			b.ResetTimer()
			for b.Loop() {
				_, err := g.Generate()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Benchmark different character set sizes
func BenchmarkGenerator_CharacterSetSize(b *testing.B) {
	testCases := []struct {
		name   string
		config *config.Config
	}{
		{
			name: "digits_only_10_chars",
			config: &config.Config{
				Length: 16, NoLower: true, NoUpper: true, NoDigits: false, NoSymbols: true,
			},
		},
		{
			name: "lowercase_only_26_chars",
			config: &config.Config{
				Length: 16, NoLower: false, NoUpper: true, NoDigits: true, NoSymbols: true,
			},
		},
		{
			name: "alphanumeric_62_chars",
			config: &config.Config{
				Length: 16, NoLower: false, NoUpper: false, NoDigits: false, NoSymbols: true,
			},
		},
		{
			name: "all_chars_95_chars",
			config: &config.Config{
				Length: 16, NoLower: false, NoUpper: false, NoDigits: false, NoSymbols: false,
			},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			g := New(tc.config)

			b.ResetTimer()
			for b.Loop() {
				_, err := g.Generate()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Benchmark crypto/rand performance for comparison
func BenchmarkCryptoRand_SingleByte(b *testing.B) {
	buf := make([]byte, 1)

	for b.Loop() {
		_, err := rand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRand_BatchBytes(b *testing.B) {
	buf := make([]byte, 16)

	for b.Loop() {
		_, err := rand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark to measure memory allocation patterns
func BenchmarkGenerator_MemoryAllocation(b *testing.B) {
	cfg := &config.Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	b.ReportAllocs()

	for b.Loop() {
		_, err := g.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark specifically the random number generation bottleneck
func BenchmarkGenerator_RandomGeneration(b *testing.B) {
	charSet := LowerChars + UpperChars + DigitChars + SymbolChars
	charSetLen := len(charSet)
	password := make([]byte, 16)

	for b.Loop() {
		for j := range 16 {
			randomBytes := make([]byte, 1)
			if _, err := rand.Read(randomBytes); err != nil {
				b.Fatal(err)
			}
			password[j] = charSet[int(randomBytes[0])%charSetLen]
		}
	}
}

// Performance baseline test - useful for regression testing
func TestGenerator_PerformanceBaseline(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	cfg := &config.Config{
		Length:    16,
		NoLower:   false,
		NoUpper:   false,
		NoDigits:  false,
		NoSymbols: false,
	}
	g := New(cfg)

	// Generate passwords to test performance
	const iterations = 1000

	for i := 0; i < iterations; i++ {
		_, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Logf("Generated %d passwords successfully", iterations)
}
