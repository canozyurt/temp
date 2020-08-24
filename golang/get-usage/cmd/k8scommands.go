package cmd

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/cmd/top"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
	"strconv"
)

type PodAttributes struct {
	name string
	namespace string
	cpuLimit int64
	cpuUsage int64
	cpuRatio float64
	cpuInfo string
	memoryLimit int64
	memoryUsage int64
	memoryInfo string
	memoryRatio float64
	nodeName string

}


type MyClientSet struct{
	*kubernetes.Clientset
}

var(
	configFlags = genericclioptions.NewConfigFlags(true)
)

func newClientSetFromConfigFile() *MyClientSet {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		panic(err.Error())
	}

	created, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &MyClientSet{created}
}

func (cs MyClientSet) getPods() (pods *v1.PodList){

	pods, err := cs.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return
}

func (cs MyClientSet) topPods() (metrics *v1beta1.PodMetricsList){

	o := top.TopPodOptions{}

	config, err := configFlags.ToRESTConfig()
	if err != nil {
		panic(err.Error())
	}

	o.MetricsClient, err = metricsclientset.NewForConfig(config)

	metrics, err = o.MetricsClient.MetricsV1beta1().PodMetricses(o.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return

}

func parsePods (pods *v1.PodList, MyList *map[string]*PodAttributes) error {

	for _, pod := range pods.Items {
		(*MyList)[pod.Name] = &PodAttributes{
			name: pod.Name,
			namespace: pod.Namespace,
			cpuLimit: pod.Spec.Containers[0].Resources.Limits.Cpu().MilliValue(),
			memoryLimit: pod.Spec.Containers[0].Resources.Limits.Memory().MilliValue()/10240000000,
			nodeName: pod.Spec.NodeName,
		}
	}
	return nil
}

func parseMetrics (metrics *v1beta1.PodMetricsList, MyList *map[string]*PodAttributes) error {

	var sumCpu, sumMemory int64

	for _, metric := range metrics.Items {
		sumCpu = 0
		sumMemory = 0
		for _, container := range metric.Containers {
			sumCpu += container.Usage.Cpu().MilliValue()
			sumMemory += container.Usage.Memory().MilliValue()/10240000000
		}

		if (*MyList)[metric.Name].cpuLimit == 0 {
			(*MyList)[metric.Name].cpuInfo = "LimitNotSet"
		} else {
			(*MyList)[metric.Name].cpuRatio = float64(sumCpu) / float64((*MyList)[metric.Name].cpuLimit) * 100
			(*MyList)[metric.Name].cpuInfo = strconv.FormatFloat((*MyList)[metric.Name].cpuRatio, 'f', 2, 64)
		}

		if (*MyList)[metric.Name].memoryLimit == 0 {
			(*MyList)[metric.Name].memoryInfo = "LimitNotSet"
		} else {
			(*MyList)[metric.Name].memoryRatio = float64(sumMemory) / float64((*MyList)[metric.Name].memoryLimit) * 100
			(*MyList)[metric.Name].memoryInfo = strconv.FormatFloat((*MyList)[metric.Name].memoryRatio, 'f', 2, 64)
		}

	}
	return nil
}

func (p PodAttributes) String() string {

	return fmt.Sprintf("%-15v " +
		"%-65v " +
		"cpu: %-15v " +
		"memory: %-15v " +
		"%v ", p.namespace, p.name, p.cpuInfo, p.memoryInfo, p.nodeName)
}