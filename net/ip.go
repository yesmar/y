// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package net

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"

	ymath "github.com/yesmar/y/math"
)

// Nipsv4 computes the number of IP addresses in the given IPv4 CIDR.
func Nipsv4(cidr *net.IPNet) (uint64, error) {
	if cidr == nil {
		return 0, errors.New("nil net.IPNet")
	}
	if cidr.IP.To4() == nil {
		return 0, fmt.Errorf("%s is not an IPv4 address", cidr.IP)
	}
	netbits, bits := cidr.Mask.Size()
	hostbits := bits - netbits
	c := math.Pow(2, float64(hostbits))
	return uint64(c), nil
}

// Nipsv6 computes the number of IP addresses in the given IPv6 CIDR.
func Nipsv6(cidr *net.IPNet) (*big.Int, error) {
	if cidr == nil {
		return nil, errors.New("nil net.IPNet")
	}
	if cidr.IP.To4() != nil {
		return nil, fmt.Errorf("%s is not an IPv6 address", cidr.IP)
	}
	netbits, bits := cidr.Mask.Size()
	hostbits := bits - netbits
	c := math.Pow(2, float64(hostbits))
	n, err := ymath.BigFloatToBigInt(big.NewFloat(c))
	if err != nil {
		return nil, err
	}
	return n, nil
}
