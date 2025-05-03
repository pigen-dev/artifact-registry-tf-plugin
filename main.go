package main

import (
	"github.com/hashicorp/go-plugin"
	//"fmt"
	"github.com/pigen-dev/artifact-registry-tf-plugin.git/pkg"
	shared "github.com/pigen-dev/shared"
)

func main() {
	// data := map[string]any{
	// 		"location":    "europe-west1",
	// 		"repo_id":     "pi-gen-testing2",
	// 		"description": "My Artifact Registry Repository",
	// 		"project_id":  "aidodev",
	// }

	// ar := pkg.ArtifactRegistry{}
	// plugin := shared.Plugin{
	// 	Label: "artifact-registry",
	// 	Config: data,
	// }
	// // err := ar.SetupPlugin(plugin)
	// // if err != nil {
	// // 	fmt.Println("Error:", err)
	// // }
	// // fmt.Println("artifact registry label:", ar.Label)
	// // out := ar.GetOutput(plugin)
	// // if out.Error != nil {
	// // 	fmt.Println("Error:", out.Error)
	// // 	return
	// // }
	// // fmt.Println("output:", out.Output)

	// err := ar.Destroy(plugin)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("\n--- All Terraform operations finished ---")

	artifactPlugin := &pkg.ArtifactRegistry{}
	pluginMap := map[string]plugin.Plugin{"pigenPlugin": &shared.PigenPlugin{Impl: artifactPlugin}}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         pluginMap,
	})
}