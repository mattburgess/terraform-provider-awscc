---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ecs_service Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ECS::Service
---

# awscc_ecs_service (Data Source)

Data Source schema for AWS::ECS::Service



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **capacity_provider_strategy** (Attributes List) (see [below for nested schema](#nestedatt--capacity_provider_strategy))
- **cluster** (String)
- **deployment_configuration** (Attributes) (see [below for nested schema](#nestedatt--deployment_configuration))
- **deployment_controller** (Attributes) (see [below for nested schema](#nestedatt--deployment_controller))
- **desired_count** (Number)
- **enable_ecs_managed_tags** (Boolean)
- **enable_execute_command** (Boolean)
- **health_check_grace_period_seconds** (Number)
- **launch_type** (String)
- **load_balancers** (Attributes List) (see [below for nested schema](#nestedatt--load_balancers))
- **name** (String)
- **network_configuration** (Attributes) (see [below for nested schema](#nestedatt--network_configuration))
- **placement_constraints** (Attributes List) (see [below for nested schema](#nestedatt--placement_constraints))
- **placement_strategies** (Attributes List) (see [below for nested schema](#nestedatt--placement_strategies))
- **platform_version** (String)
- **propagate_tags** (String)
- **role** (String)
- **scheduling_strategy** (String)
- **service_arn** (String)
- **service_name** (String)
- **service_registries** (Attributes List) (see [below for nested schema](#nestedatt--service_registries))
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))
- **task_definition** (String)

<a id="nestedatt--capacity_provider_strategy"></a>
### Nested Schema for `capacity_provider_strategy`

Read-Only:

- **base** (Number)
- **capacity_provider** (String)
- **weight** (Number)


<a id="nestedatt--deployment_configuration"></a>
### Nested Schema for `deployment_configuration`

Read-Only:

- **deployment_circuit_breaker** (Attributes) (see [below for nested schema](#nestedatt--deployment_configuration--deployment_circuit_breaker))
- **maximum_percent** (Number)
- **minimum_healthy_percent** (Number)

<a id="nestedatt--deployment_configuration--deployment_circuit_breaker"></a>
### Nested Schema for `deployment_configuration.deployment_circuit_breaker`

Read-Only:

- **enable** (Boolean)
- **rollback** (Boolean)



<a id="nestedatt--deployment_controller"></a>
### Nested Schema for `deployment_controller`

Read-Only:

- **type** (String)


<a id="nestedatt--load_balancers"></a>
### Nested Schema for `load_balancers`

Read-Only:

- **container_name** (String)
- **container_port** (Number)
- **load_balancer_name** (String)
- **target_group_arn** (String)


<a id="nestedatt--network_configuration"></a>
### Nested Schema for `network_configuration`

Read-Only:

- **awsvpc_configuration** (Attributes) (see [below for nested schema](#nestedatt--network_configuration--awsvpc_configuration))

<a id="nestedatt--network_configuration--awsvpc_configuration"></a>
### Nested Schema for `network_configuration.awsvpc_configuration`

Read-Only:

- **assign_public_ip** (String)
- **security_groups** (List of String)
- **subnets** (List of String)



<a id="nestedatt--placement_constraints"></a>
### Nested Schema for `placement_constraints`

Read-Only:

- **expression** (String)
- **type** (String)


<a id="nestedatt--placement_strategies"></a>
### Nested Schema for `placement_strategies`

Read-Only:

- **field** (String)
- **type** (String)


<a id="nestedatt--service_registries"></a>
### Nested Schema for `service_registries`

Read-Only:

- **container_name** (String)
- **container_port** (Number)
- **port** (Number)
- **registry_arn** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

