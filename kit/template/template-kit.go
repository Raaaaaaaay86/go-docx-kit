package template

import (
	"bytes"
	"errors"
	"go-docx-kit/docx"
	"strings"
)

type TemplateKit struct {
}

func NewTemplateKit() *TemplateKit {
	return new(TemplateKit)
}

func (t *TemplateKit) Render(templateDocx *docx.DocxFile, model TemplateModel) error {
	if templateDocx == nil {
		return errors.New("template docx file cannot be nil")
	}

	buffer := bytes.NewBuffer(nil)
	reader := bytes.NewReader(templateDocx.WordDirectory.DocumentXml.Data)
	originalStringBuilder := strings.Builder{}
	rootToken := model.GetTrie()
	currentToken := rootToken
	curlyBracketCount := 0
	isCheckingTree := false
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
		} else if b == '>' {
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

				_, err := buffer.Write(popBytes)
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
					_, err := buffer.Write(currentToken.Value)
					if err != nil {
						return err
					}
				} else {
					_, err := buffer.Write([]byte(originalStringBuilder.String()))
					if err != nil {
						return err
					}
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

		_, err = buffer.Write([]byte{b})
		if err != nil {
			return err
		}

		currentToken = rootToken
	}

	templateDocx.WordDirectory.DocumentXml.Data = buffer.Bytes()

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
