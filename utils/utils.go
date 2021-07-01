package utils

import (
	"fmt"
)

func Assert1(guard bool, text string) {
	if !guard {
		panic(fmt.Errorf("Error: %s", text))
	}
}