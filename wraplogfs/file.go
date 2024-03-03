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
	_stdlog     *log.Logger
	_base       expsys.File
	_writeBytes bool
}

// Close implements expsys.File
func (_d fileWithLog) Close() (e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Close")
	defer func() {
		_results := []interface{}{"FileWithLog: Close returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Close()
}

// Datasync implements expsys.File
func (_d fileWithLog) Datasync() (e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Datasync")
	defer func() {
		_results := []interface{}{"FileWithLog: Datasync returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Datasync()
}

// Dev implements expsys.File
func (_d fileWithLog) Dev() (u1 uint64, e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Dev")
	defer func() {
		_results := []interface{}{"FileWithLog: Dev returned results:", u1, e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Dev()
}

// Ino implements expsys.File
func (_d fileWithLog) Ino() (i1 wasys.Inode, e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Ino")
	defer func() {
		_results := []interface{}{"FileWithLog: Ino returned results:", i1, e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Ino()
}

// IsAppend implements expsys.File
func (_d fileWithLog) IsAppend() (b1 bool) {
	_d._stdlog.Println("FileWithLog: calling IsAppend")
	defer func() {
		_results := []interface{}{"FileWithLog: IsAppend returned results:", b1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.IsAppend()
}

// IsDir implements expsys.File
func (_d fileWithLog) IsDir() (b1 bool, e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling IsDir")
	defer func() {
		_results := []interface{}{"FileWithLog: IsDir returned results:", b1, e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.IsDir()
}

// Pread implements expsys.File
func (_d fileWithLog) Pread(buf []byte, off int64) (n int, errno expsys.Errno) {
	if _d._writeBytes {
		_d._stdlog.Println("FileWithLog: calling Pread with params:", buf, off)
	} else {
		_d._stdlog.Println("FileWithLog: calling Pread with params:", "(data)", off)
	}
	defer func() {
		if _d._writeBytes {
			_d._stdlog.Println("FileWithLog: Pread returned results:", n, errno, ", buffer:", buf)
		} else {
			_d._stdlog.Println("FileWithLog: Pread returned results:", n, errno)
		}
	}()
	return _d._base.Pread(buf, off)
}

// Pwrite implements expsys.File
func (_d fileWithLog) Pwrite(buf []byte, off int64) (n int, errno expsys.Errno) {
	if _d._writeBytes {
		_d._stdlog.Println("FileWithLog: calling Pwrite with params:", buf, off)
	} else {
		_d._stdlog.Println("FileWithLog: calling Pwrite with params:", "(data)", off)
	}
	defer func() {
		_results := []interface{}{"FileWithLog: Pwrite returned results:", n, errno}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Pwrite(buf, off)
}

// Read implements expsys.File
func (_d fileWithLog) Read(buf []byte) (n int, errno expsys.Errno) {
	if _d._writeBytes {
		_d._stdlog.Println("FileWithLog: calling Read with params:", buf)
	} else {
		_d._stdlog.Println("FileWithLog: calling Read with params:", "(data)")
	}
	defer func() {
		if _d._writeBytes {
			_d._stdlog.Println("FileWithLog: Read returned results:", n, errno, ", buffer:", buf)
		} else {
			_d._stdlog.Println("FileWithLog: Read returned results:", n, errno)
		}
	}()
	return _d._base.Read(buf)
}

// Readdir implements expsys.File
func (_d fileWithLog) Readdir(n int) (dirents []expsys.Dirent, errno expsys.Errno) {
	_params := []interface{}{"FileWithLog: calling Readdir with params:", n}
	_d._stdlog.Println(_params...)
	defer func() {
		_d._stdlog.Printf("FileWithLog: Readdir returned results: %+v %s", dirents, errno)
	}()
	return _d._base.Readdir(n)
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
func (_d fileWithLog) Seek(offset int64, whence int) (newOffset int64, errno expsys.Errno) {
	_params := []interface{}{"FileWithLog: calling Seek with params:", offset, printWhence(whence)}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"FileWithLog: Seek returned results:", newOffset, errno}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Seek(offset, whence)
}

// SetAppend implements expsys.File
func (_d fileWithLog) SetAppend(enable bool) (e1 expsys.Errno) {
	_params := []interface{}{"FileWithLog: calling SetAppend with params:", enable}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"FileWithLog: SetAppend returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.SetAppend(enable)
}

// Stat implements expsys.File
func (_d fileWithLog) Stat() (s1 wasys.Stat_t, e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Stat")
	defer func() {
		_d._stdlog.Printf("FileWithLog: Stat returned results: %+v %s", s1, e1)
	}()
	return _d._base.Stat()
}

// Sync implements expsys.File
func (_d fileWithLog) Sync() (e1 expsys.Errno) {
	_d._stdlog.Println("FileWithLog: calling Sync")
	defer func() {
		_results := []interface{}{"FileWithLog: Sync returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Sync()
}

// Truncate implements expsys.File
func (_d fileWithLog) Truncate(size int64) (e1 expsys.Errno) {
	_params := []interface{}{"FileWithLog: calling Truncate with params:", size}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"FileWithLog: Truncate returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Truncate(size)
}

// Utimens implements expsys.File
func (_d fileWithLog) Utimens(atim int64, mtim int64) (e1 expsys.Errno) {
	_params := []interface{}{"FileWithLog: calling Utimens with params:", atim, mtim}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"FileWithLog: Utimens returned results:", e1}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Utimens(atim, mtim)
}

// Write implements expsys.File
func (_d fileWithLog) Write(buf []byte) (n int, errno expsys.Errno) {
	if _d._writeBytes {
		_d._stdlog.Println("FileWithLog: calling Write with params:", buf)
	} else {
		_d._stdlog.Println("FileWithLog: calling Write with params:", "(data)")
	}
	defer func() {
		_results := []interface{}{"FileWithLog: Write returned results:", n, errno}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Write(buf)
}
