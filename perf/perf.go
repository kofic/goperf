package perf

import (
	"encoding/json"
	"fmt"
	"github.com/gnulnx/goperf/request"
	"time"
)

type Init struct {
	Url        string
	Threads    int
	Seconds    int
	Iterations int
	Output     int
	Index      int // Also the channel number
	Verbose    bool
	Results    *request.IterateReqRespAll
}

func (input *Init) Basic() request.IterateReqRespAll {
	// Create slice of channels to hold results
	// Fire off anonymous go routine using newly created channel
	chanslice := []chan request.IterateReqRespAll{}
	for i := 0; i < input.Threads; i++ {
		chanslice = append(chanslice, make(chan request.IterateReqRespAll))
		go func(c chan request.IterateReqRespAll) {
			c <- iterateRequest(input.Url, input.Seconds)
		}(chanslice[i])
	}

	// Wait on all the channels
	results := []request.IterateReqRespAll{}
	for _, ch := range chanslice {
		results = append(results, <-ch)
	}

	// Combine all the results into 1
	_ = request.Combine(results)

	input.Results = request.Combine(results)
	return *input.Results

}

func (input Init) Json() {
	tmp, _ := json.MarshalIndent(input.Results, "", "  ")
	fmt.Println(string(tmp))
}

func iterateRequest(url string, sec int) request.IterateReqRespAll {
	/*
		Continuously fetch 'url' for 'sec' second and return the results.
	*/
	start := time.Now()
	maxTime := time.Duration(sec) * time.Second
	elapsedTime := maxTime

	resp := request.IterateReqResp{
		Url: url,
	}
	jsMap := map[string]*request.IterateReqResp{}
	cssMap := map[string]*request.IterateReqResp{}
	imgMap := map[string]*request.IterateReqResp{}

	for {
		//Fetch the url and all the js, css, and img assets
		fetchAllResp := request.FetchAll(url, false)

		// Set base resp properties
		resp.Status = append(resp.Status, fetchAllResp.BaseUrl.Status)
		resp.RespTimes = append(resp.RespTimes, fetchAllResp.BaseUrl.Time)
		resp.Bytes = fetchAllResp.TotalBytes

		gatherAllStats(fetchAllResp, jsMap, cssMap, imgMap)

		elapsedTime = time.Now().Sub(start)
		if elapsedTime > maxTime {
			break
		}
	}

	// TODO Clean this up.  Perhaps some benchmark tests
	// to see if its faster as go routines or not
	jsResps := []request.IterateReqResp{}
	for _, val := range jsMap {
		jsResps = append(jsResps, *val)
	}

	cssResps := []request.IterateReqResp{}
	for _, val := range cssMap {
		cssResps = append(cssResps, *val)
	}

	imgResps := []request.IterateReqResp{}
	for _, val := range imgMap {
		imgResps = append(imgResps, *val)
	}

	output := request.IterateReqRespAll{
		BaseUrl:  resp,
		JSResps:  jsResps,
		CSSResps: cssResps,
		IMGResps: imgResps,
	}
	return output
}

func gatherAllStats(resp *request.FetchAllResponse, jsMap map[string]*request.IterateReqResp, cssMap map[string]*request.IterateReqResp, imgMap map[string]*request.IterateReqResp) {
	/*
		Gather all the asset stuff.
		Note;  You benchmarked this and the 3 go routine method was way slower so you removed the method
		BenchmarkGatherAllStatsGo-8   	  500000	      2764 ns/op
		BenchmarkGatherAllStats-8     	 2000000	       638 ns/op
	*/
	gatherStats(resp.JSResponses, jsMap)
	gatherStats(resp.CSSResponses, cssMap)
	gatherStats(resp.IMGResponses, imgMap)
}

func gatherStats(Resps []request.FetchResponse, respMap map[string]*request.IterateReqResp) {
	// gather all the responses
	for resp := 0; resp < len(Resps); resp++ {
		url2 := Resps[resp].Url
		bytes := Resps[resp].Bytes
		status := Resps[resp].Status
		respTime := Resps[resp].Time
		_, ok := respMap[url2]
		if !ok {
			respMap[url2] = &request.IterateReqResp{
				Url:         url2,
				Bytes:       bytes,
				Status:      []int{status},
				RespTimes:   []time.Duration{respTime},
				NumRequests: 1,
			}
		} else {
			respMap[url2].Status = append(respMap[url2].Status, status)
			respMap[url2].RespTimes = append(respMap[url2].RespTimes, respTime)
			respMap[url2].NumRequests += 1
		}
	}
}