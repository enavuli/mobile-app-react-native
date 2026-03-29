package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err!= nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func GetOSName() string {
	switch runtime.GOOS {
	case "darwin":
		return "macOS"
	case "windows":
		return "Windows"
	default:
		return "Linux"
	}
}

func GetOSArchitecture() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86-64"
	case "arm64":
		return "aarch64"
	default:
		return runtime.GOARCH
	}
}

func GetCommitHash() string {
	path, err := os.Executable()
	if err!= nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256([]byte(path))
	return hex.EncodeToString(hash[:])
}

func GetProjectRoot() string {
	_, filename, _, ok := runtime.Caller(1)
	if!ok {
		log.Fatal("Failed to determine project root")
	}
	return filepath.Dir(filepath.Dir(filename))
}

func GetVersion() string {
	return "1.0.0" // Replace with actual versioning logic
}

func GetAppID() string {
	return "com.example.mobile-app-react-native"
}

func GetAppVersion() string {
	return "1.0.0" // Replace with actual versioning logic
}

func GetBuildNumber() string {
	return "1234" // Replace with actual build number logic
}

func GetBuildTime() string {
	return "2023-02-20 14:30:00" // Replace with actual build time logic
}

func GetBuildOS() string {
	return GetOSName()
}

func GetBuildArchitecture() string {
	return GetOSArchitecture()
}

func GetBuildMachine() string {
	return "example-machine" // Replace with actual machine name logic
}

func GetBuildCompiler() string {
	return "go" // Replace with actual compiler logic
}

func GetBuildHash() string {
	return GetCommitHash()
}

func GetBuildID() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", GetBuildOS(), GetBuildArchitecture(), GetBuildMachine(), GetBuildCompiler(), GetBuildHash())
}