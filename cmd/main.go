package main

import (
	"fmt"

	"github.com/0xruffy/gotint"
)

func main() {
	tint := gotint.NewTint().Bold().AddUnderline().Red().OnWhite()

	fmt.Println(gotint.NewTintString("Hello").WithTint(tint))
}
