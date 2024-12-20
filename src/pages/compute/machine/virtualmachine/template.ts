export const virtualmachineYaml = `
apiVersion: kubevirt.io/v1
kind: VirtualMachine
metadata:
  name: ""
  namespace: ""
spec:
  dataVolumeTemplates:
    - metadata:
        name: ""
        labels: {}
        annotations: {}
      spec:
        pvc:
          accessModes:
            - ReadWriteOnce
          resources:
            requests:
              storage: ""
          storageClassName: local-path
        source:
          pvc:
            name: ""
            namespace: ""
  runStrategy: Always
  template:
    metadata:
      annotations: {}
    spec:
      architecture: amd64
      domain:
        cpu:
          cores: 1
          model: host-model
          sockets: 1
          threads: 2
        memory:
          guest: 1Gi
        devices:
          disks: []
          interfaces: []
        features:
          acpi:
            enabled: true
        machine:
          type: q35
        resources:
          requests:
            memory: 1Gi
            cpu: "0"
          limits: {}
      networks: []
      volumes: []
`

export const defaultCloudInit = `
#cloud-config
ssh_pwauth: true
disable_root: false
chpasswd: {"list": "root:dangerous", expire: False}

runcmd:
- dhclient -r && dhclient
- sed -i "/#\?PermitRootLogin/s/^.*$/PermitRootLogin yes/g" /etc/ssh/sshd_config
- systemctl restart sshd.service
`
