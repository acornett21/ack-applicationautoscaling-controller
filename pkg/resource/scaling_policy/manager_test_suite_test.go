// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package scaling_policy

import (
	"errors"
	"fmt"

	"path/filepath"
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/applicationautoscaling-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/applicationautoscaling-controller/pkg/testutil"
	mocksvcsdkapi "github.com/aws-controllers-k8s/applicationautoscaling-controller/test/mocks/aws-sdk-go/applicationautoscaling"
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"
	svcsdk "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/zap/zapcore"
	ctrlrtzap "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// provideResourceManagerWithMockSDKAPI accepts MockApplicationAutoScalingAPI and returns pointer to resourceManager
// the returned resourceManager is configured to use mockapi api.
func provideResourceManagerWithMockSDKAPI(mockApplicationAutoScalingAPI *mocksvcsdkapi.ApplicationAutoScalingAPI) *resourceManager {
	zapOptions := ctrlrtzap.Options{
		Development: true,
		Level:       zapcore.InfoLevel,
	}
	fakeLogger := ctrlrtzap.New(ctrlrtzap.UseFlagOptions(&zapOptions))
	return &resourceManager{
		rr:           nil,
		awsAccountID: "",
		awsRegion:    "",
		sess:         nil,
		sdkapi:       mockApplicationAutoScalingAPI,
		log:          fakeLogger,
		metrics:      ackmetrics.NewMetrics("applicationautoscaling"),
	}
}

// TestScalingPolicyTestSuite runs the test suite for scaling policy
func TestScalingPolicyTestSuite(t *testing.T) {
	var ts = testutil.TestSuite{}
	testutil.LoadFromFixture(filepath.Join("testdata", "test_suite.yaml"), &ts)
	var delegate = testRunnerDelegate{t: t}
	var runner = testutil.TestSuiteRunner{TestSuite: &ts, Delegate: &delegate}
	runner.RunTests()
}

// testRunnerDelegate implements testutil.TestRunnerDelegate
type testRunnerDelegate struct {
	t *testing.T
}

func (d *testRunnerDelegate) ResourceDescriptor() acktypes.AWSResourceDescriptor {
	return &resourceDescriptor{}
}

func (d *testRunnerDelegate) ResourceManager(mocksdkapi *mocksvcsdkapi.ApplicationAutoScalingAPI) acktypes.AWSResourceManager {
	return provideResourceManagerWithMockSDKAPI(mocksdkapi)
}

func (d *testRunnerDelegate) GoTestRunner() *testing.T {
	return d.t
}

func (d *testRunnerDelegate) EmptyServiceAPIOutput(apiName string) (interface{}, error) {
	if apiName == "" {
		return nil, errors.New("no API name specified")
	}
	//TODO: use reflection, template to auto generate this block/method.
	switch apiName {
	case "DescribeScalingPoliciesWithContext":
		var output svcsdk.DescribeScalingPoliciesOutput
		return &output, nil
	case "PutScalingPolicyWithContext":
		var output svcsdk.PutScalingPolicyOutput
		return &output, nil
	case "DeleteScalingPolicyWithContext":
		var output svcsdk.DeleteScalingPolicyOutput
		return &output, nil
	}
	return nil, errors.New(fmt.Sprintf("no matching API name found for: %s", apiName))
}

func (d *testRunnerDelegate) Equal(a acktypes.AWSResource, b acktypes.AWSResource) bool {
	ac := a.(*resource)
	bc := b.(*resource)
	// Ignore LastTransitionTime since it gets updated each run.
	opts := []cmp.Option{cmpopts.EquateEmpty(), cmpopts.IgnoreFields(ackv1alpha1.Condition{}, "LastTransitionTime"),
		cmpopts.IgnoreFields(svcapitypes.ScalingPolicyStatus{}, "CreationTime"),
		cmpopts.IgnoreFields(svcapitypes.ScalingPolicyStatus{}, "LastModifiedTime")}

	var specMatch = false
	if cmp.Equal(ac.ko.Spec, bc.ko.Spec, opts...) {
		specMatch = true
	} else {
		fmt.Printf("Difference ko.Spec (-expected +actual):\n\n")
		fmt.Println(cmp.Diff(ac.ko.Spec, bc.ko.Spec, opts...))
		specMatch = false
	}

	var statusMatch = false
	if cmp.Equal(ac.ko.Status, bc.ko.Status, opts...) {
		statusMatch = true
	} else {
		fmt.Printf("Difference ko.Status (-expected +actual):\n\n")
		fmt.Println(cmp.Diff(ac.ko.Status, bc.ko.Status, opts...))
		statusMatch = false
	}

	return statusMatch && specMatch
}
