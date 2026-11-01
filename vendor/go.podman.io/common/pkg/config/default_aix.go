//go:build aix

package config

// DefaultInitPath is the default path to the container-init binary.
// Inherited from default_common.go (//go:build !freebsd && !netbsd) which also
// compiles on AIX; this file does not redefine it.

func getDefaultCgroupsMode() string {
	// AIX has no cgroups; return "disabled" so callers get a predictable value.
	return "disabled"
}

func getDefaultLockType() string {
	return "file"
}

func getLibpodTmpDir() string {
	return "/var/run/libpod"
}

// getDefaultMachineVolumes returns default mounted volumes.
func getDefaultMachineVolumes() []string {
	return []string{"$HOME:$HOME"}
}
