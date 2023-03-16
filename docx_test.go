package godocxkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDocxFile(t *testing.T) {
	docx, err := ReadDocxFile("template.docx")

	assert.NoError(t, err)
	assert.Equal(t, "word/document.xml", docx.WordDirectory.DocumentXml)
	assert.Equal(t, "word/endnotes.xml", docx.WordDirectory.EndNotesXml)
	assert.Equal(t, "word/fontTable.xml", docx.WordDirectory.FontTableXml)
	assert.Equal(t, "word/numbering.xml", docx.WordDirectory.NumberingXml)
	assert.Equal(t, "word/settings.xml", docx.WordDirectory.SettingsXml)
	assert.Equal(t, "word/styles.xml", docx.WordDirectory.StylesXml)
	assert.Equal(t, "word/webSettings.xml", docx.WordDirectory.WebSettingsXml)
	assert.Equal(t, "word/_rels/document.xml.rels", docx.WordDirectory.RelsDirectory.DocumentXmlRels)
	assert.Equal(t, "word/theme/theme1.xml", docx.WordDirectory.ThemeDirectory.Theme1Xml)
	assert.Equal(t, "docProps/app.xml", docx.DocPropsDirectory.AppXml)
	assert.Equal(t, "docProps/core.xml", docx.DocPropsDirectory.CoreXml)
	assert.Equal(t, "docProps/custom.xml", docx.DocPropsDirectory.CustomXml)
	assert.Equal(t, "rels_/.rels", docx.RelsDirectory.Rels)
	assert.Equal(t, "[Content_Types].xml", docx.ContentTypesXml)
}