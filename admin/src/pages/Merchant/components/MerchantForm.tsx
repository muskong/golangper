import { ModalForm, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import type { Merchant } from '@/types';

interface MerchantFormProps {
  visible: boolean;
  onVisibleChange: (visible: boolean) => void;
  onFinish: (values: Partial<Merchant>) => Promise<boolean>;
  initialValues?: Partial<Merchant>;
}

const MerchantForm: React.FC<MerchantFormProps> = ({
  visible,
  onVisibleChange,
  onFinish,
  initialValues,
}) => {
  return (
    <ModalForm
      title={initialValues ? '编辑商户' : '新增商户'}
      visible={visible}
      onVisibleChange={onVisibleChange}
      onFinish={onFinish}
      initialValues={initialValues}
    >
      <ProFormText
        name="name"
        label="商户名称"
        rules={[{ required: true, message: '请输入商户名称' }]}
      />
      <ProFormText
        name="contact"
        label="联系人"
        rules={[{ required: true, message: '请输入联系人' }]}
      />
      <ProFormText
        name="phone"
        label="联系电话"
        rules={[
          { required: true, message: '请输入联系电话' },
          { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号' },
        ]}
      />
      <ProFormText
        name="address"
        label="商户地址"
        rules={[{ required: true, message: '请输入商户地址' }]}
      />
      <ProFormTextArea
        name="ipWhitelist"
        label="IP白名单"
        placeholder="每行一个IP地址"
        rules={[{ required: true, message: '请输入IP白名单' }]}
        fieldProps={{
          autoSize: { minRows: 2, maxRows: 6 },
        }}
      />
      <ProFormTextArea
        name="remark"
        label="备注"
        fieldProps={{
          autoSize: { minRows: 2, maxRows: 6 },
        }}
      />
    </ModalForm>
  );
};

export default MerchantForm; 