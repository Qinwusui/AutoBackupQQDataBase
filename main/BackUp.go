package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
)

type Config struct {
	CopyTime  string   `json:"copyTime"`
	FromPath  string   `json:"fromPath"`
	ToPath    string   `json:"toPath"`
	FileNames []string `json:"fileNames"`
}

const configPath = "/sdcard/Android/BackConfig.json"

var config Config
var backPath = "/sdcard/Download/BackUps/"
var fromPath = "/data/data/com.tencent.mobileqq/databases/"

func init() {
	readConfig()
}

// 读取Config
func readConfig() {
	isExist, _ := PathExists(configPath)
	cmd := exec.Command("su", "-c", "mkdir", backPath)
	cmd.Run()
	if !isExist {
		config := Config{
			CopyTime: "08:00:00",
			FromPath: fromPath,
			ToPath:   backPath,
			FileNames: []string{
				"fileName.db",
			},
		}
		bytes, _ := json.MarshalIndent(&config, "", "\t")
		cmd := exec.Command("su", "-c", "touch", configPath)
		cmd.Run()
		os.WriteFile(configPath, bytes, 0777)
	}
	configBytes, _ := os.ReadFile(configPath)
	e := json.Unmarshal(configBytes, &config)
	checkErr(e)
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 主函数，执行定时任务
func main() {
	cronTab := gocron.NewScheduler(time.FixedZone("UTC", +8*60*60))
	cronTab.Every(1).Day().At(config.CopyTime).Do(func() {
		readConfig()
		backUp()
	})
	cronTab.StartBlocking()
}

func backUp() {
	for _, v := range config.FileNames {
		cmd := exec.Command("su", "-c", "cp", config.FromPath+v,
			config.ToPath+v)
		e := cmd.Run()
		checkErr(e)
	}
}
func checkErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
