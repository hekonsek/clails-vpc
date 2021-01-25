package clails_vpc

import (
	"testing"
)
import "github.com/gruntwork-io/terratest/modules/terraform"
import "github.com/stretchr/testify/assert"

func TestOutputs(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "test",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Then
	vpcId := terraform.OutputRequired(t, terraformOptions, "vpc_id")
	assert.Contains(t, vpcId, "vpc-")

	vpcPrivateSubnets := terraform.OutputRequired(t, terraformOptions, "vpc_private_subnets")
	assert.Contains(t, vpcPrivateSubnets, "subnet-")
}
