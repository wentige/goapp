package asi

import (
	"encoding/xml"
)

type Branch struct {
	Code string `xml:"Code,attr"`
	Name string `xml:"Name,attr"`
	Qty  string `xml:",chardata"`
}

type BranchQty struct {
	Branches []Branch `xml:"Branch"`
}

type Item struct {
	SKU         string    `xml:"SKU,attr"`
	ItemId      string    `xml:"ItemId"`
	Description string    `xml:"Description"`
	Vendor      string    `xml:"Vendor"`
	Category    string    `xml:"Category"`
	SubCategory string    `xml:"SubCategory"`
	UPC         string    `xml:"UPC"`
	Price       string    `xml:"Price"`
	Rebate      string    `xml:"Rebate"`
	Term        string    `xml:"Term"`
	Weight      string    `xml:"Weight"`
	Status      string    `xml:"Status"`
	Qty         BranchQty `xml:"Qty"`
}

type Error struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

type ASIInventory struct {
	XMLName  xml.Name `xml:"ASIInventory"`
	DateTime string   `xml:"request,attr"`
	Excute   string   `xml:"excute,attr"`
	Items    []Item   `xml:"Inventory"`
	Error    Error    `xml:"error"`
}
