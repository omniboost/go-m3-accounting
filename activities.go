package m3

type Activity struct {
	BusinessDate      Date        `json:"businessDate"`
	PropertyCode      string      `json:"propertyCode"`
	CustomerShortName string      `json:"customerShortNamek"`
	SourceSystem      string      `json:"sourceSystem"`
	DataRecords       DataRecords `json:"dataRecords"`
}

type DataRecords []DataRecord

type DataRecord struct {
	AccountDescription string  `json:"accountDescription"`
	AccountingID       string  `json:"accountingId"`
	Amount             float64 `json:"amount"`
	RecordType         string  `json:"recordType"`
}
