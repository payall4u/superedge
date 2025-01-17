/*
Copyright 2021 The SuperEdge Authors.

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

// +k8s:deepcopy-gen=package
// +groupName=superedge.io
package v1alpha2

import (
	"github.com/superedge/superedge/pkg/site-manager/constant"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "site.superedge.io", Version: "v1alpha2"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, registerDefaults)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&NodeGroup{},
		&NodeGroupList{},
		&NodeUnit{},
		&NodeUnitList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func registerDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&NodeUnit{}, func(obj interface{}) { SetObjectDefaults_NodeUnit(obj.(*NodeUnit)) })
	return nil
}

func SetObjectDefaults_NodeUnit(in *NodeUnit) {
	SetDefaults_NodeUnitSpec(in)
	SetDefaults_NodeUnitStatus(&in.Status)
}

func SetDefaults_NodeUnitSpec(in *NodeUnit) {
	if _, ok := in.Spec.SetNode.Labels[in.Name]; !ok {
		in.Spec.SetNode.Labels[in.Name] = constant.NodeUnitSuperedge
	}
}

func SetDefaults_NodeUnitStatus(in *NodeUnitStatus) {}
