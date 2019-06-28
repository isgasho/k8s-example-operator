/*
Copyright 2016 The Kubernetes Authors.

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

package dynamic

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type Interface interface {
	Resource(resource schema.GroupVersionResource) NamespaceableResourceInterface
}

type ResourceInterface interface {
	Create(obj *unstructured.Unstructured, options metav1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error)
	Update(obj *unstructured.Unstructured, options metav1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error)
	UpdateStatus(obj *unstructured.Unstructured, options metav1.UpdateOptions) (*unstructured.Unstructured, error)
	Delete(name string, options *metav1.DeleteOptions, subresources ...string) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error)
	List(opts metav1.ListOptions) (*unstructured.UnstructuredList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error)
}

type NamespaceableResourceInterface interface {
	Namespace(string) ResourceInterface
	ResourceInterface
}

// APIPathResolverFunc knows how to convert a groupVersion to its API path. The Kind field is optional.
// TODO find a better place to move this for existing callers
type APIPathResolverFunc func(kind schema.GroupVersionKind) string

// LegacyAPIPathResolverFunc can resolve paths properly with the legacy API.
// TODO find a better place to move this for existing callers
func LegacyAPIPathResolverFunc(kind schema.GroupVersionKind) string {
	if len(kind.Group) == 0 {
		return "/api"
	}
	return "/apis"
}
