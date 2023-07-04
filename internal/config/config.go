package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func CreateConfigFile() error {
	file, err := CreateFile("Config", "json")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(file)

	vp := viper.New()

	vp.SetConfigName("Config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	err = vp.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}

func CreateFile(name string, fileType string) (string, error) {
	f := "%s.%s"

	formatedFilename := fmt.Sprintf(f, name, fileType)

	file, err := os.Create(formatedFilename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer file.Close()

	return formatedFilename, err
}
