package docx

import (
	"archive/zip"
)

type DocxFile struct {
	RelsDirectory     DocxRelsDirectory
	DocPropsDirectory DocxDocPropsDirectory
	WordDirectory     DocxWordDirectory
	ContentTypesXml   DocxContentTypesXml
	Files             []*DocxZipFile
}

func newDocxFile() *DocxFile {
	return new(DocxFile)
}

func ReadDocxFile(filePath string) (*DocxFile, error) {
	zipReader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}

	docx := newDocxFile()

	for _, file := range zipReader.File {
		docxZipFile, err := NewDocxZipFile(file)
		if err != nil {
			return nil, err
		}

		docx.Files = append(docx.Files, docxZipFile)

		switch file.Name {
		case "word/document.xml":
			docx.WordDirectory.DocumentXml = docxZipFile
		case "word/endnotes.xml":
			docx.WordDirectory.EndNotesXml = docxZipFile
		case "word/fontTable.xml":
			docx.WordDirectory.FontTableXml = docxZipFile
		case "word/footnotes.xml":
			docx.WordDirectory.FootNotesXml = docxZipFile
		case "word/numbering.xml":
			docx.WordDirectory.NumberingXml = docxZipFile
		case "word/settings.xml":
			docx.WordDirectory.SettingsXml = docxZipFile
		case "word/styles.xml":
			docx.WordDirectory.StylesXml = docxZipFile
		case "word/webSettings.xml":
			docx.WordDirectory.WebSettingsXml = docxZipFile
		case "word/_rels/document.xml.rels":
			docx.WordDirectory.RelsDirectory.DocumentXmlRels = docxZipFile
		case "word/theme/theme1.xml":
			docx.WordDirectory.ThemeDirectory.Theme1Xml = docxZipFile
		case "docProps/app.xml":
			docx.DocPropsDirectory.AppXml = docxZipFile
		case "docProps/core.xml":
			docx.DocPropsDirectory.CoreXml = docxZipFile
		case "docProps/custom.xml":
			docx.DocPropsDirectory.CustomXml = docxZipFile
		case "_rels/.rels":
			docx.RelsDirectory.Rels = docxZipFile
		case "[Content_Types].xml":
			docx.ContentTypesXml = docxZipFile
		}
	}

	return docx, nil
}
