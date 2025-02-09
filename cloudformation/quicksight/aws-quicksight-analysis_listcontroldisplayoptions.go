// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_ListControlDisplayOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.ListControlDisplayOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-listcontroldisplayoptions.html
type Analysis_ListControlDisplayOptions struct {

	// SearchOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-listcontroldisplayoptions.html#cfn-quicksight-analysis-listcontroldisplayoptions-searchoptions
	SearchOptions *Analysis_ListControlSearchOptions `json:"SearchOptions,omitempty"`

	// SelectAllOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-listcontroldisplayoptions.html#cfn-quicksight-analysis-listcontroldisplayoptions-selectalloptions
	SelectAllOptions *Analysis_ListControlSelectAllOptions `json:"SelectAllOptions,omitempty"`

	// TitleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-listcontroldisplayoptions.html#cfn-quicksight-analysis-listcontroldisplayoptions-titleoptions
	TitleOptions *Analysis_LabelOptions `json:"TitleOptions,omitempty"`

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
func (r *Analysis_ListControlDisplayOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ListControlDisplayOptions"
}
