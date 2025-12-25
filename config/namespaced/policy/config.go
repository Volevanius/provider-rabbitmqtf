package policy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_policy", func(r *config.Resource) {
		r.ShortGroup = "policy"
		r.References["vhost"] = config.Reference{
			Type: "github.com/Volevanius/provider-rabbitmqtf/apis/namespaced/vhost/v1alpha1.Vhost",
		}
	})
}
