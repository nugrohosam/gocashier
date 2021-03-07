package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	database "github.com/nugrohosam/gocashier/services/databases"
	grpcConn "github.com/nugrohosam/gocashier/services/grpc"
	httpConn "github.com/nugrohosam/gocashier/services/http"
	infrastructure "github.com/nugrohosam/gocashier/services/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	loadConfigFile()

	infrastructure.PrepareSentry()

	if err := database.ConnOrm(); err != nil {
		panic(err)
	}

	runGrpc := func() {
		if err := grpcConn.Serve(); err != nil {
			panic(err)
		}
	}

	runHTTP := func() {
		if err := httpConn.Serve(); err != nil {
			panic(err)
		}
	}

	go runGrpc()
	go runHTTP()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func initiateRedisCache() {
	cacheRedisPrefixKey := viper.GetString("cache.redis.prefix-key")
	configCacheRedis := viper.GetStringMap("cache.redis.hosts")
	redisHostsCache := make(map[string]string)

	for key, value := range configCacheRedis {
		keyRedis := cacheRedisPrefixKey + key
		valueReal := value.(map[string]string)
		redisHostsCache[keyRedis] = valueReal["host"] + ":" + valueReal["port"]
	}

	infrastructure.InitiateRedisCache(redisHostsCache)
}

func loadConfigFile() {
	viper.SetConfigType("yaml")

	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Load all files in config folders
	var files []string

	configFolderName := "config"
	root := "./" + configFolderName
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() != configFolderName {
			files = append(files, info.Name())
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var nameConfig string

	for _, file := range files {
		nameConfig = strings.ReplaceAll(file, ".yaml", "")

		viper.SetConfigName(nameConfig)
		viper.AddConfigPath(root)

		if err := viper.MergeInConfig(); err != nil {
			panic(err)
		}
	}
}