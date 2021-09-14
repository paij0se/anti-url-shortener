package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var (
	proxy *proxyData = nil
)

const (
	proxysApi = "https://proxylist.geonode.com/api/proxy-list?limit=50&page=1&sort_by=lastChecked&sort_type=desc&speed=fast&protocols=http%2Chttps"
)

func GetProxy() {

	rand.Seed(time.Now().UnixNano())
	if proxy != nil {

		u := fmt.Sprintf("%s:%s/", proxy.IP, proxy.Port)
		urlProxy, _ := url.Parse(u)
		http.DefaultClient.Transport = &http.Transport{Proxy: http.ProxyURL(urlProxy)}
	}
	resp, err := http.Get(proxysApi)
	if err != nil {
		return
	}

	var proxysApi proxyList
	json.NewDecoder((resp.Body)).Decode(&proxysApi)
	proxy = &proxysApi.Data[rand.Intn(len(proxysApi.Data))]

}
