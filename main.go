package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/k8scat/wechat-steps/pkg/zepp"
	"golang.org/x/exp/rand"
)

var (
	account    string
	password   string
	steps      int
	stepsRange bool
	stepsMin   int
	stepsMax   int
)

func main() {
	flag.StringVar(&account, "zepp-account", "", "Zepp 账号")
	flag.StringVar(&password, "zepp-password", "", "Zepp 密码")
	flag.IntVar(&steps, "steps", 10000, "步数")
	flag.BoolVar(&stepsRange, "steps-range", false, "是否使用步数范围")
	flag.IntVar(&stepsMin, "steps-min", 0, "步数开始")
	flag.IntVar(&stepsMax, "steps-max", 10000, "步数结束")
	flag.Parse()

	if stepsRange {
		rand.Seed(uint64(time.Now().UnixNano()))
		steps = rand.Intn(stepsMax-stepsMin) + stepsMin
	}

	zepp := zepp.NewZeppLife(account, password)
	err := zepp.SetSteps(steps)
	if err != nil {
		fmt.Printf("set steps error: %+v", err)
	}
}
