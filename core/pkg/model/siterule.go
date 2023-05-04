package model

type SiteRule struct {
	Key  string `mapstructure:"key"`
	Rule *Rule  `mapstructure:"rule"`
}

type CatalogRule struct {
	Enabled     bool `mapstructure:"enabled"`
	AutoCreated bool `mapstructure:"autoCreated"`
}

type Rule struct {
	Home    *HomeRule    `mapstructure:"home"`
	Novel   *NovelRule   `mapstructure:"novel"`
	Chapter *ChapterRule `mapstructure:"chapter"`
}

type HomeRule struct {
	CatalogTitleSelector string `mapstructure:"catalog_title_selector"`
	CatalogUrlAttr       string `mapstructure:"catalog_url_attr"`
	CatalogTitleName     string `mapstructure:"catalog_title_name"`
	CatalogPageSelector  string `mapstructure:"catalog_page_selector"`
}

type NovelRule struct {
	NovelUrlSelector              string `mapstructure:"novel_url_selector"`
	NovelUrlAttr                  string `mapstructure:"novel_url_attr"`
	NovelTitleSelector            string `mapstructure:"novel_title_selector"`
	NovelAuthorSelector           string `mapstructure:"novel_author_selector"`
	NovelCateSelector             string `mapstructure:"novel_cate_selector"`
	NovelDescSelector             string `mapstructure:"novel_desc_selector"`
	NovelCoverSelector            string `mapstructure:"novel_cover_selector"`
	NovelNoCover                  string `mapstructure:"novel_no_cover"`
	NovelChapterURLSelector       string `mapstructure:"novel_chapter_url_selector"`
	NovelLastChapterTitleSelector string `mapstructure:"novel_last_chapter_title_selector"`
}

type ChapterRule struct {
	ChapterCatalogSelector string `mapstructure:"chapter_catalog_selector"`
	ChapterAbandonNum      int    `mapstructure:"chapter_abandon_num"`
	InfoTitleSelector      string `mapstructure:"info_title_selector"`
	InfoDescSelector       string `mapstructure:"info_desc_selector"`
	InfoDescFilter         string `mapstructure:"info_desc_filter"`
	InfoPrePageSelector    string `mapstructure:"info_pre_page_selector"`
	InfoNextPageSelector   string `mapstructure:"info_next_page_selector"`
}

func (c *SiteRule) Validate() error {
	return nil
}

func (c *SiteRule) Complete() error {
	return nil
}
