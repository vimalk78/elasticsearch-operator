package k8shandler

import (
	"testing"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	deployment                      apps.Deployment
	nodeContainer, desiredContainer v1.Container
	node                            *deploymentNode
)

func setUp() {
	nodeContainer = v1.Container{
		Resources: v1.ResourceRequirements{
			Limits: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("2Gi"),
				v1.ResourceCPU:    resource.MustParse("600m"),
			},
			Requests: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("2Gi"),
				v1.ResourceCPU:    resource.MustParse("600m"),
			},
		},
	}

	desiredContainer = v1.Container{
		Resources: v1.ResourceRequirements{
			Limits:   v1.ResourceList{},
			Requests: v1.ResourceList{},
		},
	}
	deployment = apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: apps.DeploymentSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						nodeContainer,
					},
				},
			},
		},
	}
	node = &deploymentNode{
		self: deployment,
	}
}

func TestUpdateResourcesWhenDesiredCPULimitIsZero(t *testing.T) {
	setUp()
	desiredContainer.Resources.Limits = v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("2Gi"),
	}

	desiredContainer.Resources.Requests = v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("2Gi"),
		v1.ResourceCPU:    resource.MustParse("600m"),
	}

	actual, changed := updateResources(node, nodeContainer, desiredContainer)

	if !changed {
		t.Error("Expected updating the resources would recognized as changed, but it was not")
	}
	if !areResourcesSame(actual.Resources, desiredContainer.Resources) {
		t.Errorf("Expected %v but got %v", printResource(desiredContainer.Resources), printResource(actual.Resources))
	}
}
func TestUpdateResourcesWhenDesiredMemoryLimitIsZero(t *testing.T) {
	setUp()
	desiredContainer.Resources.Limits = v1.ResourceList{
		v1.ResourceCPU: resource.MustParse("600m"),
	}

	desiredContainer.Resources.Requests = v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("2Gi"),
		v1.ResourceCPU:    resource.MustParse("600m"),
	}
	actual, changed := updateResources(node, nodeContainer, desiredContainer)

	if !changed {
		t.Error("Expected updating the resources would recognized as changed, but it was not")
	}
	if !areResourcesSame(actual.Resources, desiredContainer.Resources) {
		t.Errorf("Expected %v but got %v", printResource(desiredContainer.Resources), printResource(actual.Resources))
	}
}
func TestUpdateResourcesWhenDesiredCPURequestIsZero(t *testing.T) {
	setUp()
	desiredContainer.Resources.Limits = v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("2Gi"),
		v1.ResourceCPU:    resource.MustParse("600m"),
	}

	desiredContainer.Resources.Requests = v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("2Gi"),
	}

	actual, changed := updateResources(node, nodeContainer, desiredContainer)

	if !changed {
		t.Error("Expected updating the resources would recognized as changed, but it was not")
	}
	if !areResourcesSame(actual.Resources, desiredContainer.Resources) {
		t.Errorf("Expected %v but got %v", printResource(desiredContainer.Resources), printResource(actual.Resources))
	}
}
func TestUpdateResourcesWhenDesiredMemoryRequestIsZero(t *testing.T) {
	setUp()
	desiredContainer.Resources.Limits = v1.ResourceList{
		v1.ResourceCPU:    resource.MustParse("600m"),
		v1.ResourceMemory: resource.MustParse("2Gi"),
	}

	desiredContainer.Resources.Requests = v1.ResourceList{
		v1.ResourceCPU: resource.MustParse("600m"),
	}
	actual, changed := updateResources(node, nodeContainer, desiredContainer)

	if !changed {
		t.Error("Expected updating the resources would recognized as changed, but it was not")
	}
	if !areResourcesSame(actual.Resources, desiredContainer.Resources) {
		t.Errorf("Expected %v but got %v", printResource(desiredContainer.Resources), printResource(actual.Resources))
	}
}
