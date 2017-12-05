package docker

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Client struct {
	// scheme sets the scheme for the client
	scheme string
	// host holds the server address to connect to
	host string
	// proto holds the client protocol i.e. unix.
	proto string
	// addr holds the client address.
	addr string
	// basePath holds the path to prepend to the requests.
	basePath string
	// client used to send and receive http requests.
	client *http.Client
	// version of the server to talk to.
	version string
}

func NewClient(host string, version string, client *http.Client) (*Client, error) {
	hostURL, err := ParseHostURL(host)
	if err != nil {
		return nil, err
	}

	if client != nil {
		if _, ok := client.Transport.(http.RoundTripper); !ok {
			return nil, fmt.Errorf("unable to verify TLS configuration, invalid transport %v", client.Transport)
		}
	} else {
		transport := new(http.Transport)
		client = &http.Client{
			Transport: transport,
		}
	}

	scheme := "http"

	return &Client{
		scheme:   scheme,
		host:     host,
		proto:    hostURL.Scheme,
		addr:     hostURL.Host,
		basePath: hostURL.Path,
		client:   client,
		version:  version,
	}, nil
}

// ParseHostURL parses a url string, validates the string is a host url, and
// returns the parsed URL
func ParseHostURL(host string) (*url.URL, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return nil, fmt.Errorf("unable to parse docker host `%s`", host)
	}

	var basePath string
	proto, addr := protoAddrParts[0], protoAddrParts[1]
	if proto == "tcp" {
		parsed, err := url.Parse("tcp://" + addr)
		if err != nil {
			return nil, err
		}
		addr = parsed.Host
		basePath = parsed.Path
	}
	return &url.URL{
		Scheme: proto,
		Host:   addr,
		Path:   basePath,
	}, nil
}

// getAPIPath returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func (cli *Client) getAPIPath(p string, query url.Values) string {
	var apiPath string
	if cli.version != "" {
		v := strings.TrimPrefix(cli.version, "v")
		apiPath = path.Join(cli.basePath, "/v"+v, p)
	} else {
		apiPath = path.Join(cli.basePath, p)
	}
	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}
