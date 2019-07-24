package parser

import (
	"golang/learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items,m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: nil,
		})
	}
	return result
}
