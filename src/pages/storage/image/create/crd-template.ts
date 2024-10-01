export const imageYaml = `
apiVersion: cdi.kubevirt.io/v1beta1
kind: DataVolume
metadata:
  name: ""
  namespace: ""
  labels: {}
  annotations:
    cdi.kubevirt.io/storage.bind.immediate.requested: "true"
spec:
  pvc:
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: ""
    storageClassName: local-path
  source:
    registry:
      url: ""
`

// url: docker://quay.io/kubevirt/cirros-container-disk-demo
// url: docker://release.daocloud.io/virtnest/system-images/centos-7.9-x86_64:v1