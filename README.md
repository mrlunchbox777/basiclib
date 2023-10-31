# github.com/mrlunchbox777/basiclib

Some common go functionality strapped into a library.

## Usage

### Initialize your module

```sh
go mod init example.com/my-golib-demo
```

### Get the go-lib module

Note that you need to include the **v** in the version tag.

Note that you need to replace `X` with the desired target.

```sh
go get github.com/mrlunchbox777/basiclib@vX.X.X
```

```go
package main

import (
    "fmt"

    "github.com/mrlunchbox777/basiclib"
)

func main() {
    fmt.Println(golib.Add(2,3))
    fmt.Println(golib.Subtract(2,3))
}
```

## Testing

This should provide sufficient testing.

```sh
go test
```

## Tagging

X should be replaced by the desired target.

```sh
git tag vX.X.X
git push origin --tags
```

## Attribution

[See docs/ATTRIBUTION](/docs/ATTRIBUTION)
