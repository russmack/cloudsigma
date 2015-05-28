# cloudsigma

A Golang REST client for CloudSigma IaaS.

[![Build Status](https://travis-ci.org/russmack/cloudsigma.svg?branch=master)](https://travis-ci.org/russmack/cloudsigma)

---
#### Status: Started.
---

## Usage
```
Set your login in the examples/config.json.
cd examples
go run example.go
```
Api usage:
```
c := cloudsigma.NewCloudStatus()
args := c.NewList()
args.Location = "zrh"
args.Format = "xml"
client := &cloudsigma.Client{}
resp, err := client.Call(args)
```
See another example at:
https://github.com/russmack/cloudsigmarepl


## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
