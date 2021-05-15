# go-license-report

Generates a CSV report of dependencies and their licenses.

Example output:
```csv
Dependency,Version,URL,License,LicenseURL
github.com/cenkalti/backoff,v2.2.1+incompatible,https://github.com/cenkalti/backoff,MIT License,https://api.github.com/licenses/mit
github.com/gin-gonic/gin,v1.6.3,https://github.com/gin-gonic/gin,MIT License,https://api.github.com/licenses/mit
github.com/google/uuid,v1.1.1,https://github.com/google/uuid,"BSD 3-Clause "New"" or ""Revised"" License",https://api.github.com/licenses/bsd-3-clause
github.com/kelseyhightower/envconfig,v1.4.0,https://github.com/kelseyhightower/envconfig,MIT License,https://api.github.com/licenses/mit
github.com/matryer/moq,v0.0.0-20191106032847-0e0395200ade,https://github.com/matryer/moq,MIT License,https://api.github.com/licenses/mit
github.com/pkg/errors,v0.9.1,https://github.com/pkg/errors,"BSD 2-Clause "Simplified"" License",https://api.github.com/licenses/bsd-2-clause
github.com/sirupsen/logrus,v1.4.2,https://github.com/sirupsen/logrus,MIT License,https://api.github.com/licenses/mit
github.com/stretchr/testify,v1.7.0,https://github.com/stretchr/testify,MIT License,https://api.github.com/licenses/mit
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
use the `-i` flag:

```
go-license-report -i github.com/myUser
```

Use the `-help` flag for detailed usage information.

## Similar projects

The below table compares go-license-report to similar projects.

| Project | Module Name | Module Version | Module URL | Module License | Module License URL |
| ------- | ----------- | -------------- | ---------- | -------------- | ------------------ |
| [go-license-report](https://github.com/bndw/go-license-report) | X | X | X | X | X |
| [Glice](https://github.com/ribice/glice) | X |  | X | X |  |
| [go-licenses](https://github.com/google/go-licenses/) | X |  |  | X | X |
