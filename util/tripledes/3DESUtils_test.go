package tripledes

import (
	"fmt"
	"testing"
)

func TestEncryptToStringWithEnv(t *testing.T) {
	str := EncryptToStringWithEnv("test_string")
	fmt.Println(str)
}
