import React from 'react'
import type { MenuProps } from 'antd'
import { Flex, Menu } from 'antd'
import { Link, useLocation } from 'react-router-dom'
import { IconFont } from '@/components/icon/index'
import { Header } from 'antd/es/layout/layout'
import styles from '@/components/main-menu/index.module.less'

interface MainMenuProps {
  onSelect?: (e: any) => void
}

type MenuItem = Required<MenuProps>['items'][number]

const items: MenuItem[] = [
  {
    key: '/vink',
    label: (
      <Flex>
        <IconFont type="icon-linux" className={styles.logo} />
        <Link to="/" className={styles['logo-text']}>Vink</Link>
      </Flex>
    ),
  },
  {
    key: '/virtual',
    label: <Link to="/virtual">虚拟资源</Link>,
  },
  {
    key: '/physical',
    label: <Link to="/physical">物理资源</Link>,
  }
]

const VinkMenu: React.FC<MainMenuProps> = ({ onSelect }) => {
  const location = useLocation()
  const currentPath = location.pathname

  const defaultSelectedKey = items.find(item => {
    return currentPath.startsWith(item?.key as string)
  })?.key as string

  return (
    <Header className={styles.header}>
      <Menu
        className={styles.menu}
        mode="horizontal"
        defaultSelectedKeys={[defaultSelectedKey]}
        items={items}
        onSelect={onSelect}
      />
    </Header>
  )
}

export default VinkMenu
