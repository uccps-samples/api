package v1

import corev1 "k8s.io/api/core/v1"

const (
	// ManagedByOpenShiftAnnotation indicates that an image is managed by OpenShift's registry.
	ManagedByOpenShiftAnnotation = "uccp.io/image.managed"

	// DockerImageRepositoryCheckAnnotation indicates that OpenShift has
	// attempted to import tag and image information from an external Docker
	// image repository.
	DockerImageRepositoryCheckAnnotation = "uccp.io/image.dockerRepositoryCheck"

	// InsecureRepositoryAnnotation may be set true on an image stream to allow insecure access to pull content.
	InsecureRepositoryAnnotation = "uccp.io/image.insecureRepository"

	// ExcludeImageSecretAnnotation indicates that a secret should not be returned by imagestream/secrets.
	ExcludeImageSecretAnnotation = "uccp.io/image.excludeSecret"

	// DockerImageLayersOrderAnnotation describes layers order in the docker image.
	DockerImageLayersOrderAnnotation = "image.uccp.io/dockerLayersOrder"

	// DockerImageLayersOrderAscending indicates that image layers are sorted in
	// the order of their addition (from oldest to latest)
	DockerImageLayersOrderAscending = "ascending"

	// DockerImageLayersOrderDescending indicates that layers are sorted in
	// reversed order of their addition (from newest to oldest).
	DockerImageLayersOrderDescending = "descending"

	// ImporterPreferArchAnnotation represents an architecture that should be
	// selected if an image uses a manifest list and it should be
	// downconverted.
	ImporterPreferArchAnnotation = "importer.image.uccp.io/prefer-arch"

	// ImporterPreferOSAnnotation represents an operation system that should
	// be selected if an image uses a manifest list and it should be
	// downconverted.
	ImporterPreferOSAnnotation = "importer.image.uccp.io/prefer-os"

	// ImageManifestBlobStoredAnnotation indicates that manifest and config blobs of image are stored in on
	// storage of integrated Docker registry.
	ImageManifestBlobStoredAnnotation = "image.uccp.io/manifestBlobStored"

	// DefaultImageTag is used when an image tag is needed and the configuration does not specify a tag to use.
	DefaultImageTag = "latest"

	// ResourceImageStreams represents a number of image streams in a project.
	ResourceImageStreams corev1.ResourceName = "uccp.io/imagestreams"

	// ResourceImageStreamImages represents a number of unique references to images in all image stream
	// statuses of a project.
	ResourceImageStreamImages corev1.ResourceName = "uccp.io/images"

	// ResourceImageStreamTags represents a number of unique references to images in all image stream specs
	// of a project.
	ResourceImageStreamTags corev1.ResourceName = "uccp.io/image-tags"

	// Limit that applies to images. Used with a max["storage"] LimitRangeItem to set
	// the maximum size of an image.
	LimitTypeImage corev1.LimitType = "uccp.io/Image"

	// Limit that applies to image streams. Used with a max[resource] LimitRangeItem to set the maximum number
	// of resource. Where the resource is one of "uccp.io/images" and "uccp.io/image-tags".
	LimitTypeImageStream corev1.LimitType = "uccp.io/ImageStream"

	// The supported type of image signature.
	ImageSignatureTypeAtomicImageV1 string = "AtomicImageV1"
)
