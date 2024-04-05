package chatgpt

import (
	"io"
	"log"
	"net/http"
)

func RequestFromVue(w http.ResponseWriter, r *http.Request) {
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	log.Println("req_at_Vue_for_chatgpt :", string(bs))

	defer r.Body.Close()
}
