import { RequestConfig } from '@umijs/max';
import { message } from 'antd';

// 错误处理方案： 错误类型
enum ErrorShowType {
  SILENT = 0,
  WARN_MESSAGE = 1,
  ERROR_MESSAGE = 2,
  NOTIFICATION = 3,
  REDIRECT = 9,
}

// 与后端约定的响应数据格式
interface ResponseStructure {
  success: boolean;
  data: any;
  errorCode?: number;
  errorMessage?: string;
  showType?: ErrorShowType;
}

export const request: RequestConfig = {
  timeout: 10000,
  errorConfig: {
    errorHandler: (error: any) => {
      if (error.response) {
        // 响应错误
        message.error(`请求错误 ${error.response.status}: ${error.response.data.message}`);
      } else {
        // 请求错误
        message.error(`请求错误: ${error.message}`);
      }
    },
  },
  requestInterceptors: [
    (url, options) => {
      // 在请求头中添加token
      const token = localStorage.getItem('token');
      if (token) {
        const headers = {
          ...options.headers,
          Authorization: `Bearer ${token}`,
        };
        return {
          url,
          options: { ...options, headers },
        };
      }
      return { url, options };
    },
  ],
  responseInterceptors: [
    (response) => {
      const { data } = response as unknown as { data: ResponseStructure };
      if (!data.success) {
        message.error(data.errorMessage || '请求失败');
      }
      return response;
    },
  ],
}; 