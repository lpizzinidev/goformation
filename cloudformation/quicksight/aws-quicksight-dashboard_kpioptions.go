// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_KPIOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.KPIOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html
type Dashboard_KPIOptions struct {

	// Comparison AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-comparison
	Comparison *Dashboard_ComparisonConfiguration `json:"Comparison,omitempty"`

	// PrimaryValueDisplayType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-primaryvaluedisplaytype
	PrimaryValueDisplayType *string `json:"PrimaryValueDisplayType,omitempty"`

	// PrimaryValueFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-primaryvaluefontconfiguration
	PrimaryValueFontConfiguration *Dashboard_FontConfiguration `json:"PrimaryValueFontConfiguration,omitempty"`

	// ProgressBar AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-progressbar
	ProgressBar *Dashboard_ProgressBarOptions `json:"ProgressBar,omitempty"`

	// SecondaryValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-secondaryvalue
	SecondaryValue *Dashboard_SecondaryValueOptions `json:"SecondaryValue,omitempty"`

	// SecondaryValueFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-secondaryvaluefontconfiguration
	SecondaryValueFontConfiguration *Dashboard_FontConfiguration `json:"SecondaryValueFontConfiguration,omitempty"`

	// TrendArrows AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpioptions.html#cfn-quicksight-dashboard-kpioptions-trendarrows
	TrendArrows *Dashboard_TrendArrowOptions `json:"TrendArrows,omitempty"`

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
func (r *Dashboard_KPIOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.KPIOptions"
}
