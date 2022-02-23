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
	if actual == jsonContentSP {
		t.Log("decode test passed")
	} else {
		// os.WriteFile("a", []byte(actual), 0644)
		// os.WriteFile("b", []byte(jsonContentSP), 0644)
		t.Error("decode result not match")
	}
}

func TestDecode2(t *testing.T) {
	r := strings.NewReader(hhcContent)
	objects, err := Decode(r, TYPE_TEXT_SITEMAP)
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
		t.Log("decode 2 test passed")
	} else {
		// os.WriteFile("c", []byte(actual), 0644)
		// os.WriteFile("d", []byte(jsonContent), 0644)
		t.Error("decode 2 result not match")
	}
}

func TestEncode(t *testing.T) {
	var objects Objects
	err := json.NewDecoder(strings.NewReader(jsonContentSP)).Decode(&objects)
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
		// os.WriteFile("e", []byte(actual), 0644)
		// os.WriteFile("f", []byte(hhcContent), 0644)
		t.Error("encode result not match")
	}
}

const hhcContent = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML//EN">
<HTML>
<HEAD>
<meta name="GENERATOR" content="Microsoft&reg; HTML Help Workshop 4.1">
<!-- Sitemap 1.0 -->
</HEAD><BODY>
<OBJECT type="text/site properties">
	<param name="Auto Generated" value="Yes">
</OBJECT>
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

const jsonContentSP = `[
  {
    "Type": "text/site properties",
    "Params": {
      "Auto Generated": "Yes"
    },
    "Objects": [
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
  }
]
`
