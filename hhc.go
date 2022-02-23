package hhc

import (
	"io"

	"golang.org/x/net/html"
)

const (
	TYPE_TEXT_SITE_PROPERTIES = "text/site properties"
	TYPE_TEXT_SITEMAP         = "text/sitemap"
)

type (
	Objects []Object

	Object struct {
		Type    string
		Params  Params
		Objects Objects `json:",omitempty"`
	}

	Params = map[string]string
)

// Decode parses a HHC document and put all objects into a Objects tree. If
// objectType is not empty, only objects having that type are included.
func Decode(r io.Reader, objectType string) (Objects, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	objects := Objects{}
	findObjects(&objects, doc, objectType)
	if len(objects) == 0 {
		return nil, nil
	}
	return objects, nil
}

func findObjects(objects *Objects, n *html.Node, objectType string) {
	if n.Type == html.ElementNode {
		if n.Data == "ul" {
			if len(*objects) > 0 {
				objects = &((*objects)[len(*objects)-1].Objects)
			}
		}
		if n.Data == "object" {
			var oType string
			for _, a := range n.Attr {
				if a.Key == "type" {
					oType = a.Val
				}
			}
			if objectType != "" {
				if oType != objectType {
					return
				}
			}
			params := Params{}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findParam(&params, c)
			}
			if len(params) == 0 {
				return
			}
			*objects = append(*objects, Object{
				Type:   oType,
				Params: params,
			})
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findObjects(objects, c, objectType)
	}
}

func findParam(params *Params, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "param" {
		var key, value string
		for _, a := range n.Attr {
			switch a.Key {
			case "name":
				key = a.Val
			case "value":
				value = a.Val
			}
		}
		if key != "" && value != "" {
			(*params)[key] = value
		}
	}
}