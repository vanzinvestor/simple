# Simple

Simple work with [httprouter](https://github.com/julienschmidt/httprouter). This project fork group router from [Hunsin/router](https://gist.github.com/Hunsin/26b2021757e831554d4f59a52a5c9152) and add more feature

## How to use

```go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vanzinvestor/simple"
	"github.com/vanzinvestor/simple/middleware"
)

func main() {
	r := simple.New()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("Hi Simple"))
	})

	r.GET("/hi/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		w.Write([]byte("Hi " + name))
	})

	http.ListenAndServe(":9000", r)
}
```
