package wraplogfs

import (
	"fmt"
	"io"
	"log"

	expsys "github.com/tetratelabs/wazero/experimental/sys"
	wasys "github.com/tetratelabs/wazero/sys"
)

// fileWithLog implements expsys.File that is instrumented with logging
type fileWithLog struct {
	stdlog     *log.Logger
	base       expsys.File
	writeBytes bool

	name   string
	fsName string
}

func (d fileWithLog) log(nm string) func(format string, params ...any) {
	return func(format string, params ...any) {
		txt := fmt.Sprintf(format, params...)
		txt = fmt.Sprintf("WrapLogFile %s %s %s: %s", d.fsName, d.name, nm, txt)
		d.stdlog.Println(txt)
	}
}

// Close implements expsys.File
func (d fileWithLog) Close() (e1 expsys.Errno) {
	l := d.log("Chmod")
	l("calling with params: <>")
	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Close()
}

// Datasync implements expsys.File
func (d fileWithLog) Datasync() (e1 expsys.Errno) {
	l := d.log("Datasync")
	l("calling with params: <>")
	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Datasync()
}

// Dev implements expsys.File
func (d fileWithLog) Dev() (u1 uint64, e1 expsys.Errno) {
	l := d.log("Dev")
	l("calling with params: <>")
	defer func() {
		l("returned results: %s %s", u1, e1)
	}()
	return d.base.Dev()
}

// Ino implements expsys.File
func (d fileWithLog) Ino() (i1 wasys.Inode, e1 expsys.Errno) {
	l := d.log("Ino")
	l("calling with params: <>")
	defer func() {
		l("returned results: %s %s", i1, e1)
	}()
	return d.base.Ino()
}

// IsAppend implements expsys.File
func (d fileWithLog) IsAppend() (b1 bool) {
	l := d.log("IsAppend")
	l("calling with params: <>")

	defer func() {
		l("returned results: %t", b1)
	}()
	return d.base.IsAppend()
}

// IsDir implements expsys.File
func (d fileWithLog) IsDir() (b1 bool, e1 expsys.Errno) {
	l := d.log("IsDir")
	l("calling with params: <>")

	defer func() {
		l("returned results: %t %s", b1, e1)
	}()
	return d.base.IsDir()
}

// Pread implements expsys.File
func (d fileWithLog) Pread(buf []byte, off int64) (n int, errno expsys.Errno) {
	l := d.log("Pread")

	if d.writeBytes {
		l("calling with params: %v %d", buf, off)
	} else {
		l("calling with params: (none) %d", off)
	}
	defer func() {
		if d.writeBytes {
			l("returned results: %d %s, buffer: %v", n, errno, buf)
		} else {
			l("returned results: %d %s", n, errno)
		}
	}()
	return d.base.Pread(buf, off)
}

// Pwrite implements expsys.File
func (d fileWithLog) Pwrite(buf []byte, off int64) (n int, errno expsys.Errno) {
	l := d.log("Pwrite")

	if d.writeBytes {
		l("calling with params: %v %d", buf, off)
	} else {
		l("calling with params: (none) %d", off)
	}

	defer func() {
		l("returned results: %d %s", n, errno)
	}()
	return d.base.Pwrite(buf, off)
}

// Read implements expsys.File
func (d fileWithLog) Read(buf []byte) (n int, errno expsys.Errno) {
	l := d.log("Read")

	if d.writeBytes {
		l("calling with params: %v", buf)
	} else {
		l("calling with params: (none)")
	}
	defer func() {
		if d.writeBytes {
			l("returned results: %d %s, buffer: %v", n, errno, buf)
		} else {
			l("returned results: %d %s", n, errno)
		}
	}()
	return d.base.Read(buf)
}

// Readdir implements expsys.File
func (d fileWithLog) Readdir(n int) (dirents []expsys.Dirent, errno expsys.Errno) {
	l := d.log("Readdir")
	l("calling with params: %d", n)
	defer func() {
		l("returned results: %+v %s", dirents, errno)
	}()
	return d.base.Readdir(n)
}

func printWhence(whence int) string {
	switch whence {
	case io.SeekStart:
		return "SeekStart"
	case io.SeekEnd:
		return "SeekEnd"
	case io.SeekCurrent:
		return "SeekCurrent"
	default:
		return fmt.Sprintf("Invalid whence (%d)", whence)
	}
}

// Seek implements expsys.File
func (d fileWithLog) Seek(offset int64, whence int) (newOffset int64, errno expsys.Errno) {
	l := d.log("Seek")
	l("calling with params: %d %d", offset, printWhence(whence))

	defer func() {
		l("returned results: %d %s", newOffset, errno)
	}()
	return d.base.Seek(offset, whence)
}

// SetAppend implements expsys.File
func (d fileWithLog) SetAppend(enable bool) (e1 expsys.Errno) {
	l := d.log("SetAppend")
	l("calling with params: %t", enable)

	defer func() {
		l("returned results: %d %s", e1)
	}()
	return d.base.SetAppend(enable)
}

// Stat implements expsys.File
func (d fileWithLog) Stat() (s1 wasys.Stat_t, e1 expsys.Errno) {
	l := d.log("Stat")
	l("calling with params: <>")

	defer func() {
		l("returned results: %+v %s", s1, e1)
	}()
	return d.base.Stat()
}

// Sync implements expsys.File
func (d fileWithLog) Sync() (e1 expsys.Errno) {
	l := d.log("Sync")
	l("calling with params: <>")
	defer func() {
		l("returned results: %s", e1)

	}()
	return d.base.Sync()
}

// Truncate implements expsys.File
func (d fileWithLog) Truncate(size int64) (e1 expsys.Errno) {
	l := d.log("Truncate")
	l("calling with params: %d", size)

	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Truncate(size)
}

// Utimens implements expsys.File
func (d fileWithLog) Utimens(atim int64, mtim int64) (e1 expsys.Errno) {
	l := d.log("Utimens")
	l("calling with params: %d %d", atim, mtim)
	defer func() {
		l("returned results: %s", e1)
	}()
	return d.base.Utimens(atim, mtim)
}

// Write implements expsys.File
func (d fileWithLog) Write(buf []byte) (n int, errno expsys.Errno) {
	l := d.log("Write")

	if d.writeBytes {
		l("calling with params: %v %d", buf)
	} else {
		l("calling with params: (none)")
	}

	defer func() {
		l("returned results: %d %s", n, errno)
	}()
	return d.base.Write(buf)
}
