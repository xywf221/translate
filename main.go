package main

import (
	"fmt"
	"os"
)

//现阶段只能中英互译以后

func main() {
	if len(os.Args) == 1 {
		fmt.Println("parameter errors")
		os.Exit(0)
	}
	src := os.Args[1]
	from := langDetect(src)
	var to = "en"
	if from != "zh" {
		to = "zh"
	}
	dist := translate(src, from, to)
	fmt.Printf("src : %s \r\nlange : %s\nto : %s \ndist : %s\n", src, from, to, dist)
}
