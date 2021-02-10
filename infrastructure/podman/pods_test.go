package podman_test

import (
	"github.com/ppmoon/home-service/infrastructure/podman"
	"testing"
)

func TestClient_PlayK8sYaml(t *testing.T) {
	pc := podman.NewPodmanClient()
	y := `
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: "2020-12-18T09:38:53Z"
  labels:
    app: nginx7
  name: nginx7_pod
spec:
  containers:
  - command:
    - nginx
    - -g
    - daemon off;
    env:
    - name: PATH
      value: /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
    - name: TERM
      value: xterm
    - name: container
      value: podman
    - name: NJS_VERSION
      value: 0.5.0
    - name: PKG_RELEASE
      value: 1~buster
    - name: NGINX_VERSION
      value: 1.19.6
    - name: HOSTNAME
    image: docker.io/library/nginx:1.19.6
    name: nginx7
    resources: {}
    securityContext:
      allowPrivilegeEscalation: true
      capabilities:
        drop:
        - CAP_MKNOD
        - CAP_NET_RAW
        - CAP_AUDIT_WRITE
      privileged: false
      readOnlyRootFilesystem: false
      seLinuxOptions: {}
    workingDir: /
status: {}
---
metadata:
  creationTimestamp: null
spec: {}
status:
  loadBalancer: {}
`
	// TODO complete unit test
	err := pc.PlayK8sYaml(y)
	if err != nil {
		t.Error(err)
		return
	}
}
