pprof 
======
github.com/ahrtr/pprof is a package to help a http server to enable runtime profiling data easily. It leverages golang build-in package "net/http/pprof". 

# How to use this repo
It's super easy. You only need to call pprof.RegisterPprof(h) to wrap your original http.handler. For example, assuming your http.Server definition is something like below,

```go
s := &http.Server{
	Addr:    ":8080",
	Handler: yourHandler,
}
```

then you can make the following change to enable the runtime profiling data,
```go
import (
	"github.com/ahrtr/pprof"
)

s := &http.Server{
	Addr:    ":8080",
	Handler: pprof.RegisterPprof(yourHandler),
}
```

Once your http server is running, then you can query the runtime profiling data. The following is a couple of examples.

heap profile:
```
go tool pprof http://localhost:8080/debug/pprof/heap
```

30-second CPU profile:
```
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
```

goroutine profile:
```
curl http://localhost:8080/debug/pprof/goroutine?debug=1
```

Please refer to **[examples/httpsrv.go](examples/httpsrv.go)** to get a complete example. 
