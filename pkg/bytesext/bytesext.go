package bytesext

import (
	"math"
)

func Split(content []byte, size int) [][]byte {
	totalSize := len(content)
	count := int(math.Ceil(float64(totalSize) / float64(size)))
	var chunks [][]byte
	for i := 0; i < count; i++ {
		var bytes []byte
		if i+1 == count {
			bytes = content[i*size:]
		} else {
			bytes = content[i*size : size+i*size]
		}
		chunks = append(chunks, bytes)
	}
	return chunks
}
