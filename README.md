# Shiremock - A Sensible API Mocking Framework

> Wiremock annoyed me so I made a knockoff.

[![Go](https://github.com/djpiper28/shiremock/actions/workflows/main.yml/badge.svg)](https://github.com/djpiper28/shiremock/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/djpiper28/shiremock/graph/badge.svg?token=78OFQ6I434)](https://codecov.io/gh/djpiper28/shiremock)
[![Go Reference](https://pkg.go.dev/badge/github.com/djpiper28/shiremock.svg)](https://pkg.go.dev/github.com/djpiper28/shiremock)

## Quick Start

> You can see `example\_test.go` for an example on how to use this library.


## Killer Features
### First Class JSON Support

This library comes with JSON matchers with implementation for a custom `shiremock:"required"` tag for required fields, this allows any JSON
object to be matched easily.

### Type Safety

Not a JSON config lol.

(I know you can use Java for Wiremock, but that is just not hip and trendy)
