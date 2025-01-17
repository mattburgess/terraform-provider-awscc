---
page_title: "awscc_bedrock_prompt_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Bedrock::PromptVersion Resource Type
---

# awscc_bedrock_prompt_version (Resource)

Definition of AWS::Bedrock::PromptVersion Resource Type

## Example Usage

```terraform
resource "awscc_bedrock_prompt_version" "example" {
  prompt_arn  = awscc_bedrock_prompt.example.arn
  description = "example"
}

resource "awscc_bedrock_prompt" "example" {
  name                        = "example"
  description                 = "example"
  customer_encryption_key_arn = awscc_kms_key.example.arn
  default_variant             = "variant-example"

  variants = [
    {
      name          = "variant-example"
      template_type = "TEXT"
      model_id      = "amazon.titan-text-express-v1"
      inference_configuration = {
        text = {
          temperature    = 1
          top_p          = 0.9900000095367432
          max_tokens     = 300
          stop_sequences = ["\\n\\nHuman:"]
          top_k          = 250
        }
      }
      template_configuration = {
        text = {
          input_variables = [
            {
              name = "topic"
            }
          ]
          text = "Make me a {{genre}} playlist consisting of the following number of songs: {{number}}."
        }
      }
    }

  ]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `prompt_arn` (String) ARN of a prompt resource possibly with a version

### Optional

- `description` (String) Description for a prompt version resource.

### Read-Only

- `arn` (String) ARN of a prompt version resource
- `created_at` (String) Time Stamp.
- `default_variant` (String) Name for a variant.
- `id` (String) Uniquely identifies the resource.
- `name` (String) Name for a prompt resource.
- `prompt_id` (String) Identifier for a Prompt
- `updated_at` (String) Time Stamp.
- `variants` (Attributes List) List of prompt variants (see [below for nested schema](#nestedatt--variants))
- `version` (String) Version.

<a id="nestedatt--variants"></a>
### Nested Schema for `variants`

Read-Only:

- `inference_configuration` (Attributes) Model inference configuration (see [below for nested schema](#nestedatt--variants--inference_configuration))
- `model_id` (String) ARN or name of a Bedrock model.
- `name` (String) Name for a variant.
- `template_configuration` (Attributes) Prompt template configuration (see [below for nested schema](#nestedatt--variants--template_configuration))
- `template_type` (String) Prompt template type

<a id="nestedatt--variants--inference_configuration"></a>
### Nested Schema for `variants.inference_configuration`

Read-Only:

- `text` (Attributes) Prompt model inference configuration (see [below for nested schema](#nestedatt--variants--inference_configuration--text))

<a id="nestedatt--variants--inference_configuration--text"></a>
### Nested Schema for `variants.inference_configuration.text`

Read-Only:

- `max_tokens` (Number) Maximum length of output
- `stop_sequences` (List of String) List of stop sequences
- `temperature` (Number) Controls randomness, higher values increase diversity
- `top_k` (Number) Sample from the k most likely next tokens
- `top_p` (Number) Cumulative probability cutoff for token selection



<a id="nestedatt--variants--template_configuration"></a>
### Nested Schema for `variants.template_configuration`

Read-Only:

- `text` (Attributes) Configuration for text prompt template (see [below for nested schema](#nestedatt--variants--template_configuration--text))

<a id="nestedatt--variants--template_configuration--text"></a>
### Nested Schema for `variants.template_configuration.text`

Read-Only:

- `input_variables` (Attributes List) List of input variables (see [below for nested schema](#nestedatt--variants--template_configuration--text--input_variables))
- `text` (String) Prompt content for String prompt template

<a id="nestedatt--variants--template_configuration--text--input_variables"></a>
### Nested Schema for `variants.template_configuration.text.input_variables`

Read-Only:

- `name` (String) Name for an input variable

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_bedrock_prompt_version.example <resource ID>
```