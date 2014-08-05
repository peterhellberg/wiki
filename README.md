Wiki
====

A tiny wiki using [Goji](http://goji.io/), [BoltDB](https://github.com/boltdb/bolt) and [Blackfriday](https://github.com/russross/blackfriday).

[![Build Status](https://travis-ci.org/peterhellberg/wiki.svg?branch=master)](https://travis-ci.org/peterhellberg/wiki)
[![GoDoc](https://godoc.org/github.com/peterhellberg/wiki?status.svg)](https://godoc.org/github.com/peterhellberg/wiki)

### Installation

```bash
go get -u github.com/peterhellberg/wiki
```

### Usage

You can specify two (optional) parameters `-bind` and `-db`

```bash
wiki -bind=":7272" -db="/tmp/foo.db"
```

### License

```
The MIT License (MIT)

Copyright (c) 2014 Peter Hellberg

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
