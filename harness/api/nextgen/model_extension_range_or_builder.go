/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ExtensionRangeOrBuilder struct {
	Options *ExtensionRangeOptions `json:"options,omitempty"`
	Start int32 `json:"start,omitempty"`
	End int32 `json:"end,omitempty"`
	OptionsOrBuilder *ExtensionRangeOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	DefaultInstanceForType *Message `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
}
