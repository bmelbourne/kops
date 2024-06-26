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

package subnet

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Overlap checks if two subnets overlap
func Overlap(l, r *net.IPNet) bool {
	if l == nil || r == nil {
		return false
	}
	return l.Contains(r.IP) || r.Contains(l.IP)
}

// BelongsTo checks if child is a subnet of parent
func BelongsTo(parent *net.IPNet, child *net.IPNet) bool {
	parentOnes, parentBits := parent.Mask.Size()
	childOnes, childBits := child.Mask.Size()
	if childBits != parentBits {
		return false
	}
	if parentOnes > childOnes {
		return false
	}
	childMasked := child.IP.Mask(parent.Mask)
	parentMasked := parent.IP.Mask(parent.Mask)
	return childMasked.Equal(parentMasked)
}

// SplitInto1 splits the parent IPNet into 1 subnet
func SplitInto1(parent *net.IPNet) ([]*net.IPNet, error) {
	return SplitInto(0, parent)
}

// SplitInto2 splits the parent IPNet into 2 subnets
func SplitInto2(parent *net.IPNet) ([]*net.IPNet, error) {
	return SplitInto(1, parent)
}

// SplitInto4 splits the parent IPNet into 4 subnets
func SplitInto4(parent *net.IPNet) ([]*net.IPNet, error) {
	return SplitInto(2, parent)
}

// SplitInto8 splits the parent IPNet into 8 subnets
func SplitInto8(parent *net.IPNet) ([]*net.IPNet, error) {
	return SplitInto(3, parent)
}

// SplitInto splits the parent IPNet into subnets with the specified number of additional bits in the prefix.
func SplitInto(additionalBits uint, parent *net.IPNet) ([]*net.IPNet, error) {
	networkLength, _ := parent.Mask.Size()
	networkLength += int(additionalBits)

	var subnets []*net.IPNet
	for i := 0; i < 1<<additionalBits; i++ {
		ip4 := parent.IP.To4()
		if ip4 != nil {
			n := binary.BigEndian.Uint32(ip4)
			n += uint32(i) << uint(32-networkLength)
			subnetIP := make(net.IP, len(ip4))
			binary.BigEndian.PutUint32(subnetIP, n)

			subnets = append(subnets, &net.IPNet{
				IP:   subnetIP,
				Mask: net.CIDRMask(networkLength, 32),
			})
		} else {
			return nil, fmt.Errorf("unexpected IP address type: %s", parent)
		}
	}

	return subnets, nil
}
