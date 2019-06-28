package app

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/", main)
}

func readHtmlFile(fileName string) string{
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}
	return text
}


func main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, readHtmlFile("myTrainGuide.html"))
	//fmt.Fprintf(w, "hello, world!")
}
