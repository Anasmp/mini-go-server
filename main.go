package main

import (

	"net/http"
	"encoding/json"
)

type User struct {
  Name    string
  Email   string
}

func main(){
	  
	mux:= http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		user := User{r.FormValue("username"),r.FormValue("email")}

		 js, err := json.Marshal(user)
		  if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		  }

		    w.Header().Set("Content-Type", "application/json")
  			w.Write(js)
	})

	http.ListenAndServe(":8080", mux)
}
