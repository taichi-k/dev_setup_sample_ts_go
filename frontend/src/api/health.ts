import { apiClient } from './client';

export type HealthResponse = {
  status: string;
};

export const fetchHealth = async (): Promise<HealthResponse> => {
  const res = await apiClient.get<HealthResponse>('/api/health');
  return res.data;
};
