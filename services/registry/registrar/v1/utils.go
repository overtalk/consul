package registrar

import (
	"fmt"
	"net/http"
	"os"
)

func getConsulID(pod podInfo) string {
	return fmt.Sprintf("%s_%s", pod.Namespace, pod.Name)
}

func getPodInfo() podInfo {
	keys := []string{
		"MY_POD_NAME",
		"MY_POD_IP",
		"MY_POD_NAMESPACE",
	}
	values := make(map[string]string)

	for _, key := range keys {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			if key == "MY_POD_IP" {
				value = "127.0.0.1"
			} else {
				value = "unknown"
			}
		}
		values[key] = value
	}

	return podInfo{
		Name:      values["MY_POD_NAME"],
		IP:        values["MY_POD_IP"],
		Namespace: values["MY_POD_NAMESPACE"],
	}
}

func (c *Registrar) updateStatusHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
