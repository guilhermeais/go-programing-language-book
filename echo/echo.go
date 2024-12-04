package echo

import (
	"fmt"
	"io"
	"strings"
)

func Echo1(args []string) string {
	var s, sep string

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func Echo2(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return fmt.Sprint(strings.Join(args[1:], " "))
}

func Echo3(args []string, w io.Writer) {
	if len(args) == 0 {
		w.Write([]byte(""))
		return
	}
	fmt.Fprint(w, strings.Join(args[1:], " "))
}
