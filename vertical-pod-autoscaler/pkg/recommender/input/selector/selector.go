package selector

import "k8s.io/apimachinery/pkg/labels"

// NamespacedSelector is used for selecting a set of pods from a namespaced which match a metav1.LabelSelector
type NamespacedSelector interface {
	// Namespace returns the namespace for the pods
	Namespace() string
	// Selector returns the labels.Selector for the pods
	Selector() labels.Selector
	// LabelSelector returns LabelSelector as a `string`
	LabelSelector() string
}
