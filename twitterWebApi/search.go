package twitterWebApi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// URL values option

//https://api.twitter.com/2/search/adaptive.json?
//include profile interstitial type=1&
//include blocking=1&
//include blocked by=1&
//include followed by=1&
//include want retweets=1&
//include mute edge=1&
//include can dm=1&
//include can media tag=1&
//skip status=1&
//cards platform=Web-12&
//include cards=1&
//include ext alt text=true&
//include quote count=true&
//include reply count=1&
//tweet mode=extended&
//include entities=true&
//include user entities=true&
//include ext media color=true&
//include ext media availability=true&
//send error codes=true&
//simple quoted tweet=true&
//q={search_word}&
//tweet search mode=live&
//count=20&
//query source=typed query&
//pc=1&
//spelling corrections=1&
//ext=mediaStats%2ChighlightedLabel
func (core Core) Search(keyword string,values url.Values)(*SearchJson,error){
	if values == nil{
		values = url.Values{}
	}
	values.Add("q",keyword)
	req,err := core.createRequest(http.MethodGet,"https://api.twitter.com/2/search/adaptive.json?" + values.Encode(),nil)
	if err != nil{
		return nil,err
	}
	resp,err := core.Client.Do(req)
	if err != nil{
		return nil,err
	}
	var searchJson SearchJson
	err = json.NewDecoder(resp.Body).Decode(&searchJson)
	if err != nil{
		return nil,err
	}
	return &searchJson,nil
}

type Tweets struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated"`
	Entities  struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []struct {
			URL         string `json:"url"`
			ExpandedURL string `json:"expanded_url"`
			DisplayURL  string `json:"display_url"`
			Indices     []int  `json:"indices"`
		} `json:"urls"`
	} `json:"entities"`
	Source                    string      `json:"source"`
	InReplyToStatusID         interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr      interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID           interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr        interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName       interface{} `json:"in_reply_to_screen_name"`
	UserID                    int64       `json:"user_id"`
	UserIDStr                 string      `json:"user_id_str"`
	Geo                       interface{} `json:"geo"`
	Coordinates               interface{} `json:"coordinates"`
	Place                     interface{} `json:"place"`
	Contributors              interface{} `json:"contributors"`
	IsQuoteStatus             bool        `json:"is_quote_status"`
	RetweetCount              int         `json:"retweet_count"`
	FavoriteCount             int         `json:"favorite_count"`
	ConversationID            int64       `json:"conversation_id"`
	ConversationIDStr         string      `json:"conversation_id_str"`
	Favorited                 bool        `json:"favorited"`
	Retweeted                 bool        `json:"retweeted"`
	PossiblySensitive         bool        `json:"possibly_sensitive"`
	PossiblySensitiveEditable bool        `json:"possibly_sensitive_editable"`
	Lang                      string      `json:"lang"`
	SupplementalLanguage      interface{} `json:"supplemental_language"`
}

type SearchJson struct {
	GlobalObjects struct {
		Tweets map[string]Tweets `json:"tweets"`
	} `json:"globalObjects"`
}