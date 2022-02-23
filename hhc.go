package hhc

import (
	"bytes"
	"io"
	"sort"

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

const header = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML//EN">
<HTML>
<HEAD>
<meta name="GENERATOR" content="Microsoft&reg; HTML Help Workshop 4.1">
<!-- Sitemap 1.0 -->
</HEAD><BODY>
`

const footer = `</BODY></HTML>
`

// Encode converts a Objects tree to HHC document.
func Encode(w io.Writer, objects Objects) error {
	if len(objects) == 0 {
		return nil
	}
	w.Write([]byte(header))
	if len(objects) == 1 && objects[0].Type == TYPE_TEXT_SITE_PROPERTIES {
		encodeObject(w, objects[0], 1, 0)
		encode(w, objects[0].Objects, 0)
	} else {
		encode(w, objects, 0)
	}
	_, err := w.Write([]byte(footer))
	return err
}

func encode(w io.Writer, objects Objects, lvl int) {
	if len(objects) == 0 {
		return
	}
	w.Write(bytes.Repeat([]byte{'\t'}, lvl))
	w.Write([]byte("<UL>\n"))
	for _, object := range objects {
		w.Write(bytes.Repeat([]byte{'\t'}, lvl+1))
		w.Write([]byte("<LI> "))
		encodeObject(w, object, lvl+2, lvl+2)
		encode(w, object.Objects, lvl+1)
	}
	w.Write(bytes.Repeat([]byte{'\t'}, lvl))
	w.Write([]byte("</UL>\n"))
}

func encodeObject(w io.Writer, object Object, lvl, endLvl int) {
	w.Write([]byte("<OBJECT type=\""))
	w.Write([]byte(object.Type))
	w.Write([]byte("\">\n"))
	keys := keysOf(object.Params)
	for i := len(keys) - 1; i > -1; i-- {
		w.Write(bytes.Repeat([]byte{'\t'}, lvl))
		w.Write([]byte("<param name=\""))
		w.Write([]byte(html.EscapeString(keys[i])))
		w.Write([]byte("\" value=\""))
		w.Write([]byte(html.EscapeString(object.Params[keys[i]])))
		w.Write([]byte("\">\n"))
	}
	w.Write(bytes.Repeat([]byte{'\t'}, endLvl))
	w.Write([]byte("</OBJECT>\n"))
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

func keysOf(params Params) (keys []string) {
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}
