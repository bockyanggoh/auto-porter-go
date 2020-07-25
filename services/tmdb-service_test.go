package services

import "testing"

func TestGetRecordFromTheMovieDB(t *testing.T) {
	if res, err := GetRecordFromTheMovieDB("The Martian"); err != nil {
		t.Errorf("Failed to get record from The Movie DB Api. %v\n", err)
	} else {
		if res.Id == 0 {
			t.Errorf("ID should not be 0: %v", res.Id)
		}

		if len(res.MediaType) == 0 {
			t.Errorf("MediaType should not be empty: %v", res.MediaType)
		}

		if len(res.OriginalLanguage) == 0 {
			t.Errorf("OriginalLanguage should not be empty: %v", res.OriginalLanguage)
		}

		if len(res.OriginalTitle) == 0 {
			t.Errorf("OriginalTitle should not be empty: %v", res.OriginalLanguage)
		}

		if len(res.Title) == 0 {
			t.Errorf("Title should not be empty: %v", res.OriginalLanguage)
		}

		if len(res.ReleaseDate) == 0 {
			t.Errorf("ReleaseDate should not be empty: %v", res.OriginalLanguage)
		}

	}
}
