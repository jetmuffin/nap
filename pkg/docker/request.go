package docker

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// serverResponse is a wrapper for http API responses.
type serverResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
	reqURL     *url.URL
}

type headers map[string][]string

func (cli *Client) get(path string, query url.Values, headers headers) (serverResponse, error) {
	return cli.sendRequest("GET", path, query, nil, headers)
}

func (cli *Client) post(path string, query url.Values, obj interface{}, headers headers) (serverResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return serverResponse{}, err
	}
	return cli.sendRequest("POST", path, query, body, headers)
}

func (cli *Client) put(path string, query url.Values, obj interface{}, headers headers) (serverResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return serverResponse{}, err
	}
	return cli.sendRequest("PUT", path, query, body, headers)
}

func (cli *Client) delete(path string, query url.Values, headers headers) (serverResponse, error) {
	return cli.sendRequest("DELETE", path, query, nil, headers)
}

func (cli *Client) postHijack(path string, query url.Values, obj interface{}, headers headers, input chan []byte) (chan []byte, error) {
	output := make(chan []byte)

	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return nil, err
	}

	req, err := cli.buildRequest("POST", cli.getAPIPath(path, query), body, headers)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", cli.addr)
	if err != nil {
		return nil, err
	}
	clientconn := httputil.NewClientConn(conn, nil)
	clientconn.Do(req)

	rwc, br := clientconn.Hijack()

	go func() {
		defer clientconn.Close()

		for {
			if data, ok := <-input; ok {
				rwc.Write(data)
			} else {
				break
			}
		}
	}()

	go func() {
		defer rwc.Close()

		for {
			buf := make([]byte, 1024)
			_, err := br.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					output <- []byte("EOF")
					break
				}
				break
			}

			output <- buf

			//Equal EOF
			if buf[0] == 69 && buf[1] == 79 && buf[2] == 70 {
				close(output)
				break
			}

			time.Sleep(500)
			buf = nil
		}
	}()
	return output, nil
}

func encodeBody(obj interface{}, headers headers) (io.Reader, headers, error) {
	if obj == nil {
		return nil, headers, nil
	}

	body := bytes.NewBuffer(nil)
	if obj != nil {
		if err := json.NewEncoder(body).Encode(obj); err != nil {
			return nil, headers, err
		}
	}
	if headers == nil {
		headers = make(map[string][]string)
	}

	headers["Content-Type"] = []string{"application/json"}
	return body, headers, nil
}

func (cli *Client) buildRequest(method, path string, body io.Reader, headers headers) (*http.Request, error) {
	expectedPayload := method == "POST" || method == "PUT"
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req = cli.addHeaders(req, headers)

	if cli.proto == "unix" || cli.proto == "npipe" {
		// For local communications, it doesn't matter what the host is. We just
		// need a valid and meaningful host name. (See #189)
		req.Host = "docker"
	}

	req.URL.Host = cli.addr
	req.URL.Scheme = cli.scheme

	if expectedPayload && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "text/plain")
	}
	return req, nil
}

func (cli *Client) addHeaders(req *http.Request, headers headers) *http.Request {
	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}
	return req
}

func (cli *Client) sendRequest(method, path string, query url.Values, body io.Reader, headers headers) (serverResponse, error) {
	req, err := cli.buildRequest(method, cli.getAPIPath(path, query), body, headers)
	if err != nil {
		return serverResponse{}, err
	}
	resp, err := cli.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (cli *Client) doRequest(req *http.Request) (serverResponse, error) {
	serverResp := serverResponse{statusCode: -1, reqURL: req.URL}

	resp, err := cli.client.Do(req)
	if err != nil {
		return serverResp, err
	}

	if resp != nil {
		serverResp.statusCode = resp.StatusCode
		serverResp.body = resp.Body
		serverResp.header = resp.Header
	}
	return serverResp, nil
}
