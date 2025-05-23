// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type NamespacesTopicsSubscriptionsRule_STATUS struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Properties of Rule resource
	Properties *Ruleproperties_STATUS `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Description of Rule Resource.
type Ruleproperties_STATUS struct {
	// Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
	// filter expression.
	Action *Action_STATUS `json:"action,omitempty"`

	// CorrelationFilter: Properties of correlationFilter
	CorrelationFilter *CorrelationFilter_STATUS `json:"correlationFilter,omitempty"`

	// FilterType: Filter type that is evaluated against a BrokeredMessage.
	FilterType *FilterType_STATUS `json:"filterType,omitempty"`

	// SqlFilter: Properties of sqlFilter
	SqlFilter *SqlFilter_STATUS `json:"sqlFilter,omitempty"`
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
// expression.
type Action_STATUS struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}

// Represents the correlation filter expression.
type CorrelationFilter_STATUS struct {
	// ContentType: Content type of the message.
	ContentType *string `json:"contentType,omitempty"`

	// CorrelationId: Identifier of the correlation.
	CorrelationId *string `json:"correlationId,omitempty"`

	// Label: Application specific label.
	Label *string `json:"label,omitempty"`

	// MessageId: Identifier of the message.
	MessageId *string `json:"messageId,omitempty"`

	// Properties: dictionary object for custom filters
	Properties map[string]string `json:"properties,omitempty"`

	// ReplyTo: Address of the queue to reply to.
	ReplyTo *string `json:"replyTo,omitempty"`

	// ReplyToSessionId: Session identifier to reply to.
	ReplyToSessionId *string `json:"replyToSessionId,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SessionId: Session identifier.
	SessionId *string `json:"sessionId,omitempty"`

	// To: Address to send to.
	To *string `json:"to,omitempty"`
}

// Rule filter types
type FilterType_STATUS string

const (
	FilterType_STATUS_CorrelationFilter = FilterType_STATUS("CorrelationFilter")
	FilterType_STATUS_SqlFilter         = FilterType_STATUS("SqlFilter")
)

// Mapping from string to FilterType_STATUS
var filterType_STATUS_Values = map[string]FilterType_STATUS{
	"correlationfilter": FilterType_STATUS_CorrelationFilter,
	"sqlfilter":         FilterType_STATUS_SqlFilter,
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter_STATUS struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: The SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}
