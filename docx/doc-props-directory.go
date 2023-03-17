package docx

type DocxDocPropsDirectory struct {
	AppXml    *DocxZipFile
	CoreXml   *DocxZipFile
	CustomXml *DocxZipFile
}
