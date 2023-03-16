# Go DOCX Template Kit
A Go Library provides tools for manipulate .docx file. using template literal to generate .docx file based on provided data model.

## TemplateKit
Using double curly bracket template literal `{{key}}` to specify the position which will be replaced by provided data.

Put `{{key}}` in your .docx file.  
![image](./img/word_template.png)

See it in action:  
```go
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
kit.SetTargetDocx(templateFile)
kit.SetTemplateModel(dataModel)
_, err = kit.Generate(resultFile)
if err != nil {
    panic(err)
}
```
- Create a new `TemplateModel` and put each value with a key which corresponding to the template literal key in the template file.
- Read the template file by using `docx.ReadDocxFile("to/template/path")`.
- Create a `*os.File` as the output .docx file.
- Create `TemplateKit` by `NewTemplateKit()`
    - Set the template resource file.
    - Set the model.
    - Generate the file which applied with the model data.

The result .docx file will output like this:
![image](./img/word_result.png)