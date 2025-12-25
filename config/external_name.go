package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// rabbitmq_vhost: import using name
	"rabbitmq_vhost": config.NameAsIdentifier,
	// rabbitmq_user: import using name
	"rabbitmq_user": config.NameAsIdentifier,
	// rabbitmq_permissions: import using user@vhost
	"rabbitmq_permissions": config.TemplatedStringAsIdentifier("user", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_topic_permissions: import using user@vhost
	"rabbitmq_topic_permissions": config.TemplatedStringAsIdentifier("user", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_exchange: import using name@vhost
	"rabbitmq_exchange": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_queue: import using name@vhost
	"rabbitmq_queue": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_binding: import using vhost/source/destination/destination_type/properties_key
	// properties_key is computed, so we use IdentifierFromProvider
	"rabbitmq_binding": identifierFromProvider(),
	// rabbitmq_policy: import using name@vhost
	"rabbitmq_policy": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_operator_policy: import using name@vhost
	"rabbitmq_operator_policy": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_federation_upstream: import using name@vhost
	"rabbitmq_federation_upstream": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
	// rabbitmq_shovel: import using name@vhost
	"rabbitmq_shovel": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}@{{ .parameters.vhost }}"),
}

func identifierFromProvider() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
