package helpers

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"strings"
)

func HashFile(path string) string {
	// Read the required file into a string,
	val, err := ioutil.ReadFile(path)
	fmt.Println(err)

	// Trim all leading and tailing spaces.
	s := strings.TrimSpace(string(val[:]))

	crc32q := crc32.MakeTable(0xEDB88320)
	h := (uint32)(crc32.Checksum([]byte(s), crc32q))
	hash := fmt.Sprintf("%X", h)

	return hash
}
