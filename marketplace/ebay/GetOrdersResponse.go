package ebay

type GetOrdersResponse struct {
	XMLName          xml.Name `xml:"GetOrdersResponse"`
	Xmlns            string   `xml:"xmlns,attr"`
	Timestamp        string   `xml:"Timestamp"`
	Ack              string   `xml:"Ack"`
	Version          string   `xml:"Version"`
	Build            string   `xml:"Build"`
	PaginationResult struct {
		TotalNumberOfPages   string `xml:"TotalNumberOfPages"`
		TotalNumberOfEntries string `xml:"TotalNumberOfEntries"`
	} `xml:"PaginationResult"`
	HasMoreOrders string `xml:"HasMoreOrders"`
	OrderArray    struct {
		Order []struct {
			OrderID          string `xml:"OrderID"`
			OrderStatus      string `xml:"OrderStatus"`
			AdjustmentAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"AdjustmentAmount"`
			AmountPaid struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"AmountPaid"`
			AmountSaved struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"AmountSaved"`
			CheckoutStatus struct {
				EBayPaymentStatus                   string `xml:"eBayPaymentStatus"`
				LastModifiedTime                    string `xml:"LastModifiedTime"`
				PaymentMethod                       string `xml:"PaymentMethod"`
				Status                              string `xml:"Status"`
				IntegratedMerchantCreditCardEnabled string `xml:"IntegratedMerchantCreditCardEnabled"`
				PaymentInstrument                   string `xml:"PaymentInstrument"`
			} `xml:"CheckoutStatus"`
			ShippingDetails struct {
				SalesTax struct {
					SalesTaxPercent       string `xml:"SalesTaxPercent"`
					SalesTaxState         string `xml:"SalesTaxState"`
					ShippingIncludedInTax string `xml:"ShippingIncludedInTax"`
					SalesTaxAmount        struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"SalesTaxAmount"`
				} `xml:"SalesTax"`
				ShippingServiceOptions []struct {
					ShippingService         string `xml:"ShippingService"`
					ShippingServicePriority string `xml:"ShippingServicePriority"`
					ExpeditedService        string `xml:"ExpeditedService"`
					ShippingTimeMin         string `xml:"ShippingTimeMin"`
					ShippingTimeMax         string `xml:"ShippingTimeMax"`
				} `xml:"ShippingServiceOptions"`
				InternationalShippingServiceOption []struct {
					ShippingService         string `xml:"ShippingService"`
					ShippingServicePriority string `xml:"ShippingServicePriority"`
					ShipToLocation          string `xml:"ShipToLocation"`
				} `xml:"InternationalShippingServiceOption"`
				SellingManagerSalesRecordNumber string `xml:"SellingManagerSalesRecordNumber"`
				GetItFast                       string `xml:"GetItFast"`
			} `xml:"ShippingDetails"`
			CreatedTime     string `xml:"CreatedTime"`
			PaymentMethods  string `xml:"PaymentMethods"`
			SellerEmail     string `xml:"SellerEmail"`
			ShippingAddress struct {
				Name              string `xml:"Name"`
				Street1           string `xml:"Street1"`
				Street2           string `xml:"Street2"`
				CityName          string `xml:"CityName"`
				StateOrProvince   string `xml:"StateOrProvince"`
				Country           string `xml:"Country"`
				CountryName       string `xml:"CountryName"`
				Phone             string `xml:"Phone"`
				PostalCode        string `xml:"PostalCode"`
				AddressID         string `xml:"AddressID"`
				AddressOwner      string `xml:"AddressOwner"`
				ExternalAddressID string `xml:"ExternalAddressID"`
			} `xml:"ShippingAddress"`
			ShippingServiceSelected struct {
				ShippingService     string `xml:"ShippingService"`
				ShippingServiceCost struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"ShippingServiceCost"`
			} `xml:"ShippingServiceSelected"`
			Subtotal struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"Subtotal"`
			Total struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"Total"`
			TransactionArray struct {
				Transaction []struct {
					Buyer struct {
						Email         string `xml:"Email"`
						UserFirstName string `xml:"UserFirstName"`
						UserLastName  string `xml:"UserLastName"`
					} `xml:"Buyer"`
					ShippingDetails struct {
						CalculatedShippingRate struct {
							OriginatingPostalCode string `xml:"OriginatingPostalCode"`
							PackageDepth          struct {
								Text              string `xml:",chardata"`
								MeasurementSystem string `xml:"measurementSystem,attr"`
								Unit              string `xml:"unit,attr"`
							} `xml:"PackageDepth"`
							PackageLength struct {
								Text              string `xml:",chardata"`
								MeasurementSystem string `xml:"measurementSystem,attr"`
								Unit              string `xml:"unit,attr"`
							} `xml:"PackageLength"`
							PackageWidth struct {
								Text              string `xml:",chardata"`
								MeasurementSystem string `xml:"measurementSystem,attr"`
								Unit              string `xml:"unit,attr"`
							} `xml:"PackageWidth"`
							PackagingHandlingCosts struct {
								Text       string `xml:",chardata"`
								CurrencyID string `xml:"currencyID,attr"`
							} `xml:"PackagingHandlingCosts"`
							ShippingIrregular string `xml:"ShippingIrregular"`
							ShippingPackage   string `xml:"ShippingPackage"`
							WeightMajor       struct {
								Text              string `xml:",chardata"`
								MeasurementSystem string `xml:"measurementSystem,attr"`
								Unit              string `xml:"unit,attr"`
							} `xml:"WeightMajor"`
							WeightMinor struct {
								Text              string `xml:",chardata"`
								MeasurementSystem string `xml:"measurementSystem,attr"`
								Unit              string `xml:"unit,attr"`
							} `xml:"WeightMinor"`
						} `xml:"CalculatedShippingRate"`
						SellingManagerSalesRecordNumber string `xml:"SellingManagerSalesRecordNumber"`
					} `xml:"ShippingDetails"`
					CreatedDate string `xml:"CreatedDate"`
					Item        []struct {
						ItemID               string `xml:"ItemID"`
						Site                 string `xml:"Site"`
						Title                string `xml:"Title"`
						SKU                  string `xml:"SKU"`
						ConditionID          string `xml:"ConditionID"`
						ConditionDisplayName string `xml:"ConditionDisplayName"`
					} `xml:"Item"`
					QuantityPurchased string `xml:"QuantityPurchased"`
					Status            struct {
						PaymentHoldStatus string `xml:"PaymentHoldStatus"`
						InquiryStatus     string `xml:"InquiryStatus"`
						ReturnStatus      string `xml:"ReturnStatus"`
					} `xml:"Status"`
					TransactionID    string `xml:"TransactionID"`
					TransactionPrice struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"TransactionPrice"`
					ShippingServiceSelected struct {
						ShippingPackageInfo struct {
							Text                     string `xml:",chardata"`
							EstimatedDeliveryTimeMin string `xml:"EstimatedDeliveryTimeMin"`
							EstimatedDeliveryTimeMax string `xml:"EstimatedDeliveryTimeMax"`
						} `xml:"ShippingPackageInfo"`
					} `xml:"ShippingServiceSelected"`
					TransactionSiteID string `xml:"TransactionSiteID"`
					Platform          string `xml:"Platform"`
					Variation         struct {
						SKU                string `xml:"SKU"`
						VariationSpecifics struct {
							NameValueList struct {
								Name  string `xml:"Name"`
								Value string `xml:"Value"`
							} `xml:"NameValueList"`
						} `xml:"VariationSpecifics"`
						VariationTitle       string `xml:"VariationTitle"`
						VariationViewItemURL string `xml:"VariationViewItemURL"`
					} `xml:"Variation"`
					Taxes struct {
						TotalTaxAmount struct {
							Text       string `xml:",chardata"`
							CurrencyID string `xml:"currencyID,attr"`
						} `xml:"TotalTaxAmount"`
						TaxDetails struct {
							Text           string `xml:",chardata"`
							Imposition     string `xml:"Imposition"`
							TaxDescription string `xml:"TaxDescription"`
							TaxAmount      struct {
								Text       string `xml:",chardata"`
								CurrencyID string `xml:"currencyID,attr"`
							} `xml:"TaxAmount"`
							TaxOnSubtotalAmount struct {
								Text       string `xml:",chardata"`
								CurrencyID string `xml:"currencyID,attr"`
							} `xml:"TaxOnSubtotalAmount"`
							TaxOnShippingAmount struct {
								Text       string `xml:",chardata"`
								CurrencyID string `xml:"currencyID,attr"`
							} `xml:"TaxOnShippingAmount"`
							TaxOnHandlingAmount struct {
								Text       string `xml:",chardata"`
								CurrencyID string `xml:"currencyID,attr"`
							} `xml:"TaxOnHandlingAmount"`
						} `xml:"TaxDetails"`
					} `xml:"Taxes"`
					ActualShippingCost struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"ActualShippingCost"`
					ActualHandlingCost struct {
						Text       string `xml:",chardata"`
						CurrencyID string `xml:"currencyID,attr"`
					} `xml:"ActualHandlingCost"`
					OrderLineItemID     string `xml:"OrderLineItemID"`
					ExtendedOrderID     string `xml:"ExtendedOrderID"`
					EBayPlusTransaction string `xml:"eBayPlusTransaction"`
				} `xml:"Transaction"`
			} `xml:"TransactionArray"`
			BuyerUserID                         string `xml:"BuyerUserID"`
			PaidTime                            string `xml:"PaidTime"`
			IntegratedMerchantCreditCardEnabled string `xml:"IntegratedMerchantCreditCardEnabled"`
			EIASToken                           string `xml:"EIASToken"`
			PaymentHoldStatus                   string `xml:"PaymentHoldStatus"`
			IsMultiLegShipping                  string `xml:"IsMultiLegShipping"`
			SellerUserID                        string `xml:"SellerUserID"`
			SellerEIASToken                     string `xml:"SellerEIASToken"`
			CancelStatus                        string `xml:"CancelStatus"`
			ExtendedOrderID                     string `xml:"ExtendedOrderID"`
			ContainseBayPlusTransaction         string `xml:"ContainseBayPlusTransaction"`
		} `xml:"Order"`
	} `xml:"OrderArray"`
	OrdersPerPage            string `xml:"OrdersPerPage"`
	PageNumber               string `xml:"PageNumber"`
	ReturnedOrderCountActual string `xml:"ReturnedOrderCountActual"`
}
