package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save () error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600)
}

func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func edit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func save(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w,r,"/test/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/test/", view)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/save/", save)
	http.ListenAndServe(":8000", nil)
}