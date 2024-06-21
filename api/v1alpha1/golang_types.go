package v1alpha1

import gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"

// Golang defines the configuration for Golang HTTP filter.
type Golang struct {
	// BackendRefs defines the configuration of the external processing service
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:XValidation:message="BackendRefs only supports Service and Backend kind.",rule="self.all(f, f.kind == 'Service' || f.kind == 'Backend')"
	// +kubebuilder:validation:XValidation:message="BackendRefs only supports Core and gateway.envoyproxy.io group.",rule="self.all(f, f.group == '' || f.group == 'gateway.envoyproxy.io')"
	BackendRefs []BackendRef `json:"backendRefs"`

	// MessageTimeout is the timeout for a response to be returned from the external processor
	// Default: 200ms
	//
	// +optional
	MessageTimeout *gwapiv1.Duration `json:"messageTimeout,omitempty"`

	// FailOpen defines if requests or responses that cannot be processed due to connectivity to the
	// external processor are terminated or passed-through.
	// Default: false
	//
	// +optional
	FailOpen *bool `json:"failOpen,omitempty"`

	// FIXME: Config
}
