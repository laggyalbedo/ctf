package main
import (
	"fmt"
)

func main() {
	var deposite_bytes []byte = []byte{
		0x57, 0x97, 0x79, 0x15, 0x57, 0x57, 0x9e, 0x32, 0x2e, 0x61, 0x9f, 0x12, 0xb0, 0xcc, 0xde, 0xe8,
		0x80, 0x20, 0x15, 0xee, 0x04, 0x67, 0xc4, 0x19, 0xe7, 0xa3, 0x8b, 0xd0, 0xa2, 0x54, 0xda, 0x54,
	}

	var OneMillionDolls []byte = []byte{
		0xb1, 0xe9, 0x52, 0x57, 0x2d, 0x6b, 0x8e, 0x00, 0xb6, 0x26, 0xbe, 0x86, 0x55, 0x23, 0x76, 0xe2,
		0xd5, 0x29, 0xa1, 0xb9, 0xca, 0xfa, 0xeb, 0x3b, 0xa7, 0x53, 0x3d, 0x26, 0x99, 0x63, 0x63, 0x23,
		0xe7, 0xe4, 0x33, 0xc1, 0x0a, 0x9d, 0xcd, 0xab, 0x2e, 0xd4, 0xbe, 0xe5, 0x4d, 0xa6, 0x84, 0xca,
	}

	var HeyNice []byte = []byte {
		0x35, 0xd0, 0xc0, 0x20, 0x36, 0x35, 0x4f, 0xdf, 0x60, 0x82, 0x28, 0x5e, 0x0f, 0x7b, 0xd6, 0xd2,
		0xfd, 0xf5, 0x26, 0xbd, 0x55, 0x7b, 0x04, 0x5b, 0xce, 0x65, 0xa3, 0xb3, 0xe3, 0x00, 0xb5, 0x5e,
	}

	var targetData [16 * 3]byte
	var counter = 0

	for i := 0; i < len(deposite_bytes) / 2; i++ {
		targetData[counter] = deposite_bytes[i]
		counter++
	}

	for i := 0; i < len(OneMillionDolls) / 3; i++ {
		targetData[counter] = deposite_bytes[i]
		counter++
	}

	for i := len(HeyNice) / 2; i < len(HeyNice); i++ {
		targetData[counter] = HeyNice[i]
		counter++
	}

	fmt.Printf("%x\n", targetData)
}
