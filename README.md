# qotd

__A Quote of the Day service written in Go, compliant with [RFC 865](https://tools.ietf.org/html/rfc865).__

### Build

`go build qotd.go`

### Usage

Run the server: `sudo qotd`

Connect to it via Telnet: `telnet localhost 17`

### Customizing qotd.txt

The `qotd.txt` contains a series of quotes for every day of the year. The file must contains entries in the form:

```
mmdd{quote}
```

where `mm` is a 2-digit month and begins a new line, `dd` is a 2-digit day, and "quote" can contain any character, including newline,
with the exception of `}`.

If `quote` is longer than 512 characters, it will be truncated in accordance with [RFC 865](https://tools.ietf.org/html/rfc865).

#### Notes

`qotd.txt` was originally sourced from [here](ftp://ftp.mrynet.com/USENIX/85.1/langston/sun/TODAY/qotd.txt).

### Contact

[Neil Cowburn](http://github.com/neilco)  
[@neilco](https://twitter.com/neilco)

## License

[MIT license](http://neil.mit-license.org)

Copyright (c) 2014 Neil Cowburn (http://github.com/neilco/)

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
