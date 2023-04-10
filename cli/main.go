package main

import (
	"log"

	normalize "github.com/pajlada/lidl-normalize"
)

func main() {
	in := "🅱🅡🅘🅓🅖🅔"
	out, _ := normalize.Normalize(in)
	log.Println(out)
}
