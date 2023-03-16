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