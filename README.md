# Go DOCX Template Kit
A Go Library provides tools for manipulate .docx file. using template literal to generate .docx file based on provided data model.

## TemplateKit
Using double curly bracket template literal `{{key}}` to specify the position which will be replaced by provided data.

Put `{{key}}` in your .docx file.  
![image](./img/word_template.png)

See it in action:  
```go
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