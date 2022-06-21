package parcelux

type TrackParams struct {
	TCode    string `json:"t_code"`
	TInvoice string `json:"t_invoice"`
}

type TrackResp struct {
	Complete  bool   `json:"complete"`
	Estimate  string `json:"estimate"`
	InvoiceNo string `json:"invoiceNo"`
}
