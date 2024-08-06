export type Environment = 'DEVELOPMENT' | 'TESTING' | 'PRODUCTION';

export interface Config {
  ENVIRONMENT: Environment;
  ENVOY_PROXY_ADDRESS: string;
}
