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
	Book    *BookRule    `mapstructure:"book"`
	Chapter *ChapterRule `mapstructure:"chapter"`
}

type HomeRule struct {
	CatalogTitleSelector string `mapstructure:"catalog_title_selector"`
	CatalogUrlAttr       string `mapstructure:"catalog_url_attr"`
	CatalogTitleName     string `mapstructure:"catalog_title_name"`
	CatalogPageSelector  string `mapstructure:"catalog_page_selector"`
}

type BookRule struct {
	BookUrlSelector              string `mapstructure:"book_url_selector"`
	BookUrlAttr                  string `mapstructure:"book_url_attr"`
	BookTitleSelector            string `mapstructure:"book_title_selector"`
	BookAuthorSelector           string `mapstructure:"book_author_selector"`
	BookCateSelector             string `mapstructure:"book_cate_selector"`
	BookDescSelector             string `mapstructure:"book_desc_selector"`
	BookCoverSelector            string `mapstructure:"book_cover_selector"`
	BookNoCover                  string `mapstructure:"book_no_cover"`
	BookChapterURLSelector       string `mapstructure:"book_chapter_url_selector"`
	BookLastChapterTitleSelector string `mapstructure:"book_last_chapter_title_selector"`
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
