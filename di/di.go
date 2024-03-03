package di

import (
	"fmt"
	"io"
)

func Great(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
