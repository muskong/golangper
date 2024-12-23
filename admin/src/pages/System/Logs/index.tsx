import { PageContainer } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { Tabs } from 'antd';
import { useRef } from 'react';

const LogsPage: React.FC = () => {
  const actionRef = useRef<ActionType>();

  const loginColumns: ProColumns[] = [
    {
      title: '商户ID',
      dataIndex: 'merchantId',
    },
    {
      title: '登录IP',
      dataIndex: 'ip',
    },
    {
      title: '用户代理',
      dataIndex: 'userAgent',
      ellipsis: true,
    },
    {
      title: '登录时间',
      dataIndex: 'loginTime',
      valueType: 'dateTime',
    },
    {
      title: '登录状态',
      dataIndex: 'status',
      valueEnum: {
        0: { text: '失败', status: 'Error' },
        1: { text: '成功', status: 'Success' },
      },
    },
  ];

  const queryColumns: ProColumns[] = [
    {
      title: '商户ID',
      dataIndex: 'merchantId',
    },
    {
      title: '查询内容',
      dataIndex: 'queryContent',
    },
    {
      title: '查询IP',
      dataIndex: 'ip',
    },
    {
      title: '查询时间',
      dataIndex: 'queryTime',
      valueType: 'dateTime',
    },
    {
      title: '查询结果',
      dataIndex: 'result',
      valueEnum: {
        0: { text: '未命中', status: 'Default' },
        1: { text: '命中', status: 'Warning' },
      },
    },
  ];

  return (
    <PageContainer>
      <Tabs
        items={[
          {
            label: '登录日志',
            key: 'login',
            children: (
              <ProTable
                columns={loginColumns}
                actionRef={actionRef}
                cardBordered
                request={async (params = {}) => {
                  // 实现登录日志查询接口
                  return {
                    data: [],
                    success: true,
                  };
                }}
              />
            ),
          },
          {
            label: '查询日志',
            key: 'query',
            children: (
              <ProTable
                columns={queryColumns}
                actionRef={actionRef}
                cardBordered
                request={async (params = {}) => {
                  // 实现查询日志接口
                  return {
                    data: [],
                    success: true,
                  };
                }}
              />
            ),
          },
        ]}
      />
    </PageContainer>
  );
};

export default LogsPage; 