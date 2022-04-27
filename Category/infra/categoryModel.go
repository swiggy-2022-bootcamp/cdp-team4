package infra

type CategoryDescriptionModel struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	MetaTitle       string `json:"meta_title"`
}

type CategoryModel struct {
	Id                  string                     `json:"id"`
	CategoryDescription []CategoryDescriptionModel `json:"category_description"`
}
