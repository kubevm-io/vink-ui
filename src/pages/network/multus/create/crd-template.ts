export const multusYaml = `
apiVersion: k8s.cni.cncf.io/v1
kind: NetworkAttachmentDefinition
metadata:
  name: attachnet
  namespace: default
spec:
  config: >-
    { "cniVersion": "0.3.0", "type": "kube-ovn", "server_socket":
    "/run/openvswitch/kube-ovn-daemon.sock", "provider": "attachnet.default.ovn"
    }
`

// export const multusYaml = `
// apiVersion: "k8s.cni.cncf.io/v1"
// kind: NetworkAttachmentDefinition
// metadata:
//   name: attachnet
//   namespace: default
// spec:
//   config: '{
//       "cniVersion": "0.3.0",
//       "type": "kube-ovn",
//       "server_socket": "/run/openvswitch/kube-ovn-daemon.sock",
//       "provider": "attachnet.default.ovn"
//     }'
// `
