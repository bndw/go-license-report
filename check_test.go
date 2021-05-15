package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/mod/module"
)

func TestDetailMod(t *testing.T) {
	tests := []struct {
		name       string
		mod        module.Version
		expUser    string
		expProject string
		expURL     string
	}{
		{
			name: "basic",
			mod: module.Version{
				Path:    "github.com/fatih/color",
				Version: "v1.7.0",
			},
			expUser:    "fatih",
			expProject: "color",
			expURL:     "https://github.com/fatih/color",
		},
		{
			name: "subpaths",
			mod: module.Version{
				Path:    "github.com/dgraph-io/dgo/v200",
				Version: "v200.0.0-20200401175452-e463f9234453",
			},
			expUser:    "dgraph-io",
			expProject: "dgo",
			expURL:     "https://github.com/dgraph-io/dgo",
		},
		{
			name: "toplevel",
			mod: module.Version{
				Path:    "google.golang.org/grpc",
				Version: "v1.29.1",
			},
			expUser:    "",
			expProject: "",
			expURL:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dm detailMod
			dm.Parse(tt.mod)

			assert.Equal(t, tt.mod.Path, dm.Version.Path, "unexpected module path")
			assert.Equal(t, tt.mod.Version, dm.Version.Version, "unexpected module version")
			assert.Equal(t, tt.expUser, dm.User, "unexpected User")
			assert.Equal(t, tt.expProject, dm.Project, "unexpected Project")
			assert.Equal(t, tt.expURL, dm.URL, "unexpected URL")
		})
	}
}
