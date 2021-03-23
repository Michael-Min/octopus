package engine

import (
	"Michael-Min/octopus/fetcher"
	"log"
	"strings"
)

func Worker(r Request) (ParseResult, error) {
	var body []byte
	var err error
	if strings.Contains(r.Url,"://fake"){

	}else {
		body, err = fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+
				"fetching url %s: %v",
				r.Url, err)
			return ParseResult{}, err
		}
	}


	return r.Parser.Parse(body, r.Url), nil
}
