# cloudsigma

A Golang REST client for the CloudSigma IaaS.

[![Build Status](https://travis-ci.org/russmack/cloudsigma.svg?branch=master)](https://travis-ci.org/russmack/cloudsigma)

## Usage
```
Set your login in the examples/config.json.
cd examples
go run example.go
```
API Usage:
```
c := cloudsigma.NewCloudStatus()
args := c.NewList()
args.Location = "zrh"
args.Format = "xml"
client := &cloudsigma.Client{}
resp, err := client.Call(nil, args)
```
See another example at:
https://github.com/russmack/cloudsigmarepl

## Currently Supports:

- [X] Cloud status
- [X] Locations
- [X] Capabilities
- [X] Profile
- [X] Balance
- [X] Subscriptions
- [X] Transactions
- [X] Pricing
- [X] Discounts
- [X] Licenses
- [X] Current usage
- [X] Burst usage
- [X] Daily usage
- [X] Notification preferences [list, edit]
- [X] Notification contacts [list, create, edit, delete]
- [X] Servers [list, create, edit, delete, start, stop, shutdown]
- [X] Drives [list, list detailed, get single, create, delete]
- [X] Snapshots [list, list detailed, get single, delete]
- [X] Vlans [list, list detailed, get single]
- [X] IP addresses [list, list detailed, get single]
- [X] ACLs [list, get single]
- [X] Tags [list, get single]
- [X] Logs [list]
- [X] Jobs [list]
- [X] Firewall Policies [list]
- [X] Keypairs [list]

## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
