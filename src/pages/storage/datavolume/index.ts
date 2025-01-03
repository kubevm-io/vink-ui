import * as yaml from 'js-yaml'
import { dataVolumYaml } from './template'
import { instances as labels } from "@/clients/ts/label/labels.gen"
import { NamespaceName } from '@/clients/ts/types/types'

const defaultAccessMode = "ReadWriteOnce"

const defaultStorageClass = "local-path"

export const newSystemImage = (namespaceName: NamespaceName, imageSource: string, imageCapacity: number, family: string, version: string) => {
    const instance: any = yaml.load(dataVolumYaml)

    instance.metadata.name = namespaceName.name
    instance.metadata.namespace = namespaceName.namespace
    instance.spec.pvc.resources.requests.storage = `${imageCapacity}Gi`
    instance.spec.source = { registry: { url: imageSource } }

    instance.metadata.labels[labels.VinkDatavolumeType.name] = "image"
    instance.metadata.labels[labels.VinkOperatingSystem.name] = family
    instance.metadata.labels[labels.VinkOperatingSystemVersion.name] = version

    return instance
}

export const newDataDisk = (namespaceName: NamespaceName, diskCapacity: number, storageClass?: string, accessMode?: string) => {
    const instance: any = yaml.load(dataVolumYaml)

    instance.metadata.name = namespaceName.name
    instance.metadata.namespace = namespaceName.namespace
    instance.spec.pvc.resources.requests.storage = `${diskCapacity}Gi`
    instance.spec.source = { blank: {} }

    instance.spec.pvc.accessModes = [(accessMode && accessMode.length > 0) ? accessMode : defaultAccessMode]
    instance.spec.pvc.storageClassName = (storageClass && storageClass.length > 0) ? storageClass : defaultStorageClass
    instance.metadata.labels[labels.VinkDatavolumeType.name] = "data"

    return instance
}
