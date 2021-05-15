<a href="https://github.com/ketion-so">
    <img src="https://avatars.githubusercontent.com/u/83997411?s=200&v=4" alt="Ketion.so logo" title="Ketion.so" align="right" height="90" />
</a>

# go-notion

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/ketion-so/go-notion?color=green)
[![Go Reference](https://pkg.go.dev/badge/github.com/ketion-so/go-notion.svg)](https://pkg.go.dev/github.com/ketion-so/go-notion)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ketion-so/go-notion)
![GitHub Repo stars](https://img.shields.io/github/stars/ketion-so/go-notion)

 [![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com/)
[![Test](https://github.com/ketion-so/go-notion/actions/workflows/test.yml/badge.svg)](https://github.com/ketion-so/go-notion/actions/workflows/test.yml)
[![reviewdog](https://github.com/ketion-so/go-notion/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/ketion-so/go-notion/actions/workflows/reviewdog.yml)
[![Coverage Status](https://coveralls.io/repos/github/ketion-so/go-notion/badge.svg?branch=main)](https://coveralls.io/github/ketion-so/go-notion?branch=main)

Go written [Notion](https://www.notion.so) SDK.

*Note: The [Notion API](https://developers.notion.com/) is in beta phase*

## Supported APIs

It supports all APIs for Notion API (as for 2021-05-15).

* [x] Blocks
* [x] Databases
* [x] Pages
* [x] Search
* [x] Users

Is this package needs update, please raise an issue or a PR.

## Installation

Include this  in your code as below:

```golang
import "github.com/ketion-so/go-notion/notion"
```

or using `go get`

```console
$  go get -u github.com/ketion-so/go-notion
```

## Usage

Initialize the client as below:

```golang
client := notion.NewClient("access token")
```

Here are some examples:

## List Dashboard

```golang
resp, _ := client.Databases.List(ctx)
fmt.Println(resp.Databases)
```

## Get user

```golang
user, _ := client.Users.Get(ctx, "user ID")
```


## License

This tool is released under Apache License 2.0. See details [here](./LICENSE)
