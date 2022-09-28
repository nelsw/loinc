package model

type Loinc struct {

	// ID is the unique LOINC Code is a string in the format of nnnnnnnn-n.
	ID string `json:"id" csv:"LOINC_NUM"`

	// Component is the first major axis-component or analyte.
	Component string `json:"component" csv:"COMPONENT"`

	// Property
	Property string `json:"property" csv:"PROPERTY"`

	// TimeAspect
	TimeAspect string `json:"time_aspect" csv:"TIME_ASPCT"`

	// System
	System string `json:"system" csv:"SYSTEM"`

	// ScaleType
	ScaleType string `json:"scale_type" csv:"SCALE_TYP"`

	// MethodType
	MethodType string `json:"method_type,omitempty" csv:"METHOD_TYP"`

	// Class
	Class string `json:"class" csv:"CLASS"`

	// SurveyQuestText
	SurveyQuestText string `json:"survey_quest_text,omitempty" csv:"SURVEY_QUEST_TEXT"`

	// ExampleAnswers
	ExampleAnswers string `json:"example_answers,omitempty" csv:"EXMPL_ANSWERS"`

	// SurveyQuestSource
	SurveyQuestSource string `json:"survey_quest_source,omitempty" csv:"SURVEY_QUEST_SRC"`

	// Formula
	Formula string `json:"formula,omitempty" csv:"FORMULA"`

	// UnitsRequired
	UnitsRequired string `json:"units_required,omitempty" csv:"UNITSREQUIRED"`

	// ClassType
	ClassType string `json:"class_type" csv:"CLASSTYPE"`

	// RelatedNames2
	RelatedNames2 string `json:"related_names_2,omitempty" csv:"RELATEDNAMES2"`

	// ConsumerName
	ConsumerName string `json:"consumer_name,omitempty" csv:"CONSUMER_NAME"`

	// ShortName
	ShortName string `json:"short_name" csv:"SHORTNAME"`

	// OrderObs
	OrderObs string `json:"order_obs,omitempty" csv:"ORDER_OBS"`

	// ExternalCopyRightNotice
	ExternalCopyRightNotice string `json:"external_copy_right_notice,omitempty" csv:"EXTERNAL_COPYRIGHT_NOTICE"`

	// Hl7FieldSubfieldId
	Hl7FieldSubfieldId string `json:"hl_7_field_subfield_id,omitempty" csv:"HL7_FIELD_SUBFIELD_ID"`

	// ExampleUnits
	ExampleUnits string `json:"example_units,omitempty" csv:"EXAMPLE_UNITS"`

	// LongCommonName
	LongCommonName string `json:"long_common_name" csv:"LONG_COMMON_NAME"`

	// ExampleUcumUnits
	ExampleUcumUnits string `json:"example_ucum_units,omitempty" csv:"EXAMPLE_UCUM_UNITS"`

	// StatusText
	StatusText string `json:"status_text,omitempty" csv:"STATUS_TEXT"`

	// ChangeReasonPublic
	ChangeReasonPublic string `json:"change_reason_public,omitempty" csv:"CHANGE_REASON_PUBLIC"`

	// Hl7AttachmentStructure
	Hl7AttachmentStructure string `json:"hl_7_attachment_structure,omitempty" csv:"HL7_ATTACHMENT_STRUCTURE"`

	// CommonTestRank
	CommonTestRank string `json:"common_test_rank,omitempty" csv:"COMMON_TEST_RANK"`

	// CommonSiTestRank
	CommonSiTestRank string `json:"common_si_test_rank,omitempty" csv:"COMMON_SI_TEST_RANK"`

	// CommonOrderRank
	CommonOrderRank string `json:"common_order_rank,omitempty" csv:"COMMON_ORDER_RANK"`

	// ExternalCopyrightLink
	ExternalCopyrightLink string `json:"external_copyright_link,omitempty" csv:"EXTERNAL_COPYRIGHT_LINK"`

	// PanelType
	PanelType string `json:"panel_type,omitempty"`

	// AskAtOrderEntry
	AskAtOrderEntry string `json:"ask_at_order_entry,omitempty"`

	// AssociatedObservations
	AssociatedObservations string `json:"associated_observations,omitempty"`

	// VersionFirstReleased
	VersionFirstReleased string `json:"version_first_released"`

	// ValidHL7AttachmentRequest
	ValidHL7AttachmentRequest string `json:"valid_hl_7_attachment_request,omitempty"`

	// DisplayName
	DisplayName string `json:"display_name,omitempty"`
}
