package techdata

import "encoding/xml"

type XMLOrderSubmit struct {
	XMLName xml.Name `xml:"XML_Order_Submit"`
	Header  struct {
		Username        string `xml:"UserName"`
		Password        string `xml:"Password"`
		ResponseVersion string `xml:"ResponseVersion"`
		OrderTypeCode   string `xml:"OrderTypeCode"`
		PoNum           string `xml:"PONbr"`
		RefNum          struct {
			RefIDQual string `xml:"RefIDQual"`
			RefID     string `xml:"RefID"`
		} `xml:"RefNbrs"`
		SalesRequirementCode string `xml:"SalesRequirementCode"`
		Name                 string `xml:"Name"`
		AddrInfo             struct {
			Addr []string `xml:"Addr"`
		} `xml:"AddrInfo"`
		CityName          string `xml:"CityName"`
		StateProvinceCode string `xml:"StateProvinceCode"`
		PostalCode        string `xml:"PostalCode"`
		ContactName       string `xml:"ContactName"`
		ContactPhone      string `xml:"ContactPhoneNbr"`
		EndUserInfo       struct {
			GovAgency         string `xml:"EuiGovAgency"`
			GovCabinetLevel   string `xml:"EuiGovCabinetLevel"`
			ContractNum       string `xml:"EuiContractNbr"`
			ContractType      string `xml:"EuiContractType"`
			OrderPriority     string `xml:"EuiOrderPriority"`
			MarketType        string `xml:"EuiMarketType"`
			ContactName       string `xml:"EuiContactName"`
			PhoneNum          string `xml:"EuiPhoneNbr"`
			FaxNum            string `xml:"EuiFaxNbr"`
			Name              string `xml:"EuiName"`
			Addr1             string `xml:"EuiAddr1"`
			Addr2             string `xml:"EuiAddr2"`
			Addr3             string `xml:"EuiAddr3"`
			CityName          string `xml:"EuiCityName"`
			StateProvinceCode string `xml:"EuiStateProvinceCode"`
			PostalCode        string `xml:"EuiPostalCode"`
			CountryCode       string `xml:"EuiCountryCode"`
			SicCode           string `xml:"EuiSicCode"`
			OrderPromoType    string `xml:"EuiOrderPromoType"`
			EndUserLicenseNum string `xml:"EuiEndUserLicenseNbr"`
			EndUserPODate     string `xml:"EuiEndUserPODate"`
			EndUserRef1       string `xml:"EuiEndUserRef1"`
			EndUserRef2       string `xml:"EuiEndUserRef2"`
			EndUserRef3       string `xml:"EuiEndUserRef3"`
			InstallName       string `xml:"EuiInstallName"`
			DropShipType      string `xml:"EuiDropShipType"`
			ContactEmailAddr1 string `xml:"EuiContactEmailAddr1"`
			ContactEmailAddr2 string `xml:"EuiContactEmailAddr2"`
		} `xml:"EndUserInfo"`
		MyOrderTracker struct {
			ResellerEmail  string   `xml:"ResellerEmail"`
			ResellerEvents []string `xml:"ResellerEvents"`
		} `xml:"MyOrderTracker"`
	} `xml:"Header"`
	Detail struct {
		LineInfo struct {
			QtyOrdered       string `xml:"QtyOrdered"`
			ProductIDQual    string `xml:"ProductIDQual"`
			ProductID        string `xml:"ProductID"`
			WhseCode         string `xml:"WhseCode"`
			IDCode           string `xml:"IDCode"`
			OrderMessageLine string `xml:"OrderMessageLine"`
		} `xml:"LineInfo"`
	} `xml:"Detail"`
}

type XMLOrderResponse struct {
	XMLName xml.Name `xml:"XML_Order_Response"`
	Header  struct {
		Username        string `xml:"UserName"`
		TransSetIDCode  string `xml:"TransSetIDCode"`
		TransControlID  string `xml:"TransControlID"`
		ResponseVersion string `xml:"ResponseVersion"`
		PurposeCode     string `xml:"PurposeCode"`
		PONum           string `xml:"PONbr"`
		RefID           string `xml:"RefID"`
	} `xml:"Header"`
	Summary string `xml:"Summary"`
}
