// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_s3_access_grant", accessGrantResource)
}

// accessGrantResource returns the Terraform awscc_s3_access_grant resource.
// This Terraform resource corresponds to the CloudFormation AWS::S3::AccessGrant resource.
func accessGrantResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessGrantArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified access grant.",
		//	  "type": "string"
		//	}
		"access_grant_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified access grant.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID assigned to this access grant.",
		//	  "examples": [
		//	    "7c89cbd1-0f4e-40e3-861d-afb906952b77"
		//	  ],
		//	  "type": "string"
		//	}
		"access_grant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID assigned to this access grant.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantsLocationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options of the grant location, which is the S3 path to the data to which you are granting access.",
		//	  "properties": {
		//	    "S3SubPrefix": {
		//	      "description": "The S3 sub prefix of a registered location in your S3 Access Grants instance",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3SubPrefix"
		//	  ],
		//	  "type": "object"
		//	}
		"access_grants_location_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3SubPrefix
				"s3_sub_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The S3 sub prefix of a registered location in your S3 Access Grants instance",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options of the grant location, which is the S3 path to the data to which you are granting access.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantsLocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The custom S3 location to be accessed by the grantee",
		//	  "examples": [
		//	    "125f332b-a499-4eb6-806f-8a6a1aa4cb96"
		//	  ],
		//	  "type": "string"
		//	}
		"access_grants_location_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The custom S3 location to be accessed by the grantee",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ApplicationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the application grantees will use to access the location",
		//	  "type": "string"
		//	}
		"application_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the application grantees will use to access the location",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GrantScope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.",
		//	  "type": "string"
		//	}
		"grant_scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Grantee
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The principal who will be granted permission to access S3.",
		//	  "properties": {
		//	    "GranteeIdentifier": {
		//	      "description": "The unique identifier of the Grantee",
		//	      "type": "string"
		//	    },
		//	    "GranteeType": {
		//	      "description": "Configures the transfer acceleration state for an Amazon S3 bucket.",
		//	      "enum": [
		//	        "IAM",
		//	        "DIRECTORY_USER",
		//	        "DIRECTORY_GROUP"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "GranteeType",
		//	    "GranteeIdentifier"
		//	  ],
		//	  "type": "object"
		//	}
		"grantee": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GranteeIdentifier
				"grantee_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The unique identifier of the Grantee",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: GranteeType
				"grantee_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Configures the transfer acceleration state for an Amazon S3 bucket.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"IAM",
							"DIRECTORY_USER",
							"DIRECTORY_GROUP",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The principal who will be granted permission to access S3.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Permission
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The level of access to be afforded to the grantee",
		//	  "enum": [
		//	    "READ",
		//	    "WRITE",
		//	    "READWRITE"
		//	  ],
		//	  "type": "string"
		//	}
		"permission": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The level of access to be afforded to the grantee",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"READ",
					"WRITE",
					"READWRITE",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: S3PrefixType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of S3SubPrefix.",
		//	  "enum": [
		//	    "Object"
		//	  ],
		//	  "type": "string"
		//	}
		"s3_prefix_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of S3SubPrefix.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Object",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// S3PrefixType is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
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
		Description: "The AWS::S3::AccessGrant resource is an Amazon S3 resource type representing permissions to a specific S3 bucket or prefix hosted in an S3 Access Grants instance.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::AccessGrant").WithTerraformTypeName("awscc_s3_access_grant")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_grant_arn":                     "AccessGrantArn",
		"access_grant_id":                      "AccessGrantId",
		"access_grants_location_configuration": "AccessGrantsLocationConfiguration",
		"access_grants_location_id":            "AccessGrantsLocationId",
		"application_arn":                      "ApplicationArn",
		"grant_scope":                          "GrantScope",
		"grantee":                              "Grantee",
		"grantee_identifier":                   "GranteeIdentifier",
		"grantee_type":                         "GranteeType",
		"key":                                  "Key",
		"permission":                           "Permission",
		"s3_prefix_type":                       "S3PrefixType",
		"s3_sub_prefix":                        "S3SubPrefix",
		"tags":                                 "Tags",
		"value":                                "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
		"/properties/S3PrefixType",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
