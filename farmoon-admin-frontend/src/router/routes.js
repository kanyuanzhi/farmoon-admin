const MainLayout = () => import('layouts/MainLayout.vue')
const FullLayout = () => import('layouts/FullLayout.vue')

export const constantRoutes = [
  {
    path: '/',
    component: MainLayout,
    redirect: '/home',
    children: [
      {
        path: 'home',
        name: 'home',
        component: () => import('pages/IndexPage.vue'),
        meta: {
          title: '主页',
          icon: 'home',
        },
      },
    ],
  },
  {
    path: '/',
    component: FullLayout,
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import('pages/LoginPage.vue'),
      },
      {
        path: 'register',
        name: 'register',
        component: () => import('pages/RegisterPage.vue'),
      },
    ],
  },
  {
    path: '/level1',
    name: 'level1',
    meta: {
      title: 'level1',
      icon: 'home',
    },
    redirect: '/level1/level1-1',
    component: MainLayout,
    children: [
      {
        path: 'level1-1',
        name: 'level1-1',
        redirect: '/level1/level1-1/level1-1-1',
        component: () => import('pages/test/AdminPage.vue'),
        meta: {
          title: 'level1-1',
          icon: 'mdi-home',
        },
        children: [
          {
            path: 'level1-1-1',
            name: 'level1-1-1',
            component: () => import('pages/test/AdminPage.vue'),
            meta: {
              title: 'level1-1-1',
              icon: 'mdi-home',
            },
          },
          {
            path: 'level1-1-2',
            name: 'level1-1-2',
            component: () => import('pages/test/AdminPage.vue'),
            meta: {
              title: 'level1-1-2',
              icon: 'mdi-home',
            },
          },
        ],
      }, {
        path: 'level1-2',
        name: 'level1-2',
        component: () => import('pages/test/EditorPage.vue'),
        meta: {
          title: 'level1-2',
          icon: 'mdi-home',
        },
      },
    ],
  },
]

export const asyncRoutes = [
  {
    path: '/user',
    component: MainLayout,
    children: [
      {
        path: 'manage',
        name: 'userManager',
        component: () => import('pages/user/IndexPage.vue'),
        meta: {
          title: '用户管理',
          icon: 'mdi-account-group',
          roles: ['admin'],
        },
      }, {
        path: 'profile',
        name: 'userProfile',
        hidden: true,
        component: () => import('pages/user/profile/IndexPage.vue'),
        meta: {
          title: '我的资料',
          icon: 'o_account_circle',
          roles: ['admin', 'editor', 'visitor'],
        },
      },
    ],
  },
  {
    path: '/dish',
    name: 'dish',
    meta: {
      title: '菜品管理',
      icon: 'format_list_bulleted',
    },
    redirect: '/dish/display',
    component: MainLayout,
    children: [
      {
        path: 'display',
        name: 'dishDisplay',
        component: () => import('pages/dish/IndexPage.vue'),
        meta: {
          title: '菜品浏览',
          icon: 'format_list_bulleted',
          roles: ['admin', 'editor'],
        },
      }, {
        path: 'edit',
        name: 'dishEdit',
        component: () => import('pages/dish/edit/IndexPage.vue'),
        meta: {
          title: '菜品制作',
          icon: 'edit',
          roles: ['admin', 'editor'],
        },
      }, {
        path: 'cuisine',
        name: 'dishCuisine',
        component: () => import('pages/dish/cuisine/IndexPage.vue'),
        meta: {
          title: '菜系管理',
          icon: 'format_list_bulleted',
          roles: ['admin', 'editor'],
        },
      }, {
        path: 'seasoning',
        name: 'dishSeasoning',
        component: () => import('pages/dish/seasoning/IndexPage.vue'),
        meta: {
          title: '调料及料泵管理',
          icon: 'o_heat_pump',
          roles: ['admin', 'editor'],
        },
      }, {
        path: 'ingredient',
        name: 'dishIngredient',
        redirect: '/dish/ingredient/ingredient',
        meta: {
          title: '食材管理',
          icon: 'format_list_bulleted',
          roles: ['admin', 'editor'],
        },
        children: [
          {
            path: 'ingredient',
            name: 'dishIngredientIngredient',
            component: () => import('pages/dish/ingredient/IngredientPage.vue'),
            meta: {
              title: '食材管理',
              icon: 'format_list_bulleted',
              roles: ['admin', 'editor'],
            },
          }, {
            path: 'type',
            name: 'dishIngredientType',
            component: () => import('pages/dish/ingredient/IngredientTypePage.vue'),
            meta: {
              title: '类别管理',
              icon: 'format_list_numbered',
              roles: ['admin', 'editor'],
            },
          }, {
            path: 'shape',
            name: 'dishIngredientShape',
            component: () => import('pages/dish/ingredient/IngredientShapePage.vue'),
            meta: {
              title: '形状管理',
              icon: 'o_category',
              roles: ['admin', 'editor'],
            },
          },
        ],
      },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    name: 'notFound',
    component: () => import('pages/ErrorNotFound.vue'),
    hidden: true,
  },
]
