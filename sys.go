package main

import (
	"fmt"
	"math"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("/etc/os-release")

func GetOS() string {
	return os.Getenv("NAME")
}

func GetHostname() string {
	var username string

	usr, err := user.Current()
	if err != nil {
		username = "Unknown"
	} else {
		username = usr.Username
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	return fmt.Sprintf("%s@%s", username, hostname)
}

func GetKernel() string {
	file, err := os.ReadFile("/proc/version")
	if err != nil {
		return "Unknown"
	}

	kernel := strings.Split(string(file), " ")[2]

	return kernel
}

func GetUptime() string {
	file, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "Unknown"
	}

	split := strings.Split(string(file), " ")[0]

	seconds, conversionErr := strconv.ParseFloat(split, 32)
	if conversionErr != nil {
		println(conversionErr)
		return "Unknown"
	}

	return formatDuration(int(math.Round(seconds)))
}
func formatDuration(seconds int) string {
	days := seconds / 86400
	seconds %= 86400
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	var parts []string

	if days > 0 {
		if days > 1 {
			parts = append(parts, fmt.Sprintf("%d days", days))
		} else {
			parts = append(parts, fmt.Sprintf("%d day", days))
		}
	}

	if hours > 0 {
		if hours > 1 {
			parts = append(parts, fmt.Sprintf("%d hours", hours))
		} else {
			parts = append(parts, fmt.Sprintf("%d hour", hours))
		}
	}

	if minutes > 0 {
		if minutes > 1 {
			parts = append(parts, fmt.Sprintf("%d minutes", minutes))
		} else {
			parts = append(parts, fmt.Sprintf("%d minute", minutes))
		}
	}

	return strings.Join(parts, ", ")
}
