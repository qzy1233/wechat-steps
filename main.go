package main

import (
	"flag"
	"fmt"

	"github.com/k8scat/wechat-steps/pkg/zepp"
)

var (
	account  string
	password string
	steps    int
)

func main() {
	flag.StringVar(&account, "zepp-account", "", "Zepp 账号")
	flag.StringVar(&password, "zepp-password", "", "Zepp 密码")
	flag.IntVar(&steps, "steps", 10000, "步数")
	flag.Parse()

	zepp := zepp.NewZeppLife(account, password)
	err := zepp.SetSteps(steps)
	if err != nil {
		fmt.Printf("set steps error: %+v", err)
	}
}
