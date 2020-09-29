package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

const html = `
<!doctype html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>scribe</title>
		<style>
			div#display {
				border: 1px solid #999;
				padding: 5px;
				margin-bottom: 5px;
				height: 93vh;
			}
		</style>
	</head>
	<body>
		<div contenteditable="true" id="display"></div>
		<button id="save">Save Content</button>
	</body>
	<script>
		document.getElementById("save").addEventListener("click", function() {
			alert("hello, world");
		});
	</script>
</html>
`

func main() {
	ui, err := lorca.New("data:text/html,"+url.PathEscape(html), "", 600, 800)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
