import { PageContainer } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { Button, message, Modal } from 'antd';
import { useRef, useState } from 'react';
import { getMerchantList, addMerchant, updateMerchant } from '@/services/api';
import type { Merchant } from '@/types';
import MerchantForm from './components/MerchantForm';

const MerchantList: React.FC = () => {
  const actionRef = useRef<ActionType>();
  const [formVisible, setFormVisible] = useState(false);
  const [currentRecord, setCurrentRecord] = useState<Partial<Merchant>>();

  const handleAdd = async (values: Partial<Merchant>) => {
    try {
      const result = await addMerchant(values);
      if (result.success) {
        message.success('添加成功');
        setFormVisible(false);
        actionRef.current?.reload();
        return true;
      }
      return false;
    } catch (error) {
      message.error('添加失败');
      return false;
    }
  };

  const handleEdit = async (values: Partial<Merchant>) => {
    try {
      const result = await updateMerchant({ ...values, id: currentRecord?.id });
      if (result.success) {
        message.success('更新成功');
        setFormVisible(false);
        actionRef.current?.reload();
        return true;
      }
      return false;
    } catch (error) {
      message.error('更新失败');
      return false;
    }
  };

  const handleDelete = (record: Merchant) => {
    Modal.confirm({
      title: '确认删除',
      content: `确定要删除商户"${record.name}"吗？`,
      onOk: async () => {
        try {
          const result = await updateMerchant({ id: record.id, status: -1 });
          if (result.success) {
            message.success('删除成功');
            actionRef.current?.reload();
          }
        } catch (error) {
          message.error('删除失败');
        }
      },
    });
  };

  const columns: ProColumns<Merchant>[] = [
    {
      title: '商户ID',
      dataIndex: 'id',
      search: false,
    },
    {
      title: '商户名称',
      dataIndex: 'name',
    },
    {
      title: '联系人',
      dataIndex: 'contact',
    },
    {
      title: '联系电话',
      dataIndex: 'phone',
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
      title: '创建时间',
      dataIndex: 'createTime',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '操作',
      valueType: 'option',
      render: (_, record) => [
        <a
          key="edit"
          onClick={() => {
            setCurrentRecord(record);
            setFormVisible(true);
          }}
        >
          编辑
        </a>,
        <a
          key="delete"
          onClick={() => handleDelete(record)}
        >
          删除
        </a>,
      ],
    },
  ];

  return (
    <PageContainer>
      <ProTable<Merchant>
        columns={columns}
        actionRef={actionRef}
        cardBordered
        request={async (params) => {
          const response = await getMerchantList(params);
          return {
            data: response.data,
            success: response.success,
            total: response.total,
          };
        }}
        toolBarRender={() => [
          <Button
            type="primary"
            key="add"
            onClick={() => {
              setCurrentRecord(undefined);
              setFormVisible(true);
            }}
          >
            新增商户
          </Button>,
        ]}
      />
      <MerchantForm
        visible={formVisible}
        onVisibleChange={setFormVisible}
        onFinish={currentRecord ? handleEdit : handleAdd}
        initialValues={currentRecord}
      />
    </PageContainer>
  );
};

export default MerchantList; 