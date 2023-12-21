# Install

```
go get github.com/turk/free-google-translate
```

# Example usage

``` go
package main

import (
	"fmt"
	"github.com/turk/free-google-translate"
	"net/http"
)

func main() {
	client := http.Client{}
	t := translator.NewTranslator(&client)
	result, _ := t.Translate("Hello!", "en", "ro")
	fmt.Println(result)
	// Output: "Salut!"
}
```

## Sponsors

<p align="center">
  <a target="_blank" href="https://nolocation.io/">
    <img alt="sponsors" src="https://nolocation.io/assets/img/logo.png">
  </a>
</p>
