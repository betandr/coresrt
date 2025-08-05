package main

import (
	"coresrt/receiver"
)

func main() {
	receiver.Start(9999, "0.0.0.0")
}
