package template

import (
	"archive/zip"
	"bufio"
	"errors"
	"go-docx-kit/docx"
	"io"
	"os"
	"strings"
)

type TemplateKit struct {
	sourceDocx *docx.DocxFile
	model      TemplateModel
}

func NewTemplateKit() *TemplateKit {
	return new(TemplateKit)
}

func (t *TemplateKit) SetTemplateDocx(sourceDocx *docx.DocxFile) {
	t.sourceDocx = sourceDocx
}

func (t *TemplateKit) SetTemplateModel(model TemplateModel) {
	t.model = model
}

func (t *TemplateKit) Write(targetFile *os.File) (int, error) {
	if t.sourceDocx == nil {
		return 0, errors.New("source docx file is not set")
	}
	if targetFile == nil {
		return 0, errors.New("target file can not be nil")
	}
	totalWritten := 0

	zipWriter := zip.NewWriter(targetFile)
	defer zipWriter.Close()

	for _, file := range t.sourceDocx.Files {
		fileToZip, err := zipWriter.Create(file.Name)
		if err != nil {
			return totalWritten, err
		}

		if file.Name == "word/document.xml" {
			written, err := t.applyModelToDocumentXml(&t.model, fileToZip)
			if err != nil {
				return totalWritten, err
			}
			totalWritten += written
			continue
		}

		written, err := writeZipToWriter(file, fileToZip)
		if err != nil {
			return totalWritten, err
		}
		totalWritten += written
	}

	return totalWritten, nil
}

func (t *TemplateKit) applyModelToDocumentXml(model *TemplateModel, resultFile io.Writer) (int, error) {
	totalWritten := 0

	documentReader, err := t.sourceDocx.WordDirectory.DocumentXml.Open()
	if err != nil {
		return totalWritten, err
	}
	defer documentReader.Close()

	reader := bufio.NewReader(documentReader)
	curlyBracketCount := 0
	originalStringBuilder := strings.Builder{}
	isCheckingTree := false
	rootToken := model.GetTrie()
	currentToken := rootToken
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return totalWritten, err
		}

		if b == '<' {
			isCheckingTree = false
		}

		if b == '>' {
			isCheckingTree = true
		}

		if isCheckingTree {
			if _, ok := currentToken.NextCharacterTokens[b]; ok {
				currentToken = currentToken.NextCharacterTokens[b]
			}
		}

		if b == '{' && curlyBracketCount <= 2 {
			curlyBracketCount++

			if curlyBracketCount > 2 {
				popBytes := t.popFirstBracketTexts(&originalStringBuilder)

				written, err := resultFile.Write(popBytes)
				if err != nil {
					return totalWritten, err
				}
				totalWritten += written

				curlyBracketCount--
				continue
			}

			err = originalStringBuilder.WriteByte(b)
			if err != nil {
					return totalWritten, err
			}
			continue
		}

		if b == '}' && curlyBracketCount > 0 {
			curlyBracketCount--

			err = originalStringBuilder.WriteByte(b)
			if err != nil {
				return totalWritten, err
			}

			if curlyBracketCount == 0 {
				if len(currentToken.Value) > 0 {
					written, err := resultFile.Write(currentToken.Value)
					if err != nil {
						return totalWritten, err
					}
					totalWritten += written
				} else {
					written, err := resultFile.Write([]byte(originalStringBuilder.String()))
					if err != nil {
						return totalWritten, err
					}
					totalWritten += written
				}

				originalStringBuilder.Reset()
				currentToken = rootToken
			}
			continue
		}

		if curlyBracketCount > 0 {
			err = originalStringBuilder.WriteByte(b)
			if err != nil {
				return totalWritten, err
			}
			continue
		}

		written, err := resultFile.Write([]byte{b})
		if err != nil {
			return totalWritten, err
		}
		totalWritten += written

		currentToken = rootToken
	}

	return totalWritten, nil
}

func (t *TemplateKit) popFirstBracketTexts(originalStringBuilder *strings.Builder) []byte {
	currentBytes := originalStringBuilder.String()
	originalStringBuilder.Reset()
	popStringBuilder := strings.Builder{}
	bracketCount := 0

	for _, char := range currentBytes {
		if char == '{' {
			bracketCount++
		}

		if bracketCount <= 1 {
			popStringBuilder.WriteByte(byte(char))
		} else {
			originalStringBuilder.WriteByte(byte(char))
		}
	}

	return []byte(popStringBuilder.String())
}

func writeZipToWriter(zip *zip.File, writer io.Writer) (int, error) {
	reader, err := zip.Open()
	if err != nil {
		return 0, err
	}
	defer reader.Close()

	written, err := io.Copy(writer, reader)
	if err != nil {
		return int(written), err
	}

	return int(written), nil
}
