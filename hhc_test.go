package hhc

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	r := strings.NewReader(hhcContent)
	objects, err := Decode(r, "")
	if err != nil {
		t.Fatal(err)
	}
	var w bytes.Buffer
	e := json.NewEncoder(&w)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	err = e.Encode(objects)
	if err != nil {
		t.Fatal(err)
	}
	actual := w.String()
	if actual == jsonContent {
		t.Log("decode test passed")
	} else {
		t.Error("decode result not match")
	}
}

func TestEncode(t *testing.T) {
	var objects Objects
	err := json.NewDecoder(strings.NewReader(jsonContent)).Decode(&objects)
	if err != nil {
		t.Fatal(err)
	}
	var w bytes.Buffer
	err = Encode(&w, objects)
	if err != nil {
		t.Fatal(err)
	}
	actual := w.String()
	if actual == hhcContent {
		t.Log("encode test passed")
	} else {
		t.Error("encode result not match")
	}
}

const hhcContent = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML//EN">
<HTML>
<HEAD>
<meta name="GENERATOR" content="Microsoft&reg; HTML Help Workshop 4.1">
<!-- Sitemap 1.0 -->
</HEAD><BODY>
<UL>
	<LI> <OBJECT type="text/sitemap">
		<param name="Name" value="Introduction">
		<param name="Local" value="html\Intro.htm">
		</OBJECT>
	<LI> <OBJECT type="text/sitemap">
		<param name="Name" value="Getting Started">
		<param name="Local" value="html\GettingStarted.htm">
		</OBJECT>
	<UL>
		<LI> <OBJECT type="text/sitemap">
			<param name="Name" value="Foobar">
			<param name="Local" value="html\Foobar.htm">
			</OBJECT>
		<UL>
			<LI> <OBJECT type="text/sitemap">
				<param name="Name" value="Hello World">
				<param name="Local" value="html\HelloWorld.htm">
				</OBJECT>
			<LI> <OBJECT type="text/sitemap">
				<param name="Name" value="Lorem &amp; Ipsum">
				<param name="Local" value="html\LoremIpsum.htm">
				</OBJECT>
			<UL>
				<LI> <OBJECT type="text/sitemap">
					<param name="Name" value="Go">
					<param name="Local" value="html\Go.htm">
					</OBJECT>
				<LI> <OBJECT type="text/sitemap">
					<param name="Name" value="Golang">
					<param name="Local" value="html\Golang.htm">
					</OBJECT>
			</UL>
		</UL>
	</UL>
</UL>
</BODY></HTML>
`

const jsonContent = `[
  {
    "Type": "text/sitemap",
    "Params": {
      "Local": "html\\Intro.htm",
      "Name": "Introduction"
    }
  },
  {
    "Type": "text/sitemap",
    "Params": {
      "Local": "html\\GettingStarted.htm",
      "Name": "Getting Started"
    },
    "Objects": [
      {
        "Type": "text/sitemap",
        "Params": {
          "Local": "html\\Foobar.htm",
          "Name": "Foobar"
        },
        "Objects": [
          {
            "Type": "text/sitemap",
            "Params": {
              "Local": "html\\HelloWorld.htm",
              "Name": "Hello World"
            }
          },
          {
            "Type": "text/sitemap",
            "Params": {
              "Local": "html\\LoremIpsum.htm",
              "Name": "Lorem & Ipsum"
            },
            "Objects": [
              {
                "Type": "text/sitemap",
                "Params": {
                  "Local": "html\\Go.htm",
                  "Name": "Go"
                }
              },
              {
                "Type": "text/sitemap",
                "Params": {
                  "Local": "html\\Golang.htm",
                  "Name": "Golang"
                }
              }
            ]
          }
        ]
      }
    ]
  }
]
`
