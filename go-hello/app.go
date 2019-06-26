package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", getWords)
}

func getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "<script type='text/javascript' language='javascript'>")
	fmt.Fprintf(w, "function onButtonClick() {")
	fmt.Fprintf(w, "		showField = document.getElementById('output');")
	fmt.Fprintf(w, "		text1 = document.forms.inputForm.textBox1.value;")
	fmt.Fprintf(w, "		text2 = document.forms.inputForm.textBox2.value;")
	fmt.Fprintf(w, "		var pata = '';")
	fmt.Fprintf(w, "		var l1 = text1.length, l2 = text2.length;")
	fmt.Fprintf(w, "		for (var i = 0; i < Math.max(l1, l2); i++) {")
	fmt.Fprintf(w, "			if (i < l1) pata += text1.charAt(i);")
	fmt.Fprintf(w, "			if (i < l2) pata += text2.charAt(i);")
	fmt.Fprintf(w, "		}")
	fmt.Fprintf(w, "		showField.innerText = pata;")
	fmt.Fprintf(w, "}")
	fmt.Fprintf(w, "</script>")
	fmt.Fprintf(w, "</head>")

    fmt.Fprintf(w, "<form name='form1' id='inputForm' action=''>")
    fmt.Fprintf(w, "word1:<br>")
    fmt.Fprintf(w, "<input type='text' name='word1' id='textBox1' value='パトカー' /><br>")
    fmt.Fprintf(w, "word2:<br>")
    fmt.Fprintf(w, "<input type='text' name='word2' id='textBox2' value='タクシー' /><br><br>")
    fmt.Fprintf(w, "<input type='button' value='合体' onclick='onButtonClick();' />")
    fmt.Fprintf(w, "</form>")
	fmt.Fprintf(w, "<div id='output'></div>")
}


