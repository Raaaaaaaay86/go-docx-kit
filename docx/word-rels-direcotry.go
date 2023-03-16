package docx

import "archive/zip"

type DocxWordRelsDirectory struct {
	DocumentXmlRels *zip.File
}