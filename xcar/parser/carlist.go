package parser

import (
	"regexp"

	"Michael-Min/octopus/config"
	"Michael-Min/octopus/engine"
)

const host = "http://dealer.xcar.com.cn"

var carModelRe = regexp.MustCompile(`<a href="(/\d+/price_m[0-9]+.htm)" target="_blank"[^>]*>`)
var carListRe = regexp.MustCompile(`<a href="(//dealer.xcar.com.cn/[0-9]+/)" target="_blank"[^>]*>`)

func ParseCarList(
	contents []byte, _ string) engine.ParseResult {
	matches := carModelRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: host + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarModel, config.ParseCarModel),
			})
	}

	matches = carListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: "http:" + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarList, config.ParseCarList),
			})
	}

	return result
}
