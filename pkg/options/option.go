/*
 * Copyright The Kmesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package options

import "github.com/spf13/cobra"

type BootstrapConfigs struct {
	DNSAddr    string
	XDSAddress string
	VIP        string
}

func NewBootstrapConfigs() *BootstrapConfigs {
	return &BootstrapConfigs{}
}

func (c *BootstrapConfigs) AttachFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&c.DNSAddr, "dnsAddr", "", "dns listen address: e.g. localhost:8053")
	cmd.PersistentFlags().StringVar(&c.XDSAddress, "xdsAddress", "", "controll plane address")
	cmd.PersistentFlags().StringVar(&c.VIP, "vip", "", "default value for A record, iff ServiceEntry has no Addresses")
}