package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	SetVerboseLogging()
	urls:=[]string{"https://album.zhenai.com/u/1773774851","http://www.baidu.com","www.baidu.com","http://www.scmp.com/news/china"}

	for i:=0;i<2;i++{
		for _,url:=range urls{

			bytes, err := Fetch(url)
			if err!=nil{
				fmt.Printf("fetch wrong ....,err is %s\n",err)
			}
			fmt.Printf("Response length:%d, strings:\n%s",len(bytes),bytes)

		}
	}


}
