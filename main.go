package main

import (
	"os"
)


func main()  {
	var FileLink = os.Args[1]
	if len(FileLink) == 0 {
		os.Exit(1)
	}

}