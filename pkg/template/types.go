package template

type Template struct {
	Author    string
	Name string
	Path string
}

type TemplateManager interface {
	// ListTemplatesNames(author string) []string
	ListTemplatesInfo(author string) []Template
	// GetTemplateInfo(author string, name string) Template
	// DownloadTemplate(author string, folderPath string) 
	// CreateTemplate(t Template, sourcePath string)
}
