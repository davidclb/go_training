package hello

import (
	"fmt"
	"strings"
)

func Say(names []string) string {
	return fmt.Sprintf("Hello, %s!", strings.Join(names, ", "))
}
