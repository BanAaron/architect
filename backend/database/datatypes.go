package database

// KeywordDocument is the document used in the keyword collection
type KeywordDocument struct {
	Keyword string   `bson:"keyword"`
	Tags    []string `bson:"tags"`
}

// NewKeywordDocument creates a new KeywordDocument
func NewKeywordDocument(keyword string, tags []string) *KeywordDocument {
	return &KeywordDocument{Keyword: keyword, Tags: tags}
}

func RemoveTags(document KeywordDocument) {
	document.Tags = []string{}
}
