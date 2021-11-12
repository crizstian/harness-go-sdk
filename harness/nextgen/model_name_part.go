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

type NamePart struct {
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
	ParserForType             *ParserNamePart        `json:"parserForType,omitempty"`
	DefaultInstanceForType    *NamePart              `json:"defaultInstanceForType,omitempty"`
	SerializedSize            int32                  `json:"serializedSize,omitempty"`
	IsExtension               bool                   `json:"isExtension,omitempty"`
	NamePart                  string                 `json:"namePart,omitempty"`
	NamePartBytes             *ByteString            `json:"namePartBytes,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                  `json:"memoizedSerializedSize,omitempty"`
}
