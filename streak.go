// Go wrapper for Streak's REST (HTTP/JSON) API.
package streak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
)

const API_URL = "https://www.streak.com/api/v1"

// Returns the time.Time value represented by the Streak timestamp timestampMs.
// Streak uses milliseconds since the Unix epoch.
func TimestampToTime(timestampMs int64) time.Time {
	seconds := timestampMs / 1000
	nanoseconds := (timestampMs % 1000) * 1000000
	return time.Unix(seconds, nanoseconds).UTC()
}

// Returns t as a Streak timestamp (milliseconds since the Unix epoch).
// The result is rounded to the nearest millisecond.
func TimeToTimestamp(t time.Time) int64 {
	ns := t.UnixNano()
	const MS_PER_NS = 1000000
	return (ns + (MS_PER_NS / 2)) / MS_PER_NS
}

// Client for communicating with Streak.
type Client struct {
	ApiKey     string
	httpClient http.Client
}

// Returns a new Client using apiKey as the API key.
func New(apiKey string) *Client {
	return &Client{apiKey, http.Client{}}
}

func (c *Client) request(path string, outValue interface{}) error {
	if path[0] != '/' {
		panic("path must start with /: " + path)
	}

	// Make the request and read the response
	url := API_URL + path
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth(c.ApiKey, "")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	// Parse the response
	// fmt.Println(string(responseBytes))
	err = json.Unmarshal(responseBytes, outValue)
	if err != nil {
		glog.Warningf("Parsing response from Streak failed; raw data: %s", string(responseBytes))
		return err
	}

	return err
}

// Represents a single Streak pipeline. See http://www.streak.com/api/#pipeline
type Pipeline struct {
	Key                  string `json:"key"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
}

type Box struct {
	Key                  string `json:"key"`
	Name                 string `json:"name"`
	LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
	GmailThreadCount     int64  `json:"gmailThreadCount"`
}

type Thread struct {
	Key                  string `json:"key"`
	Subject              string `json:"subject"`
	LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
	CreationTimestamp    int64  `json:"creationTimestamp"`
	LastEmailTimestamp   int64  `json:"lastEmailTimestamp"`
}

// Returns all pipelines. See: http://www.streak.com/api/#listallpipelines
func (c *Client) GetPipelines() ([]Pipeline, error) {
	var pipelines []Pipeline
	err := c.request("/pipelines", &pipelines)
	return pipelines, err
}

func (c *Client) GetBoxes(pipeline *Pipeline) ([]Box, error) {
	var boxes []Box
	err := c.request(fmt.Sprintf("/pipelines/%s/boxes", pipeline.Key), &boxes)
	return boxes, err
}

func (c *Client) GetThreads(box *Box) ([]Thread, error) {
	var threads []Thread
	err := c.request(fmt.Sprintf("/boxes/%s/threads", box.Key), &threads)
	return threads, err
}
