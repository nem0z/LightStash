package wallet

import (
	"os"
	"sort"

	"github.com/nem0z/LightStash/utils"
)

const mnemonicFile string = "words.txt"
const mnemonicLength int = 12

func GenerateMnemonic() ([]string, error) {
	file, err := os.Open(mnemonicFile)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	count, err := utils.CountLines(file)
	if err != nil {
		return []string{}, err
	}

	indexs := utils.RandSlice(mnemonicLength, count)
	sort.Ints(indexs)

	words, err := utils.ReadLines(file, indexs...)
	if err != nil {
		return []string{}, err
	}

	utils.Shuffle(words)
	return words, nil
}
