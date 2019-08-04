package utils

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type NamespacedSelector struct {
	namespace string
	selector  labels.Selector
}

func NewNamespacedSelect(namespace string, selector labels.Selector) NamespacedSelector {
	return NamespacedSelector{namespace: namespace, selector: selector}
}

func NewAllPodsSelector() NamespacedSelector {
	return NamespacedSelector{namespace: v1.NamespaceAll}
}

func (selector NamespacedSelector) Namespace() string {
	return selector.namespace
}

func (selector NamespacedSelector) Selector() labels.Selector {
	return selector.selector
}

func (selector NamespacedSelector) LabelSelector() string {
	if selector.selector != nil {
		return selector.selector.String()
	}
	return ""
}
