package main

import (
	"context"
	"net/url"
	"path"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/mod/module"
)

type detailMod struct {
	module.Version
	// Computed fields
	User    string
	Project string
	URL     string
	Host    string
	Ignore  bool
}

func (m *detailMod) Parse(mod module.Version) {
	m.Version = mod

	parts := strings.SplitN(m.Version.Path, "/", 4)
	if len(parts) < 3 {
		return
	}
	host, user, project := parts[0], parts[1], parts[2]
	if user == "" || project == "" {
		Log.Printf("unable to parse %s", m.Version.Path)
	}

	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(user, project),
	}

	m.Host = host
	m.User = user
	m.Project = project
	m.URL = u.String()
	if ignorePrefix != nil && *ignorePrefix != "" && strings.HasPrefix(mod.Path, *ignorePrefix) {
		m.Ignore = true
	}
}

type license struct {
	Name string
	URL  string
}

// License will try to fetch a module's license from Github.
// An empty license struct will be returned if a license does not exist
// or cannot be fetched.
func (m *detailMod) License(ctx context.Context, gh *github.Client) (*license, error) {
	if m.User == "" || m.Project == "" {
		Log.Printf("cannot fetch license for %s", m.Version.Path)
		return &license{}, nil
	}

	if l, ok := cacheGet(m.URL); ok {
		return l, nil
	}

	switch m.Host {
	default:
		Log.Printf("only supports github.com modules, skipping %s", m.Version.Path)
		return &license{}, nil
	case "github.com":
		resp, _, err := gh.Repositories.License(ctx, m.User, m.Project)
		if err != nil && !strings.Contains(err.Error(), "404 Not Found") {
			// Anything other than a 404
			return nil, err
		}

		var lic license
		if resp != nil && resp.License != nil && resp.License.Name != nil {
			lic.Name = *resp.License.Name
		}
		if resp != nil && resp.License != nil && resp.License.URL != nil {
			lic.URL = *resp.License.URL
		}

		if err := cacheSet(m.URL, &lic); err != nil {
			Log.Printf("failed to cache %s: %w", m.URL, err)
		}

		return &lic, nil
	}
}
