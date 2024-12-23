import { PageContainer } from '@ant-design/pro-components';
import { Card, Row, Col, Statistic } from 'antd';
import { useRequest } from '@umijs/max';
import { SystemInfo } from '@/types';

const MonitorPage: React.FC = () => {
  const { data: systemInfo = {} as SystemInfo, loading } = useRequest(() => {
    // 实现系统信息获取接口
    return Promise.resolve({
      cpu: {
        usage: 45.5,
        cores: 8,
      },
      memory: {
        total: 16384,
        used: 8192,
        usage: 50,
      },
      redis: {
        connected: true,
        keys: 1000,
        memory: 512,
      },
      postgresql: {
        connected: true,
        activeConnections: 50,
        dbSize: 1024,
      },
    });
  });

  return (
    <PageContainer>
      <Row gutter={16}>
        <Col span={6}>
          <Card title="CPU">
            <Statistic
              title="使用率"
              value={systemInfo?.cpu.usage}
              suffix="%"
              precision={1}
            />
            <Statistic title="核心数" value={systemInfo?.cpu.cores} />
          </Card>
        </Col>
        <Col span={6}>
          <Card title="内存">
            <Statistic
              title="使用率"
              value={systemInfo?.memory.usage}
              suffix="%"
            />
            <Statistic
              title="已用/总量"
              value={systemInfo?.memory.used}
              suffix={`/ ${systemInfo?.memory.total} MB`}
            />
          </Card>
        </Col>
        <Col span={6}>
          <Card title="Redis">
            <Statistic
              title="连接状态"
              value={systemInfo?.redis.connected ? '正常' : '异常'}
              valueStyle={{
                color: systemInfo?.redis.connected ? '#3f8600' : '#cf1322',
              }}
            />
            <Statistic title="键数量" value={systemInfo?.redis.keys} />
          </Card>
        </Col>
        <Col span={6}>
          <Card title="PostgreSQL">
            <Statistic
              title="连接状态"
              value={systemInfo?.postgresql.connected ? '正常' : '异常'}
              valueStyle={{
                color: systemInfo?.postgresql.connected ? '#3f8600' : '#cf1322',
              }}
            />
            <Statistic
              title="活跃连接数"
              value={systemInfo?.postgresql.activeConnections}
            />
          </Card>
        </Col>
      </Row>
    </PageContainer>
  );
};

export default MonitorPage; 