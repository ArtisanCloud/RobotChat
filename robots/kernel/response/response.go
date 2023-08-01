package response

import (
	"bytes"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"net/http"
	"net/http/httputil"
)

func LogResponse(logger contract.LoggerInterface, response *http.Response) {
	var output bytes.Buffer
	output.Write([]byte("------------------"))
	output.Write([]byte("response content:"))
	dumpRes, _ := httputil.DumpResponse(response, true)
	output.Write(dumpRes)
	logger.Info(output.String())
}
