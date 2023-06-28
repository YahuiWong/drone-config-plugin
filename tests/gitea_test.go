package tests

import (
	"encoding/base64"
	"log"
	"net/url"
	"testing"

	"code.gitea.io/sdk/gitea"
)

func TestBasicAuth(t *testing.T) {
	client, _ := gitea.NewClient("http://git.local.lan:3000")

	client.SetBasicAuth("adminuser", "admin@123")
	// client.CreateAccessToken(gitea.CreateAccessTokenOption{})
	res, _, _ := client.SearchRepos(gitea.SearchRepoOptions{
		Keyword: "he",
	})
	for i, count := 0, len(res); i < count; i++ {
		t.Log((*res[i]).Owner.UserName)
		t.Log((*res[i]).Name)
		contentsres, gres, gerr := client.GetContents(*&res[i].Owner.UserName, *&res[i].Name, *&res[i].DefaultBranch, ".drone.yml")
		if gerr != nil {
			t.Log(gerr)
			t.Log(gres)
		}
		t.Log(*(*contentsres).Content)
		sDec, base64err := base64.StdEncoding.DecodeString(*(*contentsres).Content)
		if base64err != nil {
			t.Logf("Error decoding string: %s ", base64err.Error())
		} else {
			t.Log(string(sDec))
		}
	}
}

func TestTokenAuth(t *testing.T) {
	clilentOption := gitea.SetToken("a9e569dbb06d301267f17ec9a78e71b04f0dbe0c")

	client, _ := gitea.NewClient("http://git.local.lan:3000", clilentOption)

	res, _, _ := client.SearchRepos(gitea.SearchRepoOptions{
		Keyword: "he",
	})
	for i, count := 0, len(res); i < count; i++ {
		t.Log((*res[i]).Owner.UserName)
		t.Log((*res[i]).Name)
		contentsres, gres, gerr := client.GetContents(*&res[i].Owner.UserName, *&res[i].Name, *&res[i].DefaultBranch, ".drone.yml")
		if gerr != nil {
			t.Log(gerr)
			t.Log(gres)
		}
		t.Log(*(*contentsres).Content)
		sDec, base64err := base64.StdEncoding.DecodeString(*(*contentsres).Content)
		if base64err != nil {
			t.Logf("Error decoding string: %s ", base64err.Error())
		} else {
			t.Log(string(sDec))
		}
	}
}

func TestUrl(t *testing.T) {

	input_url := "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"
	u, err := url.Parse(input_url)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(u.Scheme)
	t.Log(u.User)
	t.Log(u.Host)
	t.Log(u.Hostname())
	t.Log(u.Port())
	t.Log(u.Path)
	t.Log(u.RawQuery)
	t.Log(u.Fragment)
	t.Log(u.String())
	t.Logf("%s://%s", u.Scheme, u.Host)

}
