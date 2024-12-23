import { request } from '@umijs/max';
import type { Merchant, Blacklist, LoginLog, QueryLog, SystemInfo, Admin } from '@/types';

// 商户相关接口
export async function getMerchantList(params: any) {
  return request<{
    data: Merchant[];
    total: number;
    success: boolean;
  }>('/api/merchant/list', {
    method: 'GET',
    params,
  });
}

export async function addMerchant(data: Partial<Merchant>) {
  return request<{
    success: boolean;
  }>('/api/merchant/add', {
    method: 'POST',
    data,
  });
}

export async function updateMerchant(data: Partial<Merchant>) {
  return request<{
    success: boolean;
  }>('/api/merchant/update', {
    method: 'PUT',
    data,
  });
}

// 黑名单相关接口
export async function getBlacklistList(params: any) {
  return request<{
    data: Blacklist[];
    total: number;
    success: boolean;
  }>('/api/blacklist/list', {
    method: 'GET',
    params,
  });
}

export async function updateBlacklistStatus(id: number, status: number) {
  return request<{
    success: boolean;
  }>('/api/blacklist/status', {
    method: 'PUT',
    data: { id, status },
  });
}

// 系统日志相关接口
export async function getLoginLogs(params: any) {
  return request<{
    data: LoginLog[];
    total: number;
    success: boolean;
  }>('/api/system/logs/login', {
    method: 'GET',
    params,
  });
}

export async function getQueryLogs(params: any) {
  return request<{
    data: QueryLog[];
    total: number;
    success: boolean;
  }>('/api/system/logs/query', {
    method: 'GET',
    params,
  });
}

// 系统监控相关接口
export async function getSystemInfo() {
  return request<{
    data: SystemInfo;
    success: boolean;
  }>('/api/system/monitor', {
    method: 'GET',
  });
}

// 管理员相关接口
export async function getAdminList(params: any) {
  return request<{
    data: Admin[];
    total: number;
    success: boolean;
  }>('/api/system/admin/list', {
    method: 'GET',
    params,
  });
}

export async function addAdmin(data: Partial<Admin>) {
  return request<{
    success: boolean;
  }>('/api/system/admin/add', {
    method: 'POST',
    data,
  });
}

export async function updateAdmin(data: Partial<Admin>) {
  return request<{
    success: boolean;
  }>('/api/system/admin/update', {
    method: 'PUT',
    data,
  });
}

export async function resetAdminPassword(id: number) {
  return request<{
    success: boolean;
  }>('/api/system/admin/reset-password', {
    method: 'PUT',
    data: { id },
  });
} 