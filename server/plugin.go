package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ShiKaiWi/mattermost-coderunner-plugin/coderunner"
	"github.com/ShiKaiWi/mattermost-coderunner-plugin/httputil"
	"github.com/mattermost/mattermost-server/plugin"
)

const (
	postMethod = "POST"
)

// CodeRunnerPlugin defines the plugins for running code
type CodeRunnerPlugin struct {
	plugin.MattermostPlugin
}

// CodeRunReq defines the code run request
type CodeRunReq struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

// CodeRunResp defines the repsonse to CodeRunReq
type CodeRunResp struct {
	Ok     bool   `json:"ok"`
	Result string `json:"result"`
}

// ServeHTTP serves the route for /<mattermost-server>/plugins/com.mattermost.server-coderunner
func (p *CodeRunnerPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != postMethod {
		httputil.H405(w)
		return
	}
	if r.Body == nil {
		httputil.H400(w)
		return
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.H400(w)
		return
	}

	var req CodeRunReq
	err = json.Unmarshal(bs, &req)
	if err != nil {
		httputil.H400(w)
		return
	}

	codeRunner, err := coderunner.GetCodeRunner(coderunner.Lang(req.Lang))
	if err != nil {
		log.Printf("fail to get code runner, err=%v", err)
		httputil.H400(w)
		return
	}

	out, err := codeRunner.Run(req.Code)
	resp := CodeRunResp{
		Ok:     err == nil,
		Result: out,
	}

	httputil.H200(w, resp)
}

func main() {
	plugin.ClientMain(&CodeRunnerPlugin{})
}
