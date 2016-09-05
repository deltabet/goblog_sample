package main

import (
	"net/http"
)

//in Flux data type is hard coded in and passed in, dunno if this is what we want
func handleViewAction(action string, newComment Comment, w http.ResponseWriter, 
	r *http.Request){
	dispatchEmitter.Emit("newComment", action, newComment, w, r) 
}

