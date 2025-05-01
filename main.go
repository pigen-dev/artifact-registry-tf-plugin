package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/pigen-dev/artifact-registry-tf-plugin.git/pkg"
	shared "github.com/pigen-dev/shared"
)

func main() {
	// data := map[string]any{
	// 	"location":    "europe-west1",
	// 	"repo_id":     "my-repo-id",
	// 	"description": "My Artifact Registry Repository",
	// 	"project_id":  "aidodev",
	// }

	// ar := pkg.ArtifactRegistry{
	// 	Label: "my-artifact-registry",
	// }
	// // err := ar.SetupPlugin(data)
	// // if err != nil {
	// // 	fmt.Println("Error:", err)
	// // }

	// out := ar.GetOutput(data)
	// if out.Error != nil {
	// 	fmt.Println("Error:", out.Error)
	// 	return
	// }
	// fmt.Println("output:", out.Output)

	// // err := ar.Destroy(data)
	// // if err != nil {
	// // 	fmt.Println("Error:", err)
	// // }
	// fmt.Println("\n--- All Terraform operations finished ---")

	artifactPlugin := &pkg.ArtifactRegistry{}
	pluginMap := map[string]plugin.Plugin{"pigenPlugin": &shared.PigenPlugin{Impl: artifactPlugin}}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         pluginMap,
	})
}