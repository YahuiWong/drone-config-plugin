// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugin

import (
	"context"
	"drone-config-plugin/src/utils"
	"encoding/base64"
	"errors"
	"fmt"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"net/url"
	"strings"

	"code.gitea.io/sdk/gitea"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

// New returns a new configuration plugin.
func New(DroneConfigNamespaceTemp, DroneConfigRepoNameTemp, DroneConfigPathTemp, DroneConfigBranchTemp, servertype, token string) config.Plugin {
	return &plugin{
		DroneConfigNamespaceTemp: DroneConfigNamespaceTemp,
		DroneConfigRepoNameTemp:  DroneConfigRepoNameTemp,
		DroneConfigPathTemp:      DroneConfigPathTemp,
		DroneConfigBranchTemp:    DroneConfigBranchTemp,
		token:                    token,
		servertype:               servertype,
	}
}

type plugin struct {
	DroneConfigNamespaceTemp string
	DroneConfigRepoNameTemp  string
	DroneConfigPathTemp      string
	DroneConfigBranchTemp    string
	token                    string
	servertype               string
}

func (p *plugin) Find(ctx context.Context, req *config.Request) (*drone.Config, error) {

	namespace, _ := utils.GetTempString(p.DroneConfigNamespaceTemp, &req)
	reponame, _ := utils.GetTempString(p.DroneConfigRepoNameTemp, &req)
	branch, _ := utils.GetTempString(p.DroneConfigBranchTemp, &req)
	path, _ := utils.GetTempString(p.DroneConfigPathTemp, &req)
	logrus.Debugf("namespace:%s reponame:%s branch:%s path:%s", namespace, reponame, branch, path)
	switch strings.ToLower(p.servertype) {
	case "github":
		return p.HandeGithub(ctx, req)
		break
	case "gitlab":
		return p.HandeGitlab(ctx, req)
		break
	case "gitea":
		return p.HandeGitea(ctx, req)
		break
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported servertype:%s", p.servertype))
		break
	}

	return &drone.Config{
		Data: "",
	}, nil
}

// HandeGithub
func (p *plugin) HandeGithub(ctx context.Context, req *config.Request) (*drone.Config, error) {
	namespace, _ := utils.GetTempString(p.DroneConfigNamespaceTemp, &req)
	reponame, _ := utils.GetTempString(p.DroneConfigRepoNameTemp, &req)
	branch, _ := utils.GetTempString(p.DroneConfigBranchTemp, &req)
	path, _ := utils.GetTempString(p.DroneConfigPathTemp, &req)
	logrus.Debugf("namespace:%s reponame:%s branch:%s path:%s", namespace, reponame, branch, path)
	// creates a github client used to fetch the yaml.
	trans := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: p.token},
	))
	client := github.NewClient(trans)

	// HACK: the drone-go library does not currently work
	// with 0.9 which means the configuration file path is
	// always empty. default to .drone.yml. This can be
	// removed as soon as drone-go is fully updated for 0.9.

	// get the configuration file from the github
	// repository for the build ref.
	data, _, _, err := client.Repositories.GetContents(ctx, namespace, reponame, path, &github.RepositoryContentGetOptions{Ref: branch})
	if err == nil && data != nil {
		// get the file contents.
		content, err := data.GetContent()
		if err != nil {
			return nil, err
		}
		return &drone.Config{
			Data: content,
		}, nil
	}

	// if the configuration file does not exist,
	// we should fallback to a global configuration
	// file stored in a central repository.
	data, _, _, err = client.Repositories.GetContents(ctx, namespace, reponame, path, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		return nil, err
	}
	// get the file contents.
	content, err := data.GetContent()
	if err != nil {
		return nil, err
	}
	return &drone.Config{
		Data: content,
	}, nil
}

// HandeGitlab
func (p *plugin) HandeGitlab(ctx context.Context, req *config.Request) (*drone.Config, error) {
	namespace, _ := utils.GetTempString(p.DroneConfigNamespaceTemp, &req)
	reponame, _ := utils.GetTempString(p.DroneConfigRepoNameTemp, &req)
	branch, _ := utils.GetTempString(p.DroneConfigBranchTemp, &req)
	path, _ := utils.GetTempString(p.DroneConfigPathTemp, &req)
	logrus.Debugf("namespace:%s reponame:%s branch:%s path:%s", namespace, reponame, branch, path)
	httpUrl, urlerr := url.Parse(req.Repo.HTTPURL)
	if urlerr != nil {
		logrus.Debugf("urlerr:%s ", urlerr)
		return nil, urlerr
	}

	git, err := gitlab.NewOAuthClient(
		p.token,
		gitlab.WithBaseURL(fmt.Sprintf("%s://%s", httpUrl.Scheme, httpUrl.Host)),
	)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	gf := &gitlab.GetFileOptions{
		Ref: gitlab.Ptr(branch),
	}
	f, _, err := git.RepositoryFiles.GetFile(fmt.Sprintf("%s/%s", namespace, reponame), path, gf)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	logrus.Printf("File contains: %s", f.Content)
	sDec, base64err := base64.StdEncoding.DecodeString(f.Content)
	if base64err != nil {
		logrus.Debugf("Error decoding string: %s ", base64err.Error())
	} else {
		logrus.Debugf("Data:%s ", string(sDec))
		return &drone.Config{
			Data: string(sDec),
		}, nil
	}
	return &drone.Config{
		Data: "",
	}, nil
}

// HandeGitea
func (p *plugin) HandeGitea(ctx context.Context, req *config.Request) (*drone.Config, error) {
	namespace, _ := utils.GetTempString(p.DroneConfigNamespaceTemp, &req)
	reponame, _ := utils.GetTempString(p.DroneConfigRepoNameTemp, &req)
	branch, _ := utils.GetTempString(p.DroneConfigBranchTemp, &req)
	path, _ := utils.GetTempString(p.DroneConfigPathTemp, &req)
	logrus.Debugf("namespace:%s reponame:%s branch:%s path:%s", namespace, reponame, branch, path)
	clilentOption := gitea.SetToken(p.token)
	httpUrl, urlerr := url.Parse(req.Repo.HTTPURL)
	if urlerr != nil {
		logrus.Debugf("urlerr:%s ", urlerr)
		return nil, urlerr
	}

	client, newClientERR := gitea.NewClient(fmt.Sprintf("%s://%s", httpUrl.Scheme, httpUrl.Host), clilentOption)
	if newClientERR != nil {
		logrus.Debugf("newClientERR:%s ", newClientERR)
		return nil, newClientERR
	}
	_, _, repoerr := client.GetRepo(namespace, reponame)
	if repoerr != nil {
		logrus.Debugf("repoerr:%s ", repoerr)
		return nil, repoerr
	}
	contentsres, _, getcerr := client.GetContents(namespace, reponame, branch, path)
	if getcerr != nil {
		logrus.Debugf("getcerr:%s ", getcerr)
		return nil, getcerr
	}
	sDec, base64err := base64.StdEncoding.DecodeString(*(*contentsres).Content)
	if base64err != nil {
		logrus.Debugf("base64err:%s ", base64err)
		return nil, base64err
	} else {
		logrus.Debugf("Data:%s ", string(sDec))
		return &drone.Config{
			Data: string(sDec),
		}, nil
	}
	return &drone.Config{
		Data: "",
	}, nil
}
