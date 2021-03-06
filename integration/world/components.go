/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package world

import (
	"os"

	"github.com/hyperledger/fabric/integration/runner"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

type Components struct {
	Paths map[string]string
}

func (c *Components) Build() {
	if c.Paths == nil {
		c.Paths = map[string]string{}
	}
	cryptogen, err := gexec.Build("github.com/hyperledger/fabric/common/tools/cryptogen")
	Expect(err).NotTo(HaveOccurred())
	c.Paths["cryptogen"] = cryptogen

	configtxgen, err := gexec.Build("github.com/hyperledger/fabric/common/tools/configtxgen")
	Expect(err).NotTo(HaveOccurred())
	c.Paths["configtxgen"] = configtxgen

	orderer, err := gexec.Build("github.com/hyperledger/fabric/orderer")
	Expect(err).NotTo(HaveOccurred())
	c.Paths["orderer"] = orderer
}

func (c *Components) Cleanup() {
	for _, path := range c.Paths {
		err := os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	}
}

func (c *Components) Cryptogen() *runner.Cryptogen {
	return &runner.Cryptogen{
		Path: c.Paths["cryptogen"],
	}
}

func (c *Components) Configtxgen() *runner.Configtxgen {
	return &runner.Configtxgen{
		Path: c.Paths["configtxgen"],
	}
}

func (c *Components) Orderer() *runner.Orderer {
	return &runner.Orderer{
		Path: c.Paths["orderer"],
	}
}
