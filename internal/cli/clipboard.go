package cli

import (
	"fmt"
	"runtime"

	"github.com/atotto/clipboard"
)

// copyToClipboard copies the given text to the system clipboard
func copyToClipboard(text string) error {
	if !isClipboardSupported() {
		return fmt.Errorf("clipboard not supported on this system")
	}

	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	return nil
}

// isClipboardSupported checks if clipboard operations are supported
func isClipboardSupported() bool {
	// Basic check for headless environments
	// This is a simplified check - in production you might want more sophisticated detection
	switch runtime.GOOS {
	case "linux":
		// On Linux, clipboard might not be available in headless environments
		// The clipboard package will handle the actual detection
		return true
	case "darwin", "windows":
		// macOS and Windows generally support clipboard
		return true
	default:
		return false
	}
}
