package graphql // import "flamingo.me/graphql"

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strings"

	"flamingo.me/graphql/templates"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/spf13/cobra"
)

const (
	basePath       = "graphql"
	schemaBasePath = "graphql/schema"
)

func command(
	services []Service,
) *cobra.Command {
	return &cobra.Command{
		Use:     "graphql",
		Short:   "(Re-)Generate graphql module",
		Aliases: []string{"g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Generate(services, basePath, schemaBasePath)
		},
	}
}

// Generate runs the graphql generation for the defined services
func Generate(services []Service, basePath string, schemaBasePath string) error {
	schemaPath := path.Join(schemaBasePath, "schema.graphql")

	cfg := config.DefaultConfig()
	cfg.SchemaFilename = []string{schemaPath}
	cfg.Models = make(map[string]config.TypeMapEntry)

	if err := os.MkdirAll(schemaBasePath, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	if err := ioutil.WriteFile(schemaPath, templates.MustAsset("schema.graphql"), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(basePath, "module.go"), templates.MustAsset("module.go.tpl"), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(basePath, "emptymodule.go"), templates.MustAsset("emptymodule.go.tpl"), 0644); err != nil {
		return err
	}

	resolverPath := path.Join(basePath, "resolver.go")

	if _, err := os.Stat(resolverPath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(resolverPath, templates.MustAsset("resolver.go.tpl"), 0644); err != nil {
			return err
		}
	}

	for _, service := range services {
		rt := reflect.TypeOf(service).Elem()
		fname := strings.Replace(rt.PkgPath(), "/", "_", -1) + "-" + rt.Name() + ".graphql"
		fpath := path.Join(schemaBasePath, fname)

		log.Printf("Writing %s", fname)
		if err := ioutil.WriteFile(fpath, service.Schema(), 0644); err != nil {
			return err
		}
		cfg.SchemaFilename = append(cfg.SchemaFilename, fpath)

		// merge models into config models
		for k, v := range service.Models() {
			cfg.Models[k] = v
		}
	}

	float := cfg.Models["Float"]
	float.Model = append(float.Model, "flamingo.me/graphql.Float", "github.com/99designs/gqlgen/graphql.Float")
	cfg.Models["Float"] = float

	cfg.Model = config.PackageConfig{Filename: path.Join(basePath, "model_gen.go")}
	cfg.Exec = config.PackageConfig{Filename: path.Join(basePath, "generated.go")}

	return api.Generate(cfg)

}
