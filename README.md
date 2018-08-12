cjknum
=======

[![CircleCI](https://circleci.com/gh/ryo-endo/cjknum.svg?style=svg)](https://circleci.com/gh/ryo-endo/cjknum)
[![Coverage Status](https://coveralls.io/repos/github/ryo-endo/cjknum/badge.svg)](https://coveralls.io/github/ryo-endo/cjknum)

## Description

Package cjknum implements conversions to the CJK numeric string representations of integer.

## Installation

```sh
% go get github.com/ryo-endo/cjknum
```

## Usage

```go
import "cjknum"

s := cjknum.Itoc(12345)
log.Println(s) // 一万二千三百四十五
```

## Author

[ryo-endo](https://github.com/ryo-endo/cjknum)
