# Parcelux (파시룩스 - Parcel + Lux)

<div align="center">

[![Go Version](https://img.shields.io/github/go-mod/go-version/luxrobo/luxpay)](https://github.com/LUXROBO/parcelux)
[![Build Workflow Status (Github Actions)](https://img.shields.io/github/workflow/status/LUXROBO/parcelux/Go)](https://github.com/LUXROBO/parcelux/actions)
[![CodeClimate Maintainability](https://img.shields.io/codeclimate/maintainability/LUXROBO/parcelux)](https://github.com/LUXROBO/parcelux/tree/main)
[![CodeClimate Issues](https://img.shields.io/codeclimate/issues/LUXROBO/parcelux)](https://github.com/LUXROBO/parcelux/tree/main)
[![CodeClimate Coverage](https://img.shields.io/codeclimate/coverage/LUXROBO/parcelux)](https://github.com/LUXROBO/parcelux/tree/main/test)
[![Github LICENSE](https://img.shields.io/github/license/LUXROBO/parcelux)](https://github.com/LUXROBO/parcelux/blob/main/LICENSE)
[![Lines of Code](https://img.shields.io/tokei/lines/github/LUXROBO/parcelux)](https://github.com/LUXROBO/parcelux/tree/main/src)

</div>

## Description
> A parcel delivery tracking Go library for LUXROBO, utilizing [sweettracker API](https://tracking.sweettracker.co.kr/).

## Features
- Given a tracking number (i.e. 운송장번호), track current status of corresponding parcel.

## Installation
```bash
go get github.com/luxrobo/parcelux
```

## Usage
```go
import "github.com/luxrobo/parcelux/client"

// Add your own API key from sweet tracker
API_KEY := ""
parceluxClient := client.NewParceluxClient(API_KEY)

// Init your own tracking number (i.e. 운송장번호)
trackNum := ""
trackResp := parceluxClient.trackParcel(trackNum)
fmt.Println("trackResp.Status:", trackResp.Status)
```
