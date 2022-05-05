package infra

import (
	"fmt"
	"time"
)

const apiBaseUrl = "https://edinet-proxy.tk/"
const corporations = "api/corporations/"
const documents = "api/documents/"

type IHttp interface {
	Get(url string, queries map[string]string) ([]byte, error)
}

type EdiNet struct {
	IHttp
}

type CorporationQuery struct {
	Name         string `json:"name"`
	CapitalGte   int64  `json:"capital_gte"`
	CapitalIte   int64  `json:"capital_ite"`
	Location     string `json:"location"`
	ListingType  *bool  `json:"listing_type"`
	IndustryType string `json:"industry_type"`
	Page         int64  `json:"page"`
}

func (cq CorporationQuery) toHashMap() map[string]string {
	m := map[string]string{}
	if cq.Name != "" {
		m["name"] = cq.Name
	}
	if cq.CapitalGte != 0 {
		m["capital_gte"] = fmt.Sprintf("%d", cq.CapitalGte)
	}
	if cq.CapitalIte != 0 {
		m["capital_ite"] = fmt.Sprintf("%d", cq.CapitalIte)
	}
	if cq.Location != "" {
		m["location"] = cq.Location
	}
	if cq.ListingType != nil {
		lt := "2"
		if *cq.ListingType {
			lt = "1"
		}
		m["listing_type"] = lt
	}
	if cq.IndustryType != "" {
		m["industry_type"] = cq.IndustryType
	}
	if cq.Page != 0 {
		m["page"] = fmt.Sprintf("%d", cq.Page)
	}
	return m
}

type DocumentQuery struct {
	EdiNetCode  string     `json:"edinet_code"`
	SecCode     string     `json:"sec_code"`
	Jcn         string     `json:"jcn"`
	ListingType *bool      `json:"listing_type"`
	Date        *time.Time `json:"date"`
	Page        int64      `json:"page"`
}

func (dq DocumentQuery) toHashMap() map[string]string {
	m := map[string]string{}
	if dq.EdiNetCode != "" {
		m["edinet_code"] = dq.EdiNetCode
	}
	if dq.SecCode != "" {
		m["sec_code"] = dq.SecCode
	}
	if dq.Jcn != "" {
		m["jcn"] = dq.Jcn
	}
	if dq.ListingType != nil {
		lt := "2"
		if *dq.ListingType {
			lt = "1"
		}
		m["listing_type"] = lt
	}
	if dq.Date != nil {
		m["date"] = dq.Date.Format("2006-06-21")
	}
	if dq.Page != 0 {
		m["page"] = fmt.Sprintf("%d", dq.Page)
	}
	return m
}

func NewEdiNet(h IHttp) *EdiNet {
	return &EdiNet{h}
}

func (e EdiNet) SearchCorporations(c CorporationQuery) {
	b, _ := e.IHttp.Get(apiBaseUrl+corporations, c.toHashMap())
	fmt.Println(string(b))
}

func (e EdiNet) SearchDocuments(d DocumentQuery) {
	b, _ := e.IHttp.Get(apiBaseUrl+documents, d.toHashMap())
	fmt.Println(string(b))
}
