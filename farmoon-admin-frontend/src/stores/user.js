import { defineStore } from "pinia";
import { Cookies, SessionStorage } from "quasar";
import { postAPI } from "src/api";
import { usePermissionStore } from "stores/permission";

export const useUserStore = defineStore("user", {
  state: () => ({
    token: undefined,
    username: undefined,
    nickname: undefined,
    realName: undefined,
    avatar: undefined,
    roles: undefined,
    rememberMe: true,
  }),
  getters: {},
  actions: {
    async handleLogin(loginForm) {
      const res = await postAPI("/public/login", loginForm);
      if (res.code === 1) {
        const {
          token,
          username,
          nickname,
          realName,
          avatar,
          roles
        } = res.data;
        this.setToken(token);
        this.username = username;
        Cookies.set("farmoon-username", username);
        this.nickname = nickname;
        Cookies.set("farmoon-nickname", nickname);
        this.realName = realName;
        Cookies.set("farmoon-realName", realName);
        this.avatar = avatar;
        Cookies.set("farmoon-avatar", avatar);
        this.roles = roles;
        Cookies.set("farmoon-roles", roles.join(","));
        return true;
      } else {
        return false;
      }
    },
    setToken(token) {
      this.token = token;
      if (this.rememberMe) {
        Cookies.set("farmoon-token", token);
      } else {
        SessionStorage.set("farmoon-token", token);
      }
    },
    changeRememberMe(type) {
      this.rememberMe = type;
    },
    handleLogout() {
      const permissionStore = usePermissionStore();
      permissionStore.clearRoutes();
      SessionStorage.remove("farmoon-token");
      Cookies.remove("farmoon-token");
      Cookies.remove("farmoon-username");
      Cookies.remove("farmoon-nickname");
      Cookies.remove("farmoon-realName");
      Cookies.remove("farmoon-avatar");
      Cookies.remove("farmoon-roles");
      // dont delete dict
      // LocalStorage.remove('farmoon-dict')
      this.token = undefined;
      this.username = undefined;
      this.nickname = undefined;
      this.realName = undefined;
      this.avatar = undefined;
      this.roles = undefined;
    },
    getToken() {
      if (SessionStorage.getItem("farmoon-token")) {
        return SessionStorage.getItem("farmoon-token");
      } else if (Cookies.get("farmoon-token")) {
        return Cookies.get("farmoon-token");
      } else {
        return this.token;
      }
    },
    getUsername() {
      if (this.username) {
        return this.username;
      } else if (Cookies.get("farmoon-username")) {
        return Cookies.get("farmoon-username");
      } else {
        return "";
      }
    },
    getNickname() {
      if (this.nickname) {
        return this.nickname;
      } else if (Cookies.get("farmoon-nickname")) {
        return Cookies.get("farmoon-nickname");
      } else {
        return "";
      }
    },
    getRealName() {
      if (this.realName) {
        return this.realName;
      } else if (Cookies.get("farmoon-realName")) {
        return Cookies.get("farmoon-realName");
      } else {
        return "";
      }
    },
    getAvatar() {
      if (this.avatar) {
        return this.avatar;
      } else if (Cookies.get("farmoon-avatar")) {
        return Cookies.get("farmoon-avatar");
      } else {
        return "";
      }
    },
    getRoles() {
      if (this.roles) {
        return this.roles;
      } else if (Cookies.get("farmoon-roles")) {
        return Cookies.get("farmoon-roles")
          .split(",");
      } else {
        return [];
      }
    },
  },
});
