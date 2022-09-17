# exhaustivenew

`exhaustivenew` is a golang analyzer that finds constructors (methods New*) with uninitialized fields.

Based on [exhaustruct](https://github.com/GaijinEntertainment/go-exhaustruct).

## Installation

https://golangci-lint.run/contributing/new-linters/#configure-a-plugin

## Example

``` go
type Svc struct {
    config SvcConfig
    repo *Repo
}

func New(config SvcConfig) *Svc{
    return &Svc{ // fails with "repo is missing in Svc"
        config: config,
    }
}
```
