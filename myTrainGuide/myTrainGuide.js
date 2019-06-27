
var linesToSt = new Array();

linesToSt.set(["山手線"], ["品川", "大崎", "五反田", "目黒", "恵比寿", "渋谷", "原宿", "代々木", "新宿", "新大久保", "高田馬場", "目白", "池袋", "大塚", "巣鴨", "駒込", "田端", "西日暮里", "日暮里", "鶯谷", "上野", "御徒町", "秋葉原", "神田", "東京", "有楽町", "新橋", "浜松町", "田町", "品川"]);
linesToSt.set(["東横線"], ["横浜", "反町", "東白楽", "白楽", "妙蓮寺", "菊名", "大倉山", "綱島", "日吉", "元住吉", "武蔵小杉", "新丸子", "多摩川", "田園調布", "自由が丘", "都立大学", "学芸大学", "祐天寺", "中目黒", "代官山", "渋谷"]);
linesToSt.set(["目黒線"], ["日吉", "元住吉", "武蔵小杉", "新丸子", "多摩川", "田園調布", "奥沢", "大岡山", "洗足", "西小山", "武蔵小山", "不動前", "目黒"]);

function makeAdjList(lineToSt) {
    var adjList = new Array();

    for (var i = 0; i < linesToSt.size; i++) {
        var stations = linesToSt[i];

        for (var j = 0; j < stations.size; j++) {
            var from = stations[j];

            for (var k = 0; k < stations.size; k++) {
                var to = stations[k];
                if (from != to) {
                    if !(from in adjList) {
                        adjList.push(from);
                    }
                    adjList[from].push(to);
                }
            }
        }
    }
    return adjList
}


function bfs(adjList, from, to) {
    var queue = new Array();
    queue.push(from);

    var visited = new Array();

    while queue.length != 0 {
        var route = queue[0];
        var currentStation = route[route.length - 1];
        visited[currentStation] = true;

        if currentStation == to {
            return route;
        }

        if currentStation in adjList {
            var nextStations = adjList[currentStation]
            for (i = 0; i < nextStations.length; i++) {
                var next = nextStations[i];
                if !visited[next] {
                    var newRoute = route.push(next);
                    queue = queue.push(newRoute);
                }
            }
        }
        queue = queue.pop();
    }
    return ["Not Found"]
}

function onButtonClick() {
    showField = document.getElementById('output');
    from = document.forms.inputForm.id_from.value;
    to = document.forms.inputForm.id_to.value;

    var adjList = makeAdjList(linesToSt);
    console.log(adjList);
    var route = bfs(adjList, from, to)

    showField.innerText = route;
}