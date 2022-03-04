package plugins

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/davecgh/go-spew/spew"
)

// Say is a demo to show how to return data directly instead of proxying
// it to the upstream.
type BasicAuthPlugin struct {
}

type BasicAuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *BasicAuthPlugin) Name() string {
	return "go-plugin-basic-auth"
}

func (b *BasicAuthPlugin) ParseConf(in []byte) (interface{}, error) {
	conf := BasicAuthConfig{}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

func (b *BasicAuthPlugin) Filter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	baseAuthConfig := conf.(BasicAuthConfig)

	username, password, ok := basicAuth(r)
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`No basic auth present`))
		return
	}

	if username != baseAuthConfig.Username || password != baseAuthConfig.Password {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Username/Password is not correct`))
		return
	}
}

// BasicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(r pkgHTTP.Request) (username, password string, ok bool) {
	auth := r.Header().Get("Authorization")
	if auth == "" {
		return
	}
	return parseBasicAuth(auth)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		spew.Dump("not match")
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
