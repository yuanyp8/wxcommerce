package main

import (
	album "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album/albumservice"
	"log"
)

func main() {
	svr := album.NewServer(new(AlbumServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
