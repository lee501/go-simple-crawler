package process

import (
	"../crawler"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseResBody(url string, filename string) {
	outfile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer outfile.Close()
	body := crawler.Fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	htmlRe := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span>(.*?)</span>`)
	aHrefRe := regexp.MustCompile(`<a href=>`)
	items := htmlRe.FindAllStringSubmatch(body, -1)

	outWriter := bufio.NewWriter(outfile)
	for _, item := range items {
		outWriter.WriteString(titleRe.FindStringSubmatch(item[1])[1])
		outWriter.WriteString(": ")
		outWriter.WriteString(aHrefRe.FindStringSubmatch(item[1])[1])
		outWriter.WriteString("\n")
	}
}
