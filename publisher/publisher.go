package publisher

// Publisher contains a single publisher name and its total sales
type Publisher struct {
	Publisher string  `json:"publisher"`
	Sales     float32 `json:"sales"`
}
