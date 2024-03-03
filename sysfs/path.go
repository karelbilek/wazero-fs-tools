//go:build !windows

package sysfs

// toPosixPath returns the input, as only windows might return backslashes.
func toPosixPath(in string) string { return in }
