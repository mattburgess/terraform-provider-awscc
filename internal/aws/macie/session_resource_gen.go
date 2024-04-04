// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_macie_session", sessionResource)
}

// sessionResource returns the Terraform awscc_macie_session resource.
// This Terraform resource corresponds to the CloudFormation AWS::Macie::Session resource.
func sessionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS account ID of customer",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS account ID of customer",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FindingPublishingFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "SIX_HOURS",
		//	  "description": "A enumeration value that specifies how frequently finding updates are published.",
		//	  "enum": [
		//	    "FIFTEEN_MINUTES",
		//	    "ONE_HOUR",
		//	    "SIX_HOURS"
		//	  ],
		//	  "type": "string"
		//	}
		"finding_publishing_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A enumeration value that specifies how frequently finding updates are published.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"FIFTEEN_MINUTES",
					"ONE_HOUR",
					"SIX_HOURS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("SIX_HOURS"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service role used by Macie",
		//	  "type": "string"
		//	}
		"service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service role used by Macie",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "ENABLED",
		//	  "description": "A enumeration value that specifies the status of the Macie Session.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "PAUSED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A enumeration value that specifies the status of the Macie Session.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ENABLED",
					"PAUSED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("ENABLED"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::Session").WithTerraformTypeName("awscc_macie_session")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_account_id":               "AwsAccountId",
		"finding_publishing_frequency": "FindingPublishingFrequency",
		"service_role":                 "ServiceRole",
		"status":                       "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
