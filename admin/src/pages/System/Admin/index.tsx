import { PageContainer } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { Button } from 'antd';
import { useRef } from 'react';

const AdminPage: React.FC = () => {
  const actionRef = useRef<ActionType>();

  const columns: ProColumns[] = [
    {
      title: '管理员ID',
      dataIndex: 'id',
      search: false,
    },
    {
      title: '用户名',
      dataIndex: 'username',
    },
    {
      title: '姓名',
      dataIndex: 'name',
    },
    {
      title: '角色',
      dataIndex: 'role',
      valueEnum: {
        admin: { text: '超级管理员' },
        operator: { text: '普通管理员' },
      },
    },
    {
      title: '状态',
      dataIndex: 'status',
      valueEnum: {
        0: { text: '禁用', status: 'Error' },
        1: { text: '启用', status: 'Success' },
      },
    },
    {
      title: '最后登录时间',
      dataIndex: 'lastLoginTime',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '操作',
      valueType: 'option',
      render: (_, record) => [
        <a key="edit">编辑</a>,
        <a key="resetPwd">重置密码</a>,
        <a key="disable">禁用</a>,
      ],
    },
  ];

  return (
    <PageContainer>
      <ProTable
        columns={columns}
        actionRef={actionRef}
        cardBordered
        request={async (params = {}) => {
          // 实现管理员列表查询接口
          return {
            data: [],
            success: true,
          };
        }}
        toolBarRender={() => [
          <Button type="primary" key="add">
            新增管理员
          </Button>,
        ]}
      />
    </PageContainer>
  );
};

export default AdminPage; 