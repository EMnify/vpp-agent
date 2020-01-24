//  Copyright (c) 2020 EMnify.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2001

import "go.ligato.io/vpp-agent/v2/plugins/vpp/ifplugin/vppcalls"

// IsFeatureSupported indicates whether or not a particular API feature is supported.
func (h *InterfaceVppHandler) IsFeatureSupported(f vppcalls.Feature) bool {
	switch f {
	case vppcalls.FeatureGtpuUpdateTunnelDst:
		return true
	}
	return false
}
