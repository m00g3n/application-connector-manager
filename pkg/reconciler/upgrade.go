package reconciler

import (
	"context"

	"github.com/kyma-project/application-connector-manager/pkg/unstructured"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	_ "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var (
	keyCompassAgentCfg = client.ObjectKey{
		Namespace: "kyma-system",
		Name:      "compass-agent-configuration",
	}
)

func sFnPreUpdate(ctx context.Context, r *fsm, _ *systemState) (stateFn, *ctrl.Result, error) {
	var secret v1.Secret
	err := r.Get(ctx, keyCompassAgentCfg, &secret)

	if errors.IsNotFound(err) {
		return switchState(sFnCompassRtAgentScaleToZero)
	}

	if err != nil {
		return stopWithErrorAndNoRequeue(err)
	}

	return switchState(sFnUpdate)
}

var compassRtAgentPredicate unstructured.Predicate = func(u unstructured.Unstructured) bool {
	gvk := u.GetObjectKind().GroupVersionKind()
	return gvk.Kind == "Deployment" && gvk.Group == "apps" && gvk.Version == "v1" && u.GetName() == "compass-runtime-agent"
}

func sFnCompassRtAgentScaleToZero(_ context.Context, r *fsm, _ *systemState) (stateFn, *ctrl.Result, error) {
	u, err := compassRtAgentPredicate.First(r.Objs)
	if err != nil {
		return stopWithErrorAndNoRequeue(err)
	}

	if err := unstructured.Update(u, 0, updateDeploymentScaling); err != nil {
		return stopWithErrorAndNoRequeue(err)
	}

	return switchState(sFnUpdate)
}

func updateDeploymentScaling(d appv1.Deployment, replicas int32) error {
	d.Spec.Replicas = ptr.To(replicas)
	return nil
}