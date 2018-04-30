// design a password generator CLI
// should do the following:
// 1) be able to specify the number of characters (6+)
// 2) specify the character sets involved and the number of each
// e.g. 12 chars, 3 numbers, 3 punctuation
// print password to stdout

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	// "math/rand"
	"os"
	"strings"
)

const alphabetEnglish string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// next step = multiple charset support
const cyrillic string = "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"
const numbers string = "0123456789"
const symbols string = `~!@#$%^&*()_+={}[]|:";<>,.?\'\`

// defaults
type settings struct {
	length int
	nums   int
	sym    int
}

func main() {
	numInput := flag.Int("num", 2, "numbers used in password")
	symInput := flag.Int("sym", 2, "symbols used in password")
	lenInput := flag.Int("len", 12, "total length of password")

	flag.Parse()
	// defaults
	config := settings{length: 12, nums: 2, sym: 2}

	if *numInput < 0 || *symInput < 0 || *lenInput < 0 {
		fmt.Println("No setting should be less than zero.")
		os.Exit(1)
	}

	if (*numInput + *symInput) > *lenInput {
		fmt.Printf("Character selections (%v) should not be greater than length (%v) specified.\n",
			*numInput+*symInput, *lenInput)
	}

	if *lenInput < 6 {
		fmt.Println("Using minimum length of 6.")
		*lenInput = 6
	}

	config.length = *lenInput
	config.nums = *numInput
	config.sym = *symInput

	alpha := config.length - config.nums - config.sym
	var pwd string
	if alpha > 0 {
		pwd = pwd + construct(alphabetEnglish, alpha)
	}

	if config.nums > 0 {
		pwd = pwd + construct(numbers, config.nums)
	}

	if config.sym > 0 {
		pwd = pwd + construct(symbols, config.sym)
	}

	fmt.Printf("%+v\n", shuffledString(pwd))

}

// eep : first attempt at crypto/rand
func construct(charset string, length int) string {
	var chunk string
	for x := 0; x < length; x++ {
		loc := getBetterRand(len(charset))
		chunk += charset[loc : loc+1]
	}
	return chunk
}

func shuffledString(pwd string) string {
	shuffle := strings.Split(pwd, "")
	for i := 0; i < len(shuffle)-1; i++ {
		loc := getBetterRand(len(shuffle))
		shuffle[i], shuffle[loc] = shuffle[loc], shuffle[i]
	}
	return strings.Join(shuffle, "")
}

// math/rand version
// func getBetterRand(max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(max)
// }

// first attempt at crypto/rand
func getBetterRand(max int) int {
	pBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Randomization error:", err)
		os.Exit(1)
	}
	return int(pBig.Int64())

}
