package cloudconfig

import (
	"errors"

	"github.com/cloudfoundry/bosh-bootloader/storage"
)

type OpsGenerator struct {
	awsCloudFormationOpsGenerator opsGenerator
	awsTerraformOpsGenerator      opsGenerator
	gcpOpsGenerator               opsGenerator
}

func NewOpsGenerator(awsCloudFormationOpsGenerator opsGenerator, awsTerraformOpsGenerator opsGenerator, gcpOpsGenerator opsGenerator) OpsGenerator {
	return OpsGenerator{
		awsCloudFormationOpsGenerator: awsCloudFormationOpsGenerator,
		awsTerraformOpsGenerator:      awsTerraformOpsGenerator,
		gcpOpsGenerator:               gcpOpsGenerator,
	}
}

func (o OpsGenerator) Generate(state storage.State) (string, error) {
	switch state.IAAS {
	case "gcp":
		return o.gcpOpsGenerator.Generate(state)
	case "aws":
		if state.TFState != "" {
			return o.awsTerraformOpsGenerator.Generate(state)
		} else {
			return o.awsCloudFormationOpsGenerator.Generate(state)
		}
	default:
		return "", errors.New("invalid iaas type")
	}
}
