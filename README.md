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

To launch the application, just type `perm`:

```sh
$ perm
```

After a welcome banner, Perm will display a prompt with your current base URL:

```sh
$ perm
perm 0.1.0
Type "help" for more information.
http://localhost:3000>
```

You interact with Perm through this prompt. You type a command, Perm executes it, and returns you to this prompt. You can exit Perm by typing `Ctrl-D` or using the `exit` command.

### Commands

`commands`
Displays a list of all available commands

`get <url>`
Performs an HTTP GET on the specified URL. This URL can be relative or absolute. If relative, Perm uses the current base URL to GET the specified URL. If absolute, Perm GETs the specified URL _and_ replaces its current base URL. Examples:

```sh
https://grailbox.com> get /uses
```
GETs <https://grailbox.com/uses> and retains `https://grailbox.com` as the base URL.

```sh
https://grailbox.com> get https://github.com/hoop33/perm
```
GETs <https://github.com/hoop33/perm> and changes the current base URL to `https://github.com`.

## TODO

- [ ] Colorize text
- [ ] Load/save environments
- [ ] Configure HTTP client (timeout, insecure, ??)
- [ ] Add S3 URLs
- [ ] GET with query string
- [ ] POST
- [ ] PUT
- [ ] DELETE
- [ ] HEAD
- [ ] Cookies
- [ ] Headers
- [ ] Open in Browser
- [ ] Why did I add environment variables?

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
