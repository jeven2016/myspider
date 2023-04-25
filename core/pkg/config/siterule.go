package config

type SiteRule struct {
	Key         string      `mapstructure:"key"`
	Name        string      `mapstructure:"name"`
	Description string      `mapstructure:"description"`
	HomeURL     string      `mapstructure:"homeUrl"`
	Catalog     CatalogRule `mapstructure:"catalog"`
	Book        BookRule    `mapstructure:"book"`
}

type CatalogRule struct {
	Enabled     bool        `mapstructure:"enabled"`
	AutoCreated bool        `mapstructure:"autoCreated"`
	Rule        interface{} `mapstructure:"rule"`
}

type Home struct {
	CatalogTitleSelector string `mapstructure:"catalog_title_selector"`
	CatalogURL           string `mapstructure:"catalog_url"`
}

type BookRule struct {
	BookTitleSelector            string `mapstructure:"book_title_selector"`
	BookAuthorSelector           string `mapstructure:"book_author_selector"`
	BookCateSelector             string `mapstructure:"book_cate_selector"`
	BookDescSelector             string `mapstructure:"book_desc_selector"`
	BookCoverSelector            string `mapstructure:"book_cover_selector"`
	BookNoCover                  string `mapstructure:"book_no_cover"`
	BookChapterURLSelector       string `mapstructure:"book_chapter_url_selector"`
	BookLastChapterTitleSelector string `mapstructure:"book_last_chapter_title_selector"`
}

type Rule struct {
	Home                   Home        `mapstructure:"home"`
	Book                   BookRule    `mapstructure:"book"`
	Chapter                interface{} `mapstructure:"chapter"`
	ChapterCatalogSelector string      `mapstructure:"chapter_catalog_selector"`
	ChapterAbandonNum      int         `mapstructure:"chapter_abandon_num"`
	InfoTitleSelector      string      `mapstructure:"info_title_selector"`
	InfoDescSelector       string      `mapstructure:"info_desc_selector"`
	InfoDescFilter         string      `mapstructure:"info_desc_filter"`
	InfoPrePageSelector    string      `mapstructure:"info_pre_page_selector"`
	InfoNextPageSelector   string      `mapstructure:"info_next_page_selector"`
}

type Book struct {
	Enabled bool `mapstructure:"enabled"`
	Rule    Rule `mapstructure:"rule"`
}

func (c *SiteRule) Validate() error {
	return nil
}

func (c *SiteRule) Complete() error {
	return nil
}
