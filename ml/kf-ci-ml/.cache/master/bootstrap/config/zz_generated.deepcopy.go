// +build !ignore_autogenerated

// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfig) DeepCopyInto(out *ComponentConfig) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ComponentParams != nil {
		in, out := &in.ComponentParams, &out.ComponentParams
		*out = make(Parameters, len(*in))
		for key, val := range *in {
			var outVal []NameValue
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]NameValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfig.
func (in *ComponentConfig) DeepCopy() *ComponentConfig {
	if in == nil {
		return nil
	}
	out := new(ComponentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameValue) DeepCopyInto(out *NameValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameValue.
func (in *NameValue) DeepCopy() *NameValue {
	if in == nil {
		return nil
	}
	out := new(NameValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Parameters) DeepCopyInto(out *Parameters) {
	{
		in := &in
		*out = make(Parameters, len(*in))
		for key, val := range *in {
			var outVal []NameValue
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]NameValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameters.
func (in Parameters) DeepCopy() Parameters {
	if in == nil {
		return nil
	}
	out := new(Parameters)
	in.DeepCopyInto(out)
	return *out
}
