// memfs implements a simple fake memory FS for Wazero.
//
// The actual implementation of the FS is from github.com/blang/vfs/memfs,
// this package is just wrapping that for Wazero.
//
// It implements only very small number of functions, because only those were needed
// for my purposes (that is, running ghostscript with WASI), as ghostscript only calls these.
//
// Feel free to make a PR if you need to implement some other functions.
package memfs

import (
	"errors"
	"io"
	"io/fs"

	wasys "github.com/tetratelabs/wazero/sys"

	"os"

	"github.com/tetratelabs/wazero/experimental/sys"

	"github.com/blang/vfs/memfs"
)

// New creates a new memory filesystem
func New() *MemFS {
	mfs := memfs.Create()
	mmfs := &MemFS{fs: mfs}
	return mmfs
}

// WriteFile is a helper function that writes a content to a file
func (m *MemFS) WriteFile(path string, content []byte) error {
	f, err := m.fs.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	return err
}

// ReadFile is a helper function that returns a content of a file
func (m *MemFS) ReadFile(path string) ([]byte, error) {
	f, err := m.fs.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

// MemFS is a memory-only wazero filesystem, implementing just some basic functions.
type MemFS struct {
	fs *memfs.MemFS

	sys.UnimplementedFS
}

// toOsOpenFlag is copied from wazero codebase
func toOsOpenFlag(oflag sys.Oflag) (flag int) {
	// First flags are exclusive
	switch oflag & (sys.O_RDONLY | sys.O_RDWR | sys.O_WRONLY) {
	case sys.O_RDONLY:
		flag |= os.O_RDONLY
	case sys.O_RDWR:
		flag |= os.O_RDWR
	case sys.O_WRONLY:
		flag |= os.O_WRONLY
	}

	// Run down the flags defined in the os package
	if oflag&sys.O_APPEND != 0 {
		flag |= os.O_APPEND
	}
	if oflag&sys.O_CREAT != 0 {
		flag |= os.O_CREATE
	}
	if oflag&sys.O_EXCL != 0 {
		flag |= os.O_EXCL
	}
	if oflag&sys.O_SYNC != 0 {
		flag |= os.O_SYNC
	}
	if oflag&sys.O_TRUNC != 0 {
		flag |= os.O_TRUNC
	}
	return flag
}

// OpenFile opens a file as defined in sys.File
func (m *MemFS) OpenFile(path string, flag sys.Oflag, perm fs.FileMode) (sys.File, sys.Errno) {
	f, err := m.fs.OpenFile(path, toOsOpenFlag(flag), perm)
	if err != nil {
		if errors.Is(err, memfs.ErrIsDirectory) {
			if flag&sys.O_WRONLY == 1 || flag&sys.O_RDWR == 1 {
				return nil, sys.EISDIR
			}
			// return directory as a different type
			dir := &memoryFSDir{fs: m.fs, path: path}
			return dir, 0
		}
		if errors.Is(err, os.ErrNotExist) {
			return nil, sys.ENOENT
		}
		if errors.Is(err, os.ErrExist) {
			return nil, sys.EEXIST
		}
		return nil, sys.EINVAL // just general IO error, not that important
	}
	fl := &memoryFSFile{fl: f, path: path, fs: m.fs}
	return fl, 0
}

func stat(mfs *memfs.MemFS, path string) (wasys.Stat_t, sys.Errno) {
	fst, err := mfs.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return wasys.Stat_t{}, sys.ENOENT

		}
		return wasys.Stat_t{}, sys.EIO // this should "never happen"
	}
	return wasys.NewStat_t(fst), 0
}

// Stat returns file stat as defined in sys.File
func (m *MemFS) Stat(path string) (wasys.Stat_t, sys.Errno) {
	return stat(m.fs, path)
}
