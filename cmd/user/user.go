package main

import (
	"math/rand"
	"mxshop/app/user"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())         // 设置随机种子，保证项目中取的随机数不一样，新版本go基本可以抛弃这种写法，因为已经内置。
	if len((os.Getenv("GOMAXPROCS"))) == 0 { // 设置 Go 程序最多同时用多少个 CPU 核心执行 Go 代码，新版本go基本可以抛弃这种写法，因为已经内置。
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	app := user.NewApp("user-server")
	app.Run()
}
