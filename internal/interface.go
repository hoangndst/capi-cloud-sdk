package internal

import "k8s.io/apimachinery/pkg/runtime/schema"

type CloudProvider interface {
	GetClusterGroupVersionKind() schema.GroupVersionKind
	GetClusterGroupVersionResource() schema.GroupVersionResource
	GetMachineTemplateGroupVersionKind() schema.GroupVersionKind
	GetMachineTemplateGroupVersionResource() schema.GroupVersionResource
	GetMachineGroupVersionKind() schema.GroupVersionKind
	GetMachineGroupVersionResource() schema.GroupVersionResource
}
