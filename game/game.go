package game

// Game model all data about a game and its sales
type Game struct {
	Rank        int     `json:"rank"`
	Name        string  `json:"name"`
	Platform    string  `json:"platform"`
	Year        string  `json:"year"`
	Genre       string  `json:"genre"`
	Publisher   string  `json:"publisher"`
	NASales     float32 `json:"naSales"`
	EUSales     float32 `json:"euSales"`
	JPSales     float32 `json:"jpSales"`
	OtherSales  float32 `json:"otherSales"`
	GlobalSales float32 `json:"globalSales"`
}
