package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	"k8s.io/apimachinery/pkg/runtime"
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
		case api.ResourceKindPostgresVersion:
			var v api.PostgresVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			data, err := crane.Manifest(v.Spec.DB.Image)
			if err != nil {
				panic(err)
			}

			var m ImageManifest
			err = json.Unmarshal(data, &m)
			if err != nil {
				panic(err)
			}

			for _, manifest := range m.Manifests {
				fmt.Println(manifest.Size)
			}

			sz := m.Config.Size
			for _, layer := range m.Layers {
				sz += layer.Size
			}
			fmt.Println(sz)
			break
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
