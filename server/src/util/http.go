package util

import (
	"net/http"
)

var noProxyTransport = &http.Transport{
	Proxy: nil, // 设置为 nil 表示不使用代理
}

var noProxyClient = &http.Client{
	Transport: noProxyTransport,
}

var HTTPNoProxyGet = noProxyClient.Get
