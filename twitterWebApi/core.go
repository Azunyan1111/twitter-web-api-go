package twitterWebApi

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)


type Core struct {
	Client http.Client
	Token string
	GuestToken string
	QueryIdMap map[string]Query
}

type Query struct {
	Id string `json:"queryId"`
	OperationName string `json:"operationName"`
	OperationType string`json:"operationType"`
}

type ActivateJson struct {
	GuestToken string `json:"guest_token"`
}

// setup client
func NewCore()(err error,core Core){
	// Get main.js
	resp,err := http.Get("https://twitter.com/Twitter")
	if err != nil{
		return err,Core{}
	}
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		return err,Core{}
	}
	// find main.js
	var mainJsUrl string
	mainJsRegexp,_ := regexp.Compile(`main\.[0-9a-z]*\.js`)
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		src,ok := selection.Attr("src")
		if !ok{
			return
		}
		if mainJsRegexp.MatchString(src){
			mainJsUrl = src
		}
	})
	if mainJsUrl == ""{
		return errors.New("main.js not found"),Core{}
	}
	// Get Token
	resp,err = http.Get(mainJsUrl)
	if err != nil{
		return err,Core{}
	}
	bodyByte,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err,Core{}
	}
	body := string(bodyByte)
	tokenStart := strings.Index(body,"AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3")
	tokenEnd := strings.Index(body[tokenStart:],`"`)
	token := body[tokenStart : tokenStart+tokenEnd]

	if token == ""{
		return errors.New("token not found"),Core{}
	}
	core.Token = token
	core.QueryIdMap = make(map[string]Query)
	// SetQueryId
	r,_ := regexp.Compile(`{queryId:"[A-Za-z0-9_\-]*",operationName:"[A-Za-z0-9]*",operationType:"[A-Za-z0-9]*"}`)
	for _,q := range r.FindAllString(body,-1){
		q = strings.ReplaceAll(q,`queryId`,`"queryId"`)
		q = strings.ReplaceAll(q,`operationName`,`"operationName"`)
		q = strings.ReplaceAll(q,`operationType`,`"operationType"`)
		var query Query
		err := json.Unmarshal([]byte(q),&query)
		if err != nil{
			return err,Core{}
		}
		core.QueryIdMap[query.OperationName] = query
	}

	// Get GuestToken
	req,err := core.createRequest(http.MethodPost,"https://api.twitter.com/1.1/guest/activate.json",nil)
	if err != nil{
		return err,Core{}
	}
	resp,err = core.Client.Do(req)
	if err != nil{
		return err,Core{}
	}
	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err,Core{}
	}

	var activateJson ActivateJson
	err = json.Unmarshal(b,&activateJson)
	if err != nil{
		return err,Core{}
	}
	core.GuestToken = activateJson.GuestToken
	return nil,core
}

func (core Core)createRequest(method string,url string, body io.Reader)(*http.Request,error){
	req,err := http.NewRequest(method,url,body)
	if err != nil{
		return nil,err
	}
	req.Header.Set("authorization", "Bearer " + core.Token)
	req.Header.Set("x-guest-token",core.GuestToken)
	return req,nil
}