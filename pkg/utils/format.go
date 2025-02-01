package utils

import "fmt"

// FormatBytes formats a byte count into a human-readable string with units (B, KB, MB, GB, TB).
func FormatBytes(bytes uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0
	for bytes >= 1024 && i < len(units)-1 {
		bytes /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %s", float64(bytes), units[i])
}
