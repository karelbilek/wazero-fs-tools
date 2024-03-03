# wazero-fs-tools
Helper tools for custom Wazero FS

Go doc: https://pkg.go.dev/github.com/karelbilek/wazero-fs-tools

## memfs

MemFS is a in-memory filesystem. Note that minimal amount of functionality is actually done;
feel free to add a PR.

The actual underlying implementation is github.com/blang/vfs/memfs.
Here it's just a tiny wrapper around it.

## sysfs

SysFS is just a verbatim copy of wazero internal sysfs. Useful for mixing with wraplogfs.

## wraplogfs

WrapLogFS is a wrapper around existing filesystem that logs all inputs/outputs

# Example - log FS

```go
import (
    "github.com/tetratelabs/wazero"
    // note: either this or wazero-fs-tools/sysfs needs to be renamed in import
    expsysfs "github.com/tetratelabs/wazero/experimental/sysfs"
    
    "github.com/karelbilek/wazero-fs-tools/sysfs"
    "github.com/karelbilek/wazero-fs-tools/wraplogfs"
)

// ...
func main() {
    rootFS := sysfs.DirFS("/")
    wrappedFS := wraplogfs.New(rootFS, os.Stdout, false)

    fsConfig := wazero.NewFSConfig()
    fsConfig = fsConfig.(expsysfs.FSConfig).WithSysFSMount(wrappedFS, "/") 
    // now all / file operations will be logged

    moduleConfig := wazero.NewModuleConfig().WithFSConfig(fsConfig)./*...*/
}
```

# Example - memory FS

```go
import (
    "log"

    "github.com/tetratelabs/wazero"
    expsysfs "github.com/tetratelabs/wazero/experimental/sysfs"
    
    "github.com/karelbilek/wazero-fs-tools/memfs"
)

// ...
func main() {
    memFS := memfs.New()

    // can write some files for start
    err := rootFS.WriteFile("tmp/foo.txt", []byte("this is content"))
    if err != nil {
        log.Fatal(err)
    }

    fsConfig := wazero.NewFSConfig()
    fsConfig = fsConfig.(expsysfs.FSConfig).WithSysFSMount(memFS, "/") 
    // all now happens in memory
}
```