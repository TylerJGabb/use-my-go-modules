package main

import (
	lib "github.com/TylerJGabb/my-go-module/lib"
	"github.com/TylerJGabb/new-go-priv-mod/priv"
)

func main() {
	lib.Hello()
	priv.PrivateFunc()
}
