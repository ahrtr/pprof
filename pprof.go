package pprof

import (
	"net/http"
	"net/http/pprof"
	"runtime"
)

const pprofHTTPPrefix = "/debug/pprof"

// pprofHandlers returns a map of pprof handlers keyed by the HTTP path.
func pprofHandlers() map[string]http.Handler {
	// set only when there's no existing setting
	if runtime.SetMutexProfileFraction(-1) == 0 {
		// 1 out of 5 mutex events are reported, on average
		runtime.SetMutexProfileFraction(5)
	}

	m := make(map[string]http.Handler)

	m[pprofHTTPPrefix+"/"] = http.HandlerFunc(pprof.Index)
	m[pprofHTTPPrefix+"/profile"] = http.HandlerFunc(pprof.Profile)
	m[pprofHTTPPrefix+"/symbol"] = http.HandlerFunc(pprof.Symbol)
	m[pprofHTTPPrefix+"/cmdline"] = http.HandlerFunc(pprof.Cmdline)
	m[pprofHTTPPrefix+"/trace "] = http.HandlerFunc(pprof.Trace)
	m[pprofHTTPPrefix+"/heap"] = pprof.Handler("heap")
	m[pprofHTTPPrefix+"/goroutine"] = pprof.Handler("goroutine")
	m[pprofHTTPPrefix+"/threadcreate"] = pprof.Handler("threadcreate")
	m[pprofHTTPPrefix+"/allocs"] = pprof.Handler("allocs")
	m[pprofHTTPPrefix+"/block"] = pprof.Handler("block")
	m[pprofHTTPPrefix+"/mutex"] = pprof.Handler("mutex")

	return m
}

// RegisterPprof registers pprof handlers and return a new http.ServeMux,
// which wraps the original http.Handler.
func RegisterPprof(handler http.Handler) *http.ServeMux {
	httpmux := http.NewServeMux()
	for path, h := range pprofHandlers() {
		httpmux.Handle(path, h)
	}

	httpmux.Handle("/", handler)

	return httpmux
}
