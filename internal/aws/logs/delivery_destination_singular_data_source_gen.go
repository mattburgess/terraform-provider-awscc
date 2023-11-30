// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_delivery_destination", deliveryDestinationDataSource)
}

// deliveryDestinationDataSource returns the Terraform awscc_logs_delivery_destination data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::DeliveryDestination resource.
func deliveryDestinationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeliveryDestinationPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 51200",
		//	  "items": {
		//	    "$ref": "#/definitions/DestinationPolicy"
		//	  },
		//	  "type": "object"
		//	}
		"delivery_destination_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 51200",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeliveryDestinationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
		//	  "maxLength": 12,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z]+$",
		//	  "type": "string"
		//	}
		"delivery_destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the AWS resource that will receive the logs.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"destination_resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the AWS resource that will receive the logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of this delivery destination.",
		//	  "maxLength": 60,
		//	  "minLength": 1,
		//	  "pattern": "[\\w-]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of this delivery destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags that have been assigned to this delivery destination.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags that have been assigned to this delivery destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::DeliveryDestination",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::DeliveryDestination").WithTerraformTypeName("awscc_logs_delivery_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"delivery_destination_policy": "DeliveryDestinationPolicy",
		"delivery_destination_type":   "DeliveryDestinationType",
		"destination_resource_arn":    "DestinationResourceArn",
		"key":                         "Key",
		"name":                        "Name",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
