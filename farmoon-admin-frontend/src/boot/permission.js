import { useUserStore } from "stores/user";
import { usePermissionStore } from "stores/permission";
import { boot } from "quasar/wrappers";

const whiteList = ["/login", "/register"];

export default boot(({ router }) => {
  router.beforeEach(async (to, from, next) => {
    const userStore = useUserStore();
    const token = userStore.getToken();
    const permissionStore = usePermissionStore();
    if (token) {
      if (whiteList.indexOf(to.path) !== -1) {
        next({ path: "/" });
      } else {
        if (permissionStore.routes.length === 0) {
          const accessRoutes = await permissionStore.generateRoutes(userStore.getRoles());
          accessRoutes.forEach(route => {
            router.addRoute(route);
          });
          next({
            ...to,
            replace: true
          });
        } else {
          next();
          // userStore.handleLogout();
          // next({
          //   path: "/",
          //   replace: true
          // });
        }
      }
    } else {
      if (whiteList.indexOf(to.path) !== -1) {
        // in the free login whitelist, go directly
        next();
      } else {
        // other pages that do not have permission to access are redirected to the login page.
        next(`/login?redirect=${to.path}`);
      }
    }
  });
});
