# Perm

> A REPL for HTTP(S)

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](code_of_conduct.md) 

## Overview

## Installation

You can install Perm a few different ways.

### Install with Go

If you have a working Go installation, type:

```sh
$ go get -u github.com/hoop33/perm
```

### Download Binaries

Binaries for macOS, Linux, and Windows are available on the Releases page.

### Build from Source

Perm is written in Go and uses Go Modules.

Prerequisites:

* A working Go installation
* `make`

To build:

```sh
$ git clone https://github.com/hoop33/perm.git
$ cd perm
$ make deps
$ make && make install
```

## Usage

## Contributing

Please note that this project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms. The code of contact is available [here](code_of_conduct.md).

## Credits

Perm uses the following open source libraries &mdash; thank you!

* [Cobra](https://github.com/spf13/cobra.git)
* [Liner](https://github.com/peterh/liner)
* [Testify](https://github.com/stretchr/testify)
* [xdg](https://github.com/adrg/xdg)

Apologies for any inadvertent omissions.

## License

Copyright &copy; 2020 Rob Warner

Licensed under the [MIT License](https://hoop33.mit-license.org/).
