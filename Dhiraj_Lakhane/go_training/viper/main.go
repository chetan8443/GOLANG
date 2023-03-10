package main

import (
	"fmt"
	"viper/util"
	// "github.com/fsnotify/fsnotify"
	// "github.com/spf13/viper"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}

	// vp := viper.New()

	// // test.json
	// vp.SetConfigName("test")
	// vp.SetConfigType("json")
	// vp.AddConfigPath(".")
	// err := vp.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(vp.GetString("foo"))

	// vp.Set("name", "James")

	// vp.WriteConfig()

	// vp.OnConfigChange(func (in fsnotify.Event)  {
	// 	fmt.Printf("file changed: %s\n", in.Name)
	// })

	// vp.WatchConfig()

	// for {}

}
