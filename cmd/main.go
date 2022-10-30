package main

import (
	web "1536509937/ku-bbs/cmd/webserver"
	_ "1536509937/ku-bbs/pkg/config"
	_ "1536509937/ku-bbs/pkg/db"
	_ "1536509937/ku-bbs/pkg/redis"
)

func main() {
	web.Run()
}
