package docx

import "archive/zip"

type DocxRelsDirectory struct {
	Rels *zip.File
}