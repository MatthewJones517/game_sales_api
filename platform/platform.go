package platform

// Platform contains the platform name and total sales for a given game platform
type Platform struct {
	Platform string  `json:"platform"`
	Sales    float32 `json:"sales"`
}
