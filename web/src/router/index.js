import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
 roles: ['admin','editor']    control the page roles (you can set multiple roles)
 title: 'title'               the name show in sidebar and breadcrumb (recommend set)
 icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
 breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
 activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
 }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/workbench',
    children: [{
      path: 'workbench',
      name: 'Workbench',
      component: () => import('@/views/workbench/index'),
      meta: { title: '工作台', icon: 'workbench' }
    }]
  },

  {
    path: '/dataSource',
    component: Layout,
    meta: { title: '数据源管理', icon: 'dataSource' },
    redirect: '/dataSource/manage',
    children: [{
      path: '/dataSource/manage',
      name: 'DataSource',
      component: () => import('@/views/dataSource/index'),
      meta: { title: '数据源' }
    }, {
      path: '/dataSource/manage/edit',
      name: 'DataSourceEdit',
      hidden: true,
      component: () => import('@/views/dataSource/edit'),
      meta: { title: '编辑数据源' }
    }]
  },
  {
    path: '/dataReplication',
    component: Layout,
    redirect: '/dataReplication/task',
    meta: { title: '数据复制', icon: 'data_exchange' },
    children: [{
      path: 'task',
      name: 'DataReplication',
      component: () => import('@/views/dataReplication/index'),
      meta: { title: '数据复制' }
    }, {
      path: 'task/edit',
      name: 'DataReplicationEdit',
      hidden: true, // 不在侧边栏显示
      component: () => import('@/views/dataReplication/replication/index'),
      meta: { title: '编辑', icon: 'dashboard' }
    }]
  },
  {
    path: '/backup',
    component: Layout,
    redirect: '/backup/task',
    meta: { title: '备份与恢复', icon: 'data_exchange' },
    children: [{
      path: 'task',
      name: 'DataReplication',
      component: () => import('@/views/dataReplication/index'),
      meta: { title: '数据备份' }
    }, {
      path: '/backup/task/edit',
      name: 'DataReplicationEdit',
      hidden: true, // 不在侧边栏显示
      component: () => import('@/views/dataReplication/replication/index'),
      meta: { title: '编辑', icon: 'dashboard' }
    }, {
      path: '/backup/host/list',
      name: 'BackupHostList',
      component: () => import('@/views/backupAndRecovery/backupHost/index'),
      meta: { title: '备份主机' }
    }, {
      path: '/backup/host/manage/edit',
      name: 'BackupHostEdit',
      hidden: true,
      component: () => import('@/views/backupAndRecovery/backupHost/backupHostEdit'),
      meta: { title: '编辑备份主机' }
    }]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
