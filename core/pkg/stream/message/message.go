package message

import "encoding/json"

type HomeMessage struct {
	SiteKey string `json:"siteKey"`
	Name    string `json:"name"`
}

type SiteCatalogHomeMessage struct {
	SiteKey     string `json:"siteKey"`
	Name        string `json:"name"`
	CatalogLink string `json:"catalogLink"`
}

type SiteCatalogPageMessage struct {
	SiteKey         string `json:"siteKey"`
	Page            int    `json:"page"`
	CatalogPageLink string `json:"catalogPageLink"`
}

type NovelMessage struct {
	SiteKey     string `json:"siteKey"`
	CatalogId   string `json:"catalogId"`
	CatalogName string `json:"catalogName"`
	Name        string `json:"name"`
	NovelLink   string `json:"bookLink"`
}

type ChapterMessage struct {
	SiteKey   string `json:"siteKey"`
	CatalogId string `json:"catalogId"`
	Name      string `json:"name"`
	BookLink  string `json:"bookLink"`
}

func (scm SiteCatalogHomeMessage) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(scm)
	return bytes, err
}

func (scm SiteCatalogHomeMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &scm)
}
