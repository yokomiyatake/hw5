package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", getWords)
	http.HandleFunc("/test", showPata)
}

func getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

    fmt.Fprintf(w, "<form>")
    fmt.Fprintf(w, "word1:<br>")
    fmt.Fprintf(w, "<input type='text' name='word1' value='パトカー'><br>")
    fmt.Fprintf(w, "word2:<br>")
    fmt.Fprintf(w, "<input type='text' name='word2' value='タクシー'><br><br>")
    fmt.Fprintf(w, "<input type='submit' value='合体'>")
    fmt.Fprintf(w, "</form>")
}

func showPata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "<p>ぱたとくかしーー</p>")
}

