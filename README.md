# twitter-web-api-go
This is a Go language wrapper for Twitter's Web API.

## install

```
$ go get github.com/Azunyan1111/twitter-web-api-go
```

## How to use


#### search API

```
package main

import (
	"fmt"
	"github.com/Azunyan1111/twitter-web-api-go/twitterWebApi"
	"net/url"
)

func main() {
	err,core := twitterWebApi.NewCore()
	if err != nil{
		panic(err)
	}
	v := url.Values{}
	v.Add("count","10")
	v.Add("tweet_search_mode","live")
	search,err := core.Search("Hello World",v)
	if err != nil{
		panic(err)
	}
	for key,val := range search.GlobalObjects.Tweets{
		fmt.Println(key,val.Text)
	}
}

```

### Message
Don't you think Twitter's API should issue an API key without any research?
