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
			Author: "Somix",
			Path: "./ciao",
		},
	}
}