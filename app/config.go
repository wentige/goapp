package app

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// Config.Database.Host
	Database struct {
		Adapter  string `json:"adapter"`
		Dbname   string `json:"dbname"`
		Host     string `json:"host"`
		Password string `json:"password"`
		Username string `json:"username"`
	} `json:"database"`

	// Config.Amazon.US.SecretKey
	Amazon struct {
		AE amazonAccount `json:"bte-amazon-ae"`
		AU amazonAccount `json:"bte-amazon-au"`
		CA amazonAccount `json:"bte-amazon-ca"`
		CN amazonAccount `json:"bte-amazon-cn"`
		MX amazonAccount `json:"bte-amazon-mx"`
		UK amazonAccount `json:"bte-amazon-uk"`
		US amazonAccount `json:"bte-amazon-us"`
	} `json:"amazon"`

	// Config.Bestbuy.ApiKey
	Bestbuy struct {
		Apikey   string `json:"apikey"`
		Password string `json:"password"`
		Username string `json:"username"`
	} `json:"bestbuy"`

	// Config.Ebay.BTE.UserToken
	Ebay struct {
		BTE eBayAccount `json:"bte"`
		GFS eBayAccount `json:"gfs"`
		ODO eBayAccount `json:"odo"`
	} `json:"ebay"`

	// Config.Email.POP3.Hostname
	// Config.Email.SMTP.Hostname
	Email struct {
		POP3 struct {
			Hostname  string `json:"hostname"`
			Mailboxes []struct {
				Password string `json:"password"`
				Username string `json:"username"`
			} `json:"mailboxes"`
		} `json:"pop3"`

		SMTP struct {
			Hostname string `json:"hostname"`
		} `json:"smtp"`
	} `json:"email"`

	// Config.Newegg.US.SellerID
	Newegg struct {
		B2B neweggAccount `json:"B2B"`
		CA  neweggAccount `json:"CA"`
		US  neweggAccount `json:"US"`
	} `json:"newegg"`

	// Config.ShippingEasy.APIKey
	ShippingEasy struct {
		APIKey    string `json:"apiKey"`
		APISecret string `json:"apiSecret"`
	} `json:"shippingEasy"`

	// Config.Walmart.US.PrivateKey
	Walmart struct {
		CA walmartAccount `json:"CA"`
		US walmartAccount `json:"US"`
	} `json:"walmart"`

	// Config.Supplier.DH.Username
	// Config.Supplier.INGRAM.Password
	Supplier struct {
		ASI struct {
			Cert string `json:"CERT"`
			Cid  string `json:"CID"`
			URL  string `json:"url"`
		} `json:"asi"`

		DH struct {
			Backorder   string `json:"backorder"`
			Branches    string `json:"branches"`
			Dropshippw  string `json:"dropshippw"`
			Onlybranch  string `json:"onlybranch"`
			Partship    string `json:"partship"`
			Password    string `json:"password"`
			Shipcarrier string `json:"shipcarrier"`
			Shipservice string `json:"shipservice"`
			URL         string `json:"url"`
			Username    string `json:"username"`
		} `json:"dh"`

		INGRAM struct {
			AutoRelease   string `json:"autoRelease"`
			BackOrder     string `json:"backOrder"`
			CarrierCode   string `json:"carrierCode"`
			LoginID       string `json:"loginId"`
			Password      string `json:"password"`
			SplitLine     string `json:"splitLine"`
			SplitShipment string `json:"splitShipment"`
			URL           string `json:"url"`
		} `json:"ingram"`

		SYNNEX struct {
			AccountNo  string `json:"accountNo"`
			CustomerNo string `json:"customerNo"`
			Password   string `json:"password"`
			Shipmethod string `json:"shipmethod"`
			URL        string `json:"url"`
			Username   string `json:"username"`
		} `json:"synnex"`

		TECHDATA struct {
			Password string `json:"password"`
			URL      string `json:"url"`
			Username string `json:"username"`
		} `json:"techdata"`
	} `json:"xmlapi"`
}

type amazonAccount struct {
	KeyID         string `json:"keyId"`
	MarketplaceID string `json:"marketplaceId"`
	MerchantID    string `json:"merchantId"`
	SecretKey     string `json:"secretKey"`
	ServiceURL    string `json:"serviceUrl"`
	MWSAuthToken  string `json:"MWSAuthToken"`
}

type eBayAccount struct {
	AppID     string `json:"appID"`
	CertID    string `json:"certID"`
	DevID     string `json:"devID"`
	ServerURL string `json:"serverUrl"`
	SiteID    int64  `json:"siteID"`
	UserToken string `json:"userToken"`
}

type neweggAccount struct {
	Authorization string `json:"Authorization"`
	Endpoint      string `json:"Endpoint"`
	SecretKey     string `json:"SecretKey"`
	SellerID      string `json:"SellerID"`
	Version       string `json:"Version"`
}

type walmartAccount struct {
	ChannelType  string `json:"channelType"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	ConsumerID   string `json:"consumerId"`
	PrivateKey   string `json:"privateKey"`
}
