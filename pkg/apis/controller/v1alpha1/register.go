package v1alpha1

import (
	"github.com/beezy-dev/dukat/pkg/apis/controller"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: controller.GroupName, Version: "v1alpha1"}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}
