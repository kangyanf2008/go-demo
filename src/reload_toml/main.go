package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"time"
)

// Config 定义了TOML配置文件的结构
type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
}

var _currentConfig Config

func loadConfig(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config file %s not found", filePath)
	}

	if _, err := toml.DecodeFile(filePath, &_currentConfig); err != nil {
		return fmt.Errorf("error parsing config file %s: %w", filePath, err)
	}

	return nil
}
func GetConfig() Config {
	return _currentConfig
}

func watchConfig(filePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					done <- true
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("Config file changed, reloading...")
					if err := loadConfig(filePath); err != nil {
						log.Printf("Error reloading config: %v\n", err)
					} else {
						fmt.Printf("Config reloaded: %+v\n", _currentConfig)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					done <- true
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	configFilePath := "D:\\kyf\\go-workspace\\go-demo\\src\\reload_toml\\config.toml"
	if err := loadConfig(configFilePath); err != nil {
		log.Fatalf("Failed to load initial config: %v", err)
	}

	fmt.Printf("Initial config loaded: %+v\n", _currentConfig)

	// 在一个单独的goroutine中监控配置文件的变化
	go watchConfig(configFilePath)

	// 模拟应用程序的运行（这里用一个简单的循环代替）
	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("Application is running with config: %+v\n", GetConfig())
	}
}
