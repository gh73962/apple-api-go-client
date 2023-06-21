package appstoreapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gh73962/appleapis/appstore/api/v1/datatypes"
)

// AllSubscriptionStatuses see https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
// TODO Query Parameters
func (s *Service) AllSubscriptionStatuses(ctx context.Context, bearer, transactionID string,
	status datatypes.SubscriptionStatus) (*datatypes.StatusResponse, error) {

	req, err := http.NewRequest(http.MethodGet, s.BasePath+"subscriptions/"+transactionID, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("User-Agent", s.UserAgent)

	resp, err := s.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rsp datatypes.StatusResponse
	if err = json.NewDecoder(resp.Body).Decode(&rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// TODO  subscription_renewal_date_extension
