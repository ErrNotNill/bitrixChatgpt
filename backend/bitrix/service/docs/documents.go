package docs

import (
	"bitrix_app/backend/bitrix/authorize"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetDocsByDealMock() ([]byte, error) {
	// Instead of making an actual HTTP request, we'll return a mock response

	// Mock JSON response string
	mockJSON := `{"result":{"documents":[{"id":"1","title":"Счет (Россия) 1","number":"1","templateId":"3","fileId":"39","imageId":"41","pdfId":"43","createTime":"2024-03-31T14:43:25+03:00","updateTime":"2024-03-31T14:43:25+03:00","values":{"_creationMethod":"public"},"createdBy":"1","updatedBy":null,"downloadUrl":"https://b24-9f7fvg.bitrix24.ru/bitrix/services/main/ajax.php?action=crm.documentgenerator.document.download&SITE_ID=s1&id=1","pdfUrl":"https://b24-9f7fvg.bitrix24.ru/bitrix/services/main/ajax.php?action=crm.documentgenerator.document.getPdf&SITE_ID=s1&id=1","imageUrl":"https://b24-9f7fvg.bitrix24.ru/bitrix/services/main/ajax.php?action=crm.documentgenerator.document.getImage&SITE_ID=s1&id=1","stampsEnabled":false,"entityId":"79","entityTypeId":"2","downloadUrlMachine":"https://b24-9f7fvg.bitrix24.ru/rest/crm.documentgenerator.document.download.json?auth=c45a0966006c4aa0006c188c00000001000007874cfd89ae57b5f5496fc10b7b077e1b&token=crm|action=crm.documentgenerator.document.download&SITE_ID=s1&id=1&_=MJ8ZB9Y5ZQ7cPWWhcu8Qs7b8pfuw0bM","pdfUrlMachine":"https://b24-9f7fvg.bitrix24.ru/rest/crm.documentgenerator.document.getPdf.json?auth=c45a0966006c4aa0006c188c00000001000007874cfd89ae57b5f5496fc10b7b077e1b&token=crm|action=crm.documentgenerator.document.getPdf&SITE_ID=s1&id=1&_=pH1q7HXQsortDsHyrzyGu4RpDKLjviJ","imageUrlMachine":"https://b24-9f7fvg.bitrix24.ru/rest/crm.documentgenerator.document.getImage.json?auth=c45a0966006c4aa0006c188c00000001000007874cfd89ae57b5f5496fc10b7b077e1b&token=crm|action=crm.documentgenerator.document.getImage&SITE_ID=s1&id=1&_=e7hOdX6lIg8WOw8YkFzxy7CwLV8FMHJ"}]},"total":1,"time":{"start":1711885501.948452,"finish":1711885501.9738841,"duration":0.025432109832763672,"processing":0.0052890777587890625,"date_start":"2024-03-31T14:45:01+03:00","date_finish":"2024-03-31T14:45:01+03:00","operating_reset_at":1711886101,"operating":0}}`

	// Directly return the mock JSON string as a byte slice
	return []byte(mockJSON), nil
}

func DocumentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	values, err := url.ParseQuery(string(bs))
	if err != nil {
		log.Println("error parsing query:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	id, err := strconv.Atoi(values.Get("ID"))
	if err != nil {
		log.Println("error converting AUTH_EXPIRES to int:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	fmt.Println("DocumentHandler ID: ", id)

	entityId := "23"
	docs, err := GetDocsByDeal(authorize.GlobalAuthId, entityId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Directly write the mock JSON as the response
	_, err = w.Write(docs) // docs already contains the JSON
	if err != nil {
		log.Println("Error writing the mock response:", err)
	}
}

type requestBody struct {
	Select []string `json:"select"`
	Filter struct {
		EntityTypeId string `json:"entityTypeId"`
		EntityId     string `json:"entityId"`
	} `json:"filter"`
}

func GetDocsByDeal(authID string, entityId string) ([]byte, error) {
	bitrixMethod := "crm.documentgenerator.document.list"
	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/%s?auth=%s", bitrixMethod, authID)

	// Construct the request body
	body := requestBody{
		Select: []string{"*"},
		Filter: struct {
			EntityTypeId string "json:\"entityTypeId\""
			EntityId     string "json:\"entityId\""
		}{
			EntityTypeId: "2",
			EntityId:     entityId, // Use the variable value
		},
	}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Println("error marshaling request body:", err)
		return nil, err
	}

	// Create a new request with JSON body
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData)) // Switch to POST if applicable
	if err != nil {
		log.Println("error creating new request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and return the response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		return nil, err
	}

	log.Println("Response:", string(responseData))

	return responseData, nil
}
