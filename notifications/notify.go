package main

import (
	"github.com/gen2brain/beeep"
)

func Notifier() {
	err := beeep.Notify("Title", "Message Body", "")
	if err != nil {
		panic(err)
	}
}

func main() {
	Notifier()
}
