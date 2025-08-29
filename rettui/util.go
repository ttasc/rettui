package main

import "fmt"

func IntToSize(size int) string {
    units := []string{"B", "KB", "MB", "GB"}
    value := float64(size)
    unitIndex := 0

    for value >= 1024 && unitIndex < len(units)-1 {
        value /= 1024
        unitIndex++
    }

    return fmt.Sprintf("%.2f %s", value, units[unitIndex])
}
