<a href="https://github.com/ketion-so">
    <img src="https://avatars.githubusercontent.com/u/83997411?s=200&v=4" alt="Ketion.so logo" title="Ketion.so" align="right" height="90" />
</a>

# go-notion

[![Go Reference](https://pkg.go.dev/badge/github.com/ketion-so/go-notion.svg)](https://pkg.go.dev/github.com/ketion-so/go-notion)

Go written [Notion](https://www.notion.so) SDK.

*Note: The [Notion API](https://developers.notion.com/) is in beta phase*

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

## License

This tool is released under Apache License 2.0. See details [here](./LICENSE)
