package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	t "neurokin/types"
	"neurokin/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	// Resolve the absolute path for homepage_sections.json
	sectionsPath, err := filepath.Abs("homepage_sections.json")
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

	// Resolve the absolute path for homepage_tiles.json
	tilesPath, err := filepath.Abs("homepage_tiles.json")
	if err != nil {
		http.Error(w, "Error resolving path for homepage tiles", http.StatusInternalServerError)
		return err
	}

	// Load homepage tiles from JSON
	tilesData, err := ioutil.ReadFile(tilesPath)
	if err != nil {
		http.Error(w, "Error loading homepage tiles", http.StatusInternalServerError)
		return err
	}

	var tiles []t.Tile
	if err := json.Unmarshal(tilesData, &tiles); err != nil {
		http.Error(w, "Error parsing homepage tiles JSON", http.StatusInternalServerError)
		return err
	}

	// Pass the loaded data to the home.Index component
	return Render(w, r, home.Index(sections, tiles))
}

