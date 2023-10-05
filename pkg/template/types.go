package template

type Template struct {
	Author    string
	Name string
	Description string
	Path string
}

type TemplateManager interface {
	ListTemplatesInfo(author string) []Template
	GetTemplateInfo(author string, name string) Template
	DownloadTemplate(author string, name string, folderPath string) error
	CreateTemplate(t Template, sourcePath string) error
}
