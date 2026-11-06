# idp-test-app

A tiny Go web app used to test the RuruIDP / ArgoCD Image Updater flows. It serves a
page showing its build version and the pod serving it, plus `/healthz` and `/version`.

## Image

Built and pushed by GitHub Actions to the GitHub Container Registry:

```
ghcr.io/rurueuh/idp-test-app
```

Tags produced:

| Trigger | Tags |
|---|---|
| push to `main` | `latest`, `sha-<short>` |
| git tag `vX.Y.Z` | `X.Y.Z`, `X.Y`, `X` |

## Releasing a new version (for semver tests)

```sh
git tag v0.2.0
git push origin v0.2.0
```

The workflow then publishes `ghcr.io/rurueuh/idp-test-app:0.2.0` (and `0.2`, `0`).

## Testing update strategies

- **semver**: track `ghcr.io/rurueuh/idp-test-app` with a constraint (e.g. `^0.1`) and
  push new `vX.Y.Z` tags to roll forward.
- **newest-build / latest**: track the `latest` tag by digest; every push to `main`
  moves it.
- **digest**: pin a mutable tag (e.g. `latest` or `main`) and update on digest change.

## Run locally

```sh
go run .
# or
docker build -t idp-test-app --build-arg VERSION=local .
docker run -p 8080:8080 idp-test-app
```
