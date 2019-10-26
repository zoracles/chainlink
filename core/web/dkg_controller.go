package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services"
	"github.com/smartcontractkit/chainlink/core/store/presenters"
)

type DKGController struct {
	App services.Application
}

func (dkg *DKGController) GenerateKey(c *gin.Context) {
	var request struct {
		Peers     []string
		Threshold int
	}
	input, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, errors.Wrapf(err, "while reading request"))
		return
	}
	err = json.Unmarshal(input, &request)
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, errors.Wrapf(err, "while parsing json %s", input))
		return
	}
	sharedKey, err := dkg.GenerateSharedKey(request.Peers, request.Threshold)
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, errors.Wrapf(err, "while attempting to generate shared key"))
		return
	}
	jsonAPIResponse(c, presenters.SharedKey{sharedKey}, "shared key")
}
