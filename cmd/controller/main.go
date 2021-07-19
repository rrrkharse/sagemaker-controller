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

package main

import (
	"os"

	ackcfg "github.com/aws-controllers-k8s/runtime/pkg/config"
	ackrt "github.com/aws-controllers-k8s/runtime/pkg/runtime"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrlrt "sigs.k8s.io/controller-runtime"
	ctrlrtmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	svctypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svcresource "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource"

	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/data_quality_job_definition"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/endpoint"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/endpoint_config"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/feature_group"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/hyper_parameter_tuning_job"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/model"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/model_bias_job_definition"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/model_explainability_job_definition"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/model_package_group"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/model_quality_job_definition"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/monitoring_schedule"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/processing_job"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/training_job"
	_ "github.com/aws-controllers-k8s/sagemaker-controller/pkg/resource/transform_job"
)

var (
	awsServiceAPIGroup = "sagemaker.services.k8s.aws"
	awsServiceAlias    = "sagemaker"
	scheme             = runtime.NewScheme()
	setupLog           = ctrlrt.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = svctypes.AddToScheme(scheme)
	_ = ackv1alpha1.AddToScheme(scheme)
}

func main() {
	var ackCfg ackcfg.Config
	ackCfg.BindFlags()
	flag.Parse()
	ackCfg.SetupLogger()

	if err := ackCfg.Validate(); err != nil {
		setupLog.Error(
			err, "Unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	mgr, err := ctrlrt.NewManager(ctrlrt.GetConfigOrDie(), ctrlrt.Options{
		Scheme:             scheme,
		Port:               ackCfg.BindPort,
		MetricsBindAddress: ackCfg.MetricsAddr,
		LeaderElection:     ackCfg.EnableLeaderElection,
		LeaderElectionID:   awsServiceAPIGroup,
		Namespace:          ackCfg.WatchNamespace,
	})
	if err != nil {
		setupLog.Error(
			err, "unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	stopChan := ctrlrt.SetupSignalHandler()

	setupLog.Info(
		"initializing service controller",
		"aws.service", awsServiceAlias,
	)
	sc := ackrt.NewServiceController(
		awsServiceAlias, awsServiceAPIGroup,
		ackrt.VersionInfo{}, // TODO: populate version info
	).WithLogger(
		ctrlrt.Log,
	).WithResourceManagerFactories(
		svcresource.GetManagerFactories(),
	).WithPrometheusRegistry(
		ctrlrtmetrics.Registry,
	)
	if err = sc.BindControllerManager(mgr, ackCfg); err != nil {
		setupLog.Error(
			err, "unable bind to controller manager to service controller",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	setupLog.Info(
		"starting manager",
		"aws.service", awsServiceAlias,
	)
	if err := mgr.Start(stopChan); err != nil {
		setupLog.Error(
			err, "unable to start controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}
}
