// sysfs is a almost verbatim copy of internal package sysfs from wazero.
// As it's internal, it cannot be used outside of wazero; however, it's useful to be
// able to create AdaptFS and SysFS and instrument them and experiment on them further.
//
// Network and polling was removed, as it was dependent on other internal packages.
// As was all testing.
//
// path.go is copied over from internal platform package.
//
// Original wazero docs:
//
// Package sysfs includes a low-level filesystem interface and utilities needed
// for WebAssembly host functions (ABI) such as WASI and runtime.GOOS=js.
//
// The name sysfs was chosen because wazero's public API has a "sys" package,
// which was named after https://github.com/golang/sys.
package sysfs
