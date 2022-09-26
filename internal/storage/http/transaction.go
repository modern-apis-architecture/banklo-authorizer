package http

type ExternalRequestTransaction struct {
	AcquirerCode      string `json:"acquirer_code,omitempty"`
	AuthorizationCode string `json:"authorization_code,omitempty"`

	// The country code from ISO 3166 find more at https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
	CountryCode string `json:"country_code,omitempty"`

	// The currency code from ISO 4217 find more at https://en.wikipedia.org/wiki/ISO_4217
	CurrencyCode            string                  `json:"currency_code,omitempty"`
	MerchantCode            string                  `json:"merchant_code,omitempty"`
	ExternalTransactionData ExternalTransactionData `json:"transaction_data,omitempty"`
}

type ExternalTransactionData struct {
	Amount          float32 `json:"amount,omitempty"`
	CardId          string  `json:"card_id,omitempty"`
	TransactionId   string  `json:"transaction_id,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	WithPassword    bool    `json:"with_password,omitempty"`
}
