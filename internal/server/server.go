package server

import "github.com/triangletodd/gort/internal/config"

func Init() {
	config := config.GetConfig()
	hostString := config.GetString("host") + ":" + config.GetString("port")

	r := NewRouter()
	r.Run(hostString)
}
