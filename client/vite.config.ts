import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

export default ({ mode }: any) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [react()],
    resolve: {
      alias: {
        "@assets": path.resolve(__dirname, "./src/assets"),
        "@atoms": path.resolve(__dirname, "./src/components/atoms"),
        "@molecules": path.resolve(__dirname, "./src/components/molecules"),
        "@organisms": path.resolve(__dirname, "./src/components/organisms"),
        "@templates": path.resolve(__dirname, "./src/components/templates"),
        "@pages": path.resolve(__dirname, "./src/pages"),
        "@config": path.resolve(__dirname, "./src/config"),
        "@lib": path.resolve(__dirname, "./src/lib"),
        "@providers": path.resolve(__dirname, "./src/providers"),
        "@hooks": path.resolve(__dirname, "./src/hooks"),
        "@schemas": path.resolve(__dirname, "./src/schemas"),
        "@services": path.resolve(__dirname, "./src/services"),
        "@types": path.resolve(__dirname, "./src/types"),
      },
    },
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
