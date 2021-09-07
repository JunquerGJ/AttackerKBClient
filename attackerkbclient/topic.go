package attackerkbclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Topic struct {
	Id                        string
	EditorId                  string
	Name                      string
	Created                   string
	RevisionDate              string
	DisclosureDate            string
	Document                  string
	Metadata                  string
	Score                     Score
	Tags                      []Tag
	References                []Reference
	RapidAnalysis             string
	RapidAnalysisCreated      string
	RapidAnalysisRevisionDate string
}

type Score struct {
	AttackerValue  float64
	Exploitability float64
}

type Tag struct {
	Id       string
	Name     string
	Kind     string
	Code     string
	Metadata Metadata
}
type Metadata struct {
	Value      string
	Source     string
	TacticId   string
	TacticName string
}
type Reference struct {
	Id        string
	EditorId  string
	Created   string
	Name      string
	Url       string
	RefType   string
	RefSource string
}

type Links struct {
	Next Link
	Prev Link
	Self Link
}

type Link struct {
	Href string
}

type Error struct {
	Message string
	Status  string
}

type Contributor struct {
	Id       string
	Username string
	Avatar   string
	Created  string
	Score    int
}

type TopicSearch struct {
	Links Links
	Data  []Topic
}

type Assesment struct {
	Id           string
	EditorId     string
	TopicId      string
	Created      string
	RevisionDate string
	Document     string
	Score        int
	Metadata     string
	//	tags         []Tag
}

func (s *Client) TopicSearch(q string) (*TopicSearch, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/topics?q=%s", BaseUrl, q), nil)
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
