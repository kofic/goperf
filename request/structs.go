package request

import (
	"net/http"
	"time"
)

type IterateReqResp struct {
	Url         string          `json:"url"`
	Status      []int           `json:"status"`
	RespTimes   []time.Duration `json:"resp_times"`
	NumRequests int             `json:"num_requests"`
	Bytes       int             `json:"bytes"`
}

type IterateReqRespAll struct {
	AvgTotalRespTime       time.Duration    `json:"avg_total_resp_time"`
	AvgTotalLinearRespTime time.Duration    `json:"avg_total_linear_resp_time"`
	BaseUrl                IterateReqResp   `json:"baseUrl"`
	JSResps                []IterateReqResp `json:"js_resps"`
	CSSResps               []IterateReqResp `json:"css_resps"`
	IMGResps               []IterateReqResp `json:"img_resps"`
}

/*
   Structure used to create web request Channel.  This is how we get the results
   back from the 'go Run(...) method call
*/
type Result struct {
	Total     time.Duration
	Average   time.Duration
	Channel   int
	Responses []*http.Response
}
