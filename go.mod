module github.com/gandazgul/mnemosyne

go 1.26.0

require (
	github.com/asg017/sqlite-vec-go-bindings v0.1.6
	github.com/daulet/tokenizers v1.26.0
	github.com/fatih/color v1.18.0
	github.com/mattn/go-sqlite3 v1.14.34
	github.com/spf13/cobra v1.10.2
	// this is just the bindings and it doesnt match the version of onnxruntime. the version wrapped here is 1.23.1
	github.com/yalue/onnxruntime_go v1.16.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	golang.org/x/sys v0.42.0 // indirect
)
