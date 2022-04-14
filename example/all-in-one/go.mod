module github.com/uptrace/bunrouter/example/all-in-one

go 1.17

replace github.com/uptrace/bunrouter => ../..

replace github.com/uptrace/bunrouter/extra/reqlog => ../../extra/reqlog

require (
	github.com/klauspost/compress v1.15.1
	github.com/rs/cors v1.8.2
	github.com/uptrace/bunrouter v1.0.14
	github.com/uptrace/bunrouter/extra/reqlog v1.0.14
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)
