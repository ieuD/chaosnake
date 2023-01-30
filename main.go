package main

import (
	"github.com/ieud/chaosnake/internal"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	internal.Execute()
}
