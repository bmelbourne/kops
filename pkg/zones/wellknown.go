/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package zones

import (
	"sort"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"k8s.io/kops/pkg/apis/kops"
)

// These lists allow us to infer from certain well-known zones to a cloud
// Note it is safe to "overmap" zones that don't exist: we'll check later if the zones actually exist

var gceZones = []string{
	"asia-east1-a",
	"asia-east1-b",
	"asia-east1-c",
	"asia-east1-d",

	"asia-east2-a",
	"asia-east2-b",
	"asia-east2-c",

	"asia-northeast1-a",
	"asia-northeast1-b",
	"asia-northeast1-c",
	"asia-northeast1-d",

	"asia-northeast2-a",
	"asia-northeast2-b",
	"asia-northeast2-c",

	"asia-northeast3-a",
	"asia-northeast3-b",
	"asia-northeast3-c",

	"asia-south1-a",
	"asia-south1-b",
	"asia-south1-c",

	"asia-southeast1-a",
	"asia-southeast1-b",

	"asia-southeast2-a",
	"asia-southeast2-b",
	"asia-southeast2-c",

	"australia-southeast1-a",
	"australia-southeast1-b",
	"australia-southeast1-c",

	"europe-north1-a",
	"europe-north1-b",
	"europe-north1-c",

	"europe-west1-a",
	"europe-west1-b",
	"europe-west1-c",
	"europe-west1-d",
	"europe-west1-e",

	"europe-west2-a",
	"europe-west2-b",
	"europe-west2-c",

	"europe-west3-a",
	"europe-west3-b",
	"europe-west3-c",

	"europe-west4-a",
	"europe-west4-b",
	"europe-west4-c",

	"europe-west6-a",
	"europe-west6-b",
	"europe-west6-c",

	"us-central1-a",
	"us-central1-b",
	"us-central1-c",
	"us-central1-d",
	"us-central1-e",
	"us-central1-f",
	"us-central1-g",
	"us-central1-h",

	"us-east1-a",
	"us-east1-b",
	"us-east1-c",
	"us-east1-d",

	"us-east4-a",
	"us-east4-b",
	"us-east4-c",

	"us-west1-a",
	"us-west1-b",
	"us-west1-c",
	"us-west1-d",

	"us-west2-a",
	"us-west2-b",
	"us-west2-c",

	"us-west3-a",
	"us-west3-b",
	"us-west3-c",

	"us-west4-a",
	"us-west4-b",
	"us-west4-c",

	"northamerica-northeast1-a",
	"northamerica-northeast1-b",
	"northamerica-northeast1-c",

	"southamerica-east1-a",
	"southamerica-east1-b",
	"southamerica-east1-c",
}

var doZones = []string{
	"nyc1",
	"nyc3",

	"sfo3",

	"ams3",

	"tor1",

	"sgp1",

	"lon1",

	"fra1",

	"blr1",
}

var hetznerZones = []string{
	// eu-central
	"fsn1",
	"nbg1",
	"hel1",
	// us-east
	"ash",
	// us-west
	"hil",
}

var awsZones = []string{
	"af-south-1a",
	"af-south-1b",
	"af-south-1c",
	"ap-east-1a",
	"ap-east-1b",
	"ap-east-1c",
	"ap-northeast-1a",
	"ap-northeast-1b",
	"ap-northeast-1c",
	"ap-northeast-1d",
	"ap-northeast-2a",
	"ap-northeast-2b",
	"ap-northeast-2c",
	"ap-northeast-2d",
	"ap-northeast-3a",
	"ap-northeast-3b",
	"ap-northeast-3c",
	"ap-south-1a",
	"ap-south-1b",
	"ap-south-1c",
	"ap-south-2a",
	"ap-south-2b",
	"ap-south-2c",
	"ap-southeast-1a",
	"ap-southeast-1b",
	"ap-southeast-1c",
	"ap-southeast-2a",
	"ap-southeast-2b",
	"ap-southeast-2c",
	"ap-southeast-3a",
	"ap-southeast-3b",
	"ap-southeast-3c",
	"ap-southeast-4a",
	"ap-southeast-4b",
	"ap-southeast-4c",
	"ca-central-1a",
	"ca-central-1b",
	"ca-central-1c",
	"ca-west-1a",
	"ca-west-1b",
	"ca-west-1c",
	"eu-central-1a",
	"eu-central-1b",
	"eu-central-1c",
	"eu-central-2a",
	"eu-central-2b",
	"eu-central-2c",
	"eu-north-1a",
	"eu-north-1b",
	"eu-north-1c",
	"eu-south-1a",
	"eu-south-1b",
	"eu-south-1c",
	"eu-south-2a",
	"eu-south-2b",
	"eu-south-2c",
	"eu-west-1a",
	"eu-west-1b",
	"eu-west-1c",
	"eu-west-2a",
	"eu-west-2b",
	"eu-west-2c",
	"eu-west-3a",
	"eu-west-3b",
	"eu-west-3c",
	"il-central-1a",
	"il-central-1b",
	"il-central-1c",
	"me-central-1a",
	"me-central-1b",
	"me-central-1c",
	"me-south-1a",
	"me-south-1b",
	"me-south-1c",
	"sa-east-1a",
	"sa-east-1b",
	"sa-east-1c",
	"us-east-1a",
	"us-east-1b",
	"us-east-1c",
	"us-east-1d",
	"us-east-1e",
	"us-east-1f",
	"us-east-2a",
	"us-east-2b",
	"us-east-2c",
	"us-west-1a",
	"us-west-1b",
	"us-west-1c",
	"us-west-2a",
	"us-west-2b",
	"us-west-2c",
	"us-west-2d",

	"us-gov-east-1a",
	"us-gov-east-1b",
	"us-gov-east-1c",
	"us-gov-west-1a",
	"us-gov-west-1b",
	"us-gov-west-1c",

	"cn-north-1a",
	"cn-north-1b",
	"cn-north-1c",
	"cn-northwest-1a",
	"cn-northwest-1b",
	"cn-northwest-1c",
}

var azureZones = []string{
	"asia",
	"asiapacific",
	"australia",
	"australiacentral",
	"australiacentral2",
	"australiaeast",
	"australiasoutheast",
	"brazil",
	"brazilsouth",
	"brazilsoutheast",
	"canada",
	"canadacentral",
	"canadaeast",
	"centralindia",
	"centralus",
	"centraluseuap",
	"centralusstage",
	"eastasia",
	"eastasiastage",
	"eastus",
	"eastus2",
	"eastus2euap",
	"eastus2stage",
	"eastusstage",
	"europe",
	"francecentral",
	"francesouth",
	"germanynorth",
	"germanywestcentral",
	"global",
	"india",
	"japan",
	"japaneast",
	"japanwest",
	"koreacentral",
	"koreasouth",
	"northcentralus",
	"northcentralusstage",
	"northeurope",
	"norwayeast",
	"norwaywest",
	"southafricanorth",
	"southafricawest",
	"southcentralus",
	"southcentralusstage",
	"southeastasia",
	"southeastasiastage",
	"southindia",
	"switzerlandnorth",
	"switzerlandwest",
	"uaecentral",
	"uaenorth",
	"uk",
	"uksouth",
	"ukwest",
	"unitedstates",
	"westcentralus",
	"westeurope",
	"westindia",
	"westus",
	"westus2",
	"westus2stage",
	"westusstage",
}

func scwZones() []string {
	var scwZones []string
	for _, zone := range scw.AllZones {
		scwZones = append(scwZones, string(zone))
	}
	return scwZones
}

func WellKnownZonesForCloud(matchCloud kops.CloudProviderID, prefix string) []string {
	var found []string
	switch matchCloud {
	case kops.CloudProviderAWS:
		found = awsZones
	case kops.CloudProviderAzure:
		found = azureZones
	case kops.CloudProviderDO:
		found = doZones
	case kops.CloudProviderGCE:
		found = gceZones
	case kops.CloudProviderHetzner:
		found = hetznerZones
	case kops.CloudProviderScaleway:
		found = scwZones()
	default:
		return nil
	}

	sort.Strings(found)
	return found
}
