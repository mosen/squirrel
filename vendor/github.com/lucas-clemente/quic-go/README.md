# A QUIC server implementation in pure Go

<img src="docs/quic.png" width=303 height=124>

[![Godoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/lucas-clemente/quic-go)
[![Linux Build Status](https://img.shields.io/travis/lucas-clemente/quic-go.svg?style=flat-square&label=linux+build)](https://travis-ci.org/lucas-clemente/quic-go)
[![Windows Build Status](https://img.shields.io/appveyor/ci/lucas-clemente/quic-go.svg?style=flat-square&label=windows+build)](https://ci.appveyor.com/project/lucas-clemente/quic-go/branch/master)
[![Code Coverage](https://img.shields.io/codecov/c/github/lucas-clemente/quic-go/master.svg?maxAge=2592000&style=flat-square)](https://codecov.io/gh/lucas-clemente/quic-go/)

quic-go is an implementation of the [QUIC](https://en.wikipedia.org/wiki/QUIC) protocol in Go. While we're not far from being feature complete, there's still a lot of work to do regarding performance and security. At the moment, we do not recommend use in production systems. We appreciate any feedback :)

## Roadmap

Done:

- Basic protocol with support for QUIC version 30-33
- HTTP/2 support
- Crypto (RSA / ECDSA certificates, curve25519 for key exchange, chacha20-poly1305 as cipher)
- Loss detection and retransmission (currently fast retransmission & RTO)
- Flow Control
- Congestion control using cubic

Major TODOs:

- Security, especially DOS protections
- Performance
- Better packet loss detection
- Support for QUIC version 34
- Connection migration
- QUIC client
- Public API design
- Integration into caddy (mostly to figure out the right server API)

## Guides

Installing deps:

    go get -t

Running tests:

    go test ./...

Running the example server:

    go run example/main.go -www /var/www/

Using the `quic_client` from chromium:

    quic_client --quic-version=32 --host=127.0.0.1 --port=6121 --v=1 https://quic.clemente.io

Using Chrome:

    /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --user-data-dir=/tmp/chrome --no-proxy-server --enable-quic --origin-to-force-quic-on=quic.clemente.io:443 --host-resolver-rules='MAP quic.clemente.io:443 127.0.0.1:6121' https://quic.clemente.io

## Usage

See the [example server](example/main.go) or our [fork](https://github.com/lucas-clemente/caddy) of caddy. Starting a QUIC server is very similar to the standard lib http in go:

```go
http.Handle("/", http.FileServer(http.Dir(wwwDir)))
h2quic.ListenAndServeQUIC("localhost:4242", "/path/to/cert/chain.pem", "/path/to/privkey.pem", nil)
```

## Building on Windows

Due to the low Windows timer resolution (see [StackOverflow question](http://stackoverflow.com/questions/37706834/high-resolution-timers-millisecond-precision-in-go-on-windows)) available with Go 1.6.x, some optimizations might not work when compiled with this version of the compiler. Please use Go 1.7 on Windows.
