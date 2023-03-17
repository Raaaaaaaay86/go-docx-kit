package template

import (
	"go-docx-kit/docx"
	"os"
	"testing"
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

	templateFile, err := docx.ReadDocxFile("../../template.docx")
	if err != nil {
		panic(err)
	}

	resultFile, err := os.Create("../../result.docx")
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	kit := NewTemplateKit()
	kit.SetTemplateDocx(templateFile)
	kit.SetTemplateModel(dataModel)
	_, err = kit.Generate(resultFile)
	if err != nil {
		panic(err)
	}
}
