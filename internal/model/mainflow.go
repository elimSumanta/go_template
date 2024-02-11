package model

// Variable Constants --------------------------------

// Constants --------------------------------

// Variable Structs ----------------------------------------------------------------
type GetCryptoListRes struct {
	AmountFilter  []LabelKey  `json:"amount_filter"`
	PaymentFilter []LabelKey  `json:"payment_filter"`
	CryptoList    []string    `json:"crypto"`
	Currency      []LabelIcon `json:"currency"`
}

type LabelKey struct {
	Label string
	Key   string
}

type LabelIcon struct {
	Label string
	Icon  string
}

type GetAdsListRes struct {
	PictureUrl    string        `json:"picture_url"`
	Name          string        `json:"name"`
	IsVerified    bool          `json:"is_verified"`
	UserTradeInfo UserTradeInfo `json:"user_trade_info"`
	AdsInfo       AdsInfo       `json:"ads_info"`
}

type AdsInfo struct {
	CryptoLabel          string   `json:"crypto_label"`
	Price                int64    `json:"price"`
	PriceFormat          string   `json:"price_format"`
	Qty                  float64  `json:"qty"`
	Limit                string   `json:"limit"`
	SupportedPaymentList []string `json:"supported_payment_list"`
}

// ----------------------------------------------------------------
