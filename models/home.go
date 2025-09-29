package models

// Home represents the localized content that powers the landing page.
//
// The struct keeps field names that match the historic template so that we can
// restore the classic design while still supporting the newer document
// highlights section.
type Home struct {
	Home1H1       string                `json:"slide1H1"`
	Home1H3       string                `json:"slide1H3"`
	Home2H1       string                `json:"slide2H1"`
	Home2H3       string                `json:"slide2H2"`
	Home3H1       string                `json:"slide3H1"`
	Home3H3       string                `json:"slide3H2"`
	MapHeader     string                `json:"mapHeader"`
	MapContent    string                `json:"mapContent"`
	MapEmbedURL   string                `json:"mapEmbedUrl"`
	MapEmbedNote  string                `json:"mapEmbedNote"`
	MapSourceURL  string                `json:"mapContentSource"`
	MapListTitle  string                `json:"mapListTitle"`
	MapListIntro  string                `json:"mapListIntro"`
	MapPlacements []PrefecturePlacement `json:"mapPlacements"`

	HeroTitle       string      `json:"heroTitle"`
	HeroSubtitle    string      `json:"heroSubtitle"`
	HeroButtonLabel string      `json:"heroButtonLabel"`
	HeroSupport     string      `json:"heroSupport"`
	DocURL          string      `json:"docUrl"`
	DocEmbedURL     string      `json:"docEmbedUrl"`
	DocNote         string      `json:"docNote"`
	Highlights      []Highlight `json:"highlights"`
}

// Highlight represents a short value proposition displayed beneath the hero section.
type Highlight struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// PrefecturePlacement stores the number of participants placed in a Japanese prefecture.
type PrefecturePlacement struct {
	Prefecture string `json:"prefecture"`
	Count      int    `json:"count"`
}
