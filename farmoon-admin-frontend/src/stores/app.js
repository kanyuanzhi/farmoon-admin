import { defineStore } from "pinia";
import { constantRoutes, asyncRoutes } from "src/router/routes";

export const useAppStore = defineStore("app", {
  state: () => ({
    sidebarOpened: true,
  }),
  getters: {},
  actions: {
    toggleSidebar() {
      this.sidebarOpened = !this.sidebarOpened;
    }
  }
});
