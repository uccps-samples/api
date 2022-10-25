# MachineConfig Operator Certificates

TODO need to work out who and what.

![PKI Graph](cert-flow.png)

- [Signing Certificate/Key Pairs](#signing-certificatekey-pairs)
    - [root-ca](#root-ca)
- [Serving Certificate/Key Pairs](#serving-certificatekey-pairs)
    - [mco-mystery-cert](#mco-mystery-cert)
- [Client Certificate/Key Pairs](#client-certificatekey-pairs)
- [Certificates Without Keys](#certificates-without-keys)
- [Certificate Authority Bundles](#certificate-authority-bundles)

## Signing Certificate/Key Pairs


### root-ca
![PKI Graph](subcert-root-ca8389990467108443888.png)



| Property | Value |
| ----------- | ----------- |
| Type | Signer |
| CommonName | root-ca |
| SerialNumber | 8389990467108443888 |
| Issuer CommonName | [root-ca](#root-ca) |
| Validity | 10y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment<br/>- KeyUsageCertSign |
| ExtendedUsages |  |


#### root-ca Locations
| Namespace | Secret Name |
| ----------- | ----------- |


| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| /etc/kubernetes/ca.crt/ca.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |
|  |  |  |  |  |


## Serving Certificate/Key Pairs


### mco-mystery-cert
![PKI Graph](subcert-systemmachine-config-server6740508919873829930.png)

TODO: team needs to make description

| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | system:machine-config-server |
| SerialNumber | 6740508919873829930 |
| Issuer CommonName | [root-ca](#root-ca) |
| Validity | 10y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages |  |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - api-int.ci-ln-z2l4snt-f76d1.origin-ci-int-gce.dev.openshift.com |
| IP Addresses |  |


#### mco-mystery-cert Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-config-operator | machine-config-server-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |



## Client Certificate/Key Pairs

## Certificates Without Keys

These certificates are present in certificate authority bundles, but do not have keys in the cluster.
This happens when the installer bootstrap clusters with a set of certificate/key pairs that are deleted during the
installation process.

## Certificate Authority Bundles

