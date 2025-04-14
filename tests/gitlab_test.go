package tests

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
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

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found %d projects", len(projects))

}
