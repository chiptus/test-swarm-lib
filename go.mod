module github.com/chiptus/test-swarm-lib

go 1.16

require (
	github.com/docker/docker v20.10.3+incompatible // indirect
	// github.com/docker/cli v20.10.7
	github.com/spf13/pflag v1.0.5

)

replace (
	github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200204220554-5f6d6f3f2203
	k8s.io/api => k8s.io/api v0.0.0-20191003000013-35e20aa79eb8
	// k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.2
	// k8s.io/apimachinery => k8s.io/apimachinery v0.21.2
	// k8s.io/apiserver => k8s.io/apiserver v0.21.2
	// k8s.io/cli-runtime => k8s.io/cli-runtime v0.21.2
	// k8s.io/client-go => k8s.io/client-go v0.21.2
	// k8s.io/cloud-provider => k8s.io/cloud-provider v0.21.2
	// k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.21.2
	// k8s.io/code-generator => k8s.io/code-generator v0.21.2
	// k8s.io/component-base => k8s.io/component-base v0.21.2
	// k8s.io/component-helpers => k8s.io/component-helpers v0.21.2
	// k8s.io/controller-manager => k8s.io/controller-manager v0.21.2
	// k8s.io/cri-api => k8s.io/cri-api v0.21.2
	// k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.21.2
	// k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.21.2
	// k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.21.2
	// k8s.io/kube-proxy => k8s.io/kube-proxy v0.21.2
	// k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.21.2
	// k8s.io/kubectl => k8s.io/kubectl v0.21.2
	// k8s.io/kubelet => k8s.io/kubelet v0.21.2
	// k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.21.2
	// k8s.io/metrics => k8s.io/metrics v0.21.2
	// k8s.io/mount-utils => k8s.io/mount-utils v0.21.2
	// k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.21.2
	// k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.21.2
	vbom.ml/util/sortorder => github.com/fvbommel/sortorder v1.0.2
)

// go get -u k8s.io/api@35e20aa79eb876d1014e0383c7fcda49e52c5d76  k8s.io/apimachinery@27d36303b6556f377b4f34e64705fa9024a12b0c  k8s.io/apiextensions-apiserver@49e3d608220c016ce72a3bd9524eed4dcef587df  k8s.io/apiserver@3c8b233e046cffa0f8b63fff9e3ba4014a1ef456  k8s.io/client-go@f68efa97b39e67ee47d4b4683d0037d721f077be  k8s.io/component-base@f573d376509cec44c50b8d79741bc939ec7ede56  k8s.io/klog@3ca30a56d8a775276f9cdae009ba326fdc05af7f  k8s.io/kube-openapi@0270cf2f1c1d995d34b36019a6f65d58e6e33ad4 k8s.io/kubernetes@d647ddbd755faf07169599a625faf302ffc34458  k8s.io/utils@69764acb6e8e900b7c05296c5d3c9c19545475f9
