// wraplogfs is a wazero filesystem that wraps another, existing filesystem, and logs
// all inputs/outputs to writer.
//
// Was first generated by hexdigest/gowrap, but then edited for better outputs in some cases
package wraplogfs

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"strings"

	expsys "github.com/tetratelabs/wazero/experimental/sys"
	wasys "github.com/tetratelabs/wazero/sys"
)

type fsWithLog struct {
	// intentionally does NOT embed unimplemented; I do NOT want to be forward-compatible;
	// I want to break on missing funcs

	stdlog     *log.Logger
	base       expsys.FS
	writeBytes bool
	fsName     string
}

// New returns a new filesystem on top of another filesystem.
// writeBytes controls if all bytes are written on stdout on reads/writes, or just "(data)".
func New(base expsys.FS, stdout io.Writer, writeBytes bool, name string) expsys.FS {
	return fsWithLog{
		base:       base,
		stdlog:     log.New(stdout, "", log.LstdFlags),
		writeBytes: writeBytes,
		fsName:     name,
	}
}

func (d fsWithLog) log(nm string) func(format string, params ...any) {
	return func(format string, params ...any) {
		txt := fmt.Sprintf(format, params...)
		txt = fmt.Sprintf("WrapLogFS %s %s: %s", d.fsName, nm, txt)
		d.stdlog.Println(txt)
	}
}

// Chmod implements sys.FS
func (d fsWithLog) Chmod(path string, perm fs.FileMode) (e1 expsys.Errno) {
	l := d.log("Chmod")
	l("calling with params: %q %s", path, perm)
	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Chmod(path, perm)
}

// Link implements sys.FS
func (d fsWithLog) Link(oldPath string, newPath string) (e1 expsys.Errno) {
	l := d.log("Link")
	l("calling with params: %q %q", oldPath, newPath)
	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Link(oldPath, newPath)
}

// Lstat implements sys.FS
func (d fsWithLog) Lstat(path string) (s1 wasys.Stat_t, e1 expsys.Errno) {
	l := d.log("Lstat")
	l("calling with params: %q", path)

	defer func() {
		l("returned results: %+v %s", s1, e1)
	}()
	return d.base.Lstat(path)
}

// Mkdir implements sys.FS
func (d fsWithLog) Mkdir(path string, perm fs.FileMode) (e1 expsys.Errno) {
	l := d.log("Mkdir")
	l("calling with params: %q %s", path, perm)

	defer func() {
		l("returned results: %+v %s", e1)
	}()
	return d.base.Mkdir(path, perm)
}

func printOflags(flag expsys.Oflag) string {
	st := []string{}
	flags := map[expsys.Oflag]string{

		expsys.O_RDONLY:    "O_RDONLY",
		expsys.O_RDWR:      "O_RDWR",
		expsys.O_WRONLY:    "O_WRONLY",
		expsys.O_APPEND:    "O_APPEND",
		expsys.O_CREAT:     "O_CREAT",
		expsys.O_DIRECTORY: "O_DIRECTORY",
		expsys.O_DSYNC:     "O_DSYNC",
		expsys.O_EXCL:      "O_EXCL",
		expsys.O_NOFOLLOW:  "O_NOFOLLOW",
		expsys.O_NONBLOCK:  "O_NONBLOCK",
		expsys.O_RSYNC:     "O_RSYNC",
		expsys.O_SYNC:      "O_SYNC",
		expsys.O_TRUNC:     "O_TRUNC",
	}
	for f, d := range flags {
		if flag&f != 0 {
			st = append(st, d)
		}
	}
	if len(st) == 0 {
		return "(none)"
	} else {
		return strings.Join(st, "|")
	}
}

// OpenFile implements sys.FS
func (d fsWithLog) OpenFile(path string, flag expsys.Oflag, perm fs.FileMode) (f1 expsys.File, e1 expsys.Errno) {
	l := d.log("OpenFile")
	l("calling with params: %q %s; %s", path, printOflags(flag), perm)

	defer func() {
		l("returned results: %T %+v %s", f1, f1, e1)
	}()
	fl, errno := d.base.OpenFile(path, flag, perm)
	return fileWithLog{
		base:       fl,
		stdlog:     d.stdlog,
		writeBytes: d.writeBytes,
		name:       path,
		fsName:     d.fsName,
	}, errno
}

// Readlink implements sys.FS
func (d fsWithLog) Readlink(path string) (s1 string, e1 expsys.Errno) {
	l := d.log("Readlink")
	l("calling with params: %q", path)

	defer func() {
		l("returned results: %q %s", s1, e1)
	}()
	return d.base.Readlink(path)
}

// Rename implements sys.FS
func (d fsWithLog) Rename(from string, to string) (e1 expsys.Errno) {
	l := d.log("Rename")
	l("calling with params: %q %q", from, to)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Rename(from, to)
}

// Rmdir implements sys.FS
func (d fsWithLog) Rmdir(path string) (e1 expsys.Errno) {
	l := d.log("Rmdir")
	l("calling with params: %q", path)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Rmdir(path)
}

// Stat implements sys.FS
func (d fsWithLog) Stat(path string) (s1 wasys.Stat_t, e1 expsys.Errno) {
	l := d.log("Stat")
	l("calling with params: %q", path)

	defer func() {
		l("returned results: %s", e1)
	}()

	return d.base.Stat(path)
}

// Symlink implements sys.FS
func (d fsWithLog) Symlink(oldPath string, linkName string) (e1 expsys.Errno) {
	l := d.log("Symlink")
	l("calling with params: %q %q", oldPath, linkName)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Symlink(oldPath, linkName)
}

// Unlink implements sys.FS
func (d fsWithLog) Unlink(path string) (e1 expsys.Errno) {
	l := d.log("Unlink")
	l("calling with params: %q", path)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Unlink(path)
}

// Utimens implements sys.FS
func (d fsWithLog) Utimens(path string, atim int64, mtim int64) (e1 expsys.Errno) {
	l := d.log("Symlink")
	l("calling with params: %q %d %d", path, atim, mtim)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Utimens(path, atim, mtim)
}
