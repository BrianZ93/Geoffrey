export class ServerConfig {
  api_url: string;

  constructor(api_url: string) {
    this.api_url = api_url;
  }
}

export const getServerConfig = () => {
  const server_config = new ServerConfig('http://localhost:8080');
  return server_config;
};
