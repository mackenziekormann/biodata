package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mackenziekormann/biodata/pkg/alphafold"
)

func main() {
	http.HandleFunc("/fetchProteinData", func(w http.ResponseWriter, r *http.Request) {
		// Assume the protein ID is passed as a query parameter
		proteinID := r.URL.Query().Get("proteinID")
		if proteinID == "" {
			http.Error(w, "Protein ID is required", http.StatusBadRequest)
			return
		}

		data, err := alphafold.FetchAlphaFoldData(proteinID)
		if err != nil {
			log.Println("Error fetching protein data:", err)
			http.Error(w, "Error fetching protein data", http.StatusInternalServerError)
			return
		}

		// Respond with the fetched data
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
