package godocxkit

import "archive/zip"

type DocxFile struct {
	RelsDirectory DocxRelsDirectory 
	DocPropsDirectory DocxDocPropsDirectory
	WordDirectory DocxWordDirectory
	ContentTypesXml DocxContentTypesXml
}

type DocxContentTypesXml *zip.File

type DocxWordDirectory struct {
	RelsDirectory DocxWordRelsDirectory
	ThemeDirectory DocxWordThemeDirectory
	DocumentXml *zip.File
	EndNotesXml *zip.File
	FontTableXml *zip.File
	FootNotesXml *zip.File
	NumberingXml *zip.File
	SettingsXml *zip.File
	StylesXml *zip.File
	WebSettingsXml *zip.File
}

type DocxWordRelsDirectory struct {
	DocumentXmlRels *zip.File
}

type DocxWordThemeDirectory struct {
	Theme1Xml *zip.File
}

type DocxDocPropsDirectory struct {
	AppXml *zip.File
	CoreXml *zip.File
	CustomXml *zip.File
}

type DocxRelsDirectory struct {
	Rels *zip.File
}

func NewDocxFile() *DocxFile {
	return new(DocxFile)
}

func ReadDocxFile(filePath string) (*DocxFile, error) {
	zipReader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()

	docx := NewDocxFile()

	for _, file := range zipReader.File {
		switch file.Name {
		case "word/document.xml":
			docx.WordDirectory.DocumentXml = file
		case "word/endnotes.xml":
			docx.WordDirectory.EndNotesXml = file
		case "word/fontTable.xml":
			docx.WordDirectory.FontTableXml = file
		case "word/numbering.xml":
			docx.WordDirectory.NumberingXml = file
		case "word/settings.xml":
			docx.WordDirectory.SettingsXml = file
		case "word/styles.xml":
			docx.WordDirectory.StylesXml = file
		case "word/webSettings.xml":
			docx.WordDirectory.WebSettingsXml = file
		case "word/_rels/document.xml.rels":
			docx.WordDirectory.RelsDirectory.DocumentXmlRels = file
		case "word/theme/theme1.xml":
			docx.WordDirectory.ThemeDirectory.Theme1Xml = file
		case "docProps/app.xml":
			docx.DocPropsDirectory.AppXml = file
		case "docProps/core.xml":
			docx.DocPropsDirectory.CoreXml = file
		case "docProps/custom.xml":
			docx.DocPropsDirectory.CustomXml = file
		case "_rels/.rels":
			docx.RelsDirectory.Rels = file
		case "[Content_Types].xml":
			docx.ContentTypesXml = file
		}
	}

	return docx, nil
}
