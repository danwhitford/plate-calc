package main

import (
	"fmt"
	"syscall/js"

	"github.com/danwhitford/plate-counter/plates"
)

func main() {
    fmt.Println("Go Web Assembly")
    js.Global().Set("GetPlatesForBarJS", plates.GetPlatesForBarJS())
	<-make(chan bool)
}
