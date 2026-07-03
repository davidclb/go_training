package exos

import (
	"strings"
	"testing"
)

func TestCountImageWordinHTML(t *testing.T) {

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

	want := [2]int{14, 2}
	wdcount, imgcount := CountImageWordinHTML(strings.NewReader(raw))

	if want[0] != wdcount || want[1] != imgcount {
		t.Errorf("We wanted %d and %d, we got %d and %d", want[0], want[1], wdcount, imgcount)
	}

}
