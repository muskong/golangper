import { LogoutOutlined } from '@ant-design/icons';
import { useModel, history } from '@umijs/max';
import { Button, message } from 'antd';

const LogoutButton: React.FC = () => {
  const { setInitialState } = useModel('@@initialState');

  const handleLogout = () => {
    localStorage.removeItem('token');
    setInitialState({ currentUser: undefined });
    message.success('已退出登录');
    history.push('/login');
  };

  return (
    <Button
      type="text"
      icon={<LogoutOutlined />}
      onClick={handleLogout}
    >
      退出登录
    </Button>
  );
};

export default LogoutButton; 