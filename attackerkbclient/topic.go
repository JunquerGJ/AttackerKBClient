package attackerkbclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Topic struct {
	Id                        string      `json:"id"`
	EditorId                  string      `json:"editorId"`
	Name                      string      `json:"name"`
	Created                   string      `json:"created"`
	RevisionDate              string      `json:"revisionDate"`
	DisclosureDate            string      `json:"disclosureDate"`
	Document                  string      `json:"document"`
	Metadata                  string      `json:"metadata"`
	Score                     Score       `json:"score"`
	Tags                      []Tag       `json:"tags"`
	References                []Reference `json:"references"`
	RapidAnalysis             string      `json:"rapidAnalysis"`
	RapidAnalysisCreated      string      `json:"rapidAnalysisCreated"`
	RapidAnalysisRevisionDate string      `json:"rapidAnalysisRevisionDate"`
}

type Score struct {
	AttackerValue  float64 `json:"AttackerValue"`
	Exploitability float64 `json:"Exploitability"`
}

type Tag struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Kind     string   `json:"kind"`
	Code     string   `json:"code"`
	Metadata Metadata `json:"metadata"`
}
type Metadata struct {
	Value      string `json:"value"`
	Source     string `json:"source"`
	TacticId   string `json:"tacticId"`
	TacticName string `json:"tacticName"`
}
type Reference struct {
	Id        string `json:"id"`
	EditorId  string `json:"editorId"`
	Created   string `json:"created"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	RefType   string `json:"refType"`
	RefSource string `json:"refSource"`
}

type Links struct {
	Next Link `json:"next"`
	Prev Link `json:"prev"`
	Self Link `json:"self"`
}

type Link struct {
	Href string `json:"href"`
}

type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Contributor struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Created  string `json:"created"`
	Score    int    `json:"score"`
}

type TopicSearch struct {
	Links Links   `json:"links"`
	Data  []Topic `json:"data"`
}

type Assesment struct {
	Id           string `json:"id"`
	EditorId     string `json:"editorId"`
	TopicId      string `json:"topicId"`
	Created      string `json:"created"`
	RevisionDate string `json:"revisionDate"`
	Document     string `json:"document"`
	Score        int    `json:"score"`
	Metadata     string `json:"metadata"`
	//	tags         []Tag
}

func (s *Client) TopicSearch(q string) (*TopicSearch, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/topics?q=%s", BaseUrl, q), nil)
	req, err := http.Get(fmt.Sprintf("%s/topics?q=%s", BaseUrl, q))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authentication", fmt.Sprintf("Basic %s", s.apiKey))
	defer req.Body.Close()
	var ret TopicSearch
	if err := json.NewDecoder(req.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
