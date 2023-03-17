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

func (t *TemplateKit) Generate(targetFile *os.File) (*os.File, error) {
	if t.sourceDocx == nil {
		return targetFile, errors.New("source docx file is not set")
	}
	if targetFile == nil {
		return targetFile, errors.New("target file can not be nil")
	}

	zipWriter := zip.NewWriter(targetFile)
	defer zipWriter.Close()

	for _, file := range t.sourceDocx.Files {
		fileToZip, err := zipWriter.Create(file.Name)
		if err != nil {
			return targetFile, err
		}

		if file.Name == "word/document.xml" {
			err := t.applyModelToDocumentXml(&t.model, fileToZip)
			if err != nil {
				return targetFile, err
			}
			continue
		}

		err = writeZipToWriter(file, fileToZip)
		if err != nil {
			return targetFile, err
		}
	}

	return targetFile, nil
}

func (t *TemplateKit) applyModelToDocumentXml(model *TemplateModel, resultFile io.Writer) error {
	documentReader, err := t.sourceDocx.WordDirectory.DocumentXml.Open()
	if err != nil {
		return err
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
			return err
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
				_, err = resultFile.Write(popBytes)
				if err != nil {
					return err
				}
				curlyBracketCount--
				continue
			}

			err = originalStringBuilder.WriteByte(b)
			if err != nil {
				return err
			}
			continue
		}

		if b == '}' && curlyBracketCount > 0 {
			curlyBracketCount--

			err = originalStringBuilder.WriteByte(b)
			if err != nil {
				return err
			}

			if curlyBracketCount == 0 {
				if len(currentToken.Value) > 0 {
					resultFile.Write(currentToken.Value)
				} else {
					resultFile.Write([]byte(originalStringBuilder.String()))
				}

				originalStringBuilder.Reset()
				currentToken = rootToken
			}
			continue
		}

		if curlyBracketCount > 0 {
			err = originalStringBuilder.WriteByte(b)
			if err != nil {
				return err
			}
			continue
		}

		_, err = resultFile.Write([]byte{b})
		if err != nil {
			return err
		}
		currentToken = rootToken
	}

	return nil
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

func writeZipToWriter(zip *zip.File, writer io.Writer) error {
	reader, err := zip.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}
