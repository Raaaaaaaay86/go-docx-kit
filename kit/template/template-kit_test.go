package template

import (
	"github.com/raaaaaaaay86/go-docx-kit/docx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateGeneration(t *testing.T) {
	dataModel := NewTemplateModel()
	dataModel.Put("receiver", "RD Team")
	dataModel.Put("refNumber", "40666888")
	dataModel.Put("reporterName", "Kent Chou")
	dataModel.Put("reporterRank", "Intern")
	dataModel.Put("departmentName", "QA")
	dataModel.Put("telephone", "0911222333")
	dataModel.Put("reportDate", "2023/03/15")
	dataModel.Put("productName", "Pencil X")
	dataModel.Put("manufacturer", "Tesla")
	dataModel.Put("batchNumber", "777")
	dataModel.Put("quantity", "900")
	dataModel.Put("natureOfComplain", "The pencil is too short.")
	dataModel.Put("isSampleRetained", "■")
	dataModel.Put("noSampleRetained", "□")
	dataModel.Put("retainedCount", "900")
	dataModel.Put("reasonOfNotRetainingSample", " ")

	docxTemplate, err := docx.ReadDocxFile("../../template.docx")
	assert.NoError(t, err)

	err = NewTemplateKit().Render(docxTemplate, dataModel)
	assert.NoError(t, err)

	_, err = docxTemplate.ToFile("../../output.docx")
	assert.NoError(t, err)
}
