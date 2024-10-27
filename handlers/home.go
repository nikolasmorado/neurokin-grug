package handlers

import (
	"encoding/json"
	"net/http"
	t "neurokin/types"
	"neurokin/views/home"
	"os"
	"path/filepath"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	sectionsPath, err := filepath.Abs("homepage_sections.json")
	if err != nil {
		http.Error(w, "Error resolving path for homepage sections", http.StatusInternalServerError)
		return err
	}

	sectionsData, err := os.ReadFile(sectionsPath)
	if err != nil {
		http.Error(w, "Error loading homepage sections", http.StatusInternalServerError)
		return err
	}

	var sections []t.Section
	if err := json.Unmarshal(sectionsData, &sections); err != nil {
		http.Error(w, "Error parsing homepage sections JSON", http.StatusInternalServerError)
		return err
	}

	tilesPath, err := filepath.Abs("homepage_tiles.json")
	if err != nil {
		http.Error(w, "Error resolving path for homepage tiles", http.StatusInternalServerError)
		return err
	}

	tilesData, err := os.ReadFile(tilesPath)
	if err != nil {
		http.Error(w, "Error loading homepage tiles", http.StatusInternalServerError)
		return err
	}

	var tiles []t.Tile
	if err := json.Unmarshal(tilesData, &tiles); err != nil {
		http.Error(w, "Error parsing homepage tiles JSON", http.StatusInternalServerError)
		return err
	}

	return Render(w, r, home.Index(sections, tiles))
}

