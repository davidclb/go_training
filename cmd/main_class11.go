package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

//Count words and images in a HTML file

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

//Brouillon conception
//scan html ligne par ligne , quand la premiere pos = <h +nombre > or <p> recuperer la ligne dans un slice , faire len(slice) a la fin
//Pour les images <img et j'augmente le compteur

func countImageWordinHTML(r io.Reader) (int, int) {

	scan := bufio.NewScanner(r)

	var wdcount, imgcount int

	for scan.Scan() {

		s := scan.Text()
		regwd := regexp.MustCompile(`<h[1-6]>`)
		regimg := regexp.MustCompile(`<img`)

		if regwd.MatchString(strings.Fields(s)[0]) || strings.Fields(s)[0] == "<p>" {
			wdcount += len(strings.Fields(s))
		}

		if regimg.MatchString(strings.Fields(s)[0]) {
			imgcount += len(strings.Fields(s))
		}

	}

	return wdcount, imgcount

}

func main_class11() {

}
