package utils

import (
	"bastion/utils/humanize"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAbsDir(relPath string) (string, error) {
	dir, e := os.Getwd()
	if e != nil {
		return "", nil
	}
	configPath := filepath.Join(dir, relPath)
	return configPath + "/", nil
}

func GetAbsFile(relPath string) (string, error) {
	dir, e := os.Getwd()
	if e != nil {
		return "", nil
	}
	configPath := filepath.Join(dir, relPath)
	return configPath, nil
}

func GetAbsFileWithEnv(relPath string) string {
	env := os.Getenv("BASTION")
	if env == "" {
		panic("env BASTION not found")
	}
	configPath := filepath.Join(env, relPath)
	return configPath
}

func FileName() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

func DirName() string {
	return filepath.Dir(FileName())
}

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func PastFromNow(sometime string) (string, error) {
	Loc, _ := time.LoadLocation(viper.GetString("base.time_location"))
	st, err := time.ParseInLocation("2006-01-02 15:04:05", sometime, Loc)
	if err != nil {
		return "", err
	}
	s := humanize.TimeZh(st)
	return strings.ReplaceAll(s, " ", ""), nil
}

func FmtDate(sometime time.Time) string {
	res := sometime.Format("2006-01-02 15:04:05")
	return res
}

func ParseHtml(s string) template.HTML {
	return template.HTML(s)
}

func PrintJson(v interface{}) ([]byte, error) {
	bytes, e := json.Marshal(v)
	if e != nil {
		return nil, e
	}
	return bytes, nil
}

func Print(v interface{}) {
	fmt.Printf("%v \n", v)
}

func PrintJsonString(v interface{}) {
	bytes, e := json.Marshal(v)
	if e != nil {
		fmt.Printf("%s \n", e)
	}
	fmt.Printf("%s \n", bytes)
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

func IsDev() bool {
	return viper.GetString("base.env") == "dev"
}

func IsProd() bool {
	return viper.GetString("base.env") == "prod"
}
