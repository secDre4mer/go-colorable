package main

import (
	"io"
	"os"

	"github.com/secDre4mer/go-colorable"
)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}
