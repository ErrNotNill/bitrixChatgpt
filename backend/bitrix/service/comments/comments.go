package comments

import (
	"fmt"
	"log"
	"net/http"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	comments, err := GetCommentsDealMock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(comments))
	// Directly write the mock JSON as the response
	_, err = w.Write(comments) // comments already contains the JSON
	if err != nil {
		log.Println("Error writing the mock response:", err)
	}
}

func GetCommentsDealMock() ([]byte, error) {
	mockJSON := `{
		"result": [
			{
				"ID": "515",
				"COMMENT": "Text commentary..."
			}
		],
		"total": 1,
		"time": {
			"start": 1711892133.7379389,
			"finish": 1711892133.7726719,
			"duration": 0.034733057022094727,
			"processing": 0.012963056564331055,
			"date_start": "2024-03-31T16:35:33+03:00",
			"date_finish": "2024-03-31T16:35:33+03:00",
			"operating_reset_at": 1711892733,
			"operating": 0
		}
	}`
	return []byte(mockJSON), nil
}
