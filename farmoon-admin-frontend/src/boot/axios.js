import { boot } from "quasar/wrappers";
import axios from "axios";
import { Dialog, Notify } from "quasar";
import { useUserStore } from "stores/user";

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({
  baseURL: process.env.API,
  timeout: 5000,
  withCredentials: false
});
export default boot(({
  app,
  router
}) => {
  const userStore = useUserStore();

  api.interceptors.request.use(request => {
    const token = userStore.getToken();
    request.headers = {
      "Content-Type": "application/json;charset=utf-8",
      "Farmoon-Token": token,
    };
    return request;
  }, error => {
    Notify.create({
      type: "negative",
      message: error,
    });
    return Promise.reject(error);
  });

  api.interceptors.response.use(response => {
    // If the ExpiresAt of the JWT has expired,
    // but the RefreshAt has not expired,
    // the background will insert a Gqa Refresh Token in the headers,
    // which will be saved here to form a token replacement logic
    // if (response.headers['farmoon-refresh-token'] && response.data.data.refresh) {
    //   userStore.setToken(response.headers['farmoon-refresh-token'])
    //   // store.dispatch('user/SetToken', response.headers['gqa-refresh-token'])
    //   Notify.create({
    //     type: 'positive',
    //     message: i18n.global.t('Refresh') + 'Token' + i18n.global.t('Success'),
    //   })
    //   return api(response.config)
    // }
    const responseData = response.data;
    const {
      message,
      code
    } = responseData;

    if (code === 0) {
      Notify.create({
        message: message || "Error",
        type: "negative",
        duration: 5 * 1000
      });
      //todo:添加过期重新登陆
      return Promise.reject(new Error(message || "Error"));
    } else {
      return responseData;
    }
  }, error => {
    // 500
    if (error + "" === "Error: Request failed with status code 500") {
      Dialog.create({
        title: "错误",
        message: "数据过期，请重新登录",
        persistent: true,
        ok: {
          push: true,
          color: "negative",
          label: "登出"
        },
      })
        .onOk(() => {
          userStore.handleLogout();
          router.push({ name: "login" });
        });
    }
    // timeout
    if (error + "" === "Error: timeout of 5000ms exceeded") {
      Notify.create({
        type: "negative",
        message: "请求超时",
      });
    }
    // network error
    if (error + "" === "Error: Network Error") {
      router.push({ name: "notFound" });
    } else if (error.response && error.response.status === 404) {
      Notify.create({
        type: "negative",
        message: "请求地址无效" + " " + error.response.request.responseURL,
      });
    }
    return Promise.reject(error);
  });
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { api };
