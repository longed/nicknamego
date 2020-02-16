package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"math/rand"
)

type CharacterSet struct {
	Numbers            []string
	UpperCaseChars     []string
	LowerCaseChars     []string
	Symbols            []string
	UserSpecifiedChars []string
	OperateCharsSet    []string
}

type UserOptions struct {
	WantedNumber       bool   `toml:"wantedNumber"`
	WantedUpperCase    bool   `toml:"wantedUpperCase"`
	WantedLowerCase    bool   `toml:"wantedLowerCase"`
	WantedSymbol       bool   `toml:"wantedSymbol"`
	SaveNickNameToFile bool   `toml:"saveNickNameToFile"`
	NickNameLen        int    `toml:"nickNameLen"`
	BatchNumber        int    `toml:"batchNumber"`
	SpecifiedChars     string `toml:"specifiedChars"`
}

var (
	characterSet CharacterSet
	userOptions  UserOptions
)

const (
	ResultFileSizeUpperLimit int64 = 30 * 1024 * 1024
	ResultFileName                 = "nickname.pw.txt"
	ConfigFileName                 = "config.toml"
)

func init() {
	initCharacterSet()
	initUserOptions()
}

func main() {

	/* read flag param
	init characterSet
	process */

	out, err := json.Marshal(characterSet)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	fmt.Println("hello nicknamego.")
}

// set defalut value to characterSet
func initCharacterSet() {
	characterSet.Numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	characterSet.UpperCaseChars = []string{"A", "B", "C", "D", "E", "F", "G", "H",
		"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z"}
	characterSet.LowerCaseChars = []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w",
		"x", "y", "z"}
	characterSet.Symbols = []string{"!", "@", "#", "$", "%", "*", "(",
		")", "-", "=", "+", "."}
}

// set default value to userOptions
func initUserOptions() {
	userOptions.WantedNumber = true
	userOptions.WantedUpperCase = true
	userOptions.WantedLowerCase = true
	userOptions.WantedSymbol = true
	userOptions.SaveNickNameToFile = false
}

// generate nickname
func nickname(userOptions UserOptions) {
	if len(userOptions.SpecifiedChars) > 0 {
		// move specified chars to characterSet
	} else {
		composeOperateCharsSet(userOptions)
	}

	if userOptions.SaveNickNameToFile {

	}
}

func randomChars(userOptions UserOptions) {

}

// plain random algorithm
func plainRandom(userOptions UserOptions) []string {
	nnList := []string{}
	rand.Seed(time.Now().Unix())
	for batchNumber := 0; batchNumber < userOptions.BatchNumber; batchNumber++ {
		nn := ""
		for i := 0; i < userOptions.NickNameLen; i++ {
			nn = nn + characterSet.OperateCharsSet[
				rand.Intn(len(characterSet.OperateCharsSet))]
		}
		nnList = append(nnList, nn)
	}
	return nnList
}

// mersenne twister algorithm
func mersenneTwisterRandom(userOptions UserOptions) {

}

// save content to file
func saveContent(content string) error {
	executableDirPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	resultFile := executableDirPath + string(os.PathSeparator) + ResultFileName
	fi, err := os.Stat(resultFile)
	// file exist && size > ResultFileSizeUpperLimit back up result file
	if !os.IsNotExist(err) && fi.Size() > ResultFileSizeUpperLimit {
		os.Rename(resultFile, resultFile+".bak") // TODO: same file name repalce bug
	}
	writeToFile(resultFile, content)
	return nil
}

// write content to file
// if file is exist then append content to file
// if file is no exist then create new file then write to file
func writeToFile(path, content string) (bool, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		return false, err
	}
	return true, nil
}

func composeOperateCharsSet(userOptions UserOptions) {
	if userOptions.WantedNumber {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.Numbers...)
	}

	if userOptions.WantedUpperCase {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.UpperCaseChars...)
	}

	if userOptions.WantedLowerCase {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.LowerCaseChars...)
	}

	if userOptions.WantedSymbol {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.Symbols...)
	}
}
