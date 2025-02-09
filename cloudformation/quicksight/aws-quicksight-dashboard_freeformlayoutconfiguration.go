// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_FreeFormLayoutConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FreeFormLayoutConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html
type Dashboard_FreeFormLayoutConfiguration struct {

	// CanvasSizeOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html#cfn-quicksight-dashboard-freeformlayoutconfiguration-canvassizeoptions
	CanvasSizeOptions *Dashboard_FreeFormLayoutCanvasSizeOptions `json:"CanvasSizeOptions,omitempty"`

	// Elements AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html#cfn-quicksight-dashboard-freeformlayoutconfiguration-elements
	Elements []Dashboard_FreeFormLayoutElement `json:"Elements"`

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
func (r *Dashboard_FreeFormLayoutConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FreeFormLayoutConfiguration"
}
