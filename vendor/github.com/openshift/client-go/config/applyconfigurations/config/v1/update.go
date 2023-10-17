// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// UpdateApplyConfiguration represents an declarative configuration of the Update type for use
// with apply.
type UpdateApplyConfiguration struct {
	Version *string `json:"version,omitempty"`
	Image   *string `json:"image,omitempty"`
	Force   *bool   `json:"force,omitempty"`
}

// UpdateApplyConfiguration constructs an declarative configuration of the Update type for use with
// apply.
func Update() *UpdateApplyConfiguration {
	return &UpdateApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *UpdateApplyConfiguration) WithVersion(value string) *UpdateApplyConfiguration {
	b.Version = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *UpdateApplyConfiguration) WithImage(value string) *UpdateApplyConfiguration {
	b.Image = &value
	return b
}

// WithForce sets the Force field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Force field is set to the value of the last call.
func (b *UpdateApplyConfiguration) WithForce(value bool) *UpdateApplyConfiguration {
	b.Force = &value
	return b
}
