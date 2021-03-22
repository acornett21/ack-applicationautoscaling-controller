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

// Code generated by ack-generate. DO NOT EDIT.

package scheduled_action

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/applicationautoscaling-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApplicationAutoScaling{}
	_ = &svcapitypes.ScheduledAction{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeScheduledActionsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeScheduledActions", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.ScheduledActions {
		if elem.EndTime != nil {
			ko.Spec.EndTime = &metav1.Time{*elem.EndTime}
		}
		if elem.ResourceId != nil {
			ko.Spec.ResourceID = elem.ResourceId
		}
		if elem.ScalableDimension != nil {
			ko.Spec.ScalableDimension = elem.ScalableDimension
		}
		if elem.ScalableTargetAction != nil {
			f4 := &svcapitypes.ScalableTargetAction{}
			if elem.ScalableTargetAction.MaxCapacity != nil {
				f4.MaxCapacity = elem.ScalableTargetAction.MaxCapacity
			}
			if elem.ScalableTargetAction.MinCapacity != nil {
				f4.MinCapacity = elem.ScalableTargetAction.MinCapacity
			}
			ko.Spec.ScalableTargetAction = f4
		}
		if elem.Schedule != nil {
			ko.Spec.Schedule = elem.Schedule
		}
		if elem.ScheduledActionARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ScheduledActionARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.ScheduledActionName != nil {
			ko.Spec.ScheduledActionName = elem.ScheduledActionName
		}
		if elem.ServiceNamespace != nil {
			ko.Spec.ServiceNamespace = elem.ServiceNamespace
		}
		if elem.StartTime != nil {
			ko.Spec.StartTime = &metav1.Time{*elem.StartTime}
		}
		if elem.Timezone != nil {
			ko.Spec.Timezone = elem.Timezone
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeScheduledActionsInput, error) {
	res := &svcsdk.DescribeScheduledActionsInput{}

	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(ctx, r)
	if err != nil {
		return nil, err
	}

	_, respErr := rm.sdkapi.PutScheduledActionWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "PutScheduledAction", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PutScheduledActionInput, error) {
	res := &svcsdk.PutScheduledActionInput{}

	if r.ko.Spec.EndTime != nil {
		res.SetEndTime(r.ko.Spec.EndTime.Time)
	}
	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ScalableTargetAction != nil {
		f3 := &svcsdk.ScalableTargetAction{}
		if r.ko.Spec.ScalableTargetAction.MaxCapacity != nil {
			f3.SetMaxCapacity(*r.ko.Spec.ScalableTargetAction.MaxCapacity)
		}
		if r.ko.Spec.ScalableTargetAction.MinCapacity != nil {
			f3.SetMinCapacity(*r.ko.Spec.ScalableTargetAction.MinCapacity)
		}
		res.SetScalableTargetAction(f3)
	}
	if r.ko.Spec.Schedule != nil {
		res.SetSchedule(*r.ko.Spec.Schedule)
	}
	if r.ko.Spec.ScheduledActionName != nil {
		res.SetScheduledActionName(*r.ko.Spec.ScheduledActionName)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
	}
	if r.ko.Spec.StartTime != nil {
		res.SetStartTime(r.ko.Spec.StartTime.Time)
	}
	if r.ko.Spec.Timezone != nil {
		res.SetTimezone(*r.ko.Spec.Timezone)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {

	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteScheduledActionWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteScheduledAction", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteScheduledActionInput, error) {
	res := &svcsdk.DeleteScheduledActionInput{}

	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ScheduledActionName != nil {
		res.SetScheduledActionName(*r.ko.Spec.ScheduledActionName)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ScheduledAction,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
