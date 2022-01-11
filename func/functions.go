package functions

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// var Templates, TmplError = template.ParseFiles("templ/mainPage.html", "templ/error400.html", "templ/error404.html", "templ/error405.html", "templ/error500.html")  ИЛИ

// ИЛИ

var Templates, TmplError = template.ParseGlob("templ/*.html")

//  Так ненадо парсить
// var tmplHomePage, _ = template.ParseFiles("templ/mainPage.html")
// var tmplErr400, _ = template.ParseFiles("templ/error400.html")
// var tmplErr404, _ = template.ParseFiles("templ/error404.html")
// var tmplErr500, _ = template.ParseFiles("templ/error500.html")

func Post(w http.ResponseWriter, r *http.Request) {
	font := r.FormValue("Fonts") // the variable font stores the fonts
	text := r.FormValue("Text")  // the variable text stores the entered text
	button := r.FormValue("Submit")
	if CheckValid(text) && CheckFonts(font) && CheckValue(text) /*&& CheckButton(button)*/ {
		split := strings.Split(text, "\r\n")
		data, err := os.ReadFile(font)
		if err != nil {
			w.WriteHeader(500)
			// tmplErr500.Execute(w, nil)
			Templates.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			return
		}
		out := Separator(data)
		ascWeb := Print(split, out)
		if !(len(data) != 6623 || len(data) != 7463 || len(data) != 4703) || err != nil {
			w.WriteHeader(500)
			// tmplErr500.Execute(w, nil)
			Templates.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			return
		}

		if button == "submit" {
			// out := Separator(data)
			// ascWeb := Print(split, out)
			// tmplHomePage.Execute(w, ascWeb)
			Templates.ExecuteTemplate(w, "mainPage.html", ascWeb)
			return
		} else if button == "download" {
			b := []byte(ascWeb)
			file := strings.NewReader(ascWeb)
			FileSize := strconv.FormatInt(file.Size(), 10)
			w.Header().Set("Content-Disposition:", "attachment; filename=output.txt")
			w.Header().Set("Content-Type", "plain/text")
			w.Header().Set("Content-Length", FileSize)
			// io.Copy(w, file)
			w.Write(b)
			return
		} else {
			w.WriteHeader(400)
			Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		}
	}
	w.WriteHeader(400)
	// tmplErr400.Execute(w, nil)
	Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
	w.WriteHeader(400)
	// http.Error(w, http.StatusText(400), http.StatusBadRequest)
	return
}

func Separator(data []byte) []string {
	out := []string{}
	h := 1
	for i := 1; i < len(data); i++ {
		if data[i] == '\n' {
			out = append(out, string(data[h:i]))
			h = i + 1
		}
		if data[i] == '\n' && data[i+1] == '\n' {
			i = i + 2
			h = i
		}
	}
	out = append(out, string(data[h:len(data)-1]))

	return out

}

func Print(s []string, data []string) string {
	out := ""
	for j := 0; j < len(s); j++ {
		for k := 0; k <= 7; k++ {
			if k > 0 && len(s[j]) != 0 {

			}
			for _, el := range s[j] {
				if el >= 32 && el <= 126 {
					out += string((data[((el-32)*8)+rune(k)]))
				} else {
					return ""
				}
			}
			out += "\n"
		}

	}
	return out
}

func CheckValid(s string) bool {
	for _, el := range s {
		if el >= 32 && el <= 126 || el == 10 || el == 13 {
			continue
		}
		return false
	}
	return true
}

func CheckFonts(s string) bool {
	if !(s == "standard.txt" || s == "shadow.txt" || s == "thinkertoy.txt") {
		return false
	}
	return true
}

func CheckValue(s string) bool {
	if s == "" {
		return false
	}
	return true
}

func CheckButton(s string) bool {
	if s == "submit" || s == "download" {
		return true
	}
	return false
}

func Fun(s string) {

	data, err := ioutil.ReadFile(s + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
