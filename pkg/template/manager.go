// Templates Manager
// ------------------
// TODO aka-somix: Add documentation

package template


type templateManagerImpl struct {}


func NewTemplateManager() TemplateManager {
	return &templateManagerImpl{}
}

func (tm *templateManagerImpl) ListTemplatesInfo(author string) []Template {
	return []Template{
		{
			Name: "Test",
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		},
		{
			Name: "Test2",
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		},
	}
}

func (tm *templateManagerImpl) GetTemplateInfo(author string, name string) Template {
	return Template {
			Name: name,
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		}
}