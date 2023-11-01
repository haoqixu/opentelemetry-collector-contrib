// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sContainerCPULimit:                MetricConfig{Enabled: true},
					K8sContainerCPURequest:              MetricConfig{Enabled: true},
					K8sContainerEphemeralstorageLimit:   MetricConfig{Enabled: true},
					K8sContainerEphemeralstorageRequest: MetricConfig{Enabled: true},
					K8sContainerMemoryLimit:             MetricConfig{Enabled: true},
					K8sContainerMemoryRequest:           MetricConfig{Enabled: true},
					K8sContainerReady:                   MetricConfig{Enabled: true},
					K8sContainerRestarts:                MetricConfig{Enabled: true},
					K8sContainerStorageLimit:            MetricConfig{Enabled: true},
					K8sContainerStorageRequest:          MetricConfig{Enabled: true},
					K8sCronjobActiveJobs:                MetricConfig{Enabled: true},
					K8sDaemonsetCurrentScheduledNodes:   MetricConfig{Enabled: true},
					K8sDaemonsetDesiredScheduledNodes:   MetricConfig{Enabled: true},
					K8sDaemonsetMisscheduledNodes:       MetricConfig{Enabled: true},
					K8sDaemonsetReadyNodes:              MetricConfig{Enabled: true},
					K8sDeploymentAvailable:              MetricConfig{Enabled: true},
					K8sDeploymentDesired:                MetricConfig{Enabled: true},
					K8sHpaCurrentReplicas:               MetricConfig{Enabled: true},
					K8sHpaDesiredReplicas:               MetricConfig{Enabled: true},
					K8sHpaMaxReplicas:                   MetricConfig{Enabled: true},
					K8sHpaMinReplicas:                   MetricConfig{Enabled: true},
					K8sJobActivePods:                    MetricConfig{Enabled: true},
					K8sJobDesiredSuccessfulPods:         MetricConfig{Enabled: true},
					K8sJobFailedPods:                    MetricConfig{Enabled: true},
					K8sJobMaxParallelPods:               MetricConfig{Enabled: true},
					K8sJobSuccessfulPods:                MetricConfig{Enabled: true},
					K8sNamespacePhase:                   MetricConfig{Enabled: true},
					K8sNodeCondition:                    MetricConfig{Enabled: true},
					K8sPodPhase:                         MetricConfig{Enabled: true},
					K8sPodStatusReason:                  MetricConfig{Enabled: true},
					K8sReplicasetAvailable:              MetricConfig{Enabled: true},
					K8sReplicasetDesired:                MetricConfig{Enabled: true},
					K8sReplicationControllerAvailable:   MetricConfig{Enabled: true},
					K8sReplicationControllerDesired:     MetricConfig{Enabled: true},
					K8sResourceQuotaHardLimit:           MetricConfig{Enabled: true},
					K8sResourceQuotaUsed:                MetricConfig{Enabled: true},
					K8sStatefulsetCurrentPods:           MetricConfig{Enabled: true},
					K8sStatefulsetDesiredPods:           MetricConfig{Enabled: true},
					K8sStatefulsetReadyPods:             MetricConfig{Enabled: true},
					K8sStatefulsetUpdatedPods:           MetricConfig{Enabled: true},
					OpenshiftAppliedclusterquotaLimit:   MetricConfig{Enabled: true},
					OpenshiftAppliedclusterquotaUsed:    MetricConfig{Enabled: true},
					OpenshiftClusterquotaLimit:          MetricConfig{Enabled: true},
					OpenshiftClusterquotaUsed:           MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:                            ResourceAttributeConfig{Enabled: true},
					ContainerImageName:                     ResourceAttributeConfig{Enabled: true},
					ContainerImageTag:                      ResourceAttributeConfig{Enabled: true},
					ContainerRuntime:                       ResourceAttributeConfig{Enabled: true},
					ContainerRuntimeVersion:                ResourceAttributeConfig{Enabled: true},
					K8sContainerName:                       ResourceAttributeConfig{Enabled: true},
					K8sContainerStatusLastTerminatedReason: ResourceAttributeConfig{Enabled: true},
					K8sCronjobName:                         ResourceAttributeConfig{Enabled: true},
					K8sCronjobUID:                          ResourceAttributeConfig{Enabled: true},
					K8sDaemonsetName:                       ResourceAttributeConfig{Enabled: true},
					K8sDaemonsetUID:                        ResourceAttributeConfig{Enabled: true},
					K8sDeploymentName:                      ResourceAttributeConfig{Enabled: true},
					K8sDeploymentUID:                       ResourceAttributeConfig{Enabled: true},
					K8sHpaName:                             ResourceAttributeConfig{Enabled: true},
					K8sHpaUID:                              ResourceAttributeConfig{Enabled: true},
					K8sJobName:                             ResourceAttributeConfig{Enabled: true},
					K8sJobUID:                              ResourceAttributeConfig{Enabled: true},
					K8sKubeletVersion:                      ResourceAttributeConfig{Enabled: true},
					K8sNamespaceName:                       ResourceAttributeConfig{Enabled: true},
					K8sNamespaceUID:                        ResourceAttributeConfig{Enabled: true},
					K8sNodeAnnotations:                     ResourceAttributeConfig{Enabled: true},
					K8sNodeLabels:                          ResourceAttributeConfig{Enabled: true},
					K8sNodeName:                            ResourceAttributeConfig{Enabled: true},
					K8sNodeUID:                             ResourceAttributeConfig{Enabled: true},
					K8sPodName:                             ResourceAttributeConfig{Enabled: true},
					K8sPodQosClass:                         ResourceAttributeConfig{Enabled: true},
					K8sPodUID:                              ResourceAttributeConfig{Enabled: true},
					K8sReplicasetName:                      ResourceAttributeConfig{Enabled: true},
					K8sReplicasetUID:                       ResourceAttributeConfig{Enabled: true},
					K8sReplicationcontrollerName:           ResourceAttributeConfig{Enabled: true},
					K8sReplicationcontrollerUID:            ResourceAttributeConfig{Enabled: true},
					K8sResourcequotaName:                   ResourceAttributeConfig{Enabled: true},
					K8sResourcequotaUID:                    ResourceAttributeConfig{Enabled: true},
					K8sStatefulsetName:                     ResourceAttributeConfig{Enabled: true},
					K8sStatefulsetUID:                      ResourceAttributeConfig{Enabled: true},
					OpenshiftClusterquotaName:              ResourceAttributeConfig{Enabled: true},
					OpenshiftClusterquotaUID:               ResourceAttributeConfig{Enabled: true},
					OsDescription:                          ResourceAttributeConfig{Enabled: true},
					OsType:                                 ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sContainerCPULimit:                MetricConfig{Enabled: false},
					K8sContainerCPURequest:              MetricConfig{Enabled: false},
					K8sContainerEphemeralstorageLimit:   MetricConfig{Enabled: false},
					K8sContainerEphemeralstorageRequest: MetricConfig{Enabled: false},
					K8sContainerMemoryLimit:             MetricConfig{Enabled: false},
					K8sContainerMemoryRequest:           MetricConfig{Enabled: false},
					K8sContainerReady:                   MetricConfig{Enabled: false},
					K8sContainerRestarts:                MetricConfig{Enabled: false},
					K8sContainerStorageLimit:            MetricConfig{Enabled: false},
					K8sContainerStorageRequest:          MetricConfig{Enabled: false},
					K8sCronjobActiveJobs:                MetricConfig{Enabled: false},
					K8sDaemonsetCurrentScheduledNodes:   MetricConfig{Enabled: false},
					K8sDaemonsetDesiredScheduledNodes:   MetricConfig{Enabled: false},
					K8sDaemonsetMisscheduledNodes:       MetricConfig{Enabled: false},
					K8sDaemonsetReadyNodes:              MetricConfig{Enabled: false},
					K8sDeploymentAvailable:              MetricConfig{Enabled: false},
					K8sDeploymentDesired:                MetricConfig{Enabled: false},
					K8sHpaCurrentReplicas:               MetricConfig{Enabled: false},
					K8sHpaDesiredReplicas:               MetricConfig{Enabled: false},
					K8sHpaMaxReplicas:                   MetricConfig{Enabled: false},
					K8sHpaMinReplicas:                   MetricConfig{Enabled: false},
					K8sJobActivePods:                    MetricConfig{Enabled: false},
					K8sJobDesiredSuccessfulPods:         MetricConfig{Enabled: false},
					K8sJobFailedPods:                    MetricConfig{Enabled: false},
					K8sJobMaxParallelPods:               MetricConfig{Enabled: false},
					K8sJobSuccessfulPods:                MetricConfig{Enabled: false},
					K8sNamespacePhase:                   MetricConfig{Enabled: false},
					K8sNodeCondition:                    MetricConfig{Enabled: false},
					K8sPodPhase:                         MetricConfig{Enabled: false},
					K8sPodStatusReason:                  MetricConfig{Enabled: false},
					K8sReplicasetAvailable:              MetricConfig{Enabled: false},
					K8sReplicasetDesired:                MetricConfig{Enabled: false},
					K8sReplicationControllerAvailable:   MetricConfig{Enabled: false},
					K8sReplicationControllerDesired:     MetricConfig{Enabled: false},
					K8sResourceQuotaHardLimit:           MetricConfig{Enabled: false},
					K8sResourceQuotaUsed:                MetricConfig{Enabled: false},
					K8sStatefulsetCurrentPods:           MetricConfig{Enabled: false},
					K8sStatefulsetDesiredPods:           MetricConfig{Enabled: false},
					K8sStatefulsetReadyPods:             MetricConfig{Enabled: false},
					K8sStatefulsetUpdatedPods:           MetricConfig{Enabled: false},
					OpenshiftAppliedclusterquotaLimit:   MetricConfig{Enabled: false},
					OpenshiftAppliedclusterquotaUsed:    MetricConfig{Enabled: false},
					OpenshiftClusterquotaLimit:          MetricConfig{Enabled: false},
					OpenshiftClusterquotaUsed:           MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:                            ResourceAttributeConfig{Enabled: false},
					ContainerImageName:                     ResourceAttributeConfig{Enabled: false},
					ContainerImageTag:                      ResourceAttributeConfig{Enabled: false},
					ContainerRuntime:                       ResourceAttributeConfig{Enabled: false},
					ContainerRuntimeVersion:                ResourceAttributeConfig{Enabled: false},
					K8sContainerName:                       ResourceAttributeConfig{Enabled: false},
					K8sContainerStatusLastTerminatedReason: ResourceAttributeConfig{Enabled: false},
					K8sCronjobName:                         ResourceAttributeConfig{Enabled: false},
					K8sCronjobUID:                          ResourceAttributeConfig{Enabled: false},
					K8sDaemonsetName:                       ResourceAttributeConfig{Enabled: false},
					K8sDaemonsetUID:                        ResourceAttributeConfig{Enabled: false},
					K8sDeploymentName:                      ResourceAttributeConfig{Enabled: false},
					K8sDeploymentUID:                       ResourceAttributeConfig{Enabled: false},
					K8sHpaName:                             ResourceAttributeConfig{Enabled: false},
					K8sHpaUID:                              ResourceAttributeConfig{Enabled: false},
					K8sJobName:                             ResourceAttributeConfig{Enabled: false},
					K8sJobUID:                              ResourceAttributeConfig{Enabled: false},
					K8sKubeletVersion:                      ResourceAttributeConfig{Enabled: false},
					K8sNamespaceName:                       ResourceAttributeConfig{Enabled: false},
					K8sNamespaceUID:                        ResourceAttributeConfig{Enabled: false},
					K8sNodeAnnotations:                     ResourceAttributeConfig{Enabled: false},
					K8sNodeLabels:                          ResourceAttributeConfig{Enabled: false},
					K8sNodeName:                            ResourceAttributeConfig{Enabled: false},
					K8sNodeUID:                             ResourceAttributeConfig{Enabled: false},
					K8sPodName:                             ResourceAttributeConfig{Enabled: false},
					K8sPodQosClass:                         ResourceAttributeConfig{Enabled: false},
					K8sPodUID:                              ResourceAttributeConfig{Enabled: false},
					K8sReplicasetName:                      ResourceAttributeConfig{Enabled: false},
					K8sReplicasetUID:                       ResourceAttributeConfig{Enabled: false},
					K8sReplicationcontrollerName:           ResourceAttributeConfig{Enabled: false},
					K8sReplicationcontrollerUID:            ResourceAttributeConfig{Enabled: false},
					K8sResourcequotaName:                   ResourceAttributeConfig{Enabled: false},
					K8sResourcequotaUID:                    ResourceAttributeConfig{Enabled: false},
					K8sStatefulsetName:                     ResourceAttributeConfig{Enabled: false},
					K8sStatefulsetUID:                      ResourceAttributeConfig{Enabled: false},
					OpenshiftClusterquotaName:              ResourceAttributeConfig{Enabled: false},
					OpenshiftClusterquotaUID:               ResourceAttributeConfig{Enabled: false},
					OsDescription:                          ResourceAttributeConfig{Enabled: false},
					OsType:                                 ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				ContainerID:                            ResourceAttributeConfig{Enabled: true},
				ContainerImageName:                     ResourceAttributeConfig{Enabled: true},
				ContainerImageTag:                      ResourceAttributeConfig{Enabled: true},
				ContainerRuntime:                       ResourceAttributeConfig{Enabled: true},
				ContainerRuntimeVersion:                ResourceAttributeConfig{Enabled: true},
				K8sContainerName:                       ResourceAttributeConfig{Enabled: true},
				K8sContainerStatusLastTerminatedReason: ResourceAttributeConfig{Enabled: true},
				K8sCronjobName:                         ResourceAttributeConfig{Enabled: true},
				K8sCronjobUID:                          ResourceAttributeConfig{Enabled: true},
				K8sDaemonsetName:                       ResourceAttributeConfig{Enabled: true},
				K8sDaemonsetUID:                        ResourceAttributeConfig{Enabled: true},
				K8sDeploymentName:                      ResourceAttributeConfig{Enabled: true},
				K8sDeploymentUID:                       ResourceAttributeConfig{Enabled: true},
				K8sHpaName:                             ResourceAttributeConfig{Enabled: true},
				K8sHpaUID:                              ResourceAttributeConfig{Enabled: true},
				K8sJobName:                             ResourceAttributeConfig{Enabled: true},
				K8sJobUID:                              ResourceAttributeConfig{Enabled: true},
				K8sKubeletVersion:                      ResourceAttributeConfig{Enabled: true},
				K8sNamespaceName:                       ResourceAttributeConfig{Enabled: true},
				K8sNamespaceUID:                        ResourceAttributeConfig{Enabled: true},
				K8sNodeAnnotations:                     ResourceAttributeConfig{Enabled: true},
				K8sNodeLabels:                          ResourceAttributeConfig{Enabled: true},
				K8sNodeName:                            ResourceAttributeConfig{Enabled: true},
				K8sNodeUID:                             ResourceAttributeConfig{Enabled: true},
				K8sPodName:                             ResourceAttributeConfig{Enabled: true},
				K8sPodQosClass:                         ResourceAttributeConfig{Enabled: true},
				K8sPodUID:                              ResourceAttributeConfig{Enabled: true},
				K8sReplicasetName:                      ResourceAttributeConfig{Enabled: true},
				K8sReplicasetUID:                       ResourceAttributeConfig{Enabled: true},
				K8sReplicationcontrollerName:           ResourceAttributeConfig{Enabled: true},
				K8sReplicationcontrollerUID:            ResourceAttributeConfig{Enabled: true},
				K8sResourcequotaName:                   ResourceAttributeConfig{Enabled: true},
				K8sResourcequotaUID:                    ResourceAttributeConfig{Enabled: true},
				K8sStatefulsetName:                     ResourceAttributeConfig{Enabled: true},
				K8sStatefulsetUID:                      ResourceAttributeConfig{Enabled: true},
				OpenshiftClusterquotaName:              ResourceAttributeConfig{Enabled: true},
				OpenshiftClusterquotaUID:               ResourceAttributeConfig{Enabled: true},
				OsDescription:                          ResourceAttributeConfig{Enabled: true},
				OsType:                                 ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ContainerID:                            ResourceAttributeConfig{Enabled: false},
				ContainerImageName:                     ResourceAttributeConfig{Enabled: false},
				ContainerImageTag:                      ResourceAttributeConfig{Enabled: false},
				ContainerRuntime:                       ResourceAttributeConfig{Enabled: false},
				ContainerRuntimeVersion:                ResourceAttributeConfig{Enabled: false},
				K8sContainerName:                       ResourceAttributeConfig{Enabled: false},
				K8sContainerStatusLastTerminatedReason: ResourceAttributeConfig{Enabled: false},
				K8sCronjobName:                         ResourceAttributeConfig{Enabled: false},
				K8sCronjobUID:                          ResourceAttributeConfig{Enabled: false},
				K8sDaemonsetName:                       ResourceAttributeConfig{Enabled: false},
				K8sDaemonsetUID:                        ResourceAttributeConfig{Enabled: false},
				K8sDeploymentName:                      ResourceAttributeConfig{Enabled: false},
				K8sDeploymentUID:                       ResourceAttributeConfig{Enabled: false},
				K8sHpaName:                             ResourceAttributeConfig{Enabled: false},
				K8sHpaUID:                              ResourceAttributeConfig{Enabled: false},
				K8sJobName:                             ResourceAttributeConfig{Enabled: false},
				K8sJobUID:                              ResourceAttributeConfig{Enabled: false},
				K8sKubeletVersion:                      ResourceAttributeConfig{Enabled: false},
				K8sNamespaceName:                       ResourceAttributeConfig{Enabled: false},
				K8sNamespaceUID:                        ResourceAttributeConfig{Enabled: false},
				K8sNodeAnnotations:                     ResourceAttributeConfig{Enabled: false},
				K8sNodeLabels:                          ResourceAttributeConfig{Enabled: false},
				K8sNodeName:                            ResourceAttributeConfig{Enabled: false},
				K8sNodeUID:                             ResourceAttributeConfig{Enabled: false},
				K8sPodName:                             ResourceAttributeConfig{Enabled: false},
				K8sPodQosClass:                         ResourceAttributeConfig{Enabled: false},
				K8sPodUID:                              ResourceAttributeConfig{Enabled: false},
				K8sReplicasetName:                      ResourceAttributeConfig{Enabled: false},
				K8sReplicasetUID:                       ResourceAttributeConfig{Enabled: false},
				K8sReplicationcontrollerName:           ResourceAttributeConfig{Enabled: false},
				K8sReplicationcontrollerUID:            ResourceAttributeConfig{Enabled: false},
				K8sResourcequotaName:                   ResourceAttributeConfig{Enabled: false},
				K8sResourcequotaUID:                    ResourceAttributeConfig{Enabled: false},
				K8sStatefulsetName:                     ResourceAttributeConfig{Enabled: false},
				K8sStatefulsetUID:                      ResourceAttributeConfig{Enabled: false},
				OpenshiftClusterquotaName:              ResourceAttributeConfig{Enabled: false},
				OpenshiftClusterquotaUID:               ResourceAttributeConfig{Enabled: false},
				OsDescription:                          ResourceAttributeConfig{Enabled: false},
				OsType:                                 ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
