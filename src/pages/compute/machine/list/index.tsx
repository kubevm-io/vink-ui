import { PlusOutlined, LoadingOutlined } from '@ant-design/icons'
import { ProTable } from '@ant-design/pro-components'
import { App, Button, Modal, Select, Space } from 'antd'
import { useEffect, useRef, useState } from 'react'
import { namespaceName } from '@/utils/k8s'
import { NavLink, Params } from 'react-router-dom'
import { VirtualMachinePowerStateRequest_PowerState } from '@/apis/management/virtualmachine/v1alpha1/virtualmachine'
import { CustomResourceDefinition } from "@/apis/apiextensions/v1alpha1/custom_resource_definition"
import { ResourceUpdater } from '@/pages/compute/machine/list/resource-manager'
import { batchDeleteVirtualMachines, batchManageVirtualMachinePowerState } from '@/resource-manager/virtualmachine'
import { calcScroll, classNames, dataSource, generateMessage } from '@/utils/utils'
import { useNamespace } from '@/common/context'
import type { ActionType } from '@ant-design/pro-components'
import tableStyles from '@/common/styles/table.module.less'
import commonStyles from '@/common/styles/common.module.less'
import columnsFunc from '@/pages/compute/machine/list/table-columns.tsx'

export default () => {
    const ctrl = useRef<AbortController>()

    const { notification } = App.useApp()

    const { namespace } = useNamespace()

    const [searchFilter, setSearchFilter] = useState<string>("name")
    const [scroll, setScroll] = useState(150 * 11)
    const [selectedRows, setSelectedRows] = useState<CustomResourceDefinition[]>([])

    const actionRef = useRef<ActionType>()

    const [virtualMachine, setVirtualMachine] = useState<Map<string, CustomResourceDefinition>>(new Map<string, CustomResourceDefinition>())
    const [virtualMachineInstance, setVirtualMachineInstance] = useState<Map<string, CustomResourceDefinition>>(new Map<string, CustomResourceDefinition>())
    const [rootDisk, setRootDisk] = useState<Map<string, CustomResourceDefinition>>(new Map<string, CustomResourceDefinition>())
    const [node, setNode] = useState<Map<string, CustomResourceDefinition>>(new Map<string, CustomResourceDefinition>())

    const resource = useRef<ResourceUpdater>()

    const columns = columnsFunc(virtualMachineInstance, rootDisk, node, notification)

    useEffect(() => {
        if (resource.current) return
        resource.current = new ResourceUpdater(virtualMachine, setVirtualMachine, virtualMachineInstance, setVirtualMachineInstance, rootDisk, setRootDisk, node, setNode, notification)
    }, [])

    useEffect(() => {
        actionRef.current?.reload()
    }, [namespace])

    useEffect(() => {
        return () => {
            console.log('Component is unmounting and aborting operation')
            ctrl.current?.abort()
        }
    }, [])

    return (
        <ProTable<CustomResourceDefinition, Params>
            className={classNames(tableStyles["table-padding"], commonStyles["small-scrollbar"])}
            scroll={{ x: scroll }}
            rowSelection={{
                defaultSelectedRowKeys: [],
                onChange: (_, selectedRows) => {
                    setSelectedRows(selectedRows)
                }
            }}
            tableAlertRender={({ selectedRowKeys, onCleanSelected }) => {
                return (
                    <Space size={16}>
                        <span>已选 {selectedRowKeys.length} 项</span>
                        <a onClick={onCleanSelected}>取消选择</a>
                    </Space>
                )
            }}
            tableAlertOptionRender={() => {
                return (
                    <Space size={16}>
                        <a onClick={async () => await batchManageVirtualMachinePowerState(selectedRows, VirtualMachinePowerStateRequest_PowerState.ON, notification)}>批量开机</a>
                        <a onClick={async () =>
                            Modal.confirm({
                                title: "批量重启虚拟机？",
                                content: generateMessage(selectedRows, `即将重启 "{names}" 虚拟机，请确认`, `即将重启 "{names}" 等 {count} 台虚拟机，请确认。`),
                                okText: '确认重启',
                                okType: 'danger',
                                cancelText: '取消',
                                okButtonProps: {
                                    disabled: false,
                                },
                                onOk: async () => {
                                    await batchManageVirtualMachinePowerState(selectedRows, VirtualMachinePowerStateRequest_PowerState.REBOOT, notification)
                                }
                            })
                        }>批量重启</a>
                        <a onClick={async () => {
                            Modal.confirm({
                                title: "批量关闭虚拟机？",
                                content: generateMessage(selectedRows, `即将关闭 "{names}" 虚拟机，请确认`, `即将关闭 "{names}" 等 {count} 台虚拟机，请确认。`),
                                okText: '确认关闭',
                                okType: 'danger',
                                cancelText: '取消',
                                okButtonProps: {
                                    disabled: false,
                                },
                                onOk: async () => {
                                    await batchManageVirtualMachinePowerState(selectedRows, VirtualMachinePowerStateRequest_PowerState.OFF, notification)
                                }
                            })
                        }}>批量关机</a>
                        <a className={commonStyles["warning-color"]} onClick={async () => {
                            Modal.confirm({
                                title: "批量删除虚拟机？",
                                content: generateMessage(selectedRows, `即将删除 "{names}" 虚拟机，请确认`, `即将删除 "{names}" 等 {count} 台虚拟机，请确认。`),
                                okText: '确认删除',
                                okType: 'danger',
                                cancelText: '取消',
                                okButtonProps: {
                                    disabled: false,
                                },
                                onOk: async () => {
                                    await batchDeleteVirtualMachines(selectedRows, notification)
                                }
                            })
                        }}>批量删除</a>
                    </Space>
                )
            }}
            columns={columns}
            actionRef={actionRef}
            loading={{ indicator: <LoadingOutlined /> }}
            dataSource={dataSource(virtualMachine)}
            request={async (params) => {
                ctrl.current?.abort()
                ctrl.current = new AbortController()

                const advancedParams = { namespace: namespace, searchFilter: searchFilter, params: params }
                await resource.current?.updateResource(advancedParams, ctrl.current)
                return { success: true }
            }}
            columnsState={{
                persistenceKey: 'virtual-machine-list-table-columns',
                persistenceType: 'localStorage',
                onChange: (obj) => setScroll(calcScroll(obj))
            }}
            rowKey={(vm) => namespaceName(vm.metadata)}
            search={false}
            options={{
                fullScreen: true,
                search: {
                    allowClear: true,
                    style: { width: 280 },
                    addonBefore: <Select defaultValue="name" onChange={(value) => setSearchFilter(value)} options={[
                        { value: 'name', label: '名称' },
                    ]} />
                }
            }}
            pagination={false}
            toolbar={{
                actions: [
                    <NavLink to='/compute/machines/create'><Button icon={<PlusOutlined />}>创建虚拟机</Button></NavLink>
                ]
            }}
        />
    )
}