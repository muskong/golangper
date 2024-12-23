import { PageContainer } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { useRef } from 'react';

const BlacklistPage: React.FC = () => {
  const actionRef = useRef<ActionType>();

  const columns: ProColumns[] = [
    {
      title: '用户ID',
      dataIndex: 'id',
      search: false,
    },
    {
      title: '姓名',
      dataIndex: 'name',
    },
    {
      title: '手机号',
      dataIndex: 'phone',
    },
    {
      title: '身份证号',
      dataIndex: 'idCard',
      search: true,
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      search: false,
    },
    {
      title: '状态',
      dataIndex: 'status',
      valueEnum: {
        0: { text: '待审核', status: 'Warning' },
        1: { text: '已通过', status: 'Success' },
        2: { text: '已拒绝', status: 'Error' },
      },
    },
    {
      title: '操作',
      valueType: 'option',
      render: (_, record) => [
        <a key="approve">通过</a>,
        <a key="reject">拒绝</a>,
        <a key="delete">删除</a>,
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
          // 这里需要实现接口调用
          return {
            data: [],
            success: true,
          };
        }}
      />
    </PageContainer>
  );
};

export default BlacklistPage; 