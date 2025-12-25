package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	// Cluster-scoped resource configurations
	bindingCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/binding"
	exchangeCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/exchange"
	federationupstreamCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/federation_upstream"
	operatorpolicyCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/operator_policy"
	permissionsCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/permissions"
	policyCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/policy"
	queueCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/queue"
	shovelCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/shovel"
	topicpermissionsCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/topic_permissions"
	userCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/user"
	vhostCluster "github.com/Volevanius/provider-rabbitmqtf/config/cluster/vhost"

	// Namespaced resource configurations
	bindingNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/binding"
	exchangeNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/exchange"
	federationupstreamNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/federation_upstream"
	operatorpolicyNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/operator_policy"
	permissionsNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/permissions"
	policyNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/policy"
	queueNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/queue"
	shovelNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/shovel"
	topicpermissionsNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/topic_permissions"
	userNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/user"
	vhostNamespaced "github.com/Volevanius/provider-rabbitmqtf/config/namespaced/vhost"
)

const (
	resourcePrefix = "rabbitmqtf"
	modulePath     = "github.com/Volevanius/provider-rabbitmqtf"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rabbitmqtf.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// RabbitMQ resource configurations
		vhostCluster.Configure,
		userCluster.Configure,
		permissionsCluster.Configure,
		topicpermissionsCluster.Configure,
		exchangeCluster.Configure,
		queueCluster.Configure,
		bindingCluster.Configure,
		policyCluster.Configure,
		operatorpolicyCluster.Configure,
		federationupstreamCluster.Configure,
		shovelCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rabbitmqtf.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// RabbitMQ resource configurations
		vhostNamespaced.Configure,
		userNamespaced.Configure,
		permissionsNamespaced.Configure,
		topicpermissionsNamespaced.Configure,
		exchangeNamespaced.Configure,
		queueNamespaced.Configure,
		bindingNamespaced.Configure,
		policyNamespaced.Configure,
		operatorpolicyNamespaced.Configure,
		federationupstreamNamespaced.Configure,
		shovelNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
