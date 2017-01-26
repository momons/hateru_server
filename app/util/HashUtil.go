package util

import (
	"crypto/sha512"
	"fmt"
)

// ハッシュ化.
func Hash(inStr string, saltHeader string, saltFooter string) string {

	outStr := inStr
	// スクラッチ10回
	for i := 0; i < 10; i++ {
		outStr = fmt.Sprintf("%x", sha512.Sum512([]byte(saltHeader+outStr+saltFooter)))
	}

	return outStr
}
