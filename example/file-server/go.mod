module github.com/bobTheBuilder7/bunrouter/example/file-server

go 1.17

replace github.com/bobTheBuilder7/bunrouter => ../..

replace github.com/bobTheBuilder7/bunrouter/extra/reqlog => ../../extra/reqlog

require (
	github.com/bobTheBuilder7/bunrouter v1.0.21
	github.com/bobTheBuilder7/bunrouter/extra/reqlog v1.0.21
)

require (
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
