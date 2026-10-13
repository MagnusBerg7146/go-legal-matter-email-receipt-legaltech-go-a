package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	client := newEmailClient()
	http.HandleFunc("/matters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var matter matterIntake
		if err := json.NewDecoder(r.Body).Decode(&matter); err != nil {
			http.Error(w, "invalid matter", http.StatusBadRequest)
			return
		}
		ids, err := sendMatterEmails(client, matter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"matter_id": matter.MatterID, "message_ids": ids})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
