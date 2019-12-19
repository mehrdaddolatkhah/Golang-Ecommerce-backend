package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"cafekalaa/api/utils"
)

func getSmsToken() string {
	url := "test"

	requestBody, err := json.Marshal(map[string]string{
		"test": "test",
		"test":  "test",
	})

	if err != nil {
		panic(err)
	}

	req, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(req.Body)
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	token := data["TokenKey"].(string)
	return token
}

func SendVerficationSms(mobile string) string {

	apiUrl := "https://RestfulSms.com/api/UltraFastSend"

	code := utils.MakeRandomNumber(11111, 99999)
	strCode := strconv.FormatInt(int64(code), 10)
	strToken := getSmsToken()

	client := &http.Client{}

	type Param struct {
		Parameter      string
		ParameterValue string
	}

	type smsPayload struct {
		ParameterArray []Param
		Mobile         string
		TemplateID     string
	}

	smsData := &smsPayload{
		ParameterArray: []Param{
			{
				Parameter:      "VerificationCode",
				ParameterValue: strCode,
			},
		},
		Mobile:     mobile,
		TemplateID: "16712",
	}

	smsBuf, _ := json.Marshal(smsData)
	r, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(smsBuf))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("x-sms-ir-secure-token", strToken)

	resp, _ := client.Do(r)

	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil && body == nil {
		panic(err)
	}

	return strCode
}
