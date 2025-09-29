package models

// Home represents the localized content that powers the landing page.
type Home struct {
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
