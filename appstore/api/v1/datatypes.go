package appstoreapi

import (
	"encoding/json"
	"time"
)

// See https://developer.apple.com/documentation/appstoreserverapi/data_types

type OfferType int

const (
	IntroductoryOffer              OfferType = 1
	PromotionalOffer               OfferType = 2
	OfferWithSubscriptionOfferCode OfferType = 3
)

type Environment string

const (
	Sandbox    Environment = "Sandbox"
	Production Environment = "Production"
)

type InAppOwnershipType string

const (
	FamilyShared InAppOwnershipType = "FAMILY_SHARED"
	Purchased    InAppOwnershipType = "PURCHASED"
)

type TransactionType string

const (
	AutoRenewableSubscription TransactionType = "Auto-Renewable Subscription"
	NonConsumable             TransactionType = "Non-Consumable"
	Consumable                TransactionType = "Consumable"
	NonRenewingSubscription   TransactionType = "Non-Renewing Subscription"
)

// JWSDecodedHeader https://developer.apple.com/documentation/appstoreserverapi/jwsdecodedheader
type JWSDecodedHeader struct {
	Alg string   `json:"alg,omitempty"`
	X5c []string `json:"x5c,omitempty"`
}

type JWSTransaction struct {
	Header    JWSDecodedHeader
	Payload   JWSTransactionDecodedPayload
	Signature string
}

// JWSTransactionDecodedPayload https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	AppAccountToken             string             `json:"appAccountToken,omitempty"`
	BundleID                    string             `json:"bundleId,omitempty"`
	Environment                 Environment        `json:"environment,omitempty"`
	ExpiresDate                 int64              `json:"expiresDate,omitempty"`
	InAppOwnershipType          InAppOwnershipType `json:"inAppOwnershipType,omitempty"`
	IsUpgraded                  bool               `json:"isUpgraded,omitempty"`
	OfferIdentifier             string             `json:"offerIdentifier,omitempty"`
	OfferType                   OfferType          `json:"offerType,omitempty"`
	OriginalPurchaseDate        int64              `json:"originalPurchaseDate,omitempty"`
	OriginalTransactionID       string             `json:"originalTransactionId,omitempty"`
	ProductID                   string             `json:"productId,omitempty"`
	PurchaseDate                int64              `json:"purchaseDate,omitempty"`
	Quantity                    int                `json:"quantity,omitempty"`
	RevocationDate              int64              `json:"revocationDate,omitempty"`
	RevocationReason            string             `json:"revocationReason,omitempty"`
	SignedDate                  int64              `json:"signedDate,omitempty"`
	Storefront                  string             `json:"storefront,omitempty"`
	StorefrontID                string             `json:"storefrontId,omitempty"`
	SubscriptionGroupIdentifier string             `json:"subscriptionGroupIdentifier,omitempty"`
	TransactionID               string             `json:"transactionId,omitempty"`
	TransactionReason           string             `json:"transactionReason,omitempty"`
	Type                        TransactionType    `json:"type,omitempty"`
	WebOrderLineItemID          string             `json:"webOrderLineItemId,omitempty"`
}

func (j *JWSTransactionDecodedPayload) GetPurchaseTime() time.Time {
	return time.Unix(j.PurchaseDate/1e3, 0)
}

func (j *JWSTransactionDecodedPayload) GetExpiresTime() time.Time {
	return time.Unix(j.ExpiresDate/1e3, 0)
}

func (j *JWSTransactionDecodedPayload) GetSignedTime() time.Time {
	return time.Unix(j.SignedDate/1e3, 0)
}

type JWSRenewalInfo struct {
	Header    JWSDecodedHeader
	Payload   JWSRenewalInfoDecodedPayload
	Signature string
}

type ExpirationIntent int

const (
	CanceledSubscription    ExpirationIntent = 1
	BillingError            ExpirationIntent = 2
	NotConsentPriceIncrease ExpirationIntent = 3
	NotAvailable            ExpirationIntent = 4
	OtherReason             ExpirationIntent = 5
)

// JWSRenewalInfoDecodedPayload https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
type JWSRenewalInfoDecodedPayload struct {
	AutoRenewProductID          string           `json:"autoRenewProductId,omitempty"`
	AutoRenewStatus             int              `json:"autoRenewStatus,omitempty"`
	Environment                 Environment      `json:"environment,omitempty"`
	ExpirationIntent            ExpirationIntent `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate      int64            `json:"gracePeriodExpiresDate,omitempty"`
	IsInBillingRetryPeriod      bool             `json:"isInBillingRetryPeriod,omitempty"`
	OfferIdentifier             string           `json:"offerIdentifier,omitempty"`
	OfferType                   OfferType        `json:"offerType,omitempty"`
	OriginalTransactionID       string           `json:"originalTransactionId,omitempty"`
	PriceIncreaseStatus         int              `json:"priceIncreaseStatus,omitempty"`
	ProductID                   string           `json:"productId,omitempty"`
	RecentSubscriptionStartDate int64            `json:"recentSubscriptionStartDate,omitempty"`
	RenewalDate                 int64            `json:"renewalDate,omitempty"`
	SignedDate                  int64            `json:"signedDate,omitempty"`
}

func (j *JWSRenewalInfoDecodedPayload) GetRenewalTime() time.Time {
	return time.Unix(j.RenewalDate/1e3, 0)
}

func (j *JWSRenewalInfoDecodedPayload) IsAutoRenew() bool {
	return j.AutoRenewStatus == 1
}

func (j *JWSRenewalInfoDecodedPayload) IsConsentedPriceIncrease() bool {
	return j.PriceIncreaseStatus == 1
}

type TransactionInfoResponse struct {
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

// StatusResponse https://developer.apple.com/documentation/appstoreserverapi/statusresponse
type StatusResponse struct {
	BundleID    string                            `json:"bundleId,omitempty"`
	AppAppleID  int64                             `json:"appAppleId,omitempty"`
	Environment string                            `json:"environment,omitempty"`
	Data        []SubscriptionGroupIdentifierItem `json:"data,omitempty"`
}

type SubscriptionGroupIdentifierItem struct {
	SubscriptionGroupIdentifier string                 `json:"subscriptionGroupIdentifier,omitempty"`
	LastTransactions            []LastTransactionsItem `json:"lastTransactions,omitempty"`
}

type SubscriptionStatus int

const (
	Active             SubscriptionStatus = 1
	Expired            SubscriptionStatus = 2
	BillingRetryPeriod SubscriptionStatus = 3
	BillingGracePeriod SubscriptionStatus = 4
	Revoked            SubscriptionStatus = 5
)

type LastTransactionsItem struct {
	OriginalTransactionID string             `json:"originalTransactionId,omitempty"`
	Status                SubscriptionStatus `json:"status,omitempty"`
	SignedRenewalInfo     string             `json:"signedRenewalInfo,omitempty"`
	SignedTransactionInfo string             `json:"signedTransactionInfo,omitempty"`
}

// HistoryResponse see https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type HistoryResponse struct {
	Revision           string      `json:"revision,omitempty"`
	BundleID           string      `json:"bundleId,omitempty"`
	AppAppleID         int         `json:"appAppleId,omitempty"`
	Environment        Environment `json:"environment,omitempty"`
	HasMore            bool        `json:"hasMore,omitempty"`
	SignedTransactions []string    `json:"signedTransactions,omitempty"`
}

// OrderLookupResponse see https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	Status             int      `json:"status,omitempty"`
	SignedTransactions []string `json:"signedTransactions,omitempty"`
}

func (o *OrderLookupResponse) IsValid() bool {
	return o.Status == 0
}

// RefundHistoryResponse see https://developer.apple.com/documentation/appstoreserverapi/refundhistoryresponse
type RefundHistoryResponse struct {
	Revision           string   `json:"revision,omitempty"`
	HasMore            bool     `json:"hasMore,omitempty"`
	SignedTransactions []string `json:"signedTransactions,omitempty"`
}

// ErrorResponse See https://developer.apple.com/documentation/appstoreserverapi/error_codes
type ErrorResponse struct {
	HTTPStatus   int    `json:"httpStatus,omitempty"`
	ErrorCode    int64  `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e == nil {
		return ""
	}
	data, _ := json.Marshal(e)
	return string(data)
}
