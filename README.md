# graal SDK

Client libraries for the [graal Data Platform](https://graal.systems/en/docs) REST API
(`/api/v1`).

| Folder | Language | Content |
|---|---|---|
| [`go`](go) | Go | API client generated from the OpenAPI contract (module `github.com/graalsystems/sdk/go`). Used by the [Terraform provider](https://github.com/graalsystems/terraform-provider-graalsystems). |
| [`python`](python) | Python | API client generated from the OpenAPI contract (package `graalsystems`). |
| [`parquet-kms`](parquet-kms) | Java | `KmsClient` for Parquet modular encryption, for Spark jobs that read or write encrypted Parquet files with keys held by the platform. |

For a command-line client, use [graalctl](https://github.com/graalsystems/graalctl).

## Calling the API

Every request carries a bearer token from the installation's OpenID Connect server and the
tenant:

```
Authorization: Bearer <token>
X-Tenant: <tenant>
```

Resources use vendor media types (`application/vnd.graal.systems.v1.<type>+json`), edits use
JSON Patch (`application/json-patch+json`), and lists are paginated with `page` and `size`.
See the [API documentation](https://graal.systems/en/docs/api); each installation serves the
OpenAPI contract and a Swagger UI at `/api/v1/docs`.

### Go

```go
import (
    "context"

    sdk "github.com/graalsystems/sdk/go"
    "golang.org/x/oauth2/clientcredentials"
)

func listProjects(ctx context.Context) ([]sdk.Project, error) {
    auth := clientcredentials.Config{
        ClientID:     "my-application",
        ClientSecret: "...", // from a secret store
        TokenURL:     "https://sso.example.com/realms/acme/protocol/openid-connect/token",
    }
    client := sdk.NewAPIClient(&sdk.Configuration{
        HTTPClient: auth.Client(ctx),
        Servers:    sdk.ServerConfigurations{{URL: "https://graal.example.com/api/v1"}},
    })
    projects, _, err := client.ProjectAPI.FindProjects(ctx).XTenant("acme").Execute()
    return projects, err
}
```

### Python

```sh
pip install "git+https://github.com/graalsystems/sdk.git#subdirectory=python"
```

```python
import graalsystems
from graalsystems.api import project_api

configuration = graalsystems.Configuration(host="https://graal.example.com/api/v1")
configuration.access_token = token  # OpenID Connect access token

with graalsystems.ApiClient(configuration) as client:
    projects = project_api.ProjectApi(client).find_projects(x_tenant="acme")
```

### Parquet KMS client (Java)

```sh
mvn package
```

Then set `parquet.encryption.kms.client.class` to
`graal.systems.sdk.parquet.kms.InternalKmsClient` in the Spark job configuration.

## Generated code

The `go` and `python` clients are generated with [OpenAPI Generator](https://openapi-generator.tech)
from the platform's OpenAPI contract; do not edit them by hand. The contract only changes
additively, so existing calls keep working; regenerate the clients to use newer endpoints.

## License

Apache License 2.0, see [LICENSE](LICENSE).
