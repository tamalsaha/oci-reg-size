package main

import (
	"encoding/json"
	"flag"
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

func main__() {
	//ref := "busybox"
	//data, err := crane.Manifest(ref)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(data))

	collect("kubedb/operator:0.8.0", map[string]int{})
}

func main() {
	dir := "/Users/tamal/go/src/kubedb.dev/installer/catalog/raw"
	dir = "./catalog/raw"
	flag.StringVar(&dir, "dir", dir, "Dir name")
	flag.Parse()

	fmt.Printf(dir)

	dm := map[string]int{}

	err := parser.ProcessPath(dir, func(ri parser.ResourceInfo) error {
		switch ri.Object.GetKind() {
		case api.ResourceKindElasticsearchVersion:
			var v api.ElasticsearchVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Dashboard.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.DashboardInitContainer.YQImage, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindMemcachedVersion:
			var v api.MemcachedVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindMariaDBVersion:
			var v api.MariaDBVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Coordinator.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindMongoDBVersion:
			var v api.MongoDBVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.ReplicationModeDetector.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindMySQLVersion:
			var v api.MySQLVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.ReplicationModeDetector.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Coordinator.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Router.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.RouterInitContainer.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindPerconaXtraDBVersion:
			var v api.PerconaXtraDBVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindPgBouncerVersion:
			var v api.PgBouncerVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.PgBouncer.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindProxySQLVersion:
			var v api.ProxySQLVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.Proxysql.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindRedisVersion:
			var v api.RedisVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Coordinator.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
		case api.ResourceKindPostgresVersion:
			var v api.PostgresVersion
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(ri.Object.UnstructuredContent(), &v)
			if err != nil {
				panic(err)
			}

			if err := collect(v.Spec.DB.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Coordinator.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.Exporter.Image, dm); err != nil {
				panic(err)
			}
			if err := collect(v.Spec.InitContainer.Image, dm); err != nil {
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sz := 0
	for _, v := range dm {
		sz += v
	}
	fmt.Printf("%d\n", sz)
}

func collect(ref string, dm map[string]int) error {
	if ref == "" {
		return nil
	}
	//if strings.ContainsRune(ref, ':') {
	//	ref += ":latest"
	//}

	fmt.Printf("%s\n", ref)

	data, err := crane.Manifest(ref)
	if err != nil {
		return err
	}

	var m ImageManifest
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	for _, manifest := range m.Manifests {
		err = collect(ref+"@"+manifest.Digest, dm)
		if err != nil {
			return err
		}
	}

	dm[m.Config.Digest] = m.Config.Size
	for _, layer := range m.Layers {
		dm[layer.Digest] = layer.Size
	}
	return nil
}
