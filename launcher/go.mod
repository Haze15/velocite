module velocite-launcher

go 1.21

require (
	github.com/jchv/go-webview2 v0.0.0
	golang.org/x/sys v0.28.0
)

require github.com/jchv/go-winloader v0.0.0 // indirect

// bibliothèques copiées dans third_party (construction hors ligne possible)
replace (
	github.com/jchv/go-webview2 => ./third_party/go-webview2
	github.com/jchv/go-winloader => ./third_party/go-winloader
	golang.org/x/sys => ./third_party/sys
)
