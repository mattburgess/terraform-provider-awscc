// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudtrail

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudtrail_event_data_store", eventDataStoreResource)
}

// eventDataStoreResource returns the Terraform awscc_cloudtrail_event_data_store resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudTrail::EventDataStore resource.
func eventDataStoreResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdvancedEventSelectors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The advanced event selectors that were used to select events for the data store.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Advanced event selectors let you create fine-grained selectors for the following AWS CloudTrail event record ?elds. They help you control costs by logging only those events that are important to you.",
		//	    "properties": {
		//	      "FieldSelectors": {
		//	        "description": "Contains all selector statements in an advanced event selector.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "A single selector statement in an advanced event selector.",
		//	          "properties": {
		//	            "EndsWith": {
		//	              "description": "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "Equals": {
		//	              "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "Field": {
		//	              "description": "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
		//	              "maxLength": 1000,
		//	              "minLength": 1,
		//	              "pattern": "([\\w|\\d|\\.|_]+)",
		//	              "type": "string"
		//	            },
		//	            "NotEndsWith": {
		//	              "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "NotEquals": {
		//	              "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "NotStartsWith": {
		//	              "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "StartsWith": {
		//	              "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "minLength": 1,
		//	                "pattern": "(.+)",
		//	                "type": "string"
		//	              },
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "required": [
		//	            "Field"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "Name": {
		//	        "description": "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
		//	        "maxLength": 1000,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FieldSelectors"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"advanced_event_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FieldSelectors
					"field_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: EndsWith
								"ends_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Equals
								"equals": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Field
								"field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
									Required:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 1000),
										stringvalidator.RegexMatches(regexp.MustCompile("([\\w|\\d|\\.|_]+)"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: NotEndsWith
								"not_ends_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: NotEquals
								"not_equals": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: NotStartsWith
								"not_starts_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: StartsWith
								"starts_with": schema.SetAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										setvalidator.SizeAtLeast(1),
										setvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("(.+)"), ""),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Contains all selector statements in an advanced event selector.",
						Required:    true,
						Validators: []validator.Set{ /*START VALIDATORS*/
							setvalidator.SizeAtLeast(1),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 1000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The advanced event selectors that were used to select events for the data store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BillingMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The mode that the event data store will use to charge for event storage.",
		//	  "type": "string"
		//	}
		"billing_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The mode that the event data store will use to charge for event storage.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of the event data store's creation.",
		//	  "type": "string"
		//	}
		"created_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of the event data store's creation.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventDataStoreArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the event data store.",
		//	  "type": "string"
		//	}
		"event_data_store_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the event data store.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FederationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether federation is enabled on an event data store.",
		//	  "type": "boolean"
		//	}
		"federation_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether federation is enabled on an event data store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FederationRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the role used for event data store federation.",
		//	  "type": "string"
		//	}
		"federation_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the role used for event data store federation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngestionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the event data store is ingesting events.",
		//	  "type": "boolean"
		//	}
		"ingestion_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the event data store is ingesting events.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InsightSelectors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing event data store. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A string that contains Insights types that are logged on an event data store.",
		//	    "properties": {
		//	      "InsightType": {
		//	        "description": "The type of Insights to log on an event data store.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"insight_selectors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InsightType
					"insight_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of Insights to log on an event data store.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing event data store. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InsightsDestination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the ARN of the event data store that will collect Insights events. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store",
		//	  "type": "string"
		//	}
		"insights_destination": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the ARN of the event data store that will collect Insights events. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MultiRegionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
		//	  "type": "boolean"
		//	}
		"multi_region_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the event data store.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the event data store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that an event data store is collecting logged events for an organization.",
		//	  "type": "boolean"
		//	}
		"organization_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that an event data store is collecting logged events for an organization.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The retention period, in days.",
		//	  "type": "integer"
		//	}
		"retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The retention period, in days.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) for this event data store.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TerminationProtectionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the event data store is protected from termination.",
		//	  "type": "boolean"
		//	}
		"termination_protection_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the event data store is protected from termination.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
		//	  "type": "string"
		//	}
		"updated_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
			Computed:    true,
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
		Description: "A storage lake of event data against which you can run complex SQL-based queries. An event data store can include events that you have logged on your account from the last 7 to 2557 or 3653 days (about seven or ten years) depending on the selected BillingMode.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::EventDataStore").WithTerraformTypeName("awscc_cloudtrail_event_data_store")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"advanced_event_selectors":       "AdvancedEventSelectors",
		"billing_mode":                   "BillingMode",
		"created_timestamp":              "CreatedTimestamp",
		"ends_with":                      "EndsWith",
		"equals":                         "Equals",
		"event_data_store_arn":           "EventDataStoreArn",
		"federation_enabled":             "FederationEnabled",
		"federation_role_arn":            "FederationRoleArn",
		"field":                          "Field",
		"field_selectors":                "FieldSelectors",
		"ingestion_enabled":              "IngestionEnabled",
		"insight_selectors":              "InsightSelectors",
		"insight_type":                   "InsightType",
		"insights_destination":           "InsightsDestination",
		"key":                            "Key",
		"kms_key_id":                     "KmsKeyId",
		"multi_region_enabled":           "MultiRegionEnabled",
		"name":                           "Name",
		"not_ends_with":                  "NotEndsWith",
		"not_equals":                     "NotEquals",
		"not_starts_with":                "NotStartsWith",
		"organization_enabled":           "OrganizationEnabled",
		"retention_period":               "RetentionPeriod",
		"starts_with":                    "StartsWith",
		"status":                         "Status",
		"tags":                           "Tags",
		"termination_protection_enabled": "TerminationProtectionEnabled",
		"updated_timestamp":              "UpdatedTimestamp",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
