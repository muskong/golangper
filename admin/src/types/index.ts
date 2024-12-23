// 商户相关类型
export interface Merchant {
  id: number;
  name: string;
  address: string;
  contact: string;
  phone: string;
  remark: string;
  status: number;
  createTime: string;
  updateTime: string;
  deleteTime?: string;
  ipWhitelist: string[];
  apiKey: string;
  apiSecret: string;
  apiToken: string;
  apiTokenExpireTime: string;
}

// 黑名单相关类型
export interface Blacklist {
  id: number;
  name: string;
  phone: string;
  idCard: string;
  email: string;
  address: string;
  remark: string;
  status: number;
  createTime: string;
  updateTime: string;
}

// 系统日志相关类型
export interface LoginLog {
  id: number;
  merchantId: number;
  ip: string;
  userAgent: string;
  loginTime: string;
  status: number;
}

export interface QueryLog {
  id: number;
  merchantId: number;
  queryContent: string;
  ip: string;
  queryTime: string;
  userAgent: string;
  result: number;
}

// 系统监控相关类型
export interface SystemInfo {
  cpu: {
    usage: number;
    cores: number;
  };
  memory: {
    total: number;
    used: number;
    usage: number;
  };
  redis: {
    connected: boolean;
    keys: number;
    memory: number;
  };
  postgresql: {
    connected: boolean;
    activeConnections: number;
    dbSize: number;
  };
}

// 管理员相关类型
export interface Admin {
  id: number;
  username: string;
  name: string;
  role: 'admin' | 'operator';
  status: number;
  lastLoginTime: string;
} 