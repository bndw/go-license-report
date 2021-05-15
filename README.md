# go-license-report

Generates a CSV report of dependencies and their licenses by reading
the `go.mod` file. Dependencies missing a license are included in the output.

Example output:
```csv
Dependency,Version,URL,License,LicenseURL
github.com/google/go-github,v17.0.0+incompatible,https://github.com/google/go-github,"BSD 3-Clause "New"" or ""Revised"" License",https://api.github.com/licenses/bsd-3-clause
github.com/google/go-querystring,v1.0.0,https://github.com/google/go-querystring,"BSD 3-Clause "New"" or ""Revised"" License",https://api.github.com/licenses/bsd-3-clause
github.com/stretchr/testify,v1.3.0,https://github.com/stretchr/testify,MIT License,https://api.github.com/licenses/mit
golang.org/x/mod,v0.4.2,https://golang.org/x/mod,,
golang.org/x/oauth2,v0.0.0-20190402181905-9f3314589c9a,https://golang.org/x/oauth2,,
```

## Quickstart 

1. Create a Github Token. Without one, you're limited to 60 API calls to Github for an hour.
2. Export your Github token
    ```
    export GITHUB_TOKEN=<your_token>
    ```
3. Run it in the root of a Go project
    ```
    go-license-report
    ```

If you want to ignore certain dependencies, like internal org dependencies, 
use the `-ignore` flag:

```
go-license-report -ignore github.com/myUser
```

If you want to fail when a dependency does not have a license, use the 
`-strict` flag.

```
go-license-report -strict
```

Use the `-help` flag for detailed usage information.

## Similar projects

The below table compares go-license-report to similar projects.

| Project | Module Name | Module Version | Module URL | Module License | Module License URL |
| ------- | ----------- | -------------- | ---------- | -------------- | ------------------ |
| [go-license-report](https://github.com/bndw/go-license-report) | X | X | X | X | X |
| [Glice](https://github.com/ribice/glice) | X |  | X | X |  |
| [go-licenses](https://github.com/google/go-licenses/) | X |  |  | X | X |
