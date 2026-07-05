package exos

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

//Count words and images in a HTML file

//Brouillon conception
//scan html ligne par ligne , quand la premiere pos = <h +nombre > or <p> recuperer la ligne dans un slice , faire len(slice) a la fin
//Pour les images <img et j'augmente le compteur

//finalement supprime toutes balises puis je compte les mots
//je compte toute les occurences de <img >

func CountImageWordinHTML(r io.Reader) (int, int) {

	scan := bufio.NewScanner(r)

	var wdcount, imgcount int

	tagRe := regexp.MustCompile(`<[^>]*>`)
	tagImg := regexp.MustCompile(`<img[^>]*>`)
	for scan.Scan() {

		s := scan.Text()

		v := tagRe.ReplaceAllString(s, "")
		for _, wd := range strings.Fields(v) {
			fmt.Println(wd)
			wdcount++
		}
		imgcount += len(tagImg.FindAllString(s, -1))

	}
	return wdcount, imgcount

}
