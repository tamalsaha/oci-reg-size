package main

import (
	"kmodules.xyz/client-go/tools/parser"
	api "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
)

type ImageManifest struct {
	SchemaVersion int                     `json:"schemaVersion"`
	MediaType     string                  `json:"mediaType"`
	Manifests     []PlatformImageManifest `json:"manifests"`
	Config        ImageConfig             `json:"config"`
	Layers        []ImageLayer            `json:"layers"`
}

type PlatformImageManifest struct {
	MediaType string   `json:"mediaType"`
	Size      int      `json:"size"`
	Digest    string   `json:"digest"`
	Platform  Platform `json:"platform"`
}

type Platform struct {
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
}

type ImageConfig struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}

type ImageLayer struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}

func main() {
	dir := "/Users/tamal/go/src/kubedb.dev/installer/catalog/raw"
	err := parser.ProcessPath(dir, func(ri parser.ResourceInfo) error {
		switch ri.Object.GetKind() {
		case "PostgresVersion":
			var v api.PostgresVersion

		}

		if ri.Object.GetKind().GetNamespace() == "" {
			ri.Object.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, ri)
		return nil
	})
	if err != nil {
		panic(err)
	}

}
