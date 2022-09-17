package domain

const (
	Cancellation = "cancellation"
	Reversal     = "reversal"
	Confirmation = "confirmation"
)

type Transaction struct {
	Id                string  `json:"id" bson:"_id"`
	AuthorizationCode string  `json:"authorization_code" bson:"authorization_code"`
	AcquirerCode      string  `json:"acquirer_code" bson:"acquirer_code"`
	MerchantCode      string  `json:"merchant_code" bson:"merchant_code"`
	CurrencyCode      string  `json:"currency_code" bson:"currency_code"`
	CountryCode       string  `json:"country_code" bson:"country_code"`
	ProductId         string  `json:"product_id" bson:"product_id"`
	PosId             string  `json:"pos_id" bson:"pos_id"`
	WithPassword      *bool   `json:"with_password" bson:"with_password"`
	Type              string  `json:"type" bson:"type"`
	Amount            float32 `json:"amount" bson:"amount"`
	CardId            string  `json:"card_id"  bson:"card_id"`
	ExternalTid       string  `json:"external_tid" bson:"external_tid"`
}
