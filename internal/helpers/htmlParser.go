package helpers

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"strings"
)

func RenderNode(n *html.Node) (string, error) {
	var buf bytes.Buffer  // is a variable-sized buffer of bytes with read and write methods
	w := io.Writer(&buf)  // creates writer
	err := html.Render(w, n) // returns a flush that writes any buffered data to the underlying stream
	if err != nil{
		return "", err
	}
	return buf.String(), nil
}

func ParseNodes(htm string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		logrus.Fatal(err)
	}

	str, err := RenderNode(doc)
	if err != nil{
		return "", err
	}
	return str, nil
}