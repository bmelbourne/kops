/*
Copyright 2019 The Kubernetes Authors.

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

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/klog/v2"

	"k8s.io/kops/pkg/resources"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
)

func DeleteSecurityGroup(cloud fi.Cloud, t *resources.Resource) error {
	ctx := context.TODO()
	c := cloud.(awsup.AWSCloud)

	id := t.ID
	// First clear all inter-dependent rules
	// TODO: Move to a "pre-execute" phase?
	{
		request := &ec2.DescribeSecurityGroupsInput{
			GroupIds: []string{id},
		}
		response, err := c.EC2().DescribeSecurityGroups(ctx, request)
		if err != nil {
			if awsup.AWSErrorCode(err) == "InvalidGroup.NotFound" {
				klog.V(2).Infof("Got InvalidGroup.NotFound error describing SecurityGroup %q; will treat as already-deleted", id)
				return nil
			}
			return fmt.Errorf("error describing SecurityGroup %q: %v", id, err)
		}

		if len(response.SecurityGroups) == 0 {
			return nil
		}
		if len(response.SecurityGroups) != 1 {
			return fmt.Errorf("found multiple SecurityGroups with ID %q", id)
		}

		ruleReqest := &ec2.DescribeSecurityGroupRulesInput{
			Filters: []ec2types.Filter{
				{Name: aws.String("group-id"), Values: []string{id}},
			},
		}
		ruleResp, err := c.EC2().DescribeSecurityGroupRules(ctx, ruleReqest)
		if err != nil {
			if awsup.AWSErrorCode(err) == "InvalidGroup.NotFound" {
				klog.V(2).Infof("Got InvalidGroup.NotFound error describing rules for SecurityGroup %q; will treat as already-deleted", id)
				return nil
			}
			return fmt.Errorf("error describing SecurityGroup rules %q: %v", id, err)
		}

		ingressRuleIDs := make([]string, 0)
		for _, rule := range ruleResp.SecurityGroupRules {
			if !aws.ToBool(rule.IsEgress) {
				ingressRuleIDs = append(ingressRuleIDs, aws.ToString(rule.SecurityGroupRuleId))
			}
		}

		if len(ingressRuleIDs) != 0 {
			revoke := &ec2.RevokeSecurityGroupIngressInput{
				GroupId:              aws.String(id),
				SecurityGroupRuleIds: ingressRuleIDs,
			}
			_, err = c.EC2().RevokeSecurityGroupIngress(ctx, revoke)
			if err != nil {
				return fmt.Errorf("cannot revoke ingress for ID %q with rule IDs %v: %v", id, ingressRuleIDs, err)
			}
		}
	}

	{
		klog.V(2).Infof("Deleting EC2 SecurityGroup %q", id)
		request := &ec2.DeleteSecurityGroupInput{
			GroupId: &id,
		}
		_, err := c.EC2().DeleteSecurityGroup(ctx, request)
		if err != nil {
			if IsDependencyViolation(err) {
				return err
			}
			return fmt.Errorf("error deleting SecurityGroup %q: %v", id, err)
		}
	}
	return nil
}

func DumpSecurityGroup(op *resources.DumpOperation, r *resources.Resource) error {
	data := make(map[string]interface{})
	data["id"] = r.ID
	data["type"] = ec2types.ResourceTypeSecurityGroup
	data["raw"] = r.Obj
	op.Dump.Resources = append(op.Dump.Resources, data)
	return nil
}

func ListSecurityGroups(cloud fi.Cloud, vpcID, clusterName string) ([]*resources.Resource, error) {
	groups, err := DescribeSecurityGroups(cloud, clusterName)
	if err != nil {
		return nil, err
	}

	var resourceTrackers []*resources.Resource

	for id, sg := range groups {
		resourceTracker := &resources.Resource{
			Name:    FindName(sg.Tags),
			ID:      id,
			Type:    string(ec2types.ResourceTypeSecurityGroup),
			Deleter: DeleteSecurityGroup,
			Dumper:  DumpSecurityGroup,
			Obj:     sg,
			Shared:  !HasOwnedTag(string(ec2types.ResourceTypeSecurityGroup)+":"+id, sg.Tags, clusterName),
		}

		var blocks []string
		blocks = append(blocks, "vpc:"+aws.ToString(sg.VpcId))

		resourceTracker.Blocks = blocks

		resourceTrackers = append(resourceTrackers, resourceTracker)
	}

	return resourceTrackers, nil
}

func DescribeSecurityGroups(cloud fi.Cloud, clusterName string) (map[string]ec2types.SecurityGroup, error) {
	ctx := context.TODO()
	c := cloud.(awsup.AWSCloud)

	groups := make(map[string]ec2types.SecurityGroup)
	klog.V(2).Infof("Listing EC2 SecurityGroups")
	for _, filters := range buildEC2FiltersForCluster(clusterName) {
		request := &ec2.DescribeSecurityGroupsInput{
			Filters: filters,
		}
		response, err := c.EC2().DescribeSecurityGroups(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("error listing SecurityGroups: %v", err)
		}

		for _, group := range response.SecurityGroups {
			groups[aws.ToString(group.GroupId)] = group
		}
	}

	return groups, nil
}
