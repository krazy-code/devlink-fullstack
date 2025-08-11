const viteEnv = import.meta.env;

export const siteConfig = {
  apiBaseUrl: viteEnv.VITE_API_HOST + viteEnv.VITE_API_PREFIX,
  port: viteEnv.VITE_PORT,
};
