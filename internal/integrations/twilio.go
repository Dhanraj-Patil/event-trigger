package integration

import (
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func TwilioSMSAPI(phoneNo string, message string) (string, error) {
	godotenv.Load()
	accountSid := "AC70e5f87851605198123c68026624a312"
	authToken := "ee66236a0354839fdbec8a78c1dc394c"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+91" + phoneNo)
	params.SetFrom("15044748758")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}
	response, _ := json.Marshal(*resp)
	return string(response), err
}
