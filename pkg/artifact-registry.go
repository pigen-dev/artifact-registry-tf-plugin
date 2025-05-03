package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/pigen-dev/artifact-registry-tf-plugin.git/helpers"
	"github.com/pigen-dev/artifact-registry-tf-plugin.git/pkg/terraform"
	shared "github.com/pigen-dev/shared"
	tfengine "github.com/pigen-dev/shared/tfengine"
)


type ArtifactRegistry struct {
	Label string `yaml:"label"`
	Config Config `yaml:"config"`
	Output Output `yaml:"output"`
}



type Config struct {
	Location string `yaml:"location"`
	RepoID string `yaml:"repo_id"`
	Description string `yaml:"description"`
	ProjectId string `yaml:"project_id"`
}

type Output struct {
	RepoUrl string `yaml:"repo_url"`
}


func (ar *ArtifactRegistry) Initializer(plugin shared.Plugin) (*tfengine.Terraform ,error) {
	config := Config{}
	err:= helpers.YamlConfigParser(plugin.Config, &config)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse YAML config: %v", err)
	}
	ar.Config = config
	ar.Label = plugin.Label
	fmt.Println("Parsed config:", ar)
	// Initialize Terraform
	files := terraform.LoadTFFiles()
	tfVars := plugin.Config
	t, err := tfengine.NewTF(tfVars, files, ar.Label)
	if err != nil {
		return nil, fmt.Errorf("Failed to setup Terraform executor: %v", err)
	}
	
	return t, nil
}



func (ar *ArtifactRegistry) SetupPlugin(plugin shared.Plugin) error {
	tf, err := ar.Initializer(plugin)
	ctx := context.Background()
	if err != nil {
		return fmt.Errorf("Failed to initialize plugin: %v", err)
	}

	// 1. Initialize Terraform
	fmt.Println(ar.Label)
	if err := tf.TerraformInit(ctx, ar.Config.ProjectId, ar.Label); err != nil {
		return fmt.Errorf("Error during Terraform init: %v", err)
	}

	// 2. Plan Terraform changes
	if err := tf.TerraformPlan(ctx); err != nil {
		return fmt.Errorf("Error during Terraform plan: %v", err)
	}

	
	if err := tf.TerraformApply(ctx); err != nil {
		return fmt.Errorf("Error during Terraform apply: %v", err)
	}
	log.Println("Terraform apply completed.")
	return nil
}


func (ar *ArtifactRegistry) GetOutput(plugin shared.Plugin) shared.GetOutputResponse {
	tf, err := ar.Initializer(plugin)
	if err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Failed to initialize plugin: %v", err)}
	}
	ctx := context.Background()
	
	// 1. Initialize Terraform
	if err := tf.TerraformInit(ctx, ar.Config.ProjectId, ar.Label); err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Error during Terraform init: %v", err)}
	}

	output, err := tf.TerraformOutput(ctx)
	if err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Error during Terraform output: %v", err)}
	}
	log.Println("Terraform output retrieved successfully.")
	return shared.GetOutputResponse{Output: output, Error: nil}
}


func (ar *ArtifactRegistry) Destroy(plugin shared.Plugin) error {
	tf, err := ar.Initializer(plugin)
	if err != nil {
		return fmt.Errorf("Failed to initialize plugin: %v", err)
	}
	ctx := context.Background()
	// 1. Initialize Terraform
	if err := tf.TerraformInit(ctx, ar.Config.ProjectId, ar.Label); err != nil {
		return fmt.Errorf("Error during Terraform init: %v", err)
	}

	if err := tf.TerraformDestroy(ctx); err != nil {
		return fmt.Errorf("Error during Terraform destroy: %v", err)
	}
	log.Println("Terraform destroy completed.")
	return nil
}