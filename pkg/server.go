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

package pkg

import (
	"kmesh.net/kmesh-coredns-plugin/pkg/ads"
	"kmesh.net/kmesh-coredns-plugin/pkg/dns"
	"kmesh.net/kmesh-coredns-plugin/pkg/options"
)

type Manager struct {
	server *dns.KmeshDNSServer
	ads    *ads.AdsController
}

func NewDNSManager(cfg *options.BootstrapConfigs) (*Manager, error) {
	m := &Manager{}

	s, err := dns.NewDNSServer(cfg.DNSAddr)

	if err != nil {
		return nil, err
	}
	m.server = s

	adsController, err := ads.NewAdsController(cfg, s)

	if err != nil {
		return nil, err
	}

	m.ads = adsController

	return m, nil
}

func (m *Manager) Start(stop <-chan struct{}) error {
	m.server.Start()
	if err := m.ads.Start(); err != nil {
		return err
	}

	m.waitForShutDown(stop)
	return nil
}

func (m *Manager) waitForShutDown(stop <-chan struct{}) {
	go func() {
		<-stop
		m.server.Stop()
		m.ads.Stop()
	}()
}