package docx

import "archive/zip"

type DocxDocPropsDirectory struct {
	AppXml    *zip.File
	CoreXml   *zip.File
	CustomXml *zip.File
}