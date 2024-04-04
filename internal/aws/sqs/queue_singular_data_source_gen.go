// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sqs_queue", queueDataSource)
}

// queueDataSource returns the Terraform awscc_sqs_queue data source.
// This Terraform data source corresponds to the CloudFormation AWS::SQS::Queue resource.
func queueDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContentBasedDeduplication
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the ``ContentBasedDeduplication`` attribute for the ``CreateQueue`` action in the *API Reference*.",
		//	  "type": "boolean"
		//	}
		"content_based_deduplication": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the ``ContentBasedDeduplication`` attribute for the ``CreateQueue`` action in the *API Reference*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeduplicationScope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level. Valid values are ``messageGroup`` and ``queue``.\n To enable high throughput for a FIFO queue, set this attribute to ``messageGroup`` *and* set the ``FifoThroughputLimit`` attribute to ``perMessageGroupId``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
		//	  "type": "string"
		//	}
		"deduplication_scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level. Valid values are ``messageGroup`` and ``queue``.\n To enable high throughput for a FIFO queue, set this attribute to ``messageGroup`` *and* set the ``FifoThroughputLimit`` attribute to ``perMessageGroupId``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DelaySeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of ``0`` to ``900`` (15 minutes). The default value is ``0``.",
		//	  "type": "integer"
		//	}
		"delay_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of ``0`` to ``900`` (15 minutes). The default value is ``0``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FifoQueue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If set to true, creates a FIFO queue. If you don't specify this property, SQS creates a standard queue. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.",
		//	  "type": "boolean"
		//	}
		"fifo_queue": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If set to true, creates a FIFO queue. If you don't specify this property, SQS creates a standard queue. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FifoThroughputLimit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are ``perQueue`` and ``perMessageGroupId``.\n To enable high throughput for a FIFO queue, set this attribute to ``perMessageGroupId`` *and* set the ``DeduplicationScope`` attribute to ``messageGroup``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
		//	  "type": "string"
		//	}
		"fifo_throughput_limit": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are ``perQueue`` and ``perMessageGroupId``.\n To enable high throughput for a FIFO queue, set this attribute to ``perMessageGroupId`` *and* set the ``DeduplicationScope`` attribute to ``messageGroup``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsDataKeyReusePeriodSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time in seconds for which SQS can reuse a data key to encrypt or decrypt messages before calling KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).\n  A shorter time period provides better security, but results in more calls to KMS, which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Developer Guide*.",
		//	  "type": "integer"
		//	}
		"kms_data_key_reuse_period_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The length of time in seconds for which SQS can reuse a data key to encrypt or decrypt messages before calling KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).\n  A shorter time period provides better security, but results in more calls to KMS, which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsMasterKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an AWS Key Management Service (KMS) for SQS, or a custom KMS. To use the AWS managed KMS for SQS, specify a (default) alias ARN, alias name (e.g. ``alias/aws/sqs``), key ARN, or key ID. For more information, see the following:\n  +   [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Developer Guide* \n  +   [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *API Reference* \n  +   [Request Parameters](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the *Key Management Service API Reference* \n  +   The Key Management Service (KMS) section of the [Best Practices](https://docs.aws.amazon.com/https://d0.awsstatic.com/whitepapers/aws-kms-best-practices.pdf) whitepaper",
		//	  "type": "string"
		//	}
		"kms_master_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an AWS Key Management Service (KMS) for SQS, or a custom KMS. To use the AWS managed KMS for SQS, specify a (default) alias ARN, alias name (e.g. ``alias/aws/sqs``), key ARN, or key ID. For more information, see the following:\n  +   [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Developer Guide* \n  +   [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *API Reference* \n  +   [Request Parameters](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the *Key Management Service API Reference* \n  +   The Key Management Service (KMS) section of the [Best Practices](https://docs.aws.amazon.com/https://d0.awsstatic.com/whitepapers/aws-kms-best-practices.pdf) whitepaper",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaximumMessageSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The limit of how many bytes that a message can contain before SQS rejects it. You can specify an integer value from ``1,024`` bytes (1 KiB) to ``262,144`` bytes (256 KiB). The default value is ``262,144`` (256 KiB).",
		//	  "type": "integer"
		//	}
		"maximum_message_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The limit of how many bytes that a message can contain before SQS rejects it. You can specify an integer value from ``1,024`` bytes (1 KiB) to ``262,144`` bytes (256 KiB). The default value is ``262,144`` (256 KiB).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MessageRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of seconds that SQS retains a message. You can specify an integer value from ``60`` seconds (1 minute) to ``1,209,600`` seconds (14 days). The default value is ``345,600`` seconds (4 days).",
		//	  "type": "integer"
		//	}
		"message_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of seconds that SQS retains a message. You can specify an integer value from ``60`` seconds (1 minute) to ``1,209,600`` seconds (14 days). The default value is ``345,600`` seconds (4 days).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueueName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the ``.fifo`` suffix. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *User Guide*. \n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
		//	  "type": "string"
		//	}
		"queue_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the ``.fifo`` suffix. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *User Guide*. \n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueueUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"queue_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReceiveMessageWaitTimeSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Developer Guide*.",
		//	  "type": "integer"
		//	}
		"receive_message_wait_time_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedriveAllowPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object. The parameters are as follows:\n  +  ``redrivePermission``: The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:\n  +  ``allowAll``: (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.\n  +  ``denyAll``: No source queues can specify this queue as the dead-letter queue.\n  +  ``byQueue``: Only queues specified by the ``sourceQueueArns`` parameter can specify this queue as the dead-letter queue.\n  \n  +  ``sourceQueueArns``: The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the ``redrivePermission`` parameter is set to ``byQueue``. You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the ``redrivePermission`` parameter to ``allowAll``.",
		//	  "type": "string"
		//	}
		"redrive_allow_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object. The parameters are as follows:\n  +  ``redrivePermission``: The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:\n  +  ``allowAll``: (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.\n  +  ``denyAll``: No source queues can specify this queue as the dead-letter queue.\n  +  ``byQueue``: Only queues specified by the ``sourceQueueArns`` parameter can specify this queue as the dead-letter queue.\n  \n  +  ``sourceQueueArns``: The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the ``redrivePermission`` parameter is set to ``byQueue``. You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the ``redrivePermission`` parameter to ``allowAll``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedrivePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object. The parameters are as follows:\n  +  ``deadLetterTargetArn``: The Amazon Resource Name (ARN) of the dead-letter queue to which SQS moves messages after the value of ``maxReceiveCount`` is exceeded.\n  +  ``maxReceiveCount``: The number of times a message is delivered to the source queue before being moved to the dead-letter queue. When the ``ReceiveCount`` for a message exceeds the ``maxReceiveCount`` for a queue, SQS moves the message to the dead-letter-queue.\n  \n  The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.\n   *JSON* \n  ``{ \"deadLetterTargetArn\" : String, \"maxReceiveCount\" : Integer }`` \n  *YAML* \n  ``deadLetterTargetArn : String`` \n  ``maxReceiveCount : Integer``",
		//	  "type": "string"
		//	}
		"redrive_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object. The parameters are as follows:\n  +  ``deadLetterTargetArn``: The Amazon Resource Name (ARN) of the dead-letter queue to which SQS moves messages after the value of ``maxReceiveCount`` is exceeded.\n  +  ``maxReceiveCount``: The number of times a message is delivered to the source queue before being moved to the dead-letter queue. When the ``ReceiveCount`` for a message exceeds the ``maxReceiveCount`` for a queue, SQS moves the message to the dead-letter-queue.\n  \n  The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.\n   *JSON* \n  ``{ \"deadLetterTargetArn\" : String, \"maxReceiveCount\" : Integer }`` \n  *YAML* \n  ``deadLetterTargetArn : String`` \n  ``maxReceiveCount : Integer``",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SqsManagedSseEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)). When ``SqsManagedSseEnabled`` is not defined, ``SSE-SQS`` encryption is enabled by default.",
		//	  "type": "boolean"
		//	}
		"sqs_managed_sse_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)). When ``SqsManagedSseEnabled`` is not defined, ``SSE-SQS`` encryption is enabled by default.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags that you attach to this queue. For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *User Guide*.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			Description: "The tags that you attach to this queue. For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *User Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VisibilityTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.\n Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.\n For more information about SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Developer Guide*.",
		//	  "type": "integer"
		//	}
		"visibility_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.\n Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.\n For more information about SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Developer Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SQS::Queue",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SQS::Queue").WithTerraformTypeName("awscc_sqs_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                               "Arn",
		"content_based_deduplication":       "ContentBasedDeduplication",
		"deduplication_scope":               "DeduplicationScope",
		"delay_seconds":                     "DelaySeconds",
		"fifo_queue":                        "FifoQueue",
		"fifo_throughput_limit":             "FifoThroughputLimit",
		"key":                               "Key",
		"kms_data_key_reuse_period_seconds": "KmsDataKeyReusePeriodSeconds",
		"kms_master_key_id":                 "KmsMasterKeyId",
		"maximum_message_size":              "MaximumMessageSize",
		"message_retention_period":          "MessageRetentionPeriod",
		"queue_name":                        "QueueName",
		"queue_url":                         "QueueUrl",
		"receive_message_wait_time_seconds": "ReceiveMessageWaitTimeSeconds",
		"redrive_allow_policy":              "RedriveAllowPolicy",
		"redrive_policy":                    "RedrivePolicy",
		"sqs_managed_sse_enabled":           "SqsManagedSseEnabled",
		"tags":                              "Tags",
		"value":                             "Value",
		"visibility_timeout":                "VisibilityTimeout",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
