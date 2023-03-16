package godocxkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDocxFile(t *testing.T) {
	docx, err := ReadDocxFile("template.docx")

	assert.NoError(t, err)
	assert.Equal(t, "word/document.xml", docx.WordDirectory.DocumentXml.Name)
	assert.Equal(t, "word/endnotes.xml", docx.WordDirectory.EndNotesXml.Name)
	assert.Equal(t, "word/fontTable.xml", docx.WordDirectory.FontTableXml.Name)
	assert.Equal(t, "word/numbering.xml", docx.WordDirectory.NumberingXml.Name)
	assert.Equal(t, "word/settings.xml", docx.WordDirectory.SettingsXml.Name)
	assert.Equal(t, "word/styles.xml", docx.WordDirectory.StylesXml.Name)
	assert.Equal(t, "word/webSettings.xml", docx.WordDirectory.WebSettingsXml.Name)
	assert.Equal(t, "word/_rels/document.xml.rels", docx.WordDirectory.RelsDirectory.DocumentXmlRels.Name)
	assert.Equal(t, "word/theme/theme1.xml", docx.WordDirectory.ThemeDirectory.Theme1Xml.Name)
	assert.Equal(t, "docProps/app.xml", docx.DocPropsDirectory.AppXml.Name)
	assert.Equal(t, "docProps/core.xml", docx.DocPropsDirectory.CoreXml.Name)
	assert.Equal(t, "docProps/custom.xml", docx.DocPropsDirectory.CustomXml.Name)
	assert.Equal(t, "_rels/.rels", docx.RelsDirectory.Rels.Name)
	assert.Equal(t, "[Content_Types].xml", docx.ContentTypesXml.Name)
}