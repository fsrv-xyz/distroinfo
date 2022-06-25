# distroinfo

The `distro-info` and `distro-info-data` debian/ubuntu package provides centralized lists of code-names and release
history for the supported distributions (Currently: Debian and Ubuntu).

This package makes these information usable inside golang projects.

---

### installation
```shell
$ go get github.com/fsrv-xyz/distro-info
```

### usage
See the contents of the [example directory](https://github.com/fsrv-xyz/distroinfo/example) for basic usage.

---

### internals
#### updates
The `distroinfo` package is updated regularly. If you need to update the information on your own, you can use the generate script inside the `data` directory.

Requirements:
* python3
* bash
* curl
* an internet connection ;)

The current version of the latest debian package is pinned in [generate.go](https://github.com/fsrv-xyz/distroinfo/blob/master/generate.go).
Run `go generate` to update the information.