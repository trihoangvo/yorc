package ansible

import "novaforge.bull.com/starlings-janus/janus/registry"

const (
	implementationArtifactBash    = "tosca.artifacts.Implementation.Bash"
	implementationArtifactPython  = "tosca.artifacts.Implementation.Python"
	implementationArtifactAnsible = "tosca.artifacts.Implementation.Ansible"
)

func init() {
	reg := registry.GetRegistry()
	reg.RegisterOperationExecutor(
		[]string{
			implementationArtifactBash,
			implementationArtifactPython,
			implementationArtifactAnsible,
		}, NewExecutor(), registry.BuiltinOrigin)
}
