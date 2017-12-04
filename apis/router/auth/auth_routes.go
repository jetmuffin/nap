package auth

import (
	"errors"
	"fmt"
	"github.com/JetMuffin/nap/apis/utils"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ClientID     = "dc-njuics-cn"
	ClientSecret = "dcos-nap"
	GrantType    = "authorization_code"
	RedirectUrl  = "http://localhost:4200"
)

func (ar *authRouter) handleAuthorize(w http.ResponseWriter, req *http.Request) {
	// Parse request parameters
	if err := utils.ParseForm(req); err != nil {
		utils.WriteError(w, err)
		return
	}

	code := req.Form.Get("code")

	data := make(url.Values)
	data["client_id"] = []string{ClientID}
	data["client_secret"] = []string{ClientSecret}
	data["grant_type"] = []string{GrantType}
	data["redirect_uri"] = []string{RedirectUrl}
	data["code"] = []string{code}

	res, err := http.PostForm(fmt.Sprintf("http://%s/oauth", ar.oAuthAddr), data)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	if res.StatusCode != http.StatusOK {
		utils.WriteError(w, errors.New("Unauthorized."))
		return
	}

	defer res.Body.Close()
	token, err := ioutil.ReadAll(res.Body)

	if err != nil {
		utils.WriteError(w, err)
	}

	w.Write([]byte(token))
}
