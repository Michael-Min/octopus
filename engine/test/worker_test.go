package test

import (
	"Michael-Min/octopus/engine"
	"Michael-Min/octopus/config"
	"Michael-Min/octopus/xcar/parser"
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	urlCarList:="https://fake.sh.xcar.com.cn/"
	req:=engine.Request{
		Url: urlCarList,
		Parser: engine.NewFuncParser(parser.ParseCarDetailFake,config.ParseCarDetail),
	}
	res,err:=engine.Worker(req)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v",res)


}
