package main

import (
	"./src/process"
	"strconv"
)

func main() {
	filename := "/workspace/go/crawer.txt"
	url := "http://movie.doubai.com/top250?start="+ strconv.Itoa(1)
	process.ParseResBody(url, filename)
}
