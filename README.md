Wiki
====

A tiny wiki using [BoltDB](https://github.com/boltdb/bolt) and [Blackfriday](https://github.com/russross/blackfriday).

[![Build Status](https://travis-ci.org/peterhellberg/wiki.svg?branch=master)](https://travis-ci.org/peterhellberg/wiki)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/wiki)](https://goreportcard.com/report/github.com/peterhellberg/wiki)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/wiki)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/wiki#license-mit)

## Installation

```bash
go get -u github.com/peterhellberg/wiki
```

## Usage

```bash
PORT=2017 wiki -db="/tmp/foo.db"
```

## License (MIT)

Copyright (c) 2014-2017 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
