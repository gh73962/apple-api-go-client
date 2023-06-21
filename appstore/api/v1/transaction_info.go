package appstoreapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gh73962/appleapis/appstore/api/internal/httputils"
	"github.com/gh73962/appleapis/appstore/api/v1/datatypes"
)

type Service struct {
	client    *http.Client
	BasePath  string
	UserAgent string
}

func (s *Service) TransactionInfo(ctx context.Context, bearer, transactionID string) (*datatypes.JWSTransaction, error) {
	req, err := http.NewRequest(http.MethodGet, s.BasePath+"/transactions/"+transactionID, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+bearer)

	resp, err := httputils.SendAndRetry(ctx, s.client, req, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactionResponse datatypes.TransactionInfoResponse
	if err = json.NewDecoder(resp.Body).Decode(&transactionResponse); err != nil {
		return nil, err
	}

	return DecodeToJWSTransaction(transactionResponse.SignedTransactionInfo)
}
