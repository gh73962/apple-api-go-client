package appstoreapi

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func NewJWSTransaction(signedData string) (*JWSTransaction, error) {
	array := strings.Split(signedData, ".")
	if len(array) != 3 {
		return nil, errors.New("invalid signedData")
	}

	headerData, err := base64.StdEncoding.DecodeString(array[0])
	if err != nil {
		return nil, err
	}

	payloadData, err := base64.StdEncoding.DecodeString(array[1])
	if err != nil {
		return nil, err
	}

	var t JWSTransaction
	if err := json.Unmarshal(headerData, &t.Header); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(payloadData, &t.Payload); err != nil {
		return nil, err
	}

	t.Signature = array[2]

	return &t, nil
}

func NewJWSRenewalInfo(signedData string) (*JWSRenewalInfo, error) {
	array := strings.Split(signedData, ".")
	if len(array) != 3 {
		return nil, errors.New("invalid signedData")
	}

	headerData, err := base64.StdEncoding.DecodeString(array[0])
	if err != nil {
		return nil, err
	}

	payloadData, err := base64.StdEncoding.DecodeString(array[1])
	if err != nil {
		return nil, err
	}

	var t JWSRenewalInfo
	if err := json.Unmarshal(headerData, &t.Header); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(payloadData, &t.Payload); err != nil {
		return nil, err
	}

	t.Signature = array[2]

	return &t, nil
}
