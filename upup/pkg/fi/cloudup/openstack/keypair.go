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

package openstack

import (
	"context"
	"fmt"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/keypairs"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kops/util/pkg/vfs"
)

func (c *openstackCloud) GetKeypair(name string) (*keypairs.KeyPair, error) {
	return getKeypair(c, name)
}

func getKeypair(c OpenstackCloud, name string) (*keypairs.KeyPair, error) {
	var k *keypairs.KeyPair
	done, err := vfs.RetryWithBackoff(readBackoff, func() (bool, error) {
		rs, err := keypairs.Get(context.TODO(), c.ComputeClient(), name, nil).Extract()
		if err != nil {
			if isNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("error listing keypair: %v", err)
		}
		k = rs
		return true, nil
	})
	if err != nil {
		return k, err
	} else if done {
		return k, nil
	} else {
		return k, wait.ErrWaitTimeout
	}
}

func (c *openstackCloud) CreateKeypair(opt keypairs.CreateOptsBuilder) (*keypairs.KeyPair, error) {
	return createKeypair(c, opt)
}

func createKeypair(c OpenstackCloud, opt keypairs.CreateOptsBuilder) (*keypairs.KeyPair, error) {
	var k *keypairs.KeyPair

	done, err := vfs.RetryWithBackoff(writeBackoff, func() (bool, error) {
		v, err := keypairs.Create(context.TODO(), c.ComputeClient(), opt).Extract()
		if err != nil {
			return false, fmt.Errorf("error creating keypair: %v", err)
		}
		k = v
		return true, nil
	})
	if err != nil {
		return k, err
	} else if done {
		return k, nil
	} else {
		return k, wait.ErrWaitTimeout
	}
}

func (c *openstackCloud) DeleteKeyPair(name string) error {
	return deleteKeyPair(c, name)
}

func deleteKeyPair(c OpenstackCloud, name string) error {
	done, err := vfs.RetryWithBackoff(deleteBackoff, func() (bool, error) {
		err := keypairs.Delete(context.TODO(), c.ComputeClient(), name, nil).ExtractErr()
		if err != nil && !isNotFound(err) {
			return false, fmt.Errorf("error deleting keypair: %v", err)
		}
		if isNotFound(err) {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return err
	} else if done {
		return nil
	} else {
		return wait.ErrWaitTimeout
	}
}

func (c *openstackCloud) ListKeypairs() ([]keypairs.KeyPair, error) {
	return listKeypairs(c)
}

func listKeypairs(c OpenstackCloud) ([]keypairs.KeyPair, error) {
	var k []keypairs.KeyPair
	done, err := vfs.RetryWithBackoff(readBackoff, func() (bool, error) {
		allPages, err := keypairs.List(c.ComputeClient(), nil).AllPages(context.TODO())
		if err != nil {
			return false, fmt.Errorf("error listing keypairs: %v", err)
		}

		ks, err := keypairs.ExtractKeyPairs(allPages)
		if err != nil {
			return false, fmt.Errorf("error extracting keypairs from pages: %v", err)
		}
		k = ks
		return true, nil
	})
	if err != nil {
		return k, err
	} else if done {
		return k, nil
	} else {
		return k, wait.ErrWaitTimeout
	}
}
