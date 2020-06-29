package genre

// Genre model holds a single genre and its total sales
type Genre struct {
	Genre string  `json:"genre"`
	Sales float32 `json:"sales"`
}
