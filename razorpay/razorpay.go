package razorpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"shiftsync/pkg/config"
	"shiftsync/pkg/helper/response"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context, emp response.EmployeeDetails) error {
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

	fmt.Println(result)

	// Continue with the response handling
	resp := response.SuccessResponse(http.StatusCreated, "Successfully Account Created", nil)
	c.JSON(http.StatusCreated, resp)
	return nil
}
