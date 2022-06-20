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
  <a target="_blank" href="https://homiedevs.com/">
    <img alt="sponsors" src="https://homiedevs.com/assets/img/logo.png">
  </a>
</p>
