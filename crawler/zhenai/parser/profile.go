package parser

import (
	"ccmouce/crawler/engine"
	"ccmouce/crawler/model"
	"regexp"
	"strconv"
)

var ageRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[\d]+岁</div>`)
var heightRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[\d]+cm</div>`)
var weightRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[\d]+kg</div>`)
//var nameRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var genderRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var incomeRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var houseRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var carRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var marriageRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var xinzuoRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var hokouRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var educationRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)
var occupationRE  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">[^<]+</div>`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile:=model.Profile{}
	//int型数据
	age,err:=strconv.Atoi(extractString(contents,ageRE))
	if err!=nil{
			profile.Age=age
	}
	height,err:=strconv.Atoi(extractString(contents,heightRE))
	if err!=nil{
		profile.Height=height
	}
	weight,err:=strconv.Atoi(extractString(contents,weightRE))
	if err!=nil{
		profile.Weight=weight
	}
	//字符串类型
	//profile.Name=extractString(contents,nameRE)
	profile.Name=name
	profile.House=extractString(contents,houseRE)
	profile.Car=extractString(contents,carRE)
	profile.Marriage=extractString(contents,marriageRE)
	profile.Occupation=extractString(contents,occupationRE)
	profile.Education=extractString(contents,educationRE)
	profile.Hokou=extractString(contents,hokouRE)


	result:=engine.ParseResult{
		Iterms:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp)  string {
	match:=re.FindSubmatch(contents)
	if len(match)>=2{
		return string(match[1])
	}else {
		return ""
	}
}