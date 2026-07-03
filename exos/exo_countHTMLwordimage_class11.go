package exos

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

//Count words and images in a HTML file

//Brouillon conception
//scan html ligne par ligne , quand la premiere pos = <h +nombre > or <p> recuperer la ligne dans un slice , faire len(slice) a la fin
//Pour les images <img et j'augmente le compteur

func CountImageWordinHTML(r io.Reader) (int, int) {

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
