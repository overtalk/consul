package registrar

import (
	"net/http"
	"os"
)

func getConsulID(serverType string) string {
	return ""
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
