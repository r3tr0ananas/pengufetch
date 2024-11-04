package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
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
	out, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		return "Unknown"
	}

	splitUptime := strings.Split(string(out), " ")

	return strings.Join(splitUptime[1:], " ")
}
