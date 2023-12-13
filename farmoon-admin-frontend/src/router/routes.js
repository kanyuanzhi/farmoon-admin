const MainLayout = () => import("layouts/MainLayout.vue");
const FullLayout = () => import("layouts/FullLayout.vue");

export const constantRoutes = [
  {
    path: "/",
    component: MainLayout,
    redirect: "/home",
    children: [
      {
        path: "home",
        name: "home",
        component: () => import("pages/IndexPage.vue"),
        meta: {
          title: "主页",
          icon: "home",
        },
      },
    ],
  },
  {
    path: "/",
    component: FullLayout,
    hidden: true,
    children: [
      {
        path: "login",
        name: "login",
        component: () => import("pages/LoginPage.vue"),
      },
      {
        path: "register",
        name: "register",
        component: () => import("pages/RegisterPage.vue"),
      },
    ],
  },
  {
    path: "/level1",
    name: "level1",
    meta: {
      title: "level1",
      icon: "home",
    },
    redirect: "/level1/level1-1",
    component: MainLayout,
    children: [
      {
        path: "level1-1",
        name: "level1-1",
        redirect: "/level1/level1-1/level1-1-1",
        component: () => import("pages/test/AdminPage.vue"),
        meta: {
          title: "level1-1",
          icon: "mdi-home",
        },
        children: [
          {
            path: "level1-1-1",
            name: "level1-1-1",
            component: () => import("pages/test/AdminPage.vue"),
            meta: {
              title: "level1-1-1",
              icon: "mdi-home",
            },
          },
          {
            path: "level1-1-2",
            name: "level1-1-2",
            component: () => import("pages/test/AdminPage.vue"),
            meta: {
              title: "level1-1-2",
              icon: "mdi-home",
            },
          },
        ],
      }, {
        path: "level1-2",
        name: "level1-2",
        component: () => import("pages/test/EditorPage.vue"),
        meta: {
          title: "level1-2",
          icon: "mdi-home",
        },
      },
    ],
  },
  {
    path: "/map",
    component: MainLayout,
    children: [
      {
        path: "index",
        name: "map",
        component: () => import("pages/map/IndexPage.vue"),
        meta: {
          title: "map",
          icon: "mdi-map",
        },
      },
    ],
  },
];
export const asyncRoutes = [
  {
    path: "/user",
    component: MainLayout,
    children: [
      {
        path: "manage",
        name: "userManager",
        component: () => import("pages/user/IndexPage.vue"),
        meta: {
          title: "用户管理",
          icon: "mdi-account-group",
          roles: ["admin"],
        },
      }, {
        path: "profile",
        name: "userProfile",
        hidden: true,
        component: () => import("pages/user/profile/IndexPage.vue"),
        meta: {
          title: "我的资料",
          icon: "o_account_circle",
          roles: ["admin", "editor", "visitor"],
        },
      },
    ],
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
