package solidgate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const DefaultApiUrl = "https://pay.solidgate.com/api/v1/"
const PatternFormUrl = "form?merchant=%s&form_data=%s&signature=%s"
const PatternResignFormUrl = "form/resign?merchant=%s&form_data=%s&signature=%s"

type IApi interface {
	Charge(data []byte) ([]byte, error)
	Refund(data []byte) ([]byte, error)
	Recurring(data []byte) ([]byte, error)
	Status(data []byte) ([]byte, error)
	InitPayment(data []byte) ([]byte, error)
	Resign(data []byte) ([]byte, error)
	Auth(data []byte) ([]byte, error)
	Settle(data []byte) ([]byte, error)
	Void(data []byte) ([]byte, error)
	ArnCode(data []byte) ([]byte, error)
	ApplePay(data []byte) ([]byte, error)
	GooglePay(data []byte) ([]byte, error)
	FormUrl(data []byte) (string, error)
	ResignFormUrl(data []byte) (string, error)
	FormMerchantData(data []byte) (*MerchantData, error)
}

type Api struct {
	MerchantId string
	PrivateKey string
	BaseUri    string
}

func (api *Api) Charge(data []byte) ([]byte, error) {
	return api.makeRequest("charge", data)
}

func (api *Api) Recurring(data []byte) ([]byte, error) {
	return api.makeRequest("recurring", data)
}

func (api *Api) Refund(data []byte) ([]byte, error) {
	return api.makeRequest("refund", data)
}

func (api *Api) Status(data []byte) ([]byte, error) {
	return api.makeRequest("status", data)
}

func (api *Api) InitPayment(data []byte) ([]byte, error) {
	return api.makeRequest("init-payment", data)
}

func (api *Api) Resign(data []byte) ([]byte, error) {
	return api.makeRequest("resign", data)
}

func (api *Api) Auth(data []byte) ([]byte, error) {
	return api.makeRequest("auth", data)
}

func (api *Api) Settle(data []byte) ([]byte, error) {
	return api.makeRequest("settle", data)
}

func (api *Api) Void(data []byte) ([]byte, error) {
	return api.makeRequest("void", data)
}

func (api *Api) ArnCode(data []byte) ([]byte, error) {
	return api.makeRequest("arn-code", data)
}

func (api *Api) ApplePay(data []byte) ([]byte, error) {
	return api.makeRequest("apple-pay", data)
}

func (api *Api) GooglePay(data []byte) ([]byte, error) {
	return api.makeRequest("google-pay", data)
}

func (api *Api) FormUrl(data []byte) (string, error) {
	secretKey := []byte(api.PrivateKey)[:32]
	encryptedData, err := EncryptCBC(secretKey, data)

	if err != nil {
		return "", err
	}

	encoded := base64.URLEncoding.EncodeToString(encryptedData)
	signature := api.generateSignature([]byte(encoded))

	return fmt.Sprintf(api.BaseUri+PatternFormUrl, api.MerchantId, encoded, signature), nil
}

func (api *Api) ResignFormUrl(data []byte) (string, error) {
	secretKey := []byte(api.PrivateKey)[:32]
	encryptedData, err := EncryptCBC(secretKey, data)

	if err != nil {
		return "", err
	}

	encoded := base64.URLEncoding.EncodeToString(encryptedData)
	signature := api.generateSignature([]byte(encoded))

	return fmt.Sprintf(api.BaseUri+PatternResignFormUrl, api.MerchantId, encoded, signature), nil
}

func (api *Api) generateSignature(data []byte) string {
	payloadData := api.MerchantId + string(data) + api.MerchantId

	keyForSign := []byte(api.PrivateKey)
	h := hmac.New(sha512.New, keyForSign)
	h.Write([]byte(payloadData))

	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}

func (api *Api) makeRequest(url string, payloadJson []byte) ([]byte, error) {
	if len(payloadJson) <= 0 {
		return nil, errors.New("empty payload")
	}

	req, err := http.NewRequest("POST", api.BaseUri+url, bytes.NewBuffer(payloadJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Signature", api.generateSignature(payloadJson))
	req.Header.Set("Merchant", api.MerchantId)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	return body, nil
}

func (api *Api) FormMerchantData(data []byte) (*MerchantData, error) {
	if len(data) <= 0 {
		return nil, errors.New("empty payload")
	}

	secretKey := []byte(api.PrivateKey)[:32]
	encryptedData, err := EncryptCBC(secretKey, data)

	if err != nil {
		return nil, err
	}

	encoded := base64.URLEncoding.EncodeToString(encryptedData)
	signature := api.generateSignature([]byte(encoded))

	merchantData := MerchantData{
		PaymentIntent: encoded,
		Merchant:      api.MerchantId,
		Signature:     signature,
	}

	return &merchantData, nil
}

func NewSolidGateApi(merchantId string, privateKey string, baseUri *string) *Api {
	defaultUrl := DefaultApiUrl

	if baseUri == nil {
		baseUri = &defaultUrl
	}

	return &Api{MerchantId: merchantId, PrivateKey: privateKey, BaseUri: *baseUri}
}
