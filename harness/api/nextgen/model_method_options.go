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

type MethodOptions struct {
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
	SerializedSize int32 `json:"serializedSize,omitempty"`
	ParserForType *ParserMethodOptions `json:"parserForType,omitempty"`
	DefaultInstanceForType *MethodOptions `json:"defaultInstanceForType,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	UninterpretedOptionList []UninterpretedOption `json:"uninterpretedOptionList,omitempty"`
	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount,omitempty"`
	UninterpretedOptionOrBuilderList []UninterpretedOptionOrBuilder `json:"uninterpretedOptionOrBuilderList,omitempty"`
	IdempotencyLevel string `json:"idempotencyLevel,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	AllFieldsRaw map[string]interface{} `json:"allFieldsRaw,omitempty"`
	MemoizedSerializedSize int32 `json:"memoizedSerializedSize,omitempty"`
}
