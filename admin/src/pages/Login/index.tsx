import { LoginForm, ProFormText } from '@ant-design/pro-components';
import { message } from 'antd';
import { useModel, history } from '@umijs/max';
import { LockOutlined, UserOutlined } from '@ant-design/icons';

const LoginPage: React.FC = () => {
  const { initialState, setInitialState } = useModel('@@initialState');

  const handleSubmit = async (values: { username: string; password: string }) => {
    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });
      const result = await response.json();
      
      if (result.success) {
        localStorage.setItem('token', result.data.token);
        await setInitialState({
          ...initialState,
          currentUser: result.data.user,
        });
        message.success('登录成功！');
        history.push('/');
      } else {
        message.error(result.errorMessage || '登录失败');
      }
    } catch (error) {
      message.error('登录失败，请重试！');
    }
  };

  return (
    <div style={{ 
      height: '100vh',
      background: '#f0f2f5',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center'
    }}>
      <LoginForm
        title="黑名单管理系统"
        subTitle="管理员登录"
        onFinish={async (values) => {
          await handleSubmit(values as { username: string; password: string });
        }}
      >
        <ProFormText
          name="username"
          fieldProps={{
            size: 'large',
            prefix: <UserOutlined />,
          }}
          placeholder="用户名"
          rules={[
            {
              required: true,
              message: '请输入用户名!',
            },
          ]}
        />
        <ProFormText.Password
          name="password"
          fieldProps={{
            size: 'large',
            prefix: <LockOutlined />,
          }}
          placeholder="密码"
          rules={[
            {
              required: true,
              message: '请输入密码！',
            },
          ]}
        />
      </LoginForm>
    </div>
  );
};

export default LoginPage; 