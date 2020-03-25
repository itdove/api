package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ClientConfig = map[string]string{
	"":         "ClientConfig represents apiserver address of the spoke cluster",
	"url":      "URL is the url of apiserver endpoint of the spoke cluster",
	"caBundle": "CABundle is the ca bundle to connect to apiserver of the spoke cluster. System certs is used if it is not set.",
}

func (ClientConfig) SwaggerDoc() map[string]string {
	return map_ClientConfig
}

var map_SpokeCluster = map[string]string{
	"":         "SpokeCluster represents the current status of spoke cluster. SpokeCluster is cluster scoped resources. The name is the cluster UID. The cluster join follows the double opt-in process: 1. agent on spoke cluster creates CSR on hub with cluster UID and agent name 2. cluster admin on hub approves the CSR for the spoke's cluster UID and agent name 3. cluster admin on spoke creates credential. Once hub creates the cluster namespace, the spoke agent pushes the credential to hub to use against spoke's kube-apiserver",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata Cluster name must conform in the definition of DNS label format",
	"spec":     "Spec represents a desired configuration for the agent on the spoke cluster.",
	"status":   "Status represents the current status of joined spoke cluster",
}

func (SpokeCluster) SwaggerDoc() map[string]string {
	return map_SpokeCluster
}

var map_SpokeClusterList = map[string]string{
	"":         "SpokeClusterList is a collection of spoke cluster",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of spoke cluster.",
}

func (SpokeClusterList) SwaggerDoc() map[string]string {
	return map_SpokeClusterList
}

var map_SpokeClusterSpec = map[string]string{
	"":                  "SpokeClusterSpec provides the information to securely connect to a remote server and verify its identity",
	"spokeClientConfig": "SpokeClientConfig represents the apiserver address of the spoke cluster",
}

func (SpokeClusterSpec) SwaggerDoc() map[string]string {
	return map_SpokeClusterSpec
}

var map_SpokeClusterStatus = map[string]string{
	"":            "SpokeClusterStatus represents the current status of joined spoke cluster",
	"conditions":  "Conditions contains the different condition statuses for this spoke cluster.",
	"capacity":    "Capacity represents the total resource capacity from all nodeStatuses on the spoke cluster",
	"allocatable": "Allocatable represents the total allocatable resources on the spoke cluster",
	"version":     "Version represents the kubernetes version of the spoke cluster",
}

func (SpokeClusterStatus) SwaggerDoc() map[string]string {
	return map_SpokeClusterStatus
}

var map_SpokeVersion = map[string]string{
	"":           "SpokeVersion represents the Kubernetes version of spoke cluster",
	"kubernetes": "Kubernetes is the kubernetes version of spoke cluster",
}

func (SpokeVersion) SwaggerDoc() map[string]string {
	return map_SpokeVersion
}

var map_StatusCondition = map[string]string{
	"":                   "StatusCondition contains condition information for a spoke cluster.",
	"type":               "Type is the type of the cluster condition.",
	"status":             "Status is the status of the condition. One of True, False, Unknown.",
	"lastTransitionTime": "LastTransitionTime is the last time the condition changed from one status to another.",
	"reason":             "Reason is a (brief) reason for the condition's last status change.",
	"message":            "Message is a human-readable message indicating details about the last status change.",
}

func (StatusCondition) SwaggerDoc() map[string]string {
	return map_StatusCondition
}

// AUTO-GENERATED FUNCTIONS END HERE
