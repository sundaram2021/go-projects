// package admin
// import (
//     "platform/http/actionresults"
//     "sportsstore/models"
// )
// type CategoriesHandler struct {
//     models.Repository
// }
// func (handler CategoriesHandler) GetData() string {
//     return "This is the categories handler"
// }
// func (handler CategoriesHandler) GetSelect(current int) actionresults.ActionResult {
//     return actionresults.NewTemplateAction("select_category.html", struct {
//         Current int
//         Categories []models.Category
//     }{ Current: current, Categories: handler.GetCategories()})
// }

package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"sportsstore/models"
)

type CategoriesHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}
type CategoryTemplateContext struct {
	Categories []models.Category
	EditId     int
	EditUrl    string
	SaveUrl    string
}

const CATEGORY_EDIT_KEY string = "category_edit"

func (handler CategoriesHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_categories.html",
		CategoryTemplateContext{
			Categories: handler.Repository.GetCategories(),
			EditId:     handler.Session.GetValueDefault(CATEGORY_EDIT_KEY, 0).(int),
			EditUrl: mustGenerateUrl(handler.URLGenerator,
				CategoriesHandler.PostCategoryEdit),
			SaveUrl: mustGenerateUrl(handler.URLGenerator,
				CategoriesHandler.PostCategorySave),
		})
}
func (handler CategoriesHandler) PostCategoryEdit(ref EditReference) actionresults.ActionResult {
	handler.Session.SetValue(CATEGORY_EDIT_KEY, ref.ID)
	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator,
		AdminHandler.GetSection, "Categories"))
}
func (handler CategoriesHandler) PostCategorySave(
	c models.Category) actionresults.ActionResult {
	handler.Repository.SaveCategory(&c)
	handler.Session.SetValue(CATEGORY_EDIT_KEY, 0)
	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator,
		AdminHandler.GetSection, "Categories"))
}
func (handler CategoriesHandler) GetSelect(current int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("select_category.html", struct {
		Current    int
		Categories []models.Category
	}{Current: current, Categories: handler.GetCategories()})
}
