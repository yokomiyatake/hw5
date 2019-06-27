package app

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func init() {
	http.HandleFunc("/", getWords)
}

func makeAdjacencyList() map[string][]string {

	adjList := make(map[int][]int)

	var fp *os.File
	var err error

	// Open file
	fp, err = os.Open("stations.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// Read a line
		line := scanner.Text()

		// Split A\tB to [A, B]
		array := strings.Split(line, "\t")

		// Convert each number to int
		from, _ := strconv.Atoi(array[0])

		adjList[from] = append(adjList[from], to)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return adjList
}

func bfs(adjList map[int][]int, start int, goal int) []int {

	queue := make([][]int, 0)
	firstRoute := []int{start}
	queue = append(queue, firstRoute)

	var visited[nodeNum] bool

	for !(len(queue) == 0) {
		route := queue[0]
		now := route[len(route) - 1]
		visited[now] = true

		// For debugging
		//fmt.Println(route)
		//fmt.Println(now)

		if now == goal {
			fmt.Println("Found!")
			return route
		}

		// If next node from 'now' exists
		_, exist := adjList[now]
		if exist {
			// Make next routes and add them to queue
			for i := 0; i < len(adjList[now]); i++ {
				nextNode := adjList[now][i]
				if !visited[nextNode] {
					newRoute := append(route, nextNode)
					queue = append(queue, newRoute)
				}
			}
		}
		queue = queue[1:]
	}
	fmt.Println("Not Found")
	notFound := []int{}
	return notFound
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://fantasy-transit.appspot.com/net?format=json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "HTTP GET returned status %v", resp.Status)
}
*/

func getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	stations := map[string][]string{"山手線":{"品川", "大崎", "五反田", "目黒", "恵比寿", "渋谷"}}

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
	fmt.Fprintf(w, "		From:<br>")
	fmt.Fprintf(w, "		<input type='text' name='from' id='from' value='茗荷谷' /><br>")
	fmt.Fprintf(w, "		To:<br>")
	fmt.Fprintf(w, "		<input type='text' name='to' id='to' value='六本木' /><br><br>")
	fmt.Fprintf(w, "		<input type='button' value='search' onclick='onButtonClick();' />")
	fmt.Fprintf(w, "</form>")
	fmt.Fprintf(w, "<div id='output'></div>")
}
