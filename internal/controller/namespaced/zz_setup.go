// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	binding "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/binding/binding"
	exchange "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/exchange/exchange"
	upstream "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/federationupstream/upstream"
	policy "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/operatorpolicy/policy"
	permissions "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/permissions/permissions"
	policypolicy "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/policy/policy"
	providerconfig "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/providerconfig"
	queue "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/queue/queue"
	shovel "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/shovel/shovel"
	permissionstopicpermissions "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/topicpermissions/permissions"
	user "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/user/user"
	vhost "github.com/Volevanius/provider-rabbitmqtf/internal/controller/namespaced/vhost/vhost"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		binding.Setup,
		exchange.Setup,
		upstream.Setup,
		policy.Setup,
		permissions.Setup,
		policypolicy.Setup,
		providerconfig.Setup,
		queue.Setup,
		shovel.Setup,
		permissionstopicpermissions.Setup,
		user.Setup,
		vhost.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		binding.SetupGated,
		exchange.SetupGated,
		upstream.SetupGated,
		policy.SetupGated,
		permissions.SetupGated,
		policypolicy.SetupGated,
		providerconfig.SetupGated,
		queue.SetupGated,
		shovel.SetupGated,
		permissionstopicpermissions.SetupGated,
		user.SetupGated,
		vhost.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
