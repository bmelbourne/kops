/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi/cloudup/gce"
)

func gceValidateCluster(c *kops.Cluster) field.ErrorList {
	allErrs := field.ErrorList{}

	fieldSpec := field.NewPath("spec")

	for i, subnet := range c.Spec.Networking.Subnets {
		f := fieldSpec.Child("networking", "subnets").Index(i)
		if subnet.Zone != "" {
			allErrs = append(allErrs, field.Invalid(f.Child("zone"), subnet.Zone, "zones should not be specified for GCE subnets, as GCE subnets are regional"))
		}
	}

	return allErrs
}

func gceValidateInstanceGroup(ig *kops.InstanceGroup, cloud gce.GCECloud) field.ErrorList {
	allErrs := field.ErrorList{}

	if ig.Spec.GCPProvisioningModel != nil {
		fieldSpec := field.NewPath("spec")
		allErrs = append(allErrs, IsValidValue(fieldSpec.Child("gcpProvisioningModel"), ig.Spec.GCPProvisioningModel, []string{"STANDARD", "SPOT"})...)
	}
	return allErrs
}
