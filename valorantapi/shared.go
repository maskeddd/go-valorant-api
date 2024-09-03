package valorantapi

type ShopData struct {
	Cost              int           `json:"cost"`
	Category          string        `json:"category"`
	ShopOrderPriority int           `json:"shopOrderPriority"`
	CategoryText      string        `json:"categoryText"`
	GridPosition      *GridPosition `json:"gridPosition"`
	CanBeTrashed      bool          `json:"canBeTrashed"`
	Image             *string       `json:"image"`
	NewImage          string        `json:"newImage"`
	NewImage2         *string       `json:"newImage2"`
	AssetPath         string        `json:"assetPath"`
}
