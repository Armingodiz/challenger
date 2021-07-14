package cache

import "github.com/ArminGodiz/golang-code-challenge/pkg/models"

// it is a mock cache to test combiner

type mockCache struct {
}

func (m *mockCache) Get(string2 string) string {
	return string2
}
func (m *mockCache) Set(data models.CacheData) error {
	return nil
}

func (m *mockCache) InitializeCache(port int) error {
	return nil
}
func GetMockCache() *mockCache {
	return &mockCache{}
}
