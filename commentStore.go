package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func getComments() ([]Comment){
	commentData, err := ioutil.ReadFile("comments.json")
	if err != nil {
		fmt.Printf("Cannot read json file")
		return nil
	}
	var comments []Comment
	if err := json.Unmarshal(commentData, &comments); err != nil {
		fmt.Printf("Cannot read json file")
		return nil
	}
	return comments
	
}

func commentRegister(actionType string, newComment Comment, w http.ResponseWriter, r *http.Request){
	switch actionType{
		case "TODO_CREATE":
			//change json
			var comments []Comment = getComments()
			comments = append(comments, newComment)
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
		default:
	}
	//fire change event
	commentEmitter.Emit("commentChange", w, r)
}
