import { useState } from 'react'
import { App, ConfigProvider, Select } from 'antd'
import { PageContainer, ProLayout } from '@ant-design/pro-components'
import { NavLink, useLocation } from 'react-router-dom'
import { LayoutSettings } from '@/layout-config'
import { NamespaceProvider, useNamespace } from '@/common/context'
import { ResourceType } from '@/clients/ts/types/resource_type'
import { useListResources } from './hooks/use-resource'
import AppRouter from '@/router'
import styles from '@/styles/app.module.less'

export default () => {
  const location = useLocation()

  const [pathname, setPathname] = useState(location.pathname)

  const [collapsed, setCollapsed] = useState(false)

  const { resources: namespaces } = useListResources(ResourceType.NAMESPACE)

  return (
    <ConfigProvider
      theme={{
        components: {
          Form: {
            screenXSMax: 0
          }
        }
      }}
    >
      <App className={styles.vink}>
        <NamespaceProvider>
          <ProLayout
            {...LayoutSettings}
            location={{ pathname }}
            collapsed={collapsed}
            onCollapse={setCollapsed}
            menuItemRender={(item, dom) => <NavLink onClick={() => setPathname(item.path || '/')} to={item.path || '/'}>{dom}</NavLink>}
            menuExtraRender={() => {
              if (collapsed) {
                return
              }
              const { namespace, setNamespace } = useNamespace()
              return <Select
                placeholder="工作空间"
                variant="filled"
                value={namespace}
                style={{ width: "100%", marginTop: 8, marginBottom: 8 }}
                onChange={(value) => setNamespace(value)}
                options={[
                  { value: '', label: '全部工作空间' },
                  ...namespaces.map((ns: any) => ({
                    value: ns.metadata?.name,
                    label: ns.metadata?.name
                  }))
                ]}
              />
            }}
          >
            <PageContainer title={false}>
              <AppRouter />
            </PageContainer>
          </ProLayout>
        </NamespaceProvider>
      </App>
    </ConfigProvider>
  )
}
