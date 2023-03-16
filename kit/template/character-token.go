package template

import (
	"encoding/json"
	"fmt"
)

type CharacterToken struct {
	NextCharacterTokens map[byte]*CharacterToken
	Value          []byte
}

func NewCharacterToken() *CharacterToken {
	token := &CharacterToken{
		NextCharacterTokens: map[byte]*CharacterToken{},
		Value:               []byte{},
	}

	return token
}

func NewCharacterTokenTrie() *CharacterToken {
	token := NewCharacterToken()
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, char := range chars {
		token.NextCharacterTokens[byte(char)] = NewCharacterToken()
	}

	return token
}

func (t *CharacterToken) PrintTree() {
	data, _ := json.MarshalIndent(t, "", "  ")
	fmt.Println(string(data))
}