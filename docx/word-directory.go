package docx

type DocxWordDirectory struct {
	RelsDirectory  DocxWordRelsDirectory
	ThemeDirectory DocxWordThemeDirectory
	DocumentXml    *DocxZipFile
	EndNotesXml    *DocxZipFile
	FontTableXml   *DocxZipFile
	FootNotesXml   *DocxZipFile
	NumberingXml   *DocxZipFile
	SettingsXml    *DocxZipFile
	StylesXml      *DocxZipFile
	WebSettingsXml *DocxZipFile
}
