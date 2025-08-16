package main

import (
	"github.com/openrobosouls/rm-proto/client"
	"github.com/wintbiit/tsf4g/tdrcom"
)

func main() {
	message := client.NewFrame()
	tdrcom.Marshal(message, 0)
}
