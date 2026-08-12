package kinds

// config holds the grafana-app-sdk code generation settings, and is read by
// `grafana-app-sdk generate` (via the default `-c config` selector) when you run `mage generate`.
//
// Paths are relative to the plugin root and match the layout scaffolded by @grafana/create-plugin:
// the frontend lives in src/, the Go backend in pkg/.
config: {
	definitions: {
		// Emit the app manifest alongside the per-kind CRD definitions. The manifest is what tells
		// Grafana which kinds and capabilities your app serves.
		manifestSchemas: true
		manifestVersion: "v1alpha2"
		// CRD + manifest JSON output directory. The frontend build copies the manifest from here
		// into the plugin bundle as app-sdk-manifest.json (see .config/bundler/copyFiles.ts).
		path:     "definitions"
		encoding: "json"
	}

	kinds: {
		// Group generated Go packages by kind: pkg/generated/<kind>/<version>.
		grouping: "kind"
	}

	codegen: {
		// Generated Go types land next to the plugin backend in pkg/.
		goGenPath: "pkg/generated/"
		// Generated TypeScript types land in the frontend source dir.
		tsGenPath:                      "src/generated/"
		enableK8sPostProcessing:        false
		enableOperatorStatusGeneration: true
	}
}
