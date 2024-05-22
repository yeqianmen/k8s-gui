package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getK8sClient initializes and returns a Kubernetes client
func getK8sClient() (*kubernetes.Clientset, error) {
	// Use in-cluster config if running inside a k8s cluster
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to kubeconfig if not running inside k8s cluster
		kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	// Initialize Kubernetes client
	clientset, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	// 添加跨域中间件
	r.Use(corsMiddleware())

	r.GET("/v1/deployments", func(c *gin.Context) {
		namespace := c.DefaultQuery("ns", "default")

		deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		type DeploymentInfo struct {
			Name        string            `json:"name"`
			Namespace   string            `json:"namespace"`
			Replicas    int32             `json:"replicas"`
			Labels      map[string]string `json:"labels"`
			Annotations map[string]string `json:"annotations"`
			// 可以根据需要添加更多字段
		}

		var deploymentInfos []DeploymentInfo
		for _, deployment := range deployments.Items {
			deploymentInfo := DeploymentInfo{
				Name:        deployment.Name,
				Namespace:   deployment.Namespace,
				Replicas:    *deployment.Spec.Replicas,
				Labels:      deployment.Labels,
				Annotations: deployment.Annotations,
			}
			deploymentInfos = append(deploymentInfos, deploymentInfo)
		}

		c.JSON(http.StatusOK, gin.H{
			"deployments": deploymentInfos,
		})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic(err.Error())
	}
}

// 跨域中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
