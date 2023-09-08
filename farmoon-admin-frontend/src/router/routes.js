const MainLayout = () => import("layouts/MainLayout.vue");
const FullLayout = () => import("layouts/FullLayout.vue");

export const constantRoutes = [
  {
    path: "/",
    component: MainLayout,
    children: [{
      path: "",
      name: "index",
      component: () => import("pages/IndexPage.vue"),
    }],
  },
  {
    path: "/login",
    component: FullLayout,
    hidden: true,
    children: [{
      path: "",
      name: "login",
      component: () => import("pages/LoginPage.vue")
    }],
  },
  {
    path: "/register",
    component: FullLayout,
    hidden: true,
    children: [{
      path: "",
      name: "register",
      component: () => import("pages/RegisterPage.vue")
    }],
  },
  {
    path: "/test2",
    component: MainLayout,
    meta: {
      title: "测试",
    },
    children: [{
      path: "admin",
      name: "test2Admin",
      component: () => import("pages/test/AdminPage.vue"),
      meta: {
        title: "admin测试",
        icon: "mdi-home",
      }
    }, {
      path: "editor",
      name: "test2Editor",
      component: () => import("pages/test/EditorPage.vue"),
      meta: {
        title: "editor测试",
        icon: "mdi-home",
      }
    }],
  },
];

export const asyncRoutes = [
  {
    path: "/test",
    component: MainLayout,
    meta: {
      title: "测试",
    },
    children: [{
      path: "admin",
      name: "testAdmin",
      component: () => import("pages/test/AdminPage.vue"),
      meta: {
        title: "admin测试",
        icon: "mdi-home",
        roles: ["admin"]
      }
    }, {
      path: "editor",
      name: "testEditor",
      component: () => import("pages/test/EditorPage.vue"),
      meta: {
        title: "editor测试",
        icon: "mdi-home",
        roles: ["editor"]
      }
    }],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    name: "notFound",
    component: () => import("pages/ErrorNotFound.vue"),
    hidden: true,
  },
];
