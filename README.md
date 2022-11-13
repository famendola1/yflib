# Yahoo Fantasy API Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/famendola1/yflib.svg)](https://pkg.go.dev/github.com/famendola1/yflib)
![License](https://img.shields.io/badge/License-Apache-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/famendola1/yflib)](https://goreportcard.com/report/github.com/famendola1/yflib)

## Installation
~~~bash
go get github.com/famendola1/yflib
~~~

## Yahoo Endpoints
The Yahoo offical documentation for their [Fantasy Sports API](https://developer.yahoo.com/fantasysports/guide) is not comprehensive and incomplete, despite being the offical. For a more complete overview of the supported endpoints, see the [README](https://github.com/edwarddistel/yahoo-fantasy-baseball-reader#yahoo-fantasy-api-docs).

## Before You Start
The functionality of this package require the use of a `*http.Client` that is configured for the Yahoo Fantasy API endpoint. You can use the [github.com/famendola1/yauth](https://pkg.go.dev/github.com/famendola1/yauth) package to configure a `*http.Client` to use.

## Contributions
If you's like to add functionality to this library, feel free to submit a PR 🙂
