package reconciler

import (
	"context"
	"fmt"
	"time"

	"github.com/kyma-project/application-connector-manager/api/v1alpha1"
	modtest "github.com/kyma-project/application-connector-manager/pkg/reconciler/testing"
	"github.com/kyma-project/application-connector-manager/pkg/unstructured"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("ACM sFnUpdate", func() {

	var testData map[string][]unstructured.Unstructured
	updateTimeout := time.Second * 5

	defaultState := &systemState{
		instance: v1alpha1.ApplicationConnector{
			Spec: v1alpha1.ApplicationConnectorSpec{
				RuntimeAgentSpec: v1alpha1.RuntimeAgentSpec{
					ControllerSyncPeriod: metav1.Duration{Duration: time.Second * 10},
				},
			},
		},
	}

	testData, err := modtest.LoadTestData(modtest.SfnUpdate)
	Expect(err).Should(BeNil(), fmt.Errorf("unable to extract test data: %s", err))

	ctx, cancel := context.WithTimeout(context.Background(), updateTimeout)
	defer cancel()

	DescribeTable(
		"update state function",
		testUpdate,
		Entry(
			"happy path",
			ctx,
			&fsm{Cfg: Cfg{Objs: testData[modtest.TdUpdateAcmValid]}},
			defaultState,
			testUpdateOptions{
				MatchExpectedErr: BeNil(),
				MatchNextFnState: equalStateFunction(sFnApply),
			},
		),
		Entry(
			"no deployment",
			ctx,
			&fsm{},
			defaultState,
			testUpdateOptions{
				MatchExpectedErr: BeNil(),
				MatchNextFnState: equalStateFunction(sFnUpdateStatus(nil, nil)),
			},
		),
	)
})

type testUpdateOptions struct {
	MatchExpectedErr types.GomegaMatcher
	MatchNextFnState types.GomegaMatcher
}

func testUpdate(ctx context.Context, r *fsm, s *systemState, ops testUpdateOptions) {
	sFn, _, err := sFnUpdate(ctx, r, s)
	Expect(err).To(ops.MatchExpectedErr)
	Expect(sFn).To(ops.MatchNextFnState)
}