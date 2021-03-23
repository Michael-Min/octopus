package parser

import (
	"regexp"

	"Michael-Min/octopus/config"
	"Michael-Min/octopus/engine"
)

var carDetailRe = regexp.MustCompile(`<a href="(/\d+/price_m[0-9]+.htm)" target="_blank"[^>]*>`)

func ParseCarModel(
	contents []byte, _ string) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: host + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarDetail, config.ParseCarDetail),
			})
	}

	return result
}
