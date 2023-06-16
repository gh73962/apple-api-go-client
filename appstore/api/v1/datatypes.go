package v1

import (
	"encoding/json"
	"time"
)

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
	BundleId                    string             `json:"bundleId,omitempty"`
	Environment                 Environment        `json:"environment,omitempty"`
	ExpiresDate                 int64              `json:"expiresDate,omitempty"`
	InAppOwnershipType          InAppOwnershipType `json:"inAppOwnershipType,omitempty"`
	IsUpgraded                  bool               `json:"isUpgraded,omitempty"`
	OfferIdentifier             string             `json:"offerIdentifier,omitempty"`
	OfferType                   OfferType          `json:"offerType,omitempty"`
	OriginalPurchaseDate        int64              `json:"originalPurchaseDate,omitempty"`
	OriginalTransactionId       string             `json:"originalTransactionId,omitempty"`
	ProductId                   string             `json:"productId,omitempty"`
	PurchaseDate                int64              `json:"purchaseDate,omitempty"`
	Quantity                    int                `json:"quantity,omitempty"`
	RevocationDate              int64              `json:"revocationDate,omitempty"`
	RevocationReason            string             `json:"revocationReason,omitempty"`
	SignedDate                  int64              `json:"signedDate,omitempty"`
	Storefront                  string             `json:"storefront,omitempty"`
	StorefrontId                string             `json:"storefrontId,omitempty"`
	SubscriptionGroupIdentifier string             `json:"subscriptionGroupIdentifier,omitempty"`
	TransactionId               string             `json:"transactionId,omitempty"`
	TransactionReason           string             `json:"transactionReason,omitempty"`
	Type                        TransactionType    `json:"type,omitempty"`
	WebOrderLineItemId          string             `json:"webOrderLineItemId,omitempty"`
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
	AutoRenewProductId          string           `json:"autoRenewProductId,omitempty"`
	AutoRenewStatus             int              `json:"autoRenewStatus,omitempty"`
	Environment                 Environment      `json:"environment,omitempty"`
	ExpirationIntent            ExpirationIntent `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate      int64            `json:"gracePeriodExpiresDate,omitempty"`
	IsInBillingRetryPeriod      bool             `json:"isInBillingRetryPeriod,omitempty"`
	OfferIdentifier             string           `json:"offerIdentifier,omitempty"`
	OfferType                   OfferType        `json:"offerType,omitempty"`
	OriginalTransactionId       string           `json:"originalTransactionId,omitempty"`
	PriceIncreaseStatus         int              `json:"priceIncreaseStatus,omitempty"`
	ProductId                   string           `json:"productId,omitempty"`
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
	BundleId    string                            `json:"bundleId,omitempty"`
	AppAppleId  int64                             `json:"appAppleId,omitempty"`
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
	OriginalTransactionId string             `json:"originalTransactionId,omitempty"`
	Status                SubscriptionStatus `json:"status,omitempty"`
	SignedRenewalInfo     string             `json:"signedRenewalInfo,omitempty"`
	SignedTransactionInfo string             `json:"signedTransactionInfo,omitempty"`
}

// See https://developer.apple.com/documentation/appstoreserverapi/error_codes
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

type HistoryResponse struct {
	Revision           string      `json:"revision"`
	BundleId           string      `json:"bundleId"`
	AppAppleId         int         `json:"appAppleId"`
	Environment        Environment `json:"environment"`
	HasMore            bool        `json:"hasMore"`
	SignedTransactions []string    `json:"signedTransactions"`
}

type OrderLookupResponse struct {
	Status             int      `json:"status"`
	SignedTransactions []string `json:"signedTransactions"`
}

func (o *OrderLookupResponse) IsValid() bool {
	return o.Status == 0
}

type RefundHistoryResponse struct {
	Revision           string   `json:"revision"`
	HasMore            bool     `json:"hasMore"`
	SignedTransactions []string `json:"signedTransactions"`
}
