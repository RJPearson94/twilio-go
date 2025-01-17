// Forked code from https://github.com/ajg/form

// Copyright 2014 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package form implements encoding and decoding of application/x-www-form-urlencoded data.
package client

const (
	implicitKey = "_"
	omittedKey  = "-"

	defaultDelimiter = '.'
	defaultEscape    = '\\'
	defaultKeepZeros = false
)
