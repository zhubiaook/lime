# lime

Gin source code learning

## Install

```bash
go get -u github.com/zhubiaook/lime
```

## Examples

**quick start**

```go
func main() {
	r := lime.New()
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	slog.Info("Server is running on port 9000")
	r.Run(":9000")
}
```

**method chaining**

```go
func main() {
	r := lime.New()

	// method chaining
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).GET("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	slog.Info("Server is running on port 9000")
	r.Run(":9000")
}
```

**functional options**

```go
func main() {
	r := lime.New(withIndex(), with404())
	slog.Info("Server is running on port 9000")
	r.Run(":9000")
}

func withIndex() lime.OptionFunc {
	return func(r *lime.Engine) {
		r.GET("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("index"))
		})
	}
}

func with404() lime.OptionFunc {
	return func(r *lime.Engine) {
		r.GET("/404", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
		})
	}
}
```
