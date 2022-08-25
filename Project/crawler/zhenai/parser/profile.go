package parser

import (
	"crawler/engine"
	"crawler/model"
	"crawler/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var baseInfoRe = regexp.MustCompile(`<div class="des f-cl" data-v-4c07f04e>([^<]+)</div>`)
var nativePlaceRe = regexp.MustCompile(`<div class="m-btn pink" data[^>]+>籍贯:([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data[^>]+>([0-9]+kg)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data[^>]+>(已购房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data[^>]+>(已买车)</div>`)
var constellationRe = regexp.MustCompile(`<div class="m-btn purple" data[^>]+>([^\w]+座\(09.23-10.22\))</div>`)

var errorRe = regexp.MustCompile(`\{"buttonContent":"点击此处对该用户进行举`)

func Profile(contents []byte, other map[string]interface{}) (result engine.ParseResult) {
	var profile model.Profile
	errorMatch := errorRe.FindSubmatch(contents)
	if len(errorMatch) > 0 {
		return result
	}

	baseMatches := baseInfoRe.FindSubmatch(contents)
	//fmt.Printf("基础匹配结果: %s\n", baseMatches)
	if len(baseMatches) < 1 {
		log.Printf("contents 数据: %s\n", contents)
		return result
	}
	bases := strings.Split(string(baseMatches[1]), "|")
	profile.Age, _ = strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(bases[1]), "岁"))
	if val, ok := other["nickName"]; ok {
		profile.NickName = val.(string)
	}
	if val, ok := other["gender"]; ok {
		profile.Gender = val.(string)
	}
	profile.Education = strings.TrimSpace(bases[2])
	profile.Live = strings.TrimSpace(bases[0])
	profile.Marriage = strings.TrimSpace(bases[3])
	profile.Height = strings.TrimSpace(bases[4])
	profile.Income = strings.TrimSpace(bases[5])

	nativePlaceMatch := nativePlaceRe.FindSubmatch(contents)
	if ok := utils.OutRange(nativePlaceMatch, 1); ok {
		profile.NativePlace = string(nativePlaceMatch[1])
	}

	weightMatch := weightRe.FindSubmatch(contents)
	if ok := utils.OutRange(weightMatch, 1); ok {
		profile.Weight = string(weightMatch[1])
	}

	houseMatch := houseRe.FindSubmatch(contents)
	if ok := utils.OutRange(houseMatch, 1); ok {
		profile.House = ok
	}

	carMatch := carRe.FindSubmatch(contents)
	if ok := utils.OutRange(carMatch, 1); ok {
		profile.Car = ok
	}

	constellationMatch := constellationRe.FindSubmatch(contents)
	if ok := utils.OutRange(constellationMatch, 1); ok {
		profile.Constellation = string(constellationMatch[1])
	}
	result.Items = append(result.Items, profile)
	return
}
