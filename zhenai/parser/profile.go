package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"Michael-Min/octopus/config"
	"Michael-Min/octopus/engine"
	pb "Michael-Min/octopus/proto"
	"regexp"
	"strconv"
	"strings"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(
	contents []byte, url string,
	name string) engine.ParseResult {
	profile := pb.Profile{}
	profile.Name = name

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err!=nil{
		return engine.ParseResult{}
	}
	tmp:=dom.Find("div.des.f-cl").Map(func(i int, selection *goquery.Selection) string {
		return selection.Text()
	})
	fmt.Println(tmp)
	selections:=dom.Find("div.des.f-cl").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	profile.Marriage=selections.Eq(0).Text()
	i64,err:=strconv.ParseInt(strings.TrimRight(selections.Eq(1).Text(),"岁"),10,32)
	if err ==nil{
		profile.Age=int32(i64)
	}
	profile.Xinzuo=selections.Eq(2).Text()
	i64,err=strconv.ParseInt(strings.TrimRight(selections.Eq(3).Text(),"cm"),10,32)
	if err ==nil{
		profile.Age=int32(i64)
	}
	incomeArr:=strings.Split(selections.Eq(5).Text(),":")
	if len(incomeArr)>1{
		profile.Income=incomeArr[1]
	}
	profile.Occupation=selections.Eq(6).Text()
	profile.Education=selections.Eq(7).Text()
	fmt.Println(profile)

	//age, err := strconv.Atoi(
	//	extractString(contents, ageRe))
	//if err == nil {
	//	profile.Age = int32(age)
	//}
	//
	//height, err := strconv.Atoi(
	//	extractString(contents, heightRe))
	//if err == nil {
	//	profile.Height = int32(height)
	//}
	//
	//weight, err := strconv.Atoi(
	//	extractString(contents, weightRe))
	//if err == nil {
	//	profile.Weight = int32(weight)
	//}
	//
	//profile.Income = extractString(
	//	contents, incomeRe)
	//profile.Gender = extractString(
	//	contents, genderRe)
	//profile.Car = extractString(
	//	contents, carRe)
	//profile.Education = extractString(
	//	contents, educationRe)
	//profile.Hokou = extractString(
	//	contents, hokouRe)
	//profile.House = extractString(
	//	contents, houseRe)
	//profile.Marriage = extractString(
	//	contents, marriageRe)
	//profile.Occupation = extractString(
	//	contents, occupationRe)
	//profile.Xinzuo = extractString(
	//	contents, xinzuoRe)
	//
	//result := engine.ParseResult{
	//	Items: []*pb.Item{
	//		{
	//			Url:  url,
	//			Type: "zhenai",
	//			Id: extractString(
	//				[]byte(url), idUrlRe),
	//			Payload: &profile,
	//		},
	//	},
	//}
	//
	//matches := guessRe.FindAllSubmatch(
	//	contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(result.Requests,
	//		engine.Request{
	//			Url: string(m[1]),
	//			Parser: NewProfileParser(
	//				string(m[2])),
	//		})
	//}

	//return result
	return engine.ParseResult{}
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(
	contents []byte,
	url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (
	name string, args string) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(
	name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
