package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineNo := r.FormValue("line")
	lineInt, err := strconv.Atoi(lineNo)
	if err != nil {
		lineInt = -1
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var lines [][2]int
	if lineInt > 0 {
		lines = append(lines, [2]int{lineInt, lineInt})
	}

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)
	//_ = quick.Highlight(w, b.String(), "go", "html", "github")
}

//Reading lines from stack , parsing and decorate it to access the content
// of source code for panicing code
func makeLinks(stack string) string {
	splitLines := strings.Split(stack, "\n")

	for li, line := range splitLines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}

		//Parsing the line
		file := ""
		for i, ch := range line {
			if ch == ':' {
				file = line[1:i]
				break
			}
		}
		var lineNo = strings.Builder{}

		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}

			lineNo.WriteByte(line[i])
		}
		v := url.Values{}
		v.Set("path", file)
		v.Set("line", lineNo.String())
		splitLines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + file + ":" + lineNo.String() + "</a>" +
			line[len(file)+2+len(lineNo.String()):]
	}
	return strings.Join(splitLines, "\n")
}

//All the requests are routing through this handler and
//it recovers if panicing
//input : http.handler
//returns http.HandlerFunc
func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

//HandlerFunc that triggers panic
func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

//HandlerFunc that triggers panic after normal flow
func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

//Function that panics
func funcThatPanics() {
	panic("Oh no!")
	//println("ssss")
}

//HandlerFunc for hello word
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
