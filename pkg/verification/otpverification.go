package verification

import (
	"fmt"
	"shiftsync/pkg/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	SID        string
	AUTH       string
	SERVICE_ID string
)

var client *twilio.RestClient

func InitTwilio(cn config.Config) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cn.Twilio.Account_sid,
		Password: cn.Twilio.Auth_token,
	})
	SERVICE_ID = cn.Twilio.Service_id
}

func SendOtp(phone int64) (string, error) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91" + fmt.Sprint(phone))
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(SERVICE_ID, params)

	if err != nil {
		return "Otp not send", err
	}

	return "Otp send succesfully", err

}

func ValidateOtp(phone int64, code string) error {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91" + fmt.Sprint(phone))
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(SERVICE_ID, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}
