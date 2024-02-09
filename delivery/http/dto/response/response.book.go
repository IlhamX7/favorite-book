package response

type ResponseBookDTO struct {
	ID       int    `json:"id"`
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
	Rating   int    `json:"rating"`
}
