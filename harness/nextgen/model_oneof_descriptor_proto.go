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

type OneofDescriptorProto struct {
	UnknownFields             *UnknownFieldSet            `json:"unknownFields,omitempty"`
	Initialized               bool                        `json:"initialized,omitempty"`
	Options                   *OneofOptions               `json:"options,omitempty"`
	NameBytes                 *ByteString                 `json:"nameBytes,omitempty"`
	ParserForType             *ParserOneofDescriptorProto `json:"parserForType,omitempty"`
	DefaultInstanceForType    *OneofDescriptorProto       `json:"defaultInstanceForType,omitempty"`
	SerializedSize            int32                       `json:"serializedSize,omitempty"`
	OptionsOrBuilder          *OneofOptionsOrBuilder      `json:"optionsOrBuilder,omitempty"`
	Name                      string                      `json:"name,omitempty"`
	AllFields                 map[string]interface{}      `json:"allFields,omitempty"`
	InitializationErrorString string                      `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                 `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                       `json:"memoizedSerializedSize,omitempty"`
}
