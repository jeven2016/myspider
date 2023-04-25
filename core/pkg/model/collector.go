package model

type Collector struct {
	Rule Rule `json:"rule"`
}

// Rule 采集规则
type Rule struct {
	//主页地址
	HomeUrl string `json:"homeUrl"`

	//栏目相关

	// 小说简介页面采集规则
	BookTitleSelector            string `json:"book_title_selector"` // 小说标题
	BookTitleAttr                string `json:"book_title_attr"`     // 小说标题css选择器获取属性
	BookTitleFilter              string `json:"book_title_filter"`
	BookAuthorSelector           string `json:"book_author_selector"`
	BookAuthorAttr               string `json:"book_author_attr"`
	BookAuthorFilter             string `json:"book_author_filter"`
	BookCateSelector             string `json:"book_cate_selector"`
	BookCateAttr                 string `json:"book_cate_attr"`
	BookCateFilter               string `json:"book_cate_filter"`
	BookDescSelector             string `json:"book_desc_selector"`
	BookDescAttr                 string `json:"book_desc_attr"`
	BookDescFilter               string `json:"book_desc_filter"`
	BookCoverSelector            string `json:"book_cover_selector"`
	BookCoverAttr                string `json:"book_cover_attr"`
	BookNoCover                  string `json:"book_no_cover"`
	BookChapterURLSelector       string `json:"book_chapter_url_selector"`
	BookChapterURLAttr           string `json:"book_chapter_url_attr"`
	BookLastChapterTitleSelector string `json:"book_last_chapter_title_selector"`
	BookLastChapterTitleAttr     string `json:"book_last_chapter_title_attr"`

	// 章节目录采集规则
	ChapterCatalogSelector  string `json:"chapter_catalog_selector"`
	ChapterNextPageSelector string `json:"chapter_next_page_selector"` // 章节目录下一页选择器
	ChapterAbandonNum       int    `json:"chapter_abandon_num"`        // 丢弃章节数（最新章节）

	// 详情页面采集规则
	InfoTitleSelector    string `json:"info_title_selector"`
	InfoTitleFilter      string `json:"info_title_filter"`
	InfoDescSelector     string `json:"info_desc_selector"`
	InfoDescFilter       string `json:"info_desc_filter"`
	InfoPrePageSelector  string `json:"info_pre_page_selector"`
	InfoNextPageSelector string `json:"info_next_page_selector"`
}
