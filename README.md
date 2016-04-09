# differ-xlsx

[![Build Status](https://travis-ci.org/satooon/differ-xlsx.svg?branch=master)](https://travis-ci.org/satooon/differ-xlsx)
[![Coverage Status](https://coveralls.io/repos/github/satooon/differ-xlsx/badge.svg?branch=master)](https://coveralls.io/github/satooon/differ-xlsx?branch=master)

## Install

```
$ go get github.com/satooon/differ-xlsx
```

## Examples

```
$ differ-xlsx test3_org.xlsx test3_diff.xlsx

# output
# Sheet1 : modify
# Sheet3 : new sheet
# Sheet2 : delete sheet
```
