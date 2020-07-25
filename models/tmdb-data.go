package models

type TMDBResponseData struct {
	TotalResults int `json:"total_results"`
	Results []TMDBResultData `json:"results"`
}

type TMDBResultData struct {
	Id int `json:"id"`
	MediaType string `json:"media_type"`
	OriginalLanguage string `json:"original_language"`
	Title string `json:"title"`
	OriginalTitle string `json:"original_title"`
	ReleaseDate string `json:"release_date"`
}
