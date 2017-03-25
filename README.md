# Go Package: safer

  [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/cstockton/go-safer)
  [![Go Report Card](https://goreportcard.com/badge/github.com/cstockton/go-safer?style=flat-square)](https://goreportcard.com/report/github.com/cstockton/go-safer)
  [![Coverage Status](https://img.shields.io/codecov/c/github/cstockton/go-safer/master.svg?style=flat-square)](https://codecov.io/github/cstockton/go-safer?branch=master)
  [![Build Status](http://img.shields.io/travis/cstockton/go-safer.svg?style=flat-square)](https://travis-ci.org/cstockton/go-safer)
  [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/cstockton/go-safer/master/LICENSE)


# About

Package safer provides safer access to unsafe operations by providing simple
functions with high test coverage that will never panic, instead returning
zero values.

Primary motivation was KindOf which gives a 30x performance improvement against
the same operation in the reflect package. Useful when you only need to know
the underlying kind, not manipulate the value directly.
