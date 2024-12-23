import { ModalForm, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import type { Blacklist } from '@/types';

interface BlacklistFormProps {
  visible: boolean;
  onVisibleChange: (visible: boolean) => void;
  onFinish: (values: Partial<Blacklist>) => Promise<boolean>;
  initialValues?: Partial<Blacklist>;
}

const BlacklistForm: React.FC<BlacklistFormProps> = ({
  visible,
  onVisibleChange,
  onFinish,
  initialValues,
}) => {
  return (
    <ModalForm
      title={initialValues ? '编辑黑名单' : '新增黑名单'}
      visible={visible}
      onVisibleChange={onVisibleChange}
      onFinish={onFinish}
      initialValues={initialValues}
    >
      <ProFormText
        name="name"
        label="姓名"
        rules={[{ required: true, message: '请输入姓名' }]}
      />
      <ProFormText
        name="phone"
        label="手机号"
        rules={[
          { required: true, message: '请输入手机号' },
          { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号' },
        ]}
      />
      <ProFormText
        name="idCard"
        label="身份证号"
        rules={[
          { required: true, message: '请输入身份证号' },
          { pattern: /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/, message: '请输入正确的身份证号' },
        ]}
      />
      <ProFormText
        name="email"
        label="邮箱"
        rules={[
          { type: 'email', message: '请输入正确的邮箱地址' },
        ]}
      />
      <ProFormText
        name="address"
        label="地址"
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

export default BlacklistForm; 