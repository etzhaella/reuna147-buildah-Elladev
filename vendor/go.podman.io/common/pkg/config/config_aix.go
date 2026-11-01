//go:build aix

package config

var defaultHelperBinariesDir = []string{
	"/usr/bin",
	"/usr/local/bin",
	"/usr/local/libexec/podman",
	"/usr/local/lib/podman",
}
