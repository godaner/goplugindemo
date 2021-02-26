package main

import (
	"log"
	"path/filepath"
	"plugin"
	"time"
)

func main() {
	pluginFile, pluginFunc := "plugin.so", "Add"
	log.Printf("Plugin arg is: %v#%v", pluginFile, pluginFunc)
	for {
		<-time.After(time.Second)
		// 加载plugin
		plugins, err := filepath.Glob(pluginFile)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(plugins) == 0 {
			log.Printf("Plugin %v not found", pluginFile)
			continue
		}
		log.Printf("Loading plugin %v", pluginFile)
		p, err := plugin.Open(pluginFile)
		if err != nil {
			log.Println(err)
			continue
		}
		// 查找函数
		symbol, err := p.Lookup(pluginFunc)
		if err != nil {
			log.Println(err)
			continue
		}
		funcc, ok := symbol.(func(int, int) int)
		if !ok {
			log.Printf("Plugin has no %v function", pluginFunc)
			continue
		}

		// 调用函数
		res := funcc(3, 4)
		log.Printf("Result is: %d", res)
	}

}
