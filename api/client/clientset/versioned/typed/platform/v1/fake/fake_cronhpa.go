/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeCronHPAs implements CronHPAInterface
type FakeCronHPAs struct {
	Fake *FakePlatformV1
}

var cronhpasResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "cronhpas"}

var cronhpasKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "CronHPA"}

// Get takes name of the cronHPA, and returns the corresponding cronHPA object, and an error if there is any.
func (c *FakeCronHPAs) Get(name string, options v1.GetOptions) (result *platformv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cronhpasResource, name), &platformv1.CronHPA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CronHPA), err
}

// List takes label and field selectors, and returns the list of CronHPAs that match those selectors.
func (c *FakeCronHPAs) List(opts v1.ListOptions) (result *platformv1.CronHPAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cronhpasResource, cronhpasKind, opts), &platformv1.CronHPAList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.CronHPAList{ListMeta: obj.(*platformv1.CronHPAList).ListMeta}
	for _, item := range obj.(*platformv1.CronHPAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronHPAs.
func (c *FakeCronHPAs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cronhpasResource, opts))
}

// Create takes the representation of a cronHPA and creates it.  Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *FakeCronHPAs) Create(cronHPA *platformv1.CronHPA) (result *platformv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cronhpasResource, cronHPA), &platformv1.CronHPA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CronHPA), err
}

// Update takes the representation of a cronHPA and updates it. Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *FakeCronHPAs) Update(cronHPA *platformv1.CronHPA) (result *platformv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cronhpasResource, cronHPA), &platformv1.CronHPA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CronHPA), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCronHPAs) UpdateStatus(cronHPA *platformv1.CronHPA) (*platformv1.CronHPA, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cronhpasResource, "status", cronHPA), &platformv1.CronHPA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CronHPA), err
}

// Delete takes name of the cronHPA and deletes it. Returns an error if one occurs.
func (c *FakeCronHPAs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cronhpasResource, name), &platformv1.CronHPA{})
	return err
}

// Patch applies the patch and returns the patched cronHPA.
func (c *FakeCronHPAs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cronhpasResource, name, pt, data, subresources...), &platformv1.CronHPA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CronHPA), err
}
