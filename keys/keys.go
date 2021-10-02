package keys

import "crypto/rand"

type Generator struct {
	length int
}

func New(length int) *Generator {
	return &Generator{
		length: length,
	}
}

// Generate a random key
func (g *Generator) Generate() ([]byte, error) {
	key := make([]byte, g.length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
