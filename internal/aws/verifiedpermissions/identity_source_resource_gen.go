// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_verifiedpermissions_identity_source", identitySourceResource)
}

// identitySourceResource returns the Terraform awscc_verifiedpermissions_identity_source resource.
// This Terraform resource corresponds to the CloudFormation AWS::VerifiedPermissions::IdentitySource resource.
func identitySourceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CognitoUserPoolConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ClientIds": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 255,
		//	            "minLength": 1,
		//	            "pattern": "^.*$",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1000,
		//	          "minItems": 0,
		//	          "type": "array"
		//	        },
		//	        "UserPoolArn": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "pattern": "^arn:[a-zA-Z0-9-]+:cognito-idp:(([a-zA-Z0-9-]+:\\d{12}:userpool/[\\w-]+_[0-9a-zA-Z]+))$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "UserPoolArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "CognitoUserPoolConfiguration"
		//	  ],
		//	  "type": "object"
		//	}
		"configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CognitoUserPoolConfiguration
				"cognito_user_pool_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ClientIds
						"client_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 1000),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 255),
									stringvalidator.RegexMatches(regexp.MustCompile("^.*$"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: UserPoolArn
						"user_pool_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
								stringvalidator.RegexMatches(regexp.MustCompile("^arn:[a-zA-Z0-9-]+:cognito-idp:(([a-zA-Z0-9-]+:\\d{12}:userpool/[\\w-]+_[0-9a-zA-Z]+))$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Details
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ClientIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "^.*$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 1000,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "DiscoveryUrl": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "^https://.*$",
		//	      "type": "string"
		//	    },
		//	    "OpenIdIssuer": {
		//	      "enum": [
		//	        "COGNITO"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "UserPoolArn": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "pattern": "^arn:[a-zA-Z0-9-]+:cognito-idp:(([a-zA-Z0-9-]+:\\d{12}:userpool/[\\w-]+_[0-9a-zA-Z]+))$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ClientIds
				"client_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DiscoveryUrl
				"discovery_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OpenIdIssuer
				"open_id_issuer": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UserPoolArn
				"user_pool_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdentitySourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-]*$",
		//	  "type": "string"
		//	}
		"identity_source_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyStoreId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-]*$",
		//	  "type": "string"
		//	}
		"policy_store_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrincipalEntityType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^.*$",
		//	  "type": "string"
		//	}
		"principal_entity_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^.*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
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
		Description: "Definition of AWS::VerifiedPermissions::IdentitySource Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::VerifiedPermissions::IdentitySource").WithTerraformTypeName("awscc_verifiedpermissions_identity_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"client_ids":                      "ClientIds",
		"cognito_user_pool_configuration": "CognitoUserPoolConfiguration",
		"configuration":                   "Configuration",
		"details":                         "Details",
		"discovery_url":                   "DiscoveryUrl",
		"identity_source_id":              "IdentitySourceId",
		"open_id_issuer":                  "OpenIdIssuer",
		"policy_store_id":                 "PolicyStoreId",
		"principal_entity_type":           "PrincipalEntityType",
		"user_pool_arn":                   "UserPoolArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
