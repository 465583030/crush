package squirt

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Squirt struct {
	url string
}

func NewSquirt(url string) *Squirt {
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	return &Squirt{url: url + "/"}
}

func (s *Squirt) Messages(topic string) chan string {
	msgs := make(chan string)
	go s.Get(topic, msgs)
	return msgs
}

func (s *Squirt) Get(topic string, q chan string) {
	for {
		resp, err := http.Get(s.url + topic)
		defer resp.Body.Close()
		if err != nil || resp.StatusCode == 404 {
			<-time.After(1 * time.Second)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		q <- string(body)
	}
}

func (s *Squirt) Complete(topic, id string) {
	req, _ := http.NewRequest("DELETE", s.url+topic, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
}
