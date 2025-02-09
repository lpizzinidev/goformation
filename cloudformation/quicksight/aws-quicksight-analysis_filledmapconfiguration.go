// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_FilledMapConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.FilledMapConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html
type Analysis_FilledMapConfiguration struct {

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-fieldwells
	FieldWells *Analysis_FilledMapFieldWells `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-legend
	Legend *Analysis_LegendOptions `json:"Legend,omitempty"`

	// MapStyleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-mapstyleoptions
	MapStyleOptions *Analysis_GeospatialMapStyleOptions `json:"MapStyleOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-sortconfiguration
	SortConfiguration *Analysis_FilledMapSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions `json:"Tooltip,omitempty"`

	// WindowOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-windowoptions
	WindowOptions *Analysis_GeospatialWindowOptions `json:"WindowOptions,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Analysis_FilledMapConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.FilledMapConfiguration"
}
