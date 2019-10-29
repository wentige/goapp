package model

/**
 * Price & Avail response info
 */
type Branch struct {
	Name string
	Code string
	Qty  string
}

type PriceAvailItem struct {
	Sku      string
	Price    string
	Branches []Branch
}

type PriceAvailResult struct {
	Status  string
	Message string
	Items   []PriceAvailItem
}

/**
 * Purchase order request info
 */
type Buyer struct {
	Name     string
	Company  string
	Address  string
	City     string
	Province string
	Country  string
	ZipCode  string
	Phone    string
}

type PurchaseItem struct {
	Sku     string
	Price   string
	Qty     string
	Comment string
	ShipTo  Buyer
}

/**
 * Purchase order response info
 */
type PurchaseResult struct {
	Status   string
	Message  string
	PoNumber string
}
