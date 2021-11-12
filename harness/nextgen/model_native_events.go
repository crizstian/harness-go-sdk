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

type NativeEvents struct {
	UnknownFields             *UnknownFieldSet                       `json:"unknownFields,omitempty"`
	Initialized               bool                                   `json:"initialized,omitempty"`
	GithubOrBuilder           *GithubWebhookEventsOrBuilder          `json:"githubOrBuilder,omitempty"`
	Gitlab                    *GitlabWebhookEvents                   `json:"gitlab,omitempty"`
	GitlabOrBuilder           *GitlabWebhookEventsOrBuilder          `json:"gitlabOrBuilder,omitempty"`
	BitbucketCloud            *BitbucketCloudWebhookEvents           `json:"bitbucketCloud,omitempty"`
	BitbucketCloudOrBuilder   *BitbucketCloudWebhookEventsOrBuilder  `json:"bitbucketCloudOrBuilder,omitempty"`
	BitbucketServer           *BitbucketServerWebhookEvents          `json:"bitbucketServer,omitempty"`
	BitbucketServerOrBuilder  *BitbucketServerWebhookEventsOrBuilder `json:"bitbucketServerOrBuilder,omitempty"`
	ParserForType             *ParserNativeEvents                    `json:"parserForType,omitempty"`
	DefaultInstanceForType    *NativeEvents                          `json:"defaultInstanceForType,omitempty"`
	SerializedSize            int32                                  `json:"serializedSize,omitempty"`
	Github                    *GithubWebhookEvents                   `json:"github,omitempty"`
	NativeEventsCase          string                                 `json:"nativeEventsCase,omitempty"`
	AllFields                 map[string]interface{}                 `json:"allFields,omitempty"`
	InitializationErrorString string                                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                            `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                                  `json:"memoizedSerializedSize,omitempty"`
}
