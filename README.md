# Simple

Simple work with [httprouter](https://github.com/julienschmidt/httprouter). This project fork from [Hunsin/router](https://gist.github.com/Hunsin/26b2021757e831554d4f59a52a5c9152) and add more feature

## How to use

```go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vanzinvestor/simple"
)

func main() {
	r := simple.New()

	r.GET("/", Index)

	http.ListenAndServe(":9000", r)
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello"))
}
```
