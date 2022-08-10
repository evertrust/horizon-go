# Go Horizon

The official Go SDK for Horizon.

## Installation

With the Go module system, just reference this module anywhere in your code :

```go
import (
"github.com/evertrust/horizon-go
)
```

You can also explicitly require the package in your existing project using `go get` :

```shell
go get -u "github.com/evertrust/horizon-go"
```

## Breaking changes policy

The `horizon-go` project follows the semver conventions, meaning that once 1.y.z is reached, y and z versions will not
introduce any breaking changes. In the meantime, the API will mostly remain backwards-compatible but may break from time
to time.

When using a recent SDK version with an older Horizon version, you may encounter `NotImplementedError`s, which
indicates that the feature you're trying to use is not available on the targeted Horizon instance. The following
compatibility matrix provides you with what SDK versions have been tested against which Horizon versions :

| SDK version | Horizon version |
|-------------|-----------|
| 0.0.1 | >= 2.1.0 |
| 0.0.2 | >= 2.2.0 |
| 0.0.4 | >= 2.2.2 |
