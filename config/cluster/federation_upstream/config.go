package federationupstream

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_federation_upstream", func(r *config.Resource) {
		r.ShortGroup = "federationupstream"
		r.References["vhost"] = config.Reference{
			Type: "github.com/Volevanius/provider-rabbitmqtf/apis/cluster/vhost/v1alpha1.Vhost",
		}
	})
}
