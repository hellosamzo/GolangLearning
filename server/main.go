package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)
	fmt.Println("Server now running on localhost:8080")
	//fmt.Println(`Try running: curl -X POST -d '{"hello":"test123"}' http://localhost:8080/helloworld`)
	fmt.Println(`Try going to http://localhost:8080/ and resizing the window, pasting into the form fields and then submitting the completed form`)
	http.HandleFunc("/", dataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Data struct {
	EventType          string
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}

type helloWorldRequest struct {
	Hello string `json:"hello"`
	Testing string `json:"testing"`
}

// handler for http://localhost:8080/helloworld
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &helloWorldRequest{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}

	log.Printf("Request received %+v", req)

	w.WriteHeader(http.StatusOK)
}

// handler for http://localhost:8080/ 
func dataHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
		switch r.Method {
		// serves index.html as /
		case "GET":
	         http.ServeFile(w, r, "../client")
		// accepts POST requests from the js frontend
		case "POST":
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Unable to read body"))
					return
				}

				req := &Data{}

				if err = json.Unmarshal(body, req); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("FORM SUBMITTED / Unable to unmarshal JSON request"))
					return
				}

				log.Printf("Request received %+v", req)

				w.WriteHeader(http.StatusOK)
				
	    default:
	        fmt.Fprintf(w, "ONLY GET and POST methods are accepted at this URL.")
	    }
	}
