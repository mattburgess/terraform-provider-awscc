---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_acmpca_certificate_authority Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Private certificate authority.
---

# awscc_acmpca_certificate_authority (Resource)

Private certificate authority.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **key_algorithm** (String) Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
- **signing_algorithm** (String) Algorithm your CA uses to sign certificate requests.
- **subject** (Attributes) Structure that contains X.500 distinguished name information for your CA. (see [below for nested schema](#nestedatt--subject))
- **type** (String) The type of the certificate authority.

### Optional

- **csr_extensions** (Attributes) Structure that contains CSR pass though extensions information. (see [below for nested schema](#nestedatt--csr_extensions))
- **key_storage_security_standard** (String) KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.
- **revocation_configuration** (Attributes) Certificate Authority revocation information. (see [below for nested schema](#nestedatt--revocation_configuration))
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String)
- **certificate_signing_request** (String) The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--subject"></a>
### Nested Schema for `subject`

Required:

- **common_name** (String)
- **country** (String)
- **distinguished_name_qualifier** (String)
- **generation_qualifier** (String)
- **given_name** (String)
- **initials** (String)
- **locality** (String)
- **organization** (String)
- **organizational_unit** (String)
- **pseudonym** (String)
- **serial_number** (String)
- **state** (String)
- **surname** (String)
- **title** (String)


<a id="nestedatt--csr_extensions"></a>
### Nested Schema for `csr_extensions`

Optional:

- **key_usage** (Attributes) Structure that contains X.509 KeyUsage information. (see [below for nested schema](#nestedatt--csr_extensions--key_usage))
- **subject_information_access** (Attributes List) Array of X.509 AccessDescription. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access))

<a id="nestedatt--csr_extensions--key_usage"></a>
### Nested Schema for `csr_extensions.key_usage`

Optional:

- **crl_sign** (Boolean)
- **data_encipherment** (Boolean)
- **decipher_only** (Boolean)
- **digital_signature** (Boolean)
- **encipher_only** (Boolean)
- **key_agreement** (Boolean)
- **key_cert_sign** (Boolean)
- **key_encipherment** (Boolean)
- **non_repudiation** (Boolean)


<a id="nestedatt--csr_extensions--subject_information_access"></a>
### Nested Schema for `csr_extensions.subject_information_access`

Optional:

- **access_location** (Attributes) Structure that contains X.509 GeneralName information. Assign one and ONLY one field. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access--access_location))
- **access_method** (Attributes) Structure that contains X.509 AccessMethod information. Assign one and ONLY one field. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access--access_method))

<a id="nestedatt--csr_extensions--subject_information_access--access_location"></a>
### Nested Schema for `csr_extensions.subject_information_access.access_location`

Optional:

- **directory_name** (Attributes) Structure that contains X.500 distinguished name information for your CA. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access--access_location--directory_name))
- **dns_name** (String) String that contains X.509 DnsName information.
- **edi_party_name** (Attributes) Structure that contains X.509 EdiPartyName information. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access--access_location--edi_party_name))
- **ip_address** (String) String that contains X.509 IpAddress information.
- **other_name** (Attributes) Structure that contains X.509 OtherName information. (see [below for nested schema](#nestedatt--csr_extensions--subject_information_access--access_location--other_name))
- **registered_id** (String) String that contains X.509 ObjectIdentifier information.
- **rfc_822_name** (String) String that contains X.509 Rfc822Name information.
- **uniform_resource_identifier** (String) String that contains X.509 UniformResourceIdentifier information.

<a id="nestedatt--csr_extensions--subject_information_access--access_location--directory_name"></a>
### Nested Schema for `csr_extensions.subject_information_access.access_location.uniform_resource_identifier`

Optional:

- **common_name** (String)
- **country** (String)
- **distinguished_name_qualifier** (String)
- **generation_qualifier** (String)
- **given_name** (String)
- **initials** (String)
- **locality** (String)
- **organization** (String)
- **organizational_unit** (String)
- **pseudonym** (String)
- **serial_number** (String)
- **state** (String)
- **surname** (String)
- **title** (String)


<a id="nestedatt--csr_extensions--subject_information_access--access_location--edi_party_name"></a>
### Nested Schema for `csr_extensions.subject_information_access.access_location.uniform_resource_identifier`

Optional:

- **name_assigner** (String)
- **party_name** (String)


<a id="nestedatt--csr_extensions--subject_information_access--access_location--other_name"></a>
### Nested Schema for `csr_extensions.subject_information_access.access_location.uniform_resource_identifier`

Optional:

- **type_id** (String) String that contains X.509 ObjectIdentifier information.
- **value** (String)



<a id="nestedatt--csr_extensions--subject_information_access--access_method"></a>
### Nested Schema for `csr_extensions.subject_information_access.access_method`

Optional:

- **access_method_type** (String) Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.
- **custom_object_identifier** (String) String that contains X.509 ObjectIdentifier information.




<a id="nestedatt--revocation_configuration"></a>
### Nested Schema for `revocation_configuration`

Optional:

- **crl_configuration** (Attributes) Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked. (see [below for nested schema](#nestedatt--revocation_configuration--crl_configuration))

<a id="nestedatt--revocation_configuration--crl_configuration"></a>
### Nested Schema for `revocation_configuration.crl_configuration`

Optional:

- **custom_cname** (String)
- **enabled** (Boolean)
- **expiration_in_days** (Number)
- **s3_bucket_name** (String)
- **s3_object_acl** (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

