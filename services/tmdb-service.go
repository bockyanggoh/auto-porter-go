package services

import (
	"auto-porter-go/models"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

func GetRecordFromTheMovieDB(name string) (*models.TMDBResultData, error) {
	baseUrl, _ := url.Parse("https://api.themoviedb.org/3/search/multi")

	params := url.Values{}
	params.Add("api_key", os.Getenv("TMDB_API_KEY"));
	params.Add("query", name)
	baseUrl.RawQuery = params.Encode()

	if resp, err := http.Get(baseUrl.String()); err != nil {
		return nil, err
	} else {
		data := models.TMDBResponseData{}
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&data); err != nil {
			return nil, err
		}
		record := data.Results[0]
		return &record, nil
	}
}