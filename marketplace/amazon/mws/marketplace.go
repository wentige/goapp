package mws

type Marketplace struct {
	Name     string
	Country  string
	Endpoint string
	ID       string
}

var marketplaceList = []Marketplace{
	{
		Name:     "Brazil",
		Country:  "BR",
		Endpoint: "https://mws.amazonservices.com",
		ID:       "A2Q3Y263D00KWC",
	},
	{
		Name:     "Canada",
		Country:  "CA",
		Endpoint: "https://mws.amazonservices.ca",
		ID:       "A2EUQ1WTGCTBG2",
	},
	{
		Name:     "Mexico",
		Country:  "MX",
		Endpoint: "https://mws.amazonservices.com.mx",
		ID:       "A1AM78C64UM0Y8",
	},

	{
		Name:     "US",
		Country:  "US",
		Endpoint: "https://mws.amazonservices.com",
		ID:       "ATVPDKIKX0DER",
	},
	{
		Name:     "United Arab Emirates",
		Country:  "AE",
		Endpoint: "https://mws.amazonservices.ae",
		ID:       "A2VIGQ35RCS4UG",
	},
	{
		Name:     "Germany",
		Country:  "DE",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "A1PA6795UKMFR9",
	},
	{
		Name:     "Spain",
		Country:  "ES",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "A1RKKUPIHCS9HS",
	},
	{
		Name:     "France",
		Country:  "FR",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "A13V1IB3VIYZZH",
	},
	{
		Name:     "UK",
		Country:  "GB",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "A1F83G8C2ARO7P",
	},
	{
		Name:     "India",
		Country:  "IN",
		Endpoint: "https://mws.amazonservices.in",
		ID:       "A21TJRUUN4KGV",
	},
	{
		Name:     "Italy",
		Country:  "IT",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "APJ6JRA9NG5V4",
	},
	{
		Name:     "Turkey",
		Country:  "TR",
		Endpoint: "https://mws-eu.amazonservices.com",
		ID:       "A33AVAJ2PDY3EV",
	},
	{
		Name:     "Singapore",
		Country:  "SG",
		Endpoint: "https://mws-fe.amazonservices.com",
		ID:       "A19VAU5U5O7RUS",
	},
	{
		Name:     "Australia",
		Country:  "AU",
		Endpoint: "https://mws.amazonservices.com.au",
		ID:       "A39IBJ37TRP1C6",
	},
	{
		Name:     "Japan",
		Country:  "JP",
		Endpoint: "https://mws.amazonservices.jp",
		ID:       "A1VC38T7YXB528",
	},
	{
		Name:     "China",
		Country:  "CN",
		Endpoint: "https://mws.amazonservices.com.cn",
		ID:       "AAHKV2X7AFYLW",
	},
}

var marketplaceMap = make(map[string]Marketplace)

func initMarketplaces() {
	for _, marketplace := range marketplaceList {
		id := marketplace.ID
		country := marketplace.Country
		marketplaceMap[country] = marketplace
		marketplaceMap[id] = marketplace
	}
}

func GetMarketplaces() []Marketplace {
	return marketplaceList
}

func GetMarketplace(mkt string) Marketplace {
	marketplace, _ := marketplaceMap[mkt]
	return marketplace
}

func GetEndpoint(mkt string) string {
	marketplace, _ := marketplaceMap[mkt]
	return marketplace.Endpoint
}

func GetMarketplaceID(mkt string) string {
	marketplace, _ := marketplaceMap[mkt]
	return marketplace.ID
}

/*
func main() {
    Init();
    fmt.Println(GetMarketplace("CN"))
    fmt.Println(GetMarketplace("AAHKV2X7AFYLW"))

    fmt.Println(GetMarketplace("CA"))
    fmt.Println(GetMarketplace("A2EUQ1WTGCTBG2"))

    fmt.Println(GetMarketplace("XX"))
    fmt.Println(GetMarketplace("XXXXXXXXXXXXXX"))
}*/
