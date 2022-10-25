# Service Serving Certificates

Used to secure inter-service communication on the local cluster.

![PKI Graph](cert-flow.png)

- [Signing Certificate/Key Pairs](#signing-certificatekey-pairs)
    - [service-serving-signer](#service-serving-signer)
- [Serving Certificate/Key Pairs](#serving-certificatekey-pairs)
    - [alertmanager-main.uccp-monitoring.svc](#alertmanager-main.uccp-monitoring.svc)
    - [api.uccp-apiserver.svc](#api.uccp-apiserver.svc)
    - [api.uccp-oauth-apiserver.svc](#api.uccp-oauth-apiserver.svc)
    - [catalog-operator-metrics.uccp-operator-lifecycle-manager.svc](#catalog-operator-metrics.uccp-operator-lifecycle-manager.svc)
    - [cco-metrics.uccp-cloud-credential-operator.svc](#cco-metrics.uccp-cloud-credential-operator.svc)
    - [cluster-autoscaler-operator.uccp-machine-api.svc](#cluster-autoscaler-operator.uccp-machine-api.svc)
    - [cluster-monitoring-operator.uccp-monitoring.svc](#cluster-monitoring-operator.uccp-monitoring.svc)
    - [cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc](#cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc)
    - [console.uccp-console.svc](#console.uccp-console.svc)
    - [controller-manager.uccp-controller-manager.svc](#controller-manager.uccp-controller-manager.svc)
    - [csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc](#csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc)
    - [csi-snapshot-webhook.uccp-cluster-storage-operator.svc](#csi-snapshot-webhook.uccp-cluster-storage-operator.svc)
    - [dns-default.uccp-dns.svc](#dns-default.uccp-dns.svc)
    - [etcd.uccp-etcd.svc](#etcd.uccp-etcd.svc)
    - [grafana.uccp-monitoring.svc](#grafana.uccp-monitoring.svc)
    - [image-registry-operator.uccp-image-registry.svc](#image-registry-operator.uccp-image-registry.svc)
    - [image-registry.uccp-image-registry.svc](#image-registry.uccp-image-registry.svc)
    - [kube-controller-manager.uccp-kube-controller-manager.svc](#kube-controller-manager.uccp-kube-controller-manager.svc)
    - [kube-state-metrics.uccp-monitoring.svc](#kube-state-metrics.uccp-monitoring.svc)
    - [machine-api-controllers.uccp-machine-api.svc](#machine-api-controllers.uccp-machine-api.svc)
    - [machine-api-operator-webhook.uccp-machine-api.svc](#machine-api-operator-webhook.uccp-machine-api.svc)
    - [machine-api-operator.uccp-machine-api.svc](#machine-api-operator.uccp-machine-api.svc)
    - [machine-approver.uccp-cluster-machine-approver.svc](#machine-approver.uccp-cluster-machine-approver.svc)
    - [machine-config-daemon.uccp-machine-config-operator.svc](#machine-config-daemon.uccp-machine-config-operator.svc)
    - [marketplace-operator-metrics.uccp-marketplace.svc](#marketplace-operator-metrics.uccp-marketplace.svc)
    - [metrics.uccp-apiserver-operator.svc](#metrics.uccp-apiserver-operator.svc)
    - [metrics.uccp-authentication-operator.svc](#metrics.uccp-authentication-operator.svc)
    - [metrics.uccp-cluster-samples-operator.svc](#metrics.uccp-cluster-samples-operator.svc)
    - [metrics.uccp-config-operator.svc](#metrics.uccp-config-operator.svc)
    - [metrics.uccp-console-operator.svc](#metrics.uccp-console-operator.svc)
    - [metrics.uccp-controller-manager-operator.svc](#metrics.uccp-controller-manager-operator.svc)
    - [metrics.uccp-dns-operator.svc](#metrics.uccp-dns-operator.svc)
    - [metrics.uccp-etcd-operator.svc](#metrics.uccp-etcd-operator.svc)
    - [metrics.uccp-ingress-operator.svc](#metrics.uccp-ingress-operator.svc)
    - [metrics.uccp-insights.svc](#metrics.uccp-insights.svc)
    - [metrics.uccp-kube-apiserver-operator.svc](#metrics.uccp-kube-apiserver-operator.svc)
    - [metrics.uccp-kube-controller-manager-operator.svc](#metrics.uccp-kube-controller-manager-operator.svc)
    - [metrics.uccp-kube-scheduler-operator.svc](#metrics.uccp-kube-scheduler-operator.svc)
    - [metrics.uccp-kube-storage-version-migrator-operator.svc](#metrics.uccp-kube-storage-version-migrator-operator.svc)
    - [metrics.uccp-service-ca-operator.svc](#metrics.uccp-service-ca-operator.svc)
    - [multus-admission-controller.uccp-multus.svc](#multus-admission-controller.uccp-multus.svc)
    - [network-metrics-service.uccp-multus.svc](#network-metrics-service.uccp-multus.svc)
    - [node-exporter.uccp-monitoring.svc](#node-exporter.uccp-monitoring.svc)
    - [node-tuning-operator.uccp-cluster-node-tuning-operator.svc](#node-tuning-operator.uccp-cluster-node-tuning-operator.svc)
    - [oauth-openshift.uccp-authentication.svc](#oauth-openshift.uccp-authentication.svc)
    - [olm-operator-metrics.uccp-operator-lifecycle-manager.svc](#olm-operator-metrics.uccp-operator-lifecycle-manager.svc)
    - [uccp-state-metrics.uccp-monitoring.svc](#uccp-state-metrics.uccp-monitoring.svc)
    - [prometheus-adapter.uccp-monitoring.svc](#prometheus-adapter.uccp-monitoring.svc)
    - [prometheus-k8s-thanos-sidecar.uccp-monitoring.svc](#prometheus-k8s-thanos-sidecar.uccp-monitoring.svc)
    - [prometheus-k8s.uccp-monitoring.svc](#prometheus-k8s.uccp-monitoring.svc)
    - [prometheus-operator.uccp-monitoring.svc](#prometheus-operator.uccp-monitoring.svc)
    - [router-internal-default.uccp-ingress.svc](#router-internal-default.uccp-ingress.svc)
    - [scheduler.uccp-kube-scheduler.svc](#scheduler.uccp-kube-scheduler.svc)
    - [sdn.uccp-sdn.svc](#sdn.uccp-sdn.svc)
    - [telemeter-client.uccp-monitoring.svc](#telemeter-client.uccp-monitoring.svc)
    - [thanos-querier.uccp-monitoring.svc](#thanos-querier.uccp-monitoring.svc)
- [Client Certificate/Key Pairs](#client-certificatekey-pairs)
- [Certificates Without Keys](#certificates-without-keys)
- [Certificate Authority Bundles](#certificate-authority-bundles)
    - [service-ca](#service-ca)

## Signing Certificate/Key Pairs


### service-serving-signer
![PKI Graph](subcert-uccp-service-serving-signer16221335705446463666206287945.png)

Signer used by service-ca to sign serving certificates for internal service DNS names.

| Property | Value |
| ----------- | ----------- |
| Type | Signer |
| CommonName | uccp-service-serving-signer@1622133570 |
| SerialNumber | 5446463666206287945 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y60d |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment<br/>- KeyUsageCertSign |
| ExtendedUsages |  |


#### service-serving-signer Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-service-ca | signing-key |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |



## Serving Certificate/Key Pairs


### alertmanager-main.uccp-monitoring.svc
![PKI Graph](subcert-alertmanager-main.uccp-monitoring.svc7889519590428116393.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | alertmanager-main.uccp-monitoring.svc |
| SerialNumber | 7889519590428116393 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - alertmanager-main.uccp-monitoring.svc<br/>- alertmanager-main.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### alertmanager-main.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | alertmanager-main-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### api.uccp-apiserver.svc
![PKI Graph](subcert-api.uccp-apiserver.svc2115297822024425807.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | api.uccp-apiserver.svc |
| SerialNumber | 2115297822024425807 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - api.uccp-apiserver.svc<br/>- api.uccp-apiserver.svc.cluster.local |
| IP Addresses |  |


#### api.uccp-apiserver.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-apiserver | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### api.uccp-oauth-apiserver.svc
![PKI Graph](subcert-api.uccp-oauth-apiserver.svc485864516996010702.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | api.uccp-oauth-apiserver.svc |
| SerialNumber | 485864516996010702 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - api.uccp-oauth-apiserver.svc<br/>- api.uccp-oauth-apiserver.svc.cluster.local |
| IP Addresses |  |


#### api.uccp-oauth-apiserver.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-oauth-apiserver | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### catalog-operator-metrics.uccp-operator-lifecycle-manager.svc
![PKI Graph](subcert-catalog-operator-metrics.uccp-operator-lifecycle-manager.svc5069841447494153423.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | catalog-operator-metrics.uccp-operator-lifecycle-manager.svc |
| SerialNumber | 5069841447494153423 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - catalog-operator-metrics.uccp-operator-lifecycle-manager.svc<br/>- catalog-operator-metrics.uccp-operator-lifecycle-manager.svc.cluster.local |
| IP Addresses |  |


#### catalog-operator-metrics.uccp-operator-lifecycle-manager.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-operator-lifecycle-manager | catalog-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### cco-metrics.uccp-cloud-credential-operator.svc
![PKI Graph](subcert-cco-metrics.uccp-cloud-credential-operator.svc6893942648061066111.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | cco-metrics.uccp-cloud-credential-operator.svc |
| SerialNumber | 6893942648061066111 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - cco-metrics.uccp-cloud-credential-operator.svc<br/>- cco-metrics.uccp-cloud-credential-operator.svc.cluster.local |
| IP Addresses |  |


#### cco-metrics.uccp-cloud-credential-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cloud-credential-operator | cloud-credential-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### cluster-autoscaler-operator.uccp-machine-api.svc
![PKI Graph](subcert-cluster-autoscaler-operator.uccp-machine-api.svc8305498258745803921.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | cluster-autoscaler-operator.uccp-machine-api.svc |
| SerialNumber | 8305498258745803921 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - cluster-autoscaler-operator.uccp-machine-api.svc<br/>- cluster-autoscaler-operator.uccp-machine-api.svc.cluster.local |
| IP Addresses |  |


#### cluster-autoscaler-operator.uccp-machine-api.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-api | cluster-autoscaler-operator-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### cluster-monitoring-operator.uccp-monitoring.svc
![PKI Graph](subcert-cluster-monitoring-operator.uccp-monitoring.svc8944425219050445342.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | cluster-monitoring-operator.uccp-monitoring.svc |
| SerialNumber | 8944425219050445342 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - cluster-monitoring-operator.uccp-monitoring.svc<br/>- cluster-monitoring-operator.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### cluster-monitoring-operator.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | cluster-monitoring-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc
![PKI Graph](subcert-cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc5294915743012167366.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc |
| SerialNumber | 5294915743012167366 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc<br/>- cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc.cluster.local |
| IP Addresses |  |


#### cluster-storage-operator-metrics.uccp-cluster-storage-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-storage-operator | cluster-storage-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### console.uccp-console.svc
![PKI Graph](subcert-console.uccp-console.svc2317112508926355245.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | console.uccp-console.svc |
| SerialNumber | 2317112508926355245 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - console.uccp-console.svc<br/>- console.uccp-console.svc.cluster.local |
| IP Addresses |  |


#### console.uccp-console.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-console | console-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### controller-manager.uccp-controller-manager.svc
![PKI Graph](subcert-controller-manager.uccp-controller-manager.svc8478506552664981432.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | controller-manager.uccp-controller-manager.svc |
| SerialNumber | 8478506552664981432 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - controller-manager.uccp-controller-manager.svc<br/>- controller-manager.uccp-controller-manager.svc.cluster.local |
| IP Addresses |  |


#### controller-manager.uccp-controller-manager.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-controller-manager | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc
![PKI Graph](subcert-csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc5025200834724127258.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc |
| SerialNumber | 5025200834724127258 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc<br/>- csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc.cluster.local |
| IP Addresses |  |


#### csi-snapshot-controller-operator-metrics.uccp-cluster-storage-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-storage-operator | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### csi-snapshot-webhook.uccp-cluster-storage-operator.svc
![PKI Graph](subcert-csi-snapshot-webhook.uccp-cluster-storage-operator.svc1282769300244468729.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | csi-snapshot-webhook.uccp-cluster-storage-operator.svc |
| SerialNumber | 1282769300244468729 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - csi-snapshot-webhook.uccp-cluster-storage-operator.svc<br/>- csi-snapshot-webhook.uccp-cluster-storage-operator.svc.cluster.local |
| IP Addresses |  |


#### csi-snapshot-webhook.uccp-cluster-storage-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-storage-operator | csi-snapshot-webhook-secret |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### dns-default.uccp-dns.svc
![PKI Graph](subcert-dns-default.uccp-dns.svc495137039081958925.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | dns-default.uccp-dns.svc |
| SerialNumber | 495137039081958925 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - dns-default.uccp-dns.svc<br/>- dns-default.uccp-dns.svc.cluster.local |
| IP Addresses |  |


#### dns-default.uccp-dns.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-dns | dns-default-metrics-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### etcd.uccp-etcd.svc
![PKI Graph](subcert-etcd.uccp-etcd.svc1695572914480243966.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | etcd.uccp-etcd.svc |
| SerialNumber | 1695572914480243966 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - etcd.uccp-etcd.svc<br/>- etcd.uccp-etcd.svc.cluster.local |
| IP Addresses |  |


#### etcd.uccp-etcd.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-etcd | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### grafana.uccp-monitoring.svc
![PKI Graph](subcert-grafana.uccp-monitoring.svc5127637701693466147.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | grafana.uccp-monitoring.svc |
| SerialNumber | 5127637701693466147 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - grafana.uccp-monitoring.svc<br/>- grafana.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### grafana.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | grafana-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### image-registry-operator.uccp-image-registry.svc
![PKI Graph](subcert-image-registry-operator.uccp-image-registry.svc4967320171357519668.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | image-registry-operator.uccp-image-registry.svc |
| SerialNumber | 4967320171357519668 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - image-registry-operator.uccp-image-registry.svc<br/>- image-registry-operator.uccp-image-registry.svc.cluster.local |
| IP Addresses |  |


#### image-registry-operator.uccp-image-registry.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-image-registry | image-registry-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### image-registry.uccp-image-registry.svc
![PKI Graph](subcert-image-registry.uccp-image-registry.svc7911780555769594156.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | image-registry.uccp-image-registry.svc |
| SerialNumber | 7911780555769594156 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - image-registry.uccp-image-registry.svc<br/>- image-registry.uccp-image-registry.svc.cluster.local |
| IP Addresses |  |


#### image-registry.uccp-image-registry.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-image-registry | image-registry-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### kube-controller-manager.uccp-kube-controller-manager.svc
![PKI Graph](subcert-kube-controller-manager.uccp-kube-controller-manager.svc4706016511474554482.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | kube-controller-manager.uccp-kube-controller-manager.svc |
| SerialNumber | 4706016511474554482 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - kube-controller-manager.uccp-kube-controller-manager.svc<br/>- kube-controller-manager.uccp-kube-controller-manager.svc.cluster.local |
| IP Addresses |  |


#### kube-controller-manager.uccp-kube-controller-manager.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-controller-manager | serving-cert |
| uccp-kube-controller-manager | serving-cert-2 |
| uccp-kube-controller-manager | serving-cert-3 |
| uccp-kube-controller-manager | serving-cert-4 |
| uccp-kube-controller-manager | serving-cert-5 |
| uccp-kube-controller-manager | serving-cert-6 |
| uccp-kube-controller-manager | serving-cert-7 |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-3/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-3/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-4/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-4/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-5/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-5/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-6/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-6/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-7/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-7/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |



### kube-state-metrics.uccp-monitoring.svc
![PKI Graph](subcert-kube-state-metrics.uccp-monitoring.svc2719079659670312610.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | kube-state-metrics.uccp-monitoring.svc |
| SerialNumber | 2719079659670312610 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - kube-state-metrics.uccp-monitoring.svc<br/>- kube-state-metrics.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### kube-state-metrics.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | kube-state-metrics-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### machine-api-controllers.uccp-machine-api.svc
![PKI Graph](subcert-machine-api-controllers.uccp-machine-api.svc7828335248087138693.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | machine-api-controllers.uccp-machine-api.svc |
| SerialNumber | 7828335248087138693 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - machine-api-controllers.uccp-machine-api.svc<br/>- machine-api-controllers.uccp-machine-api.svc.cluster.local |
| IP Addresses |  |


#### machine-api-controllers.uccp-machine-api.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-api | machine-api-controllers-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### machine-api-operator-webhook.uccp-machine-api.svc
![PKI Graph](subcert-machine-api-operator-webhook.uccp-machine-api.svc2625486651396182955.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | machine-api-operator-webhook.uccp-machine-api.svc |
| SerialNumber | 2625486651396182955 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - machine-api-operator-webhook.uccp-machine-api.svc<br/>- machine-api-operator-webhook.uccp-machine-api.svc.cluster.local |
| IP Addresses |  |


#### machine-api-operator-webhook.uccp-machine-api.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-api | machine-api-operator-webhook-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### machine-api-operator.uccp-machine-api.svc
![PKI Graph](subcert-machine-api-operator.uccp-machine-api.svc5923902699021639505.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | machine-api-operator.uccp-machine-api.svc |
| SerialNumber | 5923902699021639505 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - machine-api-operator.uccp-machine-api.svc<br/>- machine-api-operator.uccp-machine-api.svc.cluster.local |
| IP Addresses |  |


#### machine-api-operator.uccp-machine-api.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-api | machine-api-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### machine-approver.uccp-cluster-machine-approver.svc
![PKI Graph](subcert-machine-approver.uccp-cluster-machine-approver.svc1831298527724831562.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | machine-approver.uccp-cluster-machine-approver.svc |
| SerialNumber | 1831298527724831562 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - machine-approver.uccp-cluster-machine-approver.svc<br/>- machine-approver.uccp-cluster-machine-approver.svc.cluster.local |
| IP Addresses |  |


#### machine-approver.uccp-cluster-machine-approver.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-machine-approver | machine-approver-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### machine-config-daemon.uccp-machine-config-operator.svc
![PKI Graph](subcert-machine-config-daemon.uccp-machine-config-operator.svc894043062816778974.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | machine-config-daemon.uccp-machine-config-operator.svc |
| SerialNumber | 894043062816778974 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - machine-config-daemon.uccp-machine-config-operator.svc<br/>- machine-config-daemon.uccp-machine-config-operator.svc.cluster.local |
| IP Addresses |  |


#### machine-config-daemon.uccp-machine-config-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-machine-config-operator | proxy-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### marketplace-operator-metrics.uccp-marketplace.svc
![PKI Graph](subcert-marketplace-operator-metrics.uccp-marketplace.svc9089605778427485628.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | marketplace-operator-metrics.uccp-marketplace.svc |
| SerialNumber | 9089605778427485628 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - marketplace-operator-metrics.uccp-marketplace.svc<br/>- marketplace-operator-metrics.uccp-marketplace.svc.cluster.local |
| IP Addresses |  |


#### marketplace-operator-metrics.uccp-marketplace.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-marketplace | marketplace-operator-metrics |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-apiserver-operator.svc
![PKI Graph](subcert-metrics.uccp-apiserver-operator.svc7400757488192498955.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-apiserver-operator.svc |
| SerialNumber | 7400757488192498955 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-apiserver-operator.svc<br/>- metrics.uccp-apiserver-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-apiserver-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-apiserver-operator | uccp-apiserver-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-authentication-operator.svc
![PKI Graph](subcert-metrics.uccp-authentication-operator.svc6301067254016768819.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-authentication-operator.svc |
| SerialNumber | 6301067254016768819 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-authentication-operator.svc<br/>- metrics.uccp-authentication-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-authentication-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-authentication-operator | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-cluster-samples-operator.svc
![PKI Graph](subcert-metrics.uccp-cluster-samples-operator.svc2686164050326297964.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-cluster-samples-operator.svc |
| SerialNumber | 2686164050326297964 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-cluster-samples-operator.svc<br/>- metrics.uccp-cluster-samples-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-cluster-samples-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-samples-operator | samples-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-config-operator.svc
![PKI Graph](subcert-metrics.uccp-config-operator.svc1282522901324235330.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-config-operator.svc |
| SerialNumber | 1282522901324235330 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-config-operator.svc<br/>- metrics.uccp-config-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-config-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-config-operator | config-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-console-operator.svc
![PKI Graph](subcert-metrics.uccp-console-operator.svc4841722672729242428.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-console-operator.svc |
| SerialNumber | 4841722672729242428 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-console-operator.svc<br/>- metrics.uccp-console-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-console-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-console-operator | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-controller-manager-operator.svc
![PKI Graph](subcert-metrics.uccp-controller-manager-operator.svc7895726074443218984.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-controller-manager-operator.svc |
| SerialNumber | 7895726074443218984 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-controller-manager-operator.svc<br/>- metrics.uccp-controller-manager-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-controller-manager-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-controller-manager-operator | uccp-controller-manager-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-dns-operator.svc
![PKI Graph](subcert-metrics.uccp-dns-operator.svc7601847597278589785.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-dns-operator.svc |
| SerialNumber | 7601847597278589785 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-dns-operator.svc<br/>- metrics.uccp-dns-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-dns-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-dns-operator | metrics-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-etcd-operator.svc
![PKI Graph](subcert-metrics.uccp-etcd-operator.svc3978805962409959490.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-etcd-operator.svc |
| SerialNumber | 3978805962409959490 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-etcd-operator.svc<br/>- metrics.uccp-etcd-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-etcd-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-etcd-operator | etcd-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-ingress-operator.svc
![PKI Graph](subcert-metrics.uccp-ingress-operator.svc8279326142192924063.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-ingress-operator.svc |
| SerialNumber | 8279326142192924063 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-ingress-operator.svc<br/>- metrics.uccp-ingress-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-ingress-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-ingress-operator | metrics-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-insights.svc
![PKI Graph](subcert-metrics.uccp-insights.svc6511383752257504790.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-insights.svc |
| SerialNumber | 6511383752257504790 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-insights.svc<br/>- metrics.uccp-insights.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-insights.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-insights | uccp-insights-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-kube-apiserver-operator.svc
![PKI Graph](subcert-metrics.uccp-kube-apiserver-operator.svc29949893468932305.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-kube-apiserver-operator.svc |
| SerialNumber | 29949893468932305 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-kube-apiserver-operator.svc<br/>- metrics.uccp-kube-apiserver-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-kube-apiserver-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-apiserver-operator | kube-apiserver-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-kube-controller-manager-operator.svc
![PKI Graph](subcert-metrics.uccp-kube-controller-manager-operator.svc1704350015219690809.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-kube-controller-manager-operator.svc |
| SerialNumber | 1704350015219690809 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-kube-controller-manager-operator.svc<br/>- metrics.uccp-kube-controller-manager-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-kube-controller-manager-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-controller-manager-operator | kube-controller-manager-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-kube-scheduler-operator.svc
![PKI Graph](subcert-metrics.uccp-kube-scheduler-operator.svc5130875085637374527.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-kube-scheduler-operator.svc |
| SerialNumber | 5130875085637374527 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-kube-scheduler-operator.svc<br/>- metrics.uccp-kube-scheduler-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-kube-scheduler-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-scheduler-operator | kube-scheduler-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-kube-storage-version-migrator-operator.svc
![PKI Graph](subcert-metrics.uccp-kube-storage-version-migrator-operator.svc5764379107559423336.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-kube-storage-version-migrator-operator.svc |
| SerialNumber | 5764379107559423336 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-kube-storage-version-migrator-operator.svc<br/>- metrics.uccp-kube-storage-version-migrator-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-kube-storage-version-migrator-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-storage-version-migrator-operator | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### metrics.uccp-service-ca-operator.svc
![PKI Graph](subcert-metrics.uccp-service-ca-operator.svc8465006101141555491.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | metrics.uccp-service-ca-operator.svc |
| SerialNumber | 8465006101141555491 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - metrics.uccp-service-ca-operator.svc<br/>- metrics.uccp-service-ca-operator.svc.cluster.local |
| IP Addresses |  |


#### metrics.uccp-service-ca-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-service-ca-operator | serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### multus-admission-controller.uccp-multus.svc
![PKI Graph](subcert-multus-admission-controller.uccp-multus.svc4660313081021989101.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | multus-admission-controller.uccp-multus.svc |
| SerialNumber | 4660313081021989101 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - multus-admission-controller.uccp-multus.svc<br/>- multus-admission-controller.uccp-multus.svc.cluster.local |
| IP Addresses |  |


#### multus-admission-controller.uccp-multus.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-multus | multus-admission-controller-secret |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### network-metrics-service.uccp-multus.svc
![PKI Graph](subcert-network-metrics-service.uccp-multus.svc1889672894063829328.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | network-metrics-service.uccp-multus.svc |
| SerialNumber | 1889672894063829328 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - network-metrics-service.uccp-multus.svc<br/>- network-metrics-service.uccp-multus.svc.cluster.local |
| IP Addresses |  |


#### network-metrics-service.uccp-multus.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-multus | metrics-daemon-secret |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### node-exporter.uccp-monitoring.svc
![PKI Graph](subcert-node-exporter.uccp-monitoring.svc6997256025626337803.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | node-exporter.uccp-monitoring.svc |
| SerialNumber | 6997256025626337803 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - node-exporter.uccp-monitoring.svc<br/>- node-exporter.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### node-exporter.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | node-exporter-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### node-tuning-operator.uccp-cluster-node-tuning-operator.svc
![PKI Graph](subcert-node-tuning-operator.uccp-cluster-node-tuning-operator.svc7832843275082956404.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | node-tuning-operator.uccp-cluster-node-tuning-operator.svc |
| SerialNumber | 7832843275082956404 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - node-tuning-operator.uccp-cluster-node-tuning-operator.svc<br/>- node-tuning-operator.uccp-cluster-node-tuning-operator.svc.cluster.local |
| IP Addresses |  |


#### node-tuning-operator.uccp-cluster-node-tuning-operator.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-cluster-node-tuning-operator | node-tuning-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### oauth-openshift.uccp-authentication.svc
![PKI Graph](subcert-oauth-openshift.uccp-authentication.svc455687257200358236.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | oauth-openshift.uccp-authentication.svc |
| SerialNumber | 455687257200358236 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - oauth-openshift.uccp-authentication.svc<br/>- oauth-openshift.uccp-authentication.svc.cluster.local |
| IP Addresses |  |


#### oauth-openshift.uccp-authentication.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-authentication | v4-0-config-system-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### olm-operator-metrics.uccp-operator-lifecycle-manager.svc
![PKI Graph](subcert-olm-operator-metrics.uccp-operator-lifecycle-manager.svc8800549647405875654.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | olm-operator-metrics.uccp-operator-lifecycle-manager.svc |
| SerialNumber | 8800549647405875654 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - olm-operator-metrics.uccp-operator-lifecycle-manager.svc<br/>- olm-operator-metrics.uccp-operator-lifecycle-manager.svc.cluster.local |
| IP Addresses |  |


#### olm-operator-metrics.uccp-operator-lifecycle-manager.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-operator-lifecycle-manager | olm-operator-serving-cert |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### uccp-state-metrics.uccp-monitoring.svc
![PKI Graph](subcert-uccp-state-metrics.uccp-monitoring.svc7882046295044958152.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | uccp-state-metrics.uccp-monitoring.svc |
| SerialNumber | 7882046295044958152 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - uccp-state-metrics.uccp-monitoring.svc<br/>- uccp-state-metrics.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### uccp-state-metrics.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | uccp-state-metrics-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### prometheus-adapter.uccp-monitoring.svc
![PKI Graph](subcert-prometheus-adapter.uccp-monitoring.svc2945501834265381842.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | prometheus-adapter.uccp-monitoring.svc |
| SerialNumber | 2945501834265381842 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - prometheus-adapter.uccp-monitoring.svc<br/>- prometheus-adapter.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### prometheus-adapter.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | prometheus-adapter-8tkqrsmu9afpe |
| uccp-monitoring | prometheus-adapter-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### prometheus-k8s-thanos-sidecar.uccp-monitoring.svc
![PKI Graph](subcert-prometheus-k8s-thanos-sidecar.uccp-monitoring.svc8595435866243050124.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | prometheus-k8s-thanos-sidecar.uccp-monitoring.svc |
| SerialNumber | 8595435866243050124 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - prometheus-k8s-thanos-sidecar.uccp-monitoring.svc<br/>- prometheus-k8s-thanos-sidecar.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### prometheus-k8s-thanos-sidecar.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | prometheus-k8s-thanos-sidecar-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### prometheus-k8s.uccp-monitoring.svc
![PKI Graph](subcert-prometheus-k8s.uccp-monitoring.svc8353222847585982884.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | prometheus-k8s.uccp-monitoring.svc |
| SerialNumber | 8353222847585982884 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - prometheus-k8s.uccp-monitoring.svc<br/>- prometheus-k8s.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### prometheus-k8s.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | prometheus-k8s-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### prometheus-operator.uccp-monitoring.svc
![PKI Graph](subcert-prometheus-operator.uccp-monitoring.svc4178324176483245369.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | prometheus-operator.uccp-monitoring.svc |
| SerialNumber | 4178324176483245369 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - prometheus-operator.uccp-monitoring.svc<br/>- prometheus-operator.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### prometheus-operator.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | prometheus-operator-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### router-internal-default.uccp-ingress.svc
![PKI Graph](subcert-router-internal-default.uccp-ingress.svc5035056584329763135.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | router-internal-default.uccp-ingress.svc |
| SerialNumber | 5035056584329763135 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - router-internal-default.uccp-ingress.svc<br/>- router-internal-default.uccp-ingress.svc.cluster.local |
| IP Addresses |  |


#### router-internal-default.uccp-ingress.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-ingress | router-metrics-certs-default |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### scheduler.uccp-kube-scheduler.svc
![PKI Graph](subcert-scheduler.uccp-kube-scheduler.svc6657891906279326247.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | scheduler.uccp-kube-scheduler.svc |
| SerialNumber | 6657891906279326247 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - scheduler.uccp-kube-scheduler.svc<br/>- scheduler.uccp-kube-scheduler.svc.cluster.local |
| IP Addresses |  |


#### scheduler.uccp-kube-scheduler.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-kube-scheduler | serving-cert |
| uccp-kube-scheduler | serving-cert-2 |
| uccp-kube-scheduler | serving-cert-3 |
| uccp-kube-scheduler | serving-cert-4 |
| uccp-kube-scheduler | serving-cert-5 |
| uccp-kube-scheduler | serving-cert-6 |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-3/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-3/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-5/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-5/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-6/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-6/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-2/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-2/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-4/secrets/serving-cert/tls.crt/tls.crt | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-scheduler-pod-4/secrets/serving-cert/tls.crt/tls.key | -rw-------. | root | root | system_u:object_r:kubernetes_file_t:s0 |



### sdn.uccp-sdn.svc
![PKI Graph](subcert-sdn.uccp-sdn.svc2625658655664643767.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | sdn.uccp-sdn.svc |
| SerialNumber | 2625658655664643767 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - sdn.uccp-sdn.svc<br/>- sdn.uccp-sdn.svc.cluster.local |
| IP Addresses |  |


#### sdn.uccp-sdn.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-sdn | sdn-metrics-certs |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### telemeter-client.uccp-monitoring.svc
![PKI Graph](subcert-telemeter-client.uccp-monitoring.svc1112524969442289252.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | telemeter-client.uccp-monitoring.svc |
| SerialNumber | 1112524969442289252 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - telemeter-client.uccp-monitoring.svc<br/>- telemeter-client.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### telemeter-client.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | telemeter-client-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |




### thanos-querier.uccp-monitoring.svc
![PKI Graph](subcert-thanos-querier.uccp-monitoring.svc2348697830836353775.png)



| Property | Value |
| ----------- | ----------- |
| Type | Serving |
| CommonName | thanos-querier.uccp-monitoring.svc |
| SerialNumber | 2348697830836353775 |
| Issuer CommonName | [service-serving-signer](#service-serving-signer) |
| Validity | 2y |
| Signature Algorithm | SHA256-RSA |
| PublicKey Algorithm | RSA 2048 bit |
| Usages | - KeyUsageDigitalSignature<br/>- KeyUsageKeyEncipherment |
| ExtendedUsages | - ExtKeyUsageServerAuth |
| DNS Names | - thanos-querier.uccp-monitoring.svc<br/>- thanos-querier.uccp-monitoring.svc.cluster.local |
| IP Addresses |  |


#### thanos-querier.uccp-monitoring.svc Locations
| Namespace | Secret Name |
| ----------- | ----------- |
| uccp-monitoring | thanos-querier-tls |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |



## Client Certificate/Key Pairs

## Certificates Without Keys

These certificates are present in certificate authority bundles, but do not have keys in the cluster.
This happens when the installer bootstrap clusters with a set of certificate/key pairs that are deleted during the
installation process.

## Certificate Authority Bundles


### service-ca
![PKI Graph](subca-3983882995.png)

CA for recognizing serving certificates for services that were signed by our service-ca controller.

**Bundled Certificates**

| CommonName | Issuer CommonName | Validity | PublicKey Algorithm |
| ----------- | ----------- | ----------- | ----------- |
| [service-serving-signer](#service-serving-signer) | [service-serving-signer](#service-serving-signer) | 2y60d | RSA 2048 bit |

#### service-ca Locations
| Namespace | ConfigMap Name |
| ----------- | ----------- |
| uccp-config-managed | service-ca |
| uccp-kube-controller-manager | service-ca |
| uccp-kube-controller-manager | service-ca-2 |
| uccp-kube-controller-manager | service-ca-3 |
| uccp-kube-controller-manager | service-ca-4 |
| uccp-kube-controller-manager | service-ca-5 |
| uccp-kube-controller-manager | service-ca-6 |
| uccp-kube-controller-manager | service-ca-7 |
| uccp-service-ca | signing-cabundle |

| File | Permissions | User | Group | SE Linux |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-3/configmaps/service-ca/ca-bundle.crt/ca-bundle.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-4/configmaps/service-ca/ca-bundle.crt/ca-bundle.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-5/configmaps/service-ca/ca-bundle.crt/ca-bundle.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-6/configmaps/service-ca/ca-bundle.crt/ca-bundle.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |
| /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-7/configmaps/service-ca/ca-bundle.crt/ca-bundle.crt | -rw-r--r--. | root | root | system_u:object_r:kubernetes_file_t:s0 |


