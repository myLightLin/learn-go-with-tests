package main

import (
	"os"
	"time"

	"github.com/learn-go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
