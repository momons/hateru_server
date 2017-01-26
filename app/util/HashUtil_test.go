package util

import (
	. "github.com/r7kamura/gospel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashUtil(t *testing.T) {

	Describe(t, "hash", func() {
		Context("ハッシュ化テスト", func() {
			password := "123456789"
			saltHeader := "*"
			saltFooter := "#"
			// 1回目
			hash1 := Hash(password, saltHeader, saltFooter)
			// 2回目
			hash2 := Hash(password, saltHeader, saltFooter)
			assert.Equal(t, hash1, hash2)
		})
	})
}
