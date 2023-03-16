package template

type TemplateModel struct {
	hashMap map[string]string
}

func NewTemplateModel() TemplateModel {
	return TemplateModel{
		hashMap: make(map[string]string),
	}
}

func (t *TemplateModel) Put(key string, value string) {
	t.hashMap[key] = value
}

func (t *TemplateModel) Get(key string) string {
	return t.hashMap[key]
}

func (t *TemplateModel) GetTrie() *CharacterToken {
	rootToken := NewCharacterTokenTrie()

	for key, value := range t.hashMap {
		t.insertTrie(rootToken, key, value)
	}

	return rootToken
}

func (t *TemplateModel) insertTrie(token *CharacterToken, key, value string) {
	if len(key) == 0 {
		token.Value = []byte(value)
		return
	}

	currentByte := key[0]
	if _, ok := token.NextCharacterTokens[currentByte]; !ok {
		token.NextCharacterTokens[currentByte] = NewCharacterToken()
	}

	t.insertTrie(token.NextCharacterTokens[currentByte], key[1:], value)
}
