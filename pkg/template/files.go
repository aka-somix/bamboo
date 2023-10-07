package template

import (
	"fmt"

	"github.com/aka-somix/bamboo/pkg/utils"
)

var baseFolder = "./"

type FilesPacker struct {
	path string
}

func (fp FilesPacker) PackAndUpload(name string) error {

	zipFilePath := fmt.Sprintf("%s/%s.zip", baseFolder, name)

	utils.ZipFolder(fp.path, zipFilePath)

	return nil

}

func (fp FilesPacker) DownloadAndUnpack() {

}
