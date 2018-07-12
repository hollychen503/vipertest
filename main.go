package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/devilsray/golang-viper-config-example/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/dhmsconfig/")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Println("ConfigFileUsed", viper.ConfigFileUsed())

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)

	///////////////////

	confMd5, err := fileMd5(viper.ConfigFileUsed())
	if err != nil {
		panic(err)
	}

	///////////////////

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		// check : identical file ?
		confNewMd5, err := fileMd5(viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}
		log.Println(confNewMd5)

		if confNewMd5 == confMd5 {
			log.Println(" same config file")
			return
		}
		confMd5 = confNewMd5

		err = viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		log.Printf("database uri is %s", configuration.Database.ConnectionUri)
		log.Printf("port for this application is %d", configuration.Server.Port)
	})

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

func fileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}
