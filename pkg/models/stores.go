package models

type Stores struct {
	Stores []StoresItem `json:"stores"`
	Config StoreConfig  `json:"config"`
}

type StoresItem struct {
	StoreNumber string `json:"storeNumber"`
	City        string `json:"city"`
	Latitude    string `json:"latitude"`
	StoreName   string `json:"storeName"`
	Enabled     bool   `json:"enabled"`
	Longitude   string `json:"longitude"`
}

type StoreConfig struct {
	UpdatedTime      int64 `json:"updatedTime"`
	SupportedDomains struct {
		Field1 string `json:"1"`
		Field2 string `json:"2"`
		Field3 string `json:"3"`
		Field4 string `json:"4"`
		Field5 string `json:"0"`
	} `json:"supportedDomains"`
	IUPEnabled             bool  `json:"iUPEnabled"`
	PostLaunchStartDate    int64 `json:"postLaunchStartDate"`
	ContractOptionsEnabled bool  `json:"contractOptionsEnabled"`
	AppleCarePreference    struct {
		U struct {
			DefaultAppleCare       bool `json:"defaultAppleCare"`
			ConsiderUserPreference bool `json:"considerUserPreference"`
		} `json:"U"`
		E struct {
			DefaultAppleCare       bool `json:"defaultAppleCare"`
			ConsiderUserPreference bool `json:"considerUserPreference"`
		} `json:"E"`
		N struct {
			DefaultAppleCare       bool `json:"defaultAppleCare"`
			ConsiderUserPreference bool `json:"considerUserPreference"`
		} `json:"N"`
	} `json:"appleCarePreference"`
	Timezone           string   `json:"timezone"`
	PreorderStartDate  int64    `json:"preorderStartDate"`
	ProductGridEnabled bool     `json:"productGridEnabled"`
	LaunchDate         int64    `json:"launchDate"`
	CurrencyFormat     string   `json:"currencyFormat"`
	ActiveChannels     []string `json:"activeChannels"`
	PrsContractOptions string   `json:"prsContractOptions"`
	DisplayState       bool     `json:"displayState"`
	PreorderEndDate    int64    `json:"preorderEndDate"`
	NumberFormat       string   `json:"numberFormat"`
	ReservableQuantity struct {
		U int `json:"U"`
		E int `json:"E"`
		N int `json:"N"`
	} `json:"reservableQuantity"`
	ProductPricingPreference struct {
		U struct {
			UseDisplayPrice     bool `json:"useDisplayPrice"`
			UseInstallmentPrice bool `json:"useInstallmentPrice"`
		} `json:"U"`
		E struct {
			UseDisplayPrice     bool `json:"useDisplayPrice"`
			UseInstallmentPrice bool `json:"useInstallmentPrice"`
		} `json:"E"`
		N struct {
			UseDisplayPrice     bool `json:"useDisplayPrice"`
			UseInstallmentPrice bool `json:"useInstallmentPrice"`
		} `json:"N"`
	} `json:"productPricingPreference"`
	ReservationURL string `json:"reservationURL"`
}
