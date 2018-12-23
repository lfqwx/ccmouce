package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"net/http"
	"regexp"
)
//匹配规则
var iframeRe=regexp.MustCompile(`<iframe id='frame_content' width="100%" height="100%" src="http://10.162.96.103:9080/whjgdb/stone/dyysk_wh/query.jsp" frameborder="0" style="overflow: hidden;"></iframe>
`)
var urlRe=regexp.MustCompile(`http://10.162.96.103:9080/whjgdb/stone/dyysk_wh/dysk.jsp`)
//读取excel
func read()[][]string  {
	xlsx, err := excelize.OpenFile("./cycle.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	// Get value from cell by given worksheet name and axis.
	rows := xlsx.GetRows("Sheet1")
	return rows
}
//获取网页数据
func HttpGet(url string)  []byte {
	resp, err1 := http.Get(url)//核心方法
	if err1 != nil {
		err := err1
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096) //切片装读取到的数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err:= err2
			fmt.Println(err)
		}

	}
	return buf
}
//正则匹配
func testregexp(result []byte,iframeRe *regexp.Regexp,urlRe *regexp.Regexp)[][][]byte{
	match:=iframeRe.FindAllSubmatch(result,-1)
	if match==nil{
		fmt.Println("没有找到")
	}else{
		for _,v1:=range match{
			for _,v2:=range v1{
				for _,v3:=range v2{
					fmt.Println(v3)
				}
			}
		}
	}
	return match
}
func main() {
	//fmt.Println(string(HttpGet("http://www.zhenai.com/zhenghun")))//有数据啊

	//循环获取url,并通过HttpGet()访问
	for _,v:=range read() {
		for _,url:=range v  {
			testregexp(HttpGet(url),iframeRe,urlRe)
		}
	}
}
