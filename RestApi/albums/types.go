package albumspkg

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumRequest struct {
	Title string
}

func (a *Album) setTitle(title string) {
	(*a).Title = title
}
