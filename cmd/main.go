package main

import (
	"github.com/chuksgpfr/uploadverse/config"
	"github.com/chuksgpfr/uploadverse/http"
	"github.com/chuksgpfr/uploadverse/ipfs"
	"github.com/chuksgpfr/uploadverse/postgres"
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	config, err := config.LoadEnv()

	if err != nil {
		panic(err)
	}

	sh := shell.NewShell(config.IpfsNode)

	db := postgres.NewDBClient(config.PostgresDSN)

	ipfsService := ipfs.IpfsService{
		Sh: sh,
	}

	fileService := postgres.FileService{
		DB:          db,
		IpfsService: ipfsService,
	}

	router, err := http.NewServer(fileService)

	if err != nil {
		panic(err)
	}

	err = router.Run(config.ServerAddress)

	if err != nil {
		panic(err)
	}
}
