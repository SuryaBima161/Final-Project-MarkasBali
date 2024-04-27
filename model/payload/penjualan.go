package payload

type ItemPenjualanRequest struct {
	Kode_Barang string `json:"kode_barang"`
	Jumlah      uint   `json:"jumlah"`
	Subtotal    uint   `json:"subtotal"`
}

type AddPenjualanRequest struct {
	Nama_Pembeli   string                 `json:"nama_pembeli" valid:"required, type(string)"`
	Subtotal       float64                `json:"subtotal" valid:"required"`
	Kode_Diskon    string                 `json:"kode_diskon"`
	Total          float64                `json:"total" valid:"optional , type(float64)"`
	CreatedBy      string                 `json:"created_by" valid:"required, type(string)"`
	Item_Penjualan []ItemPenjualanRequest `json:"item_penjualan"`
}

type GetPenjualanRespone struct {
	ID           uint    `json:"id"`
	Kode_Invoice string  `json:"kode_invoice"`
	Nama_Pembeli string  `json:"nama_pembeli"`
	Subtotal     float64 `json:"subtotal"`
	Kode_Diskon  string  `json:"kode_diskon"`
	Total        float64 `json:"total"`
	Created_at   string  `json:"created_at"`
	Updated_at   string  `json:"updated_at"`
	Deleted_at   string  `json:"deleted_at"`
	Created_By   string  `json:"created_by"`
}
