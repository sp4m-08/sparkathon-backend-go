package models

type Product struct {
	ID                   uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	EAN                  string   `json:"ean"` // Barcode
	Title                string   `json:"title"`
	Description          string   `json:"description"`
	ELID                 string   `json:"el_id"` // Possibly an external listing ID
	Brand                string   `json:"brand"`
	Model                string   `json:"model"`
	Color                string   `json:"color"`
	Size                 string   `json:"size"`
	Dimension            string   `json:"dimension"`
	Weight               string   `json:"weight"`
	Category             string   `json:"category"`
	LowestRecordedPrice  float64  `json:"lowest_recorded_price"`
	HighestRecordedPrice float64  `json:"highest_recorded_price"`
	Images               []string `json:"images" gorm:"-"`
	Offers               []string `json:"offers" gorm:"-"`
}

type UPCItemDBResponse struct {
	Code   string    `json:"code"`
	Total  int       `json:"total"`
	Offset int       `json:"offset"`
	Items  []Product `json:"items"`
}
