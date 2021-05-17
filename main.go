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
	ignorePrefix = flag.String("ignore", "", "module prefixes to ignore")
	strict       = flag.Bool("strict", false, "fail if a dependency does not have a license")
)

var Log = log.New(os.Stderr, "", 0)

func main() {
	flag.Parse()
	args := flag.Args()

	modDir := "./"
	if len(args) == 1 {
		modDir = args[0]
	}

	modFile, err := readModFile(modDir)
	if err != nil {
		Log.Fatal(err)
	}

	ctx := context.Background()
	gh := ghClient(ctx)

	table := [][]string{
		{"Dependency", "Version", "URL", "License", "LicenseURL"},
	}

	for _, require := range modFile.Require {
		var dm detailMod
		dm.Parse(require.Mod)

		if dm.Ignore {
			continue
		}

		lic, err := dm.License(ctx, gh)
		if err != nil {
			Log.Println(err)
			continue
		}

		if lic.Name == "" && *strict {
			Log.Fatalf("missing license (strict mode): %s", dm.Version.Path)
		}

		row := []string{dm.Version.Path, dm.Version.Version, dm.URL, lic.Name, lic.URL}
		table = append(table, row)
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

func readModFile(repoDir string) (*modfile.File, error) {
	goModPath := filepath.Join(repoDir, "go.mod")

	goModData, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return nil, fmt.Errorf("could not read go.mod file at [%s]: %w", goModPath, err)
	}

	return modfile.Parse(goModPath, goModData, nil)
}
