//go:build aix

package types

const (
	// these are default paths for run and graph root for rootful users
	// for rootless path is constructed via getRootlessStorageOpts
	defaultRunRoot   string = "/var/run/containers/storage"
	defaultGraphRoot string = "/var/lib/containers/storage"
)
