// Package main is the entry point for the mnemosyne CLI.
//
// Mnemosyne is a document storage and retrieval tool that uses hybrid search
// (full-text + vector similarity) with reciprocal rank fusion and local
// cross-encoder reranking. All inference runs locally via ONNX Runtime.
package main

import (
	"fmt"
	"os"

	"github.com/gandazgul/mnemosyne/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
