package main

import(
	"net/http"
)

func newCommentView(w http.ResponseWriter, r *http.Request){
	//in react, would call setState, here we refresh page.
	//Dunno if this is right
	http.Redirect(w, r, "/", http.StatusFound)
}
