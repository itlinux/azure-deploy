package test

import (
	"testing"

	//            "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"

	//            "github.com/gruntwork-io/terratest/modules/azure"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzure3NicExample(t *testing.T) {

	t.Parallel()

	terraformOptions := &terraform.Options{

		// The path to where our Terraform code is located

		TerraformDir: "../examples/bigip-3nic-deploy",

		Vars: map[string]interface{}{
			"location": "southindia",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//      vmName := terraform.Output(t, terraformOptions, "bigip_public_ip")

	bigipPublicDnsName := terraform.Output(t, terraformOptions, "bigip_public_dns_name")

	bigipUserName := terraform.Output(t, terraformOptions, "bigip_username")

	assert.Equal(t, "azureuser", bigipUserName)
	assert.Equal(t, "[\n  \"ecosysf5hyd3.southindia.cloudapp.azure.com\",\n  \"external.southindia.cloudapp.azure.com\",\n  \"internal.southindia.cloudapp.azure.com\",\n]", bigipPublicDnsName)

}
