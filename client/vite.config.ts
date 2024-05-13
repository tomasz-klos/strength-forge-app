import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default ({ mode }: any) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [react()],

    server: {
      port: 3000,
      proxy: {
        "/api": {
          target: `${process.env.VITE_API_URL}`,
          changeOrigin: true,
          secure: true,
        },
      },
    },
  });
};
