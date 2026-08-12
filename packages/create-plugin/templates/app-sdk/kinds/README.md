# Kinds (grafana-app-sdk)

This directory declares your app's API as [CUE](https://cuelang.org/) "kinds", which
[grafana-app-sdk](https://github.com/grafana/grafana-app-sdk) turns into Go types, TypeScript types,
and an app manifest.

| File | Purpose |
| ---- | ------- |
| `manifest.cue` | The app manifest: app name, and the versions/kinds your app serves. |
| `example.cue` | An example kind. Rename it and edit its `spec` to model your own resource. |
| `config.cue` | Code generation settings (output paths). You rarely need to change this. |
| `cue.mod/module.cue` | The CUE module definition. |

## Generating code

First time only, add the app-sdk as a dependency:

```bash
go get github.com/grafana/grafana-app-sdk
```

Then, and after every change under this directory:

```bash
mage generate
go mod tidy
```

> Run `mage generate` **before** `go mod tidy`. Until generated code exists, nothing imports the
> app-sdk, so `tidy` would remove the dependency.

This writes:

| Output | Path |
| ------ | ---- |
| Go types, client, codecs | `pkg/generated/<kind>/<version>/` |
| Embedded Go manifest | `pkg/generated/manifestdata/` |
| TypeScript types | `src/generated/<kind>/<version>/` |
| Per-kind CRD (JSON) | `definitions/<plural>.<group>.json` |
| App manifest (JSON) | `definitions/<appName>-manifest.json` |

Generated code is meant to be committed, so schema changes show up in review and a fresh clone builds
without a CUE toolchain.

## How the manifest reaches Grafana

The frontend build copies `definitions/*-manifest.json` into the plugin bundle as
`dist/app-sdk-manifest.json` (configured in `.config/bundler/copyFiles.ts`). Grafana reads it when the
`plugins.appSDKManifest` feature toggle is enabled — the Docker dev server in this repo enables it for
you.

## Serving your kinds

Generating types does not by itself serve them. To do that, build an `app.App` in your plugin backend
using the generated `manifestdata.LocalManifest()` and register your kinds with it. See the
[app-sdk documentation](https://github.com/grafana/grafana-app-sdk/tree/main/docs) for details.
