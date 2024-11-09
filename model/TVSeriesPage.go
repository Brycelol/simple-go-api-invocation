package model

type TVSeries struct {
	Name              string  `json:"name"`
	RuntimeOfSeries   string  `json:"runtime_of_series"`
	Certificate       string  `json:"certificate"`
	RuntimeOfEpisodes string  `json:"runtime_of_episodes"`
	Genre             string  `json:"genre"`
	IMDBRating        float32 `json:"imdb_rating"`
	Overview          string  `json:"overview"`
	NumberOfVotes     int64   `json:"no_of_votes"`
	Id                int     `json:"id"`
}

type TVSeriesPage struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []TVSeries `json:"data"`
}

func NewTVSeriesPage() *TVSeriesPage {
	return &TVSeriesPage{}
}
