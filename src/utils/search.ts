import type { ListOptions as APIListOptions } from '@kubevm.io/vink/common/common.pb'
// import { osFamilyLabel, diskTypeLabel } from '@/utils/k8s.ts'
import { instances } from "@kubevm.io/vink/label/labels.gen.ts"


export type ListOptions = {
    namespace?: string
    opts?: APIListOptions
}

export type Filter = {
    key: string
    operator: string
    value: string
}

export const LabelsSelector = (filters: Filter[]): string => {
    const labels = filters.map((filter) => {
        return `${filter.key}${filter.operator}${filter.value}`
    })
    return labels.join(',')
}

export const LabelsSelectorString = (filters: string[]): string => {
    const labels = filters.map((filter) => {
        return `${filter}`
    })
    return labels.join(',')
}

export const FieldSelector = (filter: Filter): string => {
    return `${filter.key}${filter.operator}${filter.value}`
}

export const NameFieldSelector = (name: string): string => {
    if (name.length === 0) {
        return ''
    }
    return FieldSelector({
        key: 'metadata.name',
        operator: '=',
        value: name
    })
}

export const virtualMachineOSLabelSelector = (name: string): string => {
    if (name.length === 0) {
        return ''
    }
    return LabelsSelector([{
        key: instances.VinkVirtualmachineOs.name,
        operator: '=',
        value: name
    }])
}

export const dataVolumeTypeLabelSelector = (name: string) => {
    if (name.length === 0) {
        return ''
    }
    return LabelsSelector([{
        key: instances.VinkDatavolumeType.name,
        operator: '=',
        value: name
    }])
}
