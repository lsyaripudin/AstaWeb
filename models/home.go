package models

type Home struct {
	Home1H1    string `json:"slide1H1"`
	Home1H3    string `json:"slide1H3"`
	Home2H1    string `json:"slide2H1"`
	Home2H3    string `json:"slide2H2"`
	Home3H1    string `json:"slide3H1"`
	Home3H3    string `json:"slide3H2"`
	MapHeader  string `json:"mapHeader"`
	MapContent string `json:"mapContent"`
}

type GoogleSheetResponse struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}
