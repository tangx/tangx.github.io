package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	Use:   "greeting",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		// greeting(name, age)

		// 1. 读取配置文件
		person := readConfig(config)
		// 2. 业务逻辑
		greeting(person.Name, person.Age)

		// 3. 转换配置文件
		dumpConfig(person)
	},
}

// var (
// 	name string
// 	age  int
// )

// func init() {
// 	root.Flags().StringVarP(&name, "name", "n", "", "名字")
// 	root.Flags().IntVarP(&age, "age", "a", 20, "名字")
// }

func greeting(name string, age int) {
	fmt.Printf("你好 %s 先生, 你今天 %d 岁\n", name, age)
}

var config string

func init() {
	root.Flags().StringVarP(&config, "config", "c", "config.yml", "配置文件")
}

type Person struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Age  int    `yaml:"age,omitempty" json:"age,omitempty"`
}

func readConfig(name string) *Person {

	person := &Person{}

	// 绑定参数
	b, err := os.ReadFile(config)
	if err != nil {
		panic(err)
	}

	err2 := yaml.Unmarshal(b, person)
	if err2 != nil {
		panic(err)
	}

	return person
}

// dumpConfig 保存文件
func dumpConfig(person *Person) {

	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	// os.ModePerm => folder 755, file 644
	err2 := os.WriteFile("config.json", b, os.ModePerm)
	if err2 != nil {
		panic(err)
	}
}
