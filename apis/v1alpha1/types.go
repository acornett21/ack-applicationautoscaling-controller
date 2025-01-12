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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Represents a CloudWatch alarm associated with a scaling policy.
type Alarm struct {
	AlarmARN  *string `json:"alarmARN,omitempty"`
	AlarmName *string `json:"alarmName,omitempty"`
}

// Represents a CloudWatch metric of your choosing for a target tracking scaling
// policy to use with Application Auto Scaling.
//
// For information about the available metrics for a service, see Amazon Web
// Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
// in the Amazon CloudWatch User Guide.
//
// To create your customized metric specification:
//
//    * Add values for each required parameter from CloudWatch. You can use
//    an existing metric, or a new metric that you create. To use your own metric,
//    you must first publish the metric to CloudWatch. For more information,
//    see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
//    in the Amazon CloudWatch User Guide.
//
//    * Choose a metric that changes proportionally with capacity. The value
//    of the metric should increase or decrease in inverse proportion to the
//    number of capacity units. That is, the value of the metric should decrease
//    when capacity increases, and increase when capacity decreases.
//
// For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).
type CustomizedMetricSpecification struct {
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`
	MetricName *string            `json:"metricName,omitempty"`
	Namespace  *string            `json:"namespace,omitempty"`
	Statistic  *string            `json:"statistic,omitempty"`
	Unit       *string            `json:"unit,omitempty"`
}

// Describes the dimension names and values associated with a metric.
type MetricDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Represents a predefined metric for a target tracking scaling policy to use
// with Application Auto Scaling.
//
// Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch.
// To determine whether a desired metric already exists by looking up its namespace
// and dimension using the CloudWatch metrics dashboard in the console, follow
// the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html)
// in the Application Auto Scaling User Guide.
type PredefinedMetricSpecification struct {
	PredefinedMetricType *string `json:"predefinedMetricType,omitempty"`
	ResourceLabel        *string `json:"resourceLabel,omitempty"`
}

// Represents the minimum and maximum capacity for a scheduled action.
type ScalableTargetAction struct {
	MaxCapacity *int64 `json:"maxCapacity,omitempty"`
	MinCapacity *int64 `json:"minCapacity,omitempty"`
}

// Represents a scalable target.
type ScalableTarget_SDK struct {
	CreationTime      *metav1.Time `json:"creationTime,omitempty"`
	MaxCapacity       *int64       `json:"maxCapacity,omitempty"`
	MinCapacity       *int64       `json:"minCapacity,omitempty"`
	ResourceID        *string      `json:"resourceID,omitempty"`
	RoleARN           *string      `json:"roleARN,omitempty"`
	ScalableDimension *string      `json:"scalableDimension,omitempty"`
	ServiceNamespace  *string      `json:"serviceNamespace,omitempty"`
	// Specifies whether the scaling activities for a scalable target are in a suspended
	// state.
	SuspendedState *SuspendedState `json:"suspendedState,omitempty"`
}

// Represents a scaling activity.
type ScalingActivity struct {
	ActivityID        *string      `json:"activityID,omitempty"`
	Cause             *string      `json:"cause,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Details           *string      `json:"details,omitempty"`
	EndTime           *metav1.Time `json:"endTime,omitempty"`
	ResourceID        *string      `json:"resourceID,omitempty"`
	ScalableDimension *string      `json:"scalableDimension,omitempty"`
	ServiceNamespace  *string      `json:"serviceNamespace,omitempty"`
	StartTime         *metav1.Time `json:"startTime,omitempty"`
	StatusMessage     *string      `json:"statusMessage,omitempty"`
}

// Represents a scaling policy to use with Application Auto Scaling.
//
// For more information about configuring scaling policies for a specific service,
// see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html)
// in the Application Auto Scaling User Guide.
type ScalingPolicy_SDK struct {
	Alarms            []*Alarm     `json:"alarms,omitempty"`
	CreationTime      *metav1.Time `json:"creationTime,omitempty"`
	PolicyARN         *string      `json:"policyARN,omitempty"`
	PolicyName        *string      `json:"policyName,omitempty"`
	PolicyType        *string      `json:"policyType,omitempty"`
	ResourceID        *string      `json:"resourceID,omitempty"`
	ScalableDimension *string      `json:"scalableDimension,omitempty"`
	ServiceNamespace  *string      `json:"serviceNamespace,omitempty"`
	// Represents a step scaling policy configuration to use with Application Auto
	// Scaling.
	StepScalingPolicyConfiguration *StepScalingPolicyConfiguration `json:"stepScalingPolicyConfiguration,omitempty"`
	// Represents a target tracking scaling policy configuration to use with Application
	// Auto Scaling.
	TargetTrackingScalingPolicyConfiguration *TargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration,omitempty"`
}

// Represents a scheduled action.
type ScheduledAction struct {
	CreationTime       *metav1.Time `json:"creationTime,omitempty"`
	EndTime            *metav1.Time `json:"endTime,omitempty"`
	ResourceID         *string      `json:"resourceID,omitempty"`
	ScalableDimension  *string      `json:"scalableDimension,omitempty"`
	Schedule           *string      `json:"schedule,omitempty"`
	ScheduledActionARN *string      `json:"scheduledActionARN,omitempty"`
	ServiceNamespace   *string      `json:"serviceNamespace,omitempty"`
	StartTime          *metav1.Time `json:"startTime,omitempty"`
	Timezone           *string      `json:"timezone,omitempty"`
}

// Represents a step adjustment for a StepScalingPolicyConfiguration (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html).
// Describes an adjustment based on the difference between the value of the
// aggregated CloudWatch metric and the breach threshold that you've defined
// for the alarm.
//
// For the following examples, suppose that you have an alarm with a breach
// threshold of 50:
//
//    * To trigger the adjustment when the metric is greater than or equal to
//    50 and less than 60, specify a lower bound of 0 and an upper bound of
//    10.
//
//    * To trigger the adjustment when the metric is greater than 40 and less
//    than or equal to 50, specify a lower bound of -10 and an upper bound of
//    0.
//
// There are a few rules for the step adjustments for your step policy:
//
//    * The ranges of your step adjustments can't overlap or have a gap.
//
//    * At most one step adjustment can have a null lower bound. If one step
//    adjustment has a negative lower bound, then there must be a step adjustment
//    with a null lower bound.
//
//    * At most one step adjustment can have a null upper bound. If one step
//    adjustment has a positive upper bound, then there must be a step adjustment
//    with a null upper bound.
//
//    * The upper and lower bound can't be null in the same step adjustment.
type StepAdjustment struct {
	MetricIntervalLowerBound *float64 `json:"metricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound *float64 `json:"metricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int64   `json:"scalingAdjustment,omitempty"`
}

// Represents a step scaling policy configuration to use with Application Auto
// Scaling.
type StepScalingPolicyConfiguration struct {
	AdjustmentType         *string           `json:"adjustmentType,omitempty"`
	Cooldown               *int64            `json:"cooldown,omitempty"`
	MetricAggregationType  *string           `json:"metricAggregationType,omitempty"`
	MinAdjustmentMagnitude *int64            `json:"minAdjustmentMagnitude,omitempty"`
	StepAdjustments        []*StepAdjustment `json:"stepAdjustments,omitempty"`
}

// Specifies whether the scaling activities for a scalable target are in a suspended
// state.
type SuspendedState struct {
	DynamicScalingInSuspended  *bool `json:"dynamicScalingInSuspended,omitempty"`
	DynamicScalingOutSuspended *bool `json:"dynamicScalingOutSuspended,omitempty"`
	ScheduledScalingSuspended  *bool `json:"scheduledScalingSuspended,omitempty"`
}

// Represents a target tracking scaling policy configuration to use with Application
// Auto Scaling.
type TargetTrackingScalingPolicyConfiguration struct {
	// Represents a CloudWatch metric of your choosing for a target tracking scaling
	// policy to use with Application Auto Scaling.
	//
	// For information about the available metrics for a service, see Amazon Web
	// Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	// in the Amazon CloudWatch User Guide.
	//
	// To create your customized metric specification:
	//
	//    * Add values for each required parameter from CloudWatch. You can use
	//    an existing metric, or a new metric that you create. To use your own metric,
	//    you must first publish the metric to CloudWatch. For more information,
	//    see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
	//    in the Amazon CloudWatch User Guide.
	//
	//    * Choose a metric that changes proportionally with capacity. The value
	//    of the metric should increase or decrease in inverse proportion to the
	//    number of capacity units. That is, the value of the metric should decrease
	//    when capacity increases, and increase when capacity decreases.
	//
	// For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).
	CustomizedMetricSpecification *CustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty"`
	DisableScaleIn                *bool                          `json:"disableScaleIn,omitempty"`
	// Represents a predefined metric for a target tracking scaling policy to use
	// with Application Auto Scaling.
	//
	// Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch.
	// To determine whether a desired metric already exists by looking up its namespace
	// and dimension using the CloudWatch metrics dashboard in the console, follow
	// the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html)
	// in the Application Auto Scaling User Guide.
	PredefinedMetricSpecification *PredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty"`
	ScaleInCooldown               *int64                         `json:"scaleInCooldown,omitempty"`
	ScaleOutCooldown              *int64                         `json:"scaleOutCooldown,omitempty"`
	TargetValue                   *float64                       `json:"targetValue,omitempty"`
}
