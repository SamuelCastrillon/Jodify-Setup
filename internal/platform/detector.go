package platform

import "runtime"

// Detect returns the appropriate Platform implementation for the current OS
func Detect() (Platform, error) {
	return DetectGoOS(runtime.GOOS)
}

// DetectGoOS returns the appropriate Platform implementation for the given GOOS
func DetectGoOS(goos string) (Platform, error) {
	switch goos {
	case "windows":
		return NewWindows(), nil
	case "darwin":
		return NewDarwin(), nil
	default:
		return nil, ErrUnsupportedPlatform
	}
}
