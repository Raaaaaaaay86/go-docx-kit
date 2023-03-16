package docx

import "archive/zip"

type DocxWordDirectory struct {
	RelsDirectory  DocxWordRelsDirectory
	ThemeDirectory DocxWordThemeDirectory
	DocumentXml    *zip.File
	EndNotesXml    *zip.File
	FontTableXml   *zip.File
	FootNotesXml   *zip.File
	NumberingXml   *zip.File
	SettingsXml    *zip.File
	StylesXml      *zip.File
	WebSettingsXml *zip.File
}
