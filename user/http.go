package user

import (
	"net/url"
	"net/http"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

func SetJarCook(c *http.Client) {
	urls,_ := url.Parse(`http://ttservice.toplogis.com/`)
	cook := make([]*http.Cookie,0)
	cook = append(cook, getCookies(OneConfig.SessionId))
	c.Jar.SetCookies(urls,cook)
}

func getCookies(SessionId string) *http.Cookie {
	return &http.Cookie{
		Name:   "JSESSIONID",
		Value:    SessionId,
		Path:     "/",
		HttpOnly: false,
		MaxAge:9999999999,
	}
}

func getClient() *http.Client  {
	return &http.Client{
		Jar:new(Jar),
		Timeout:9999999999,
	}
}

func setHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.117 Safari/537.36")
	req.Header.Set("Upgrade-Insecure-Requests","1")
	req.Header.Set("Cache-Control","0")
	req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
}
