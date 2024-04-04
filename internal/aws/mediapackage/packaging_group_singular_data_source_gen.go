// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediapackage_packaging_group", packagingGroupDataSource)
}

// packagingGroupDataSource returns the Terraform awscc_mediapackage_packaging_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaPackage::PackagingGroup resource.
func packagingGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the PackagingGroup.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the PackagingGroup.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Authorization
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "CDN Authorization",
		//	  "properties": {
		//	    "CdnIdentifierSecret": {
		//	      "description": "The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.",
		//	      "type": "string"
		//	    },
		//	    "SecretsRoleArn": {
		//	      "description": "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "CdnIdentifierSecret",
		//	    "SecretsRoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"authorization": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CdnIdentifierSecret
				"cdn_identifier_secret": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretsRoleArn
				"secrets_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "CDN Authorization",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The fully qualified domain name for Assets in the PackagingGroup.",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The fully qualified domain name for Assets in the PackagingGroup.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EgressAccessLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration parameters for egress access logging.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.",
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"egress_access_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration parameters for egress access logging.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the PackagingGroup.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"packaging_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the PackagingGroup.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
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
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaPackage::PackagingGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::PackagingGroup").WithTerraformTypeName("awscc_mediapackage_packaging_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"authorization":         "Authorization",
		"cdn_identifier_secret": "CdnIdentifierSecret",
		"domain_name":           "DomainName",
		"egress_access_logs":    "EgressAccessLogs",
		"key":                   "Key",
		"log_group_name":        "LogGroupName",
		"packaging_group_id":    "Id",
		"secrets_role_arn":      "SecretsRoleArn",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
