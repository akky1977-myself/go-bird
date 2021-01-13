package handler

import (
	"fmt"
	"net/http"
)

// Bird is struct
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// Birds is slide
var Birds []Bird

// CreateBirdHandler return nothing
func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {

	bird := Bird{}
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	Birds = append(Birds, bird)
	http.Redirect(w, r, "/assets_path/", http.StatusFound)
}
