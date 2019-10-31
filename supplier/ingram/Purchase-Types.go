package ingram

import "encoding/xml"

type OrderRequest struct {
	XMLName xml.Name `xml:"OrderRequest"`
	Version string   `xml:"Version"`
	Header  struct {
		SenderID      string `xml:"SenderID"`
		ReceiverID    string `xml:"ReceiverID"`
		CountryCode   string `xml:"CountryCode"`
		LoginID       string `xml:"LoginID"`
		Password      string `xml:"Password"`
		TransactionID string `xml:"TransactionID"`
	} `xml:"TransactionHeader"`
	OrderHeader struct {
		BillToSuffix   string `xml:"BillToSuffix"`
		AddressingInfo struct {
			CustomerPO string `xml:"CustomerPO"`
			ShipToAttn string `xml:"ShipToAttention"`
			EndUserPO  string `xml:"EndUserPO"`
			ShipTo     struct {
				Address struct {
					Address1   string `xml:"ShipToAddress1"`
					Address2   string `xml:"ShipToAddress2"`
					Address3   string `xml:"ShipToAddress3"`
					City       string `xml:"ShipToCity"`
					Province   string `xml:"ShipToProvince"`
					PostalCode string `xml:"ShipToPostalCode"`
				} `xml:"Address"`
			} `xml:"ShipTo"`
		} `xml:"AddressingInformation"`
		ProcessingOptions struct {
			CarrierCode     string `xml:"CarrierCode"`
			AutoRelease     string `xml:"AutoRelease"`
			ShipmentOptions struct {
				BackOrderFlag     string `xml:"BackOrderFlag"`
				SplitShipmentFlag string `xml:"SplitShipmentFlag"`
				SplitLine         string `xml:"SplitLine"`
				ShipFromBranches  string `xml:"ShipFromBranches"`
				DeliveryDate      string `xml:"DeliveryDate"`
			} `xml:"ShipmentOptions"`
		} `xml:"ProcessingOptions"`
		DynamicMessage struct {
			MessageLines string `xml:"MessageLines"`
		} `xml:"DynamicMessage"`
	} `xml:"OrderHeaderInformation"`
	OrderLine struct {
		ProductLine struct {
			SKU                string `xml:"SKU"`
			Quantity           string `xml:"Quantity"`
			CustomerLineNumber string `xml:"CustomerLineNumber"`
		} `xml:"ProductLine"`
		CommentLine struct {
			CommentText string `xml:"CommentText"`
		} `xml:"CommentLine"`
	} `xml:"OrderLineInformation"`
	ShowDetail string `xml:"ShowDetail"`
}

type OrderResponse struct {
	XMLName xml.Name `xml:"OrderResponse"`
	Version string   `xml:"Version"`
	Header  struct {
		SenderID    string `xml:"SenderID"`
		ReceiverID  string `xml:"ReceiverID"`
		ErrorStatus struct {
			Text        string `xml:",chardata"`
			ErrorNumber string `xml:"ErrorNumber,attr"`
		} `xml:"ErrorStatus"`
		DocumentID    string `xml:"DocumentID"`
		TransactionID string `xml:"TransactionID"`
		TimeStamp     string `xml:"TimeStamp"`
	} `xml:"TransactionHeader"`
	OrderInfo struct {
		OrderNumbers struct {
			BranchOrderNumber        string `xml:"BranchOrderNumber"`
			CustomerPO               string `xml:"CustomerPO"`
			ThirdPartyFreightAccount string `xml:"ThirdPartyFreightAccount"`
			ShipToAddress1           string `xml:"ShipToAddress1"`
			ShipToAddress2           string `xml:"ShipToAddress2"`
			ShipToAddress3           string `xml:"ShipToAddress3"`
			ShipToCity               string `xml:"ShipToCity"`
			ShipToProvince           string `xml:"ShipToProvince"`
			ShipToPostalCode         string `xml:"ShipToPostalCode"`
			AddressErrorMessage      struct {
				AddressErrorType string `xml:"AddressErrorType,attr"`
			} `xml:"AddressErrorMessage"`
			ContractNumber string `xml:"ContractNumber"`
			OrderSuffix    struct {
				Suffix              string `xml:"Suffix,attr"`
				DistributionWeight  string `xml:"DistributionWeight"`
				SuffixErrorResponse struct {
					SuffixErrorType string `xml:"SuffixErrorType,attr"`
				} `xml:"SuffixErrorResponse"`
				Carrier struct {
					CarrierCode string `xml:"CarrierCode,attr"`
				} `xml:"Carrier"`
				LineInfo struct {
					ProductLine struct {
						LineError           string `xml:"LineError"`
						SKU                 string `xml:"SKU"`
						MfrPartNumber       string `xml:"ManufacturerPartNumber"`
						UPC                 string `xml:"UPC"`
						UnitPrice           string `xml:"UnitPrice"`
						IngramLineNumber    string `xml:"IngramLineNumber"`
						CustomerLineNumber  string `xml:"CustomerLineNumber"`
						ShipFromBranch      string `xml:"ShipFromBranch"`
						OrderQuantity       string `xml:"OrderQuantity"`
						AllocatedQuantity   string `xml:"AllocatedQuantity"`
						BackOrderedQuantity string `xml:"BackOrderedQuantity"`
						BackOrderETADate    string `xml:"BackOrderETADate"`
						PriceDerivedFlag    string `xml:"PriceDerivedFlag"`
						ForeignCurrency     string `xml:"ForeignCurrency"`
						FreightRate         string `xml:"FreightRate"`
						TransitDays         string `xml:"TransitDays"`
						LineBillToSuffix    string `xml:"LineBillToSuffix"`
					} `xml:"ProductLine"`
				} `xml:"LineInformation"`
			} `xml:"OrderSuffix"`
		} `xml:"OrderNumbers"`
	} `xml:"OrderInfo"`
}
