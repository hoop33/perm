# Perm

> A REPL for HTTP(S)

![](https://github.com/hoop33/perm/workflows/Build/badge.svg)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](code_of_conduct.md) 

## Disclaimer &mdash; Very Early Alpha

Lots of things changing, early days, etc. Use at your own risk.

## Overview

When we interact with URLs from the command line, most of us reach for the excellent [curl](https://curl.haxx.se). And with good reason.

In a "new phone, who dis?" way, though, curl runs your request and promptly forgets it. In many cases, that's fine. You're just running a request or two, so specifying all your information in each request is no big deal.

In some cases, however, you're running multiple requests to the same domain, using the same headers, and maybe mostly the same query or form fields. Wouldn't _persistence_ be nice here?

Enter Perm. Perm is curl with persistence. You launch perm, set some headers, set some query fields, and run a query. _Then_, you run another query with a _relative URL_, and all the headers and query strings you set before are _still sent_. Example session:

```sh
$ perm
perm 0.1.0
Type "help" for more information.
http://localhost:3000> set header Accept application/json
http://localhost:3000> set header api-key abcd1234
http://localhost:3000> set var foo bar
http://localhost:3000> set var baz bat
http://localhost:3000> get https://example.com
# Results for https://example.com/?foo=bar&baz=bat with both headers set
https://example.com> unset var baz
https://example.com> get /test
# Results for https://example.com/?foo=bar with both headers set
```

## Installation

You can install Perm a few different ways.

### Install with Go

If you have a working Go installation, type:

```sh
$ go get -u github.com/hoop33/perm
```

### Download Binaries

Binaries for macOS, Linux, and Windows are available on the Releases page (coming).

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
- [ ] Add exit command
- [ ] Load/save environments
- [ ] Configure HTTP client (timeout, insecure, ??)
- [ ] Add S3 URLs
- [x] GET with query string
- [ ] POST
- [ ] PUT
- [ ] DELETE
- [ ] HEAD
- [ ] Cookies
- [x] Headers
- [ ] Open in Browser

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
