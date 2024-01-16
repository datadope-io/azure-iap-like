package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const redirectUrl = "http://10.2.0.4"

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	r.Host = p.target.Host
	// w.Header().Set("X-Ben", "radi")
	p.proxy.ServeHTTP(w, r)
}

func main() {
	// Replace 'target' with the URL of the server you want to proxy to
	target, err := url.Parse(redirectUrl)
	if err != nil {
		panic(err)
	}

	// Create a new ReverseProxy instance
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Configure the reverse proxy to use HTTPS
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create a new Proxy instance
	p := &Proxy{target: target, proxy: proxy}

	// Start the HTTP server and register the Proxy instance as the handler

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	err = http.ListenAndServe(listenAddr, p)
	if err != nil {
		panic(err)
	}
}
