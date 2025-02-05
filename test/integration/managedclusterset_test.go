package integration

import (
	"context"
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	clusterv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
)

var _ = ginkgo.Describe("ManagedClusterSet API test", func() {
	var clusterSetName string

	ginkgo.BeforeEach(func() {
		suffix := rand.String(5)
		clusterSetName = fmt.Sprintf("clusterset-%s", suffix)
	})

	ginkgo.It("Create a clusterset with empty spec", func() {
		clusterset := &clusterv1beta1.ManagedClusterSet{
			ObjectMeta: metav1.ObjectMeta{
				Name: clusterSetName,
			},
		}
		clusterset, err := hubClusterClient.ClusterV1beta1().ManagedClusterSets().Create(context.TODO(), clusterset, metav1.CreateOptions{})
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		gomega.Expect(clusterset.Spec.ClusterSelector.SelectorType).Should(gomega.Equal(clusterv1beta1.LegacyClusterSetLabel))
	})

	ginkgo.It("Create a clusterset with wrong selector", func() {
		clusterset := &clusterv1beta1.ManagedClusterSet{
			ObjectMeta: metav1.ObjectMeta{
				Name: clusterSetName,
			},
			Spec: clusterv1beta1.ManagedClusterSetSpec{
				ClusterSelector: clusterv1beta1.ManagedClusterSelector{
					SelectorType: "WrongSelector",
				},
			},
		}
		_, err := hubClusterClient.ClusterV1beta1().ManagedClusterSets().Create(context.TODO(), clusterset, metav1.CreateOptions{})
		gomega.Expect(err).To(gomega.HaveOccurred())
	})
})
