package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func main() {
	arg1 := os.Args[1]
	sample := lo.Map([]string{"hoge", "fuga"}, func(s string, i int) string {
		return fmt.Sprintf("%d: %s", i, s)
	})
	fmt.Printf("Hello, %s!, %s\n", arg1, strings.Join(sample, ", "))
}
