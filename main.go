package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CharacterSet struct {
	Numbers            []string
	UpperCaseChars     []string
	LowerCaseChars     []string
	Symbols            []string
	UserSpecifiedChars []string
	OperateCharsSet    []string
}

type Config struct {
	User   UserOptions   `toml:"UserOptions"`
	Server ServerOptions `toml:"ServerOptions"`
}

type ServerOptions struct {
	ServerIpPort string `toml:"serverIpPort"`
}

type UserOptions struct {
	WantNumber         bool   `toml:"wantNumber"`
	WantUpperCase      bool   `toml:"wantUpperCase"`
	WantLowerCase      bool   `toml:"wantLowerCase"`
	WantSymbol         bool   `toml:"wantSymbol"`
	SaveNickNameToFile bool   `toml:"saveNickNameToFile"`
	NickNameLen        int    `toml:"nickNameLen"`
	BatchNumber        int    `toml:"batchNumber"`
	SpecifiedChars     string `toml:"specifiedChars"`
}

var (
	config       Config
	characterSet CharacterSet
	userOptions  UserOptions
)

const (
	ResultFileSizeUpperLimit int64 = 30 * 1024 * 1024
	ResultFileName                 = "nickname.txt"
	ConfigFileName                 = "config.toml"
)

func init() {
	initCharacterSet()
	loadConfigAndInit()
}

func main() {
	welcome()
	nickname(userOptions)
	startServer(config.Server.ServerIpPort)
}

// print welcome message
func welcome() {
	fmt.Println("*** Welcome to nickname ! ***")
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

// init userOptions
func loadConfigAndInit() {
	// sort by priority, left priority < right priority
	pathsToFind := []string{"./", "../conf/"}
	found := false
	for _, path := range pathsToFind {
		if _, err := toml.DecodeFile(path+ConfigFileName, &config); err != nil {
			fmt.Println(err)
		} else {
			found = true
		}
	}
	if !found {
		panic("read config error. exit.")
	}
	userOptions = config.User
}

// generate nickname
func nickname(userOptions UserOptions) ApiData {
	composeOperateCharsSet(userOptions)

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	content := []string{}
	content = append(content, plainRandom(userOptions)...)
	result := format(content)
	if userOptions.SaveNickNameToFile {
		saveContent(result)
	}
	fmt.Println(timeStr + " " + result)
	apiData := defaultApiData()
	apiData.Date = timeStr
	apiData.Nn = result
	return apiData
}

// format content
func format(content []string) string {
	line := ""
	for _, item := range content {
		line = line + " " + item
	}
	line = strings.TrimSpace(line)
	line = line + "\n"
	return line
}

// plain random algorithm
func plainRandom(userOptions UserOptions) []string {
	nnList := []string{}
	rand.Seed(time.Now().Unix())
	for batchNumber := 0; batchNumber < userOptions.BatchNumber; batchNumber++ {
		nn := ""
		for i := 0; i < userOptions.NickNameLen; i++ {
			if slen := len(characterSet.OperateCharsSet); slen > 0 {
				nn = nn + characterSet.OperateCharsSet[rand.Intn(slen)]
			}
		}
		nnList = append(nnList, nn)
	}
	return nnList
}

// mersenne twister algorithm
// TODO
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

// compose operate character set by userOptions
func composeOperateCharsSet(userOptions UserOptions) {
	if len(userOptions.SpecifiedChars) > 0 {
		for _, i := range userOptions.SpecifiedChars {
			characterSet.OperateCharsSet = append(characterSet.OperateCharsSet, string(i))
		}
		return
	}

	if userOptions.WantNumber {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.Numbers...)
	}

	if userOptions.WantUpperCase {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.UpperCaseChars...)
	}

	if userOptions.WantLowerCase {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.LowerCaseChars...)
	}

	if userOptions.WantSymbol {
		characterSet.OperateCharsSet = append(characterSet.OperateCharsSet,
			characterSet.Symbols...)
	}
}
