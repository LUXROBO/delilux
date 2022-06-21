package client

type TrackParcelParams struct {
	TKey     string `json:"t_key"`
	TCode    string `json:"t_code"`
	TInvoice string `json:"t_invoice"`
}

type TrackParcelResp struct {
	Complete  bool   `json:"complete"`
	Estimate  string `json:"estimate"`
	InvoiceNo string `json:"invoiceNo"`
}
