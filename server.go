package main

import (
	"html/template"
	"net/http"
	//"io"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

import (
  "github.com/chuckpreslar/emission"
)

type Comment struct{
	Author string `json:"author"`
	Text   string `json:"text"`
}

//global emitters
var dispatchEmitter = emission.NewEmitter()
//in actual Flux each store has own emitter
var commentEmitter = emission.NewEmitter()


var templates = template.Must(template.ParseFiles("blog.html"))
//save into json file, reload for now but in final will call dispatcher
func newComment(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("text")
	author := r.FormValue("author")
	handleViewAction("TODO_CREATE", Comment{Author: author, Text:body}, w, r)
	/*var comments []Comment = getComments()
	comments = append(comments, Comment{Author: author, Text: body})
	commentData, err := json.MarshalIndent(comments, "", "    ")
		if err != nil {
			fmt.Printf("Unable to marshal comments to json")
			return
		}
	//look up proper filemode, not just 0644
	err = ioutil.WriteFile("comments.json", commentData, 0644)
		if err != nil {
			fmt.Printf("Unable to write comments to data file")
			return
		}
	http.Redirect(w, r, "/", http.StatusFound)*/
}

//json is used instead of database for now
func renderComments(w http.ResponseWriter, r *http.Request) {
	//read and decode json
	commentData, err := ioutil.ReadFile("comments.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot read json file"), http.StatusInternalServerError)
		return
	}
	var comments []Comment
		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Cannot unmarshal json"), http.StatusInternalServerError)
			return
		}
  //name of blog.html is hardcoded for now
	err = templates.ExecuteTemplate(w, "blog.html", &comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	//add Listeners, should be done in a better location?
	//callbacks should be registered with their respective stores and views
	//not tested for concurrency
	dispatchEmitter.On("newComment", commentRegister)
	commentEmitter.On("commentChange", newCommentView)
	http.HandleFunc("/", renderComments)
	http.HandleFunc("/new/", newComment)

	http.ListenAndServe(":8080", nil)
}
