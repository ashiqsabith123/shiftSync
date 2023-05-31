package razorpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"shiftsync/pkg/config"
	"shiftsync/pkg/db"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateContact(emp response.EmployeeDetails) (string, error) {
	url := "https://api.razorpay.com/v1/contacts"
	method := "POST"
	payload := map[string]interface{}{
		"name":         emp.Name,
		"email":        emp.Email,
		"contact":      emp.Phone,
		"type":         "employee",
		"reference_id": "Acme Contact ID 12345",
		"notes": map[string]string{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}

	jsonPayload, payloadErr := json.Marshal(payload)
	if payloadErr != nil {
		return "", payloadErr
	}

	client := &http.Client{}
	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if reqErr != nil {
		return "", reqErr
	}

	apiKey, apiSecret := config.GeyRazorpayKey()

	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	res, doErr := client.Do(req)
	if doErr != nil {
		return "", doErr
	}

	defer res.Body.Close()

	var result map[string]interface{}

	decodeErr := json.NewDecoder(res.Body).Decode(&result)
	if decodeErr != nil {
		return "", decodeErr
	}

	return result["id"].(string), nil
}

func CreateFundAccount(c *gin.Context, details response.AccountDetails) error {
	url := "https://api.razorpay.com/v1/fund_accounts"
	method := "POST"
	payload := map[string]interface{}{
		"contact_id":   details.Contact_id,
		"account_type": "bank_account",
		"bank_account": map[string]interface{}{
			"name":           details.Name,
			"ifsc":           details.Ifsc_code,
			"account_number": details.Account_no,
		},
	}

	jsonPayload, payloadErr := json.Marshal(payload)
	if payloadErr != nil {
		return payloadErr
	}

	client := &http.Client{}
	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if reqErr != nil {
		return reqErr
	}

	apiKey, apiSecret := config.GeyRazorpayKey()

	fmt.Println("api:", apiKey, apiSecret)

	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	res, doErr := client.Do(req)
	if doErr != nil {
		return doErr
	}

	defer res.Body.Close()

	var result map[string]interface{}

	decodeErr := json.NewDecoder(res.Body).Decode(&result)
	if decodeErr != nil {
		return decodeErr
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var razro domain.Razorpay

		fmt.Println(result)

		razro.Id = details.Id
		razro.Contact_id = result["id"].(string)

		if err := db.GetDatabaseInstance().Create(&razro).Error; err != nil {
			log.Fatal(err)
		}

		resp := response.SuccessResponse(http.StatusCreated, "Successfully Account Created", nil)
		c.JSON(http.StatusCreated, resp)
		return nil
	} else if res.StatusCode >= 400 && res.StatusCode < 500 {

		err, ok := result["error"].(map[string]interface{})
		if ok {

			return errors.New(err["description"].(string))
		}

		return errors.New("error")
	}

	return nil

}

func CreatePayouts(details response.AccDetails) error {
	fmt.Println(details.Account_no, details.Contact_id)

	payload := map[string]interface{}{
		"account_number":       "2323230030545681",
		"fund_account_id":      details.Contact_id,
		"amount":               details.Amount * 100,
		"currency":             "INR",
		"mode":                 "IMPS",
		"purpose":              "salary",
		"queue_if_low_balance": true,
		"reference_id":         details.Reference_id,
		"narration":            "Acme Corp Fund Transfer",
		"notes": map[string]string{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}

	url := "https://api.razorpay.com/v1/payouts"
	method := "POST"

	jsonPayload, payloadErr := json.Marshal(payload)
	if payloadErr != nil {
		return payloadErr
	}

	client := &http.Client{}
	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if reqErr != nil {
		return reqErr
	}

	apiKey, apiSecret := config.GeyRazorpayKey()

	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	res, doErr := client.Do(req)
	if doErr != nil {
		return doErr
	}

	defer res.Body.Close()

	var result map[string]interface{}

	decodeErr := json.NewDecoder(res.Body).Decode(&result)
	if decodeErr != nil {
		return decodeErr
	}

	var history domain.Transaction

	history.Date = time.Now()
	history.Amount = int(details.Amount)
	history.Employee_Id = details.ID
	history.Refrence_id = details.Reference_id

	if err := db.GetDatabaseInstance().Create(&history).Error; err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
