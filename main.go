package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"golang.org/x/mod/modfile"
	"golang.org/x/oauth2"
)

// Depends on GITHUB_TOKEN env var to increase re rate limit to from
// 60req/h to 5000req/h

var (
	// Args
	ignorePrefix = flag.String("i", "", "module prefixes to ignore")
)

var Log = log.New(os.Stderr, "", 0)

func main() {
	flag.Parse()

	info, err := GetGoModInfo("./")
	if err != nil {
		Log.Fatal(err)
	}

	table := [][]string{
		{"Dependency", "Version", "URL", "License", "LicenseURL"},
	}

	ctx := context.Background()
	gh := ghClient(ctx)

	for _, require := range info.Require {
		var dm detailMod
		dm.Parse(require.Mod)

		if !dm.Ignore {
			lic, err := dm.License(ctx, gh)
			if err != nil {
				Log.Println(err)
				continue
			}

			table = append(table,
				[]string{dm.Mod.Path, dm.Mod.Version, dm.URL, lic.Name, lic.URL},
			)
		}
	}

	w := csv.NewWriter(os.Stdout)
	for _, row := range table {
		if err := w.Write(row); err != nil {
			Log.Println(err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		Log.Fatal(err)
	}
}

func ghClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func GetGoModInfo(repoDir string) (*modfile.File, error) {
	goModPath := filepath.Join(repoDir, "go.mod")

	goModData, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return nil, fmt.Errorf("could not read go.mod file at [%s]: %w", goModPath, err)
	}

	return modfile.Parse(goModPath, goModData, nil)
}
