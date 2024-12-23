import { defineConfig } from '@umijs/max';

export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: '黑名单管理系统',
  },
  routes: [
    {
      path: '/login',
      component: './Login',
      layout: false,
    },
    {
      path: '/',
      redirect: '/merchant',
    },
    {
      name: '商户管理',
      path: '/merchant',
      component: './Merchant',
      icon: 'ShopOutlined',
      access: 'canAdmin',
    },
    {
      name: '黑名单管理',
      path: '/blacklist',
      component: './Blacklist',
      icon: 'StopOutlined',
      access: 'canAdmin',
    },
    {
      name: '系统管理',
      path: '/system',
      icon: 'SettingOutlined',
      access: 'isAdmin',
      routes: [
        {
          name: '系统日志',
          path: '/system/logs',
          component: './System/Logs',
        },
        {
          name: '系统监控',
          path: '/system/monitor',
          component: './System/Monitor',
        },
        {
          name: '管理员',
          path: '/system/admin',
          component: './System/Admin',
        },
      ],
    },
  ],
  npmClient: 'pnpm',
}); 