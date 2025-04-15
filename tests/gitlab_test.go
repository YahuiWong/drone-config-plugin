package tests

import (
	"encoding/base64"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"log"
	"testing"
)

// This example shows how to create a client with username and password.
func TestBasicAuthExample(t *testing.T) {
	git, err := gitlab.NewBasicAuthClient(
		"admin1",
		"Qaz@1231",
		gitlab.WithBaseURL("http://192.168.133.131:30080"),
	)
	if err != nil {
		t.Fatal(err)
	}

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found %d projects", len(projects))
}

// 添加一个token
// -/user_settings/personal_access_tokens?page=1
func TestBasicOAuthClient(t *testing.T) {
	git, err := gitlab.NewOAuthClient(
		"glpat-oxueSEK_LPsFPPznHQkF",
		gitlab.WithBaseURL("http://192.168.133.131:30080"),
	)
	if err != nil {
		t.Fatal(err)
	}

	gf := &gitlab.GetFileOptions{
		Ref: gitlab.Ptr("main"),
	}
	f, _, err := git.RepositoryFiles.GetFile("admin1/test", "README.md", gf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contains: %s", f.Content)
	sDec, base64err := base64.StdEncoding.DecodeString(f.Content)
	if base64err != nil {
		t.Logf("Error decoding string: %s ", base64err.Error())
	} else {
		t.Log(string(sDec))
	}
}
