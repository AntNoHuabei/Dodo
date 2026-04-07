import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import wails from "@wailsio/runtime/plugins/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), wails("./bindings")],
  resolve: {
    alias: {
      "./models.js": path.resolve(__dirname, "./bindings/github.com/wailsapp/wails/v3/pkg/services/sqlite/models.ts"),
      "./stmt.js": path.resolve(__dirname, "./bindings/github.com/wailsapp/wails/v3/pkg/services/sqlite/stmt.js")
    }
  }
});
