package entity

import "time"

type IP struct {
	Id         int        `json:"id,omitempty"`
	Headers    *Header    `json:"headers"`
	Origin     string     `json:"origin"`
	StatusCode int        `json:"statusCode"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}

type Header struct {
	Accept                  string `json:"Accept,omitempty"`
	AcceptEncoding          string `json:"Accept-Encoding,omitempty"`
	AcceptLanguage          string `json:"Accept-Language,omitempty"`
	Host                    string `json:"Host,omitempty"`
	SecFetchDest            string `json:"Sec-Fetch-Dest,omitempty"`
	SecFetchMode            string `json:"Sec-Fetch-Mode,omitempty"`
	SecFetchSite            string `json:"Sec-Fetch-Site,omitempty"`
	SecFetchUser            string `json:"Sec-Fetch-User,omitempty"`
	UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests,omitempty"`
	UserAgent               string `json:"User-Agent,omitempty"`
	XAmznTraceID            string `json:"X-Amzn-Trace-Id,omitempty"`
}
