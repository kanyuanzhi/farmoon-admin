import {defineStore} from "pinia";
import {SessionStorage, LocalStorage} from "quasar";
import {postAPI} from "src/api";
import {usePermissionStore} from "stores/permission";

export const useUserStore = defineStore("user", {
  state: () => ({
    token: undefined,
    id: undefined,
    username: undefined,
    nickname: undefined,
    realName: undefined,
    avatar: undefined,
    gender: undefined,
    mobile: undefined,
    email: undefined,
    roles: undefined,
    rememberMe: true,
  }),
  getters: {},
  actions: {
    async handleLogin(loginForm) {
      const res = await postAPI("/public/login", loginForm);
      if (res.code === 1) {
        const {
          id,
          token,
          username,
          nickname,
          realName,
          avatar,
          gender,
          mobile,
          email,
          roles
        } = res.data;
        this.setToken(token);

        this.id = id;
        LocalStorage.set("farmoon-id", id);
        this.username = username;
        LocalStorage.set("farmoon-username", username);
        this.nickname = nickname;
        LocalStorage.set("farmoon-nickname", nickname);
        this.realName = realName;
        LocalStorage.set("farmoon-realName", realName);
        this.avatar = avatar;
        LocalStorage.set("farmoon-avatar", avatar);
        this.gender = gender;
        LocalStorage.set("farmoon-gender", gender);
        this.mobile = mobile;
        LocalStorage.set("farmoon-mobile", mobile);
        this.email = email;
        LocalStorage.set("farmoon-email", email);
        this.roles = roles;
        LocalStorage.set("farmoon-roles", roles);

        return true;
      } else {
        return false;
      }
    },
    setToken(token) {
      this.token = token;
      if (this.rememberMe) {
        LocalStorage.set("farmoon-token", token);
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
      LocalStorage.remove("farmoon-token");
      LocalStorage.remove("farmoon-userInfo");
      // dont delete dict
      // LocalStorage.remove('farmoon-dict')
      this.token = undefined;
      this.id = undefined;
      this.username = undefined;
      this.nickname = undefined;
      this.realName = undefined;
      this.avatar = undefined;
      this.gender = undefined;
      this.mobile = undefined;
      this.email = undefined;
      this.roles = undefined;
    },
    getToken() {
      if (SessionStorage.getItem("farmoon-token")) {
        return SessionStorage.getItem("farmoon-token");
      } else if (LocalStorage.getItem("farmoon-token")) {
        return LocalStorage.getItem("farmoon-token");
      } else {
        return this.token;
      }
    },
    getId() {
      if (this.id) {
        return this.id;
      } else if (LocalStorage.getItem("farmoon-id")) {
        return LocalStorage.getItem("farmoon-id");
      } else {
        return 0
      }
    },
    getUsername() {
      if (this.username) {
        return this.username;
      } else if (LocalStorage.getItem("farmoon-username")) {
        return LocalStorage.getItem("farmoon-username");
      } else {
        return ""
      }
    },
    getNickname() {
      if (this.nickname) {
        return this.nickname;
      } else if (LocalStorage.getItem("farmoon-nickname")) {
        return LocalStorage.getItem("farmoon-nickname");
      } else {
        return ""
      }
    },
    getRealName() {
      if (this.realName) {
        return this.realName;
      } else if (LocalStorage.getItem("farmoon-realName")) {
        return LocalStorage.getItem("farmoon-realName");
      } else {
        return ""
      }
    },
    getAvatar() {
      if (this.avatar) {
        return this.avatar;
      } else if (LocalStorage.getItem("farmoon-avatar")) {
        return LocalStorage.getItem("farmoon-avatar");
      } else {
        return ""
      }
    },
    getGender() {
      if (this.gender) {
        return this.gender;
      } else if (LocalStorage.getItem("farmoon-gender")) {
        return LocalStorage.getItem("farmoon-gender");
      } else {
        return ""
      }
    },
    getMobile() {
      if (this.mobile) {
        return this.mobile;
      } else if (LocalStorage.getItem("farmoon-mobile")) {
        return LocalStorage.getItem("farmoon-mobile");
      } else {
        return ""
      }
    },
    getEmail() {
      if (this.email) {
        return this.email;
      } else if (LocalStorage.getItem("farmoon-email")) {
        return LocalStorage.getItem("farmoon-email");
      } else {
        return ""
      }
    },
    getRoles() {
      if (this.roles) {
        return this.roles;
      } else if (LocalStorage.getItem("farmoon-roles")) {
        return LocalStorage.getItem("farmoon-roles");
      } else {
        return []
      }
    },
    setId(id) {
      this.id = id;
      LocalStorage.set("farmoon-id", id)
    },
    setUsername(username) {
      this.username = username;
      LocalStorage.set("farmoon-username", username)
    },
    setNickname(nickname) {
      this.nickname = nickname;
      LocalStorage.set("farmoon-nickname", nickname)
    },
    setRealName(realName) {
      this.realName = realName;
      LocalStorage.set("farmoon-realName", realName)
    },
    setAvatar(avatar) {
      this.avatar = avatar;
      LocalStorage.set("farmoon-avatar", avatar)
    },
    setGender(gender) {
      this.gender = gender;
      LocalStorage.set("farmoon-gender", gender)
    },
    setMobile(mobile) {
      this.mobile = mobile;
      LocalStorage.set("farmoon-mobile", mobile)
    },
    setEmail(email) {
      this.email = email;
      LocalStorage.set("farmoon-email", email)
    },
    setRoles(roles) {
      this.roles = roles;
      LocalStorage.set("farmoon-roles", roles)
    }
  },
});
