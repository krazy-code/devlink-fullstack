const viteEnv = import.meta.env;

export const siteConfig = {
  apiBaseUrl: viteEnv.VITE_API_HOST + viteEnv.VITE_API_PREFIX,
  host: viteEnv.VITE_API_HOST,
  port: viteEnv.VITE_PORT,
};
