/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	kcorev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Informer Cache", func() {
	var stop chan struct{}

	BeforeEach(func() {
		stop = make(chan struct{})
		Expect(cfg).NotTo(BeNil())
	})
	AfterEach(func() {
		close(stop)
	})

	Describe("as a Reader", func() {
		It("should be able to list objects that haven't been watched previously", func() {
			By("Creating the cache")
			reader, err := cache.New(cfg, cache.Options{})
			Expect(err).NotTo(HaveOccurred())

			By("running the cache and waiting for it to sync")
			go func() {
				defer GinkgoRecover()
				Expect(reader.Start(stop)).ToNot(HaveOccurred())
			}()
			Expect(reader.WaitForCacheSync(stop)).NotTo(BeFalse())

			By("Listing all services in the cluster")
			listObj := &kcorev1.ServiceList{}
			Expect(reader.List(context.Background(), nil, listObj)).NotTo(HaveOccurred())

			By("Verifying that the returned list contains the Kubernetes service")
			// NB: there has to be at least the kubernetes service in the cluster
			Expect(listObj.Items).NotTo(BeEmpty())
			hasKubeService := false
			for _, svc := range listObj.Items {
				if svc.Namespace == "default" && svc.Name == "kubernetes" {
					hasKubeService = true
					break
				}
			}
			Expect(hasKubeService).To(BeTrue())
		})

		It("should be able to get objects that haven't been watched previously", func() {
			By("Creating the cache")
			reader, err := cache.New(cfg, cache.Options{})
			Expect(err).NotTo(HaveOccurred())

			By("running the cache and waiting for it to sync")
			go func() {
				defer GinkgoRecover()
				Expect(reader.Start(stop)).ToNot(HaveOccurred())
			}()
			Expect(reader.WaitForCacheSync(stop)).NotTo(BeFalse())

			By("Getting the Kubernetes service")
			svc := &kcorev1.Service{}
			Expect(reader.Get(context.Background(), client.ObjectKey{Namespace: "default", Name: "kubernetes"}, svc)).NotTo(HaveOccurred())

			By("Verifying that the returned service looks reasonable")
			Expect(svc.Name).To(Equal("kubernetes"))
			Expect(svc.Namespace).To(Equal("default"))
		})
	})
})

var _ = Describe("Indexers", func() {
	//three := int64(3)
	//knownPodKey := client.ObjectKey{Name: "some-pod", Namespace: "some-ns"}
	//knownPod3Key := client.ObjectKey{Name: "some-pod", Namespace: "some-other-ns"}
	//knownVolumeKey := client.ObjectKey{Name: "some-vol", Namespace: "some-ns"}
	//knownPod := &kapi.Pod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      knownPodKey.Name,
	//		Namespace: knownPodKey.Namespace,
	//	},
	//	Spec: kapi.PodSpec{
	//		RestartPolicy:         kapi.RestartPolicyNever,
	//		ActiveDeadlineSeconds: &three,
	//	},
	//}
	//knownPod2 := &kapi.Pod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      knownVolumeKey.Name,
	//		Namespace: knownVolumeKey.Namespace,
	//		Labels: map[string]string{
	//			"somelbl": "someval",
	//		},
	//	},
	//	Spec: kapi.PodSpec{
	//		RestartPolicy: kapi.RestartPolicyAlways,
	//	},
	//}
	//knownPod3 := &kapi.Pod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      knownPod3Key.Name,
	//		Namespace: knownPod3Key.Namespace,
	//		Labels: map[string]string{
	//			"somelbl": "someval",
	//		},
	//	},
	//	Spec: kapi.PodSpec{
	//		RestartPolicy: kapi.RestartPolicyNever,
	//	},
	//}
	//knownVolume := &kapi.PersistentVolume{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      knownVolumeKey.Name,
	//		Namespace: knownVolumeKey.Namespace,
	//	},
	//}
	//var multiCache *cache
	//
	//BeforeEach(func() {
	//	multiCache = &cache{
	//		cachesByType: make(map[reflect.Type]*singleObjectCache),
	//		scheme:       scheme.Scheme,
	//	}
	//	podIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
	//		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	//	})
	//	volumeIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
	//		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	//	})
	//	indexByField(podIndexer, "spec.restartPolicy", func(obj runtime.Object) []string {
	//		return []string{string(obj.(*kapi.Pod).Spec.RestartPolicy)}
	//	})
	//	Expect(podIndexer.Add(knownPod)).NotTo(HaveOccurred())
	//	Expect(podIndexer.Add(knownPod2)).NotTo(HaveOccurred())
	//	Expect(podIndexer.Add(knownPod3)).NotTo(HaveOccurred())
	//	Expect(volumeIndexer.Add(knownVolume)).NotTo(HaveOccurred())
	//	multiCache.registerCache(&kapi.Pod{}, kapi.SchemeGroupVersion.WithKind("Pod"), podIndexer)
	//	multiCache.registerCache(&kapi.PersistentVolume{}, kapi.SchemeGroupVersion.WithKind("PersistentVolume"), volumeIndexer)
	//})
	//
	//Describe("populatingClient interface wrapper around an indexer", func() {
	//	var singleCache client.Reader
	//
	//	BeforeEach(func() {
	//		var err error
	//		singleCache, err = multiCache.cacheFor(&kapi.Pod{})
	//		Expect(err).NotTo(HaveOccurred())
	//	})
	//
	//	It("should be able to fetch a particular object by key", func() {
	//		out := kapi.Pod{}
	//		Expect(singleCache.Get(context.TODO(), knownPodKey, &out)).NotTo(HaveOccurred())
	//		Expect(&out).To(Equal(knownPod))
	//	})
	//
	//	It("should error out for missing objects", func() {
	//		Expect(singleCache.Get(context.TODO(), client.ObjectKey{Name: "unknown-pod"}, &kapi.Pod{})).To(HaveOccurred())
	//	})
	//
	//	It("should be able to list objects by namespace", func() {
	//		out := kapi.PodList{}
	//		Expect(singleCache.List(context.TODO(), client.InNamespace(knownPodKey.Namespace), &out)).NotTo(HaveOccurred())
	//		Expect(out.Items).To(ConsistOf(*knownPod, *knownPod2))
	//	})
	//
	//	It("should error out if the incorrect object type is passed for this indexer", func() {
	//		Expect(singleCache.Get(context.TODO(), knownPodKey, &kapi.PersistentVolume{})).To(HaveOccurred())
	//	})
	//
	//	It("should deep copy the object unless told otherwise", func() {
	//		out := kapi.Pod{}
	//		Expect(singleCache.Get(context.TODO(), knownPodKey, &out)).NotTo(HaveOccurred())
	//		Expect(&out).To(Equal(knownPod))
	//
	//		*out.Spec.ActiveDeadlineSeconds = 4
	//		Expect(*out.Spec.ActiveDeadlineSeconds).NotTo(Equal(*knownPod.Spec.ActiveDeadlineSeconds))
	//	})
	//
	//	It("should support filtering by labels", func() {
	//		out := kapi.PodList{}
	//		Expect(singleCache.List(context.TODO(), client.InNamespace(knownPodKey.Namespace).
	// 			MatchingLabels(map[string]string{"somelbl": "someval"}), &out)).NotTo(HaveOccurred())
	//		Expect(out.Items).To(ConsistOf(*knownPod2))
	//	})
	//
	//	It("should support filtering by a single field=value specification, if previously indexed", func() {
	//		By("listing by field selector in a namespace")
	//		out := kapi.PodList{}
	//		Expect(singleCache.List(context.TODO(), client.InNamespace(knownPodKey.Namespace).MatchingField("spec.restartPolicy", "Always"), &out)).NotTo(HaveOccurred())
	//		Expect(out.Items).To(ConsistOf(*knownPod2))
	//
	//		By("listing by field selector across all namespaces")
	//		Expect(singleCache.List(context.TODO(), client.MatchingField("spec.restartPolicy", "Never"), &out)).NotTo(HaveOccurred())
	//		Expect(out.Items).To(ConsistOf(*knownPod, *knownPod3))
	//	})
	//})
	//
	//Describe("populatingClient interface wrapper around multiple indexers", func() {
	//	It("should be able to fetch any known object by key and type", func() {
	//		outPod := kapi.Pod{}
	//		Expect(multiCache.Get(context.TODO(), knownPodKey, &outPod)).NotTo(HaveOccurred())
	//		Expect(&outPod).To(Equal(knownPod))
	//
	//		outVol := kapi.PersistentVolume{}
	//		Expect(multiCache.Get(context.TODO(), knownVolumeKey, &outVol)).NotTo(HaveOccurred())
	//		Expect(&outVol).To(Equal(knownVolume))
	//	})
	//
	//	It("should error out if the object type is unknown", func() {
	//		Expect(multiCache.Get(context.TODO(), knownPodKey, &kapi.PersistentVolumeClaim{})).To(HaveOccurred())
	//	})
	//
	//	It("should deep copy the object unless told otherwise", func() {
	//		out := kapi.Pod{}
	//		Expect(multiCache.Get(context.TODO(), knownPodKey, &out)).NotTo(HaveOccurred())
	//		Expect(&out).To(Equal(knownPod))
	//
	//		*out.Spec.ActiveDeadlineSeconds = 4
	//		Expect(*out.Spec.ActiveDeadlineSeconds).NotTo(Equal(*knownPod.Spec.ActiveDeadlineSeconds))
	//	})
	//
	//	It("should be able to fetch single caches for known types", func() {
	//		indexer, ok := multiCache.cacheFor(&kapi.Pod{})
	//		Expect(ok).To(BeTrue())
	//		Expect(indexer).NotTo(BeNil())
	//
	//		_, ok2 := multiCache.cacheFor(&kapi.PersistentVolumeClaim{})
	//		Expect(ok2).To(BeFalse())
	//	})
	//})
})
