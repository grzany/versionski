package checker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/github"
)

/*
func main() {
	githubTag := &GithubTag{
		Owner:      "kubernetes",
		Repository: "kops",
	}
	githubTag.Fetch()
} */

// GithubTag is used to fetch version(tag) information from Github.
type GithubTag struct {
	// Owner and Repository are GitHub owner name and its repository name.
	// e.g., If you want to check https://github.com/tcnksm/ghr version
	// Repository is `ghr`, and Owner is `tcnksm`.
	Owner      string
	Repository string

	// URL & Token is used for GitHub Enterprise
	URL       string
	Token     string
	LatestTag string
}

func (g *GithubTag) newClient() *github.Client {
	client := github.NewClient(nil)
	if g.URL != "" {
		client.BaseURL, _ = url.Parse(g.URL)
	}
	return client
}

func (g *GithubTag) Validate() error {

	if len(g.Repository) == 0 {
		return fmt.Errorf("GitHub repository name must be set")
	}

	if len(g.Owner) == 0 {
		return fmt.Errorf("GitHub owner name must be set")
	}

	if g.URL != "" {
		if _, err := url.Parse(g.URL); err != nil {
			return fmt.Errorf("GitHub API Url invalid: %s", err)
		}
	}

	return nil
}

func (g *GithubTag) Fetch() (*FetchResponse, error) {

	//fr := newFetchResponse()

	// Create a client
	client := g.newClient()
	//tags, resp, err := client.Repositories.ListTags(context.Background(), g.Owner, g.Repository, nil)
	latest, resp, err := client.Repositories.GetLatestRelease(context.Background(), g.Owner, g.Repository)
	if err != nil {
		fmt.Println("error talking to github", resp)
	}
	fmt.Println(resp)

	/*if resp.StatusCode != 200 {
		return fr, fmt.Errorf("Unknown status: %d", resp.StatusCode)
	}

	for _, tag := range tags {
		v, err := version.NewVersion(tag.Name)
		if err != nil {
			continue
		}
		fr.Versions = append(fr.Versions, v)
	}

	return fr, nil
	*/
	g.LatestTag = *latest.TagName
	fmt.Println(*latest.TagName)
	//for _, tag := range tags {
	//		fmt.Println("Tag: ", *tag.Name)
	//	}
	return &FetchResponse{
		Version: g.LatestTag,
	}, nil
}
