package authorize

import (
	"bitrix_app/backend/bitrix/authorize/auth"
	"bitrix_app/backend/bitrix/service/events"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var GlobalAuthId string

func ConnectionBitrixLocalApp(w http.ResponseWriter, r *http.Request) {
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	log.Println("resp_at_first:", string(bs))
	defer r.Body.Close()

	auth := ParseValues(w, bs) //todo here we must to add this data in dbase?
	fmt.Printf("auth.AuthID : %s, auth.MemberID: %s", auth.AuthID, auth.MemberID)

	//w.Write([]byte(auth.AuthID))
	redirectURL := "https://b24app.rwp2.com/"

	// Use http.Redirect to redirect the client
	// The http.StatusFound status code is commonly used for redirects
	http.Redirect(w, r, redirectURL, http.StatusFound)

	fmt.Println("redirect is done...")
	GlobalAuthId = auth.AuthID

	events.OnCrmDealAddEventRegistration(auth.AuthID) //todo return this method

}

func ParseValues(w http.ResponseWriter, bs []byte) auth.Request {
	values, err := url.ParseQuery(string(bs))
	if err != nil {
		log.Println("error parsing query:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	authExpires, err := strconv.Atoi(values.Get("AUTH_EXPIRES"))
	if err != nil {
		log.Println("error converting AUTH_EXPIRES to int:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	auth := auth.Request{
		AuthID:           values.Get("AUTH_ID"),
		AuthExpires:      authExpires,
		RefreshID:        values.Get("REFRESH_ID"),
		MemberID:         values.Get("member_id"),
		Status:           values.Get("status"),
		Placement:        values.Get("PLACEMENT"),
		PlacementOptions: values.Get("PLACEMENT_OPTIONS"),
	}
	return auth
}
