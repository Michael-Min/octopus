package parser

import (
	"Michael-Min/octopus/engine"
	pb "Michael-Min/octopus/proto"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var namefake = regexp.MustCompile(`<title>【(.*)】.*</title>`)

func ParseCarDetailFake(contents []byte, _ string) engine.ParseResult {
	_ = namefake.FindAllSubmatch(contents, -1)
	id:="m" + GenValidateCode(5,false)
	value,_:=strconv.ParseFloat(GenValidateCode(2,false) +"."+ GenValidateCode(2,false),32)
	price := float32(value)
	item:=&pb.Item{
		Url:  "http://newcar.xcar.com.cn/m35001/",
		Type: "xcar",
		Id:   id,
		Car: &pb.Car{
			Name:         "奥迪TT双门2017款45 TFSI",
			Price:        price,
			ImageURL:     "http://img1.xcarimg.com/b63/s8386/m_20170616000036181753843373443.jpg-280x210.jpg",
			Size:         "4191×1832×1353mm",
			Fuel:         16.7,
			Transmission: "6挡双离合",
			Engine:       "169kW(2.0L涡轮增压)",
			Displacement: 2,
			MaxSpeed:     250,
			Acceleration: 5.9,
		},
	}

	req:=engine.Request{
		Url:    "http://fake.dealer.xcar.com.cn/12919/price_m50077.htm",
		Parser: engine.NewFuncParser(ParseCarDetailFake,"ParseCarDetailFake"),
	}

	result := engine.ParseResult{
		Items: []*pb.Item{item},
		Requests: []engine.Request{req},
	}

	return result
}

func GenValidateCode(width int,haszero bool) string {
	if width==0{
		return ""
	}
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	var num byte
	var sb strings.Builder
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < width; i++ {
		if haszero==false{
			num=numeric[ rand.Intn(r-1)+1]
		}else {
			num=numeric[ rand.Intn(r)]
		}
		_,_=fmt.Fprintf(&sb, "%d",num)
	}
	return sb.String()
}


