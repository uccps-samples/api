package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configv1 "github.com/uccps-samples/api/config/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Console provides a means to configure an operator to manage the console.
type Console struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// +required
	Spec ConsoleSpec `json:"spec,omitempty"`
	// +optional
	Status ConsoleStatus `json:"status,omitempty"`
}

// ConsoleSpec is the specification of the desired behavior of the Console.
type ConsoleSpec struct {
	OperatorSpec `json:",inline"`
	// customization is used to optionally provide a small set of
	// customization options to the web console.
	// +optional
	Customization ConsoleCustomization `json:"customization"`
	// providers contains configuration for using specific service providers.
	Providers ConsoleProviders `json:"providers"`
	// route contains hostname and secret reference that contains the serving certificate.
	// If a custom route is specified, a new route will be created with the
	// provided hostname, under which console will be available.
	// In case of custom hostname uses the default routing suffix of the cluster,
	// the Secret specification for a serving certificate will not be needed.
	// In case of custom hostname points to an arbitrary domain, manual DNS configurations steps are necessary.
	// The default console route will be maintained to reserve the default hostname
	// for console if the custom route is removed.
	// If not specified, default route will be used.
	// +optional
	Route ConsoleConfigRoute `json:"route"`
}

// ConsoleConfigRoute holds information on external route access to console.
type ConsoleConfigRoute struct {
	// hostname is the desired custom domain under which console will be available.
	Hostname string `json:"hostname"`
	// secret points to secret in the uccp-config namespace that contains custom
	// certificate and key and needs to be created manually by the cluster admin.
	// Referenced Secret is required to contain following key value pairs:
	// - "tls.crt" - to specifies custom certificate
	// - "tls.key" - to specifies private key of the custom certificate
	// If the custom hostname uses the default routing suffix of the cluster,
	// the Secret specification for a serving certificate will not be needed.
	// +optional
	Secret configv1.SecretNameReference `json:"secret"`
}

// ConsoleStatus defines the observed status of the Console.
type ConsoleStatus struct {
	OperatorStatus `json:",inline"`
}

// ConsoleProviders defines a list of optional additional providers of
// functionality to the console.
type ConsoleProviders struct {
	// statuspage contains ID for statuspage.io page that provides status info about.
	// +optional
	Statuspage *StatuspageProvider `json:"statuspage,omitempty"`
}

// StatuspageProvider provides identity for statuspage account.
type StatuspageProvider struct {
	// pageID is the unique ID assigned by Statuspage for your page. This must be a public page.
	PageID string `json:"pageID"`
}

// ConsoleCustomization defines a list of optional configuration for the console UI.
type ConsoleCustomization struct {
	// brand is the default branding of the web console which can be overridden by
	// providing the brand field.  There is a limited set of specific brand options.
	// This field controls elements of the console such as the logo.
	// Invalid value will prevent a console rollout.
	Brand Brand `json:"brand,omitempty"`
	// documentationBaseURL links to external documentation are shown in various sections
	// of the web console.  Providing documentationBaseURL will override the default
	// documentation URL.
	// Invalid value will prevent a console rollout.
	// +kubebuilder:validation:Pattern=`^$|^((https):\/\/?)[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/?))\/$`
	DocumentationBaseURL string `json:"documentationBaseURL,omitempty"`
	// customProductName is the name that will be displayed in page titles, logo alt text, and the about dialog
	// instead of the normal OpenShift product name.
	// +optional
	CustomProductName string `json:"customProductName,omitempty"`
	// customLogoFile replaces the default OpenShift logo in the masthead and about dialog. It is a reference to a
	// ConfigMap in the uccp-config namespace. This can be created with a command like
	// 'oc create configmap custom-logo --from-file=/path/to/file -n uccp-config'.
	// Image size must be less than 1 MB due to constraints on the ConfigMap size.
	// The ConfigMap key should include a file extension so that the console serves the file
	// with the correct MIME type.
	// Recommended logo specifications:
	// Dimensions: Max height of 68px and max width of 200px
	// SVG format preferred
	// +optional
	CustomLogoFile configv1.ConfigMapFileReference `json:"customLogoFile,omitempty"`
}

// Brand is a specific supported brand within the console.
// +kubebuilder:validation:Pattern=`^$|^(ocp|origin|okd|dedicated|online|azure)$`
type Brand string

const (
	// Branding for OpenShift
	BrandOpenShift Brand = "uccp"
	// Branding for The Origin Community Distribution of Kubernetes
	BrandOKD Brand = "okd"
	// Branding for OpenShift Online
	BrandOnline Brand = "online"
	// Branding for OpenShift Container Platform
	BrandOCP Brand = "ocp"
	// Branding for OpenShift Dedicated
	BrandDedicated Brand = "dedicated"
	// Branding for Azure Red Hat OpenShift
	BrandAzure Brand = "azure"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConsoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Console `json:"items"`
}
