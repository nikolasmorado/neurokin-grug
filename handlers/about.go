package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	t "neurokin/types"
	"neurokin/views/about"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) error {
	sectionsPath, err := filepath.Abs("about_sections.json")
	if err != nil {
		http.Error(w, "Error resolving path for homepage sections", http.StatusInternalServerError)
		return err
	}

	// Load homepage sections from JSON
	sectionsData, err := ioutil.ReadFile(sectionsPath)
	if err != nil {
		http.Error(w, "Error loading homepage sections", http.StatusInternalServerError)
		return err
	}

	var sections []t.Section
	if err := json.Unmarshal(sectionsData, &sections); err != nil {
		http.Error(w, "Error parsing homepage sections JSON", http.StatusInternalServerError)
		return err
	}
	return Render(w, r, about.Index(sections))
}

