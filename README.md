# Go bindings for the CPython-3 C-API

> [!IMPORTANT]
> **Currently supports python-3.12 only.**

This package provides a ``go`` package named "python" under which most of the ``PyXYZ`` functions and macros of the 
public C-API of CPython have been exposed. Theoretically, you should be able use https://docs.python.org/3/c-api
and know what to type in your ``go`` program.

## Prerequisites

- `go >= 1.24`
- `python = 3.12.x`

### MacOS
  - `brew install python@3.12`
  - `brew install pkg-config`

### Linux

We will need `pkg-config` and a working `python3.12` environment to build these bindings. Make sure you have Python 
libraries and header files installed as well (`python3.12-dev` on Debian, `brew install python@3.12` on macOS, or 
`python3-devel` on Centos for example).

By default `pkg-config` will look at the `python3` library so if you want to choose a specific version just symlink 
`python-X.Y.pc` to `python3.pc` or use the `PKG_CONFIG_PATH` environment variable.

## Install

```shell
go get go.nhat.io/cpy/v3
```

## API

Some functions mix go code and call to Python function. Those functions will return and `int` and `error` type. The 
`int` represent the Python result code and the `error` represent any issue from the Go layer.

Example:

`func PyRun_AnyFile(filename string)` open `filename` and then call CPython API function 
`int PyRun_AnyFile(FILE *fp, const char *filename)`.

Therefore, its signature is `func PyRun_AnyFile(filename string) (int, error)`, the `int` represent the error code from 
the CPython `PyRun_AnyFile` function and error will be set if we failed to open `filename`.

If an error is raise before calling th CPython function `int` default to `-1`.

Take a look at some [examples](examples) and this [tutorial blogpost](https://poweruser.blog/embedding-python-in-go-338c0399f3d5).

## Versioning

We follow the versioning of the CPython API. The version of this package is `3.12.x` which means it supports the CPython
API version `3.12`. However, the patch version `x` is used to indicate the version of this package, not the CPython API.
If you see a version `3.12.7`, it doesn't mean the module only supports CPython API version `3.12.7`, it's just the 7th
time we patch to support the CPython API version `3.12`.

## Contributing

Contributions are welcome! See [details](CONTRIBUTING.md).  

## Relations

### Relation to `DataDog/go-python3`

This project is a community maintained successor to [`DataDog/go-python3`](https://github.com/DataDog/go-python3), which will get archived in December 2021.

If you use the Go package `github.com/DataDog/go-python3` in your code, you can use `go.nhat.io/cpy/v3` as a drop-in replacement. We intend to not introduce breaking changes.

### Relation to `sbinet/go-python`

This project was inspired by [`sbinet/go-python`](https://github.com/sbinet/go-python) (Go bindings for the CPython-2 C-API).

### Relation to `sublime-security/cpy3`

This project was a fork of [`sublime-security`](https://github.com/sublime-security) (Go bindings for the CPython-3 C-API) which only supports python 3.10.  
