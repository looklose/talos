// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type V1Alpha1 -pointer-receiver -header-file ../../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package extensionservicesconfig

// DeepCopy generates a deep copy of *V1Alpha1.
func (o *V1Alpha1) DeepCopy() *V1Alpha1 {
	var cp V1Alpha1 = *o
	if o.Config != nil {
		cp.Config = make([]ExtensionServiceConfig, len(o.Config))
		copy(cp.Config, o.Config)
		for i2 := range o.Config {
			if o.Config[i2].ExtensionServiceConfigFiles != nil {
				cp.Config[i2].ExtensionServiceConfigFiles = make([]ExtensionServiceConfigFile, len(o.Config[i2].ExtensionServiceConfigFiles))
				copy(cp.Config[i2].ExtensionServiceConfigFiles, o.Config[i2].ExtensionServiceConfigFiles)
			}
		}
	}
	return &cp
}
