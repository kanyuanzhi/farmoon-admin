<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn flat dense round icon="menu" aria-label="Menu" @click="appStore.toggleSidebar"/>
      <TheBreadcrumbs/>

      <q-toolbar-title>
      </q-toolbar-title>
      <q-btn flat round icon="logout" aria-label="退出登录" @click="logout"/>
    </q-toolbar>
  </q-header>
</template>

<script setup>
import {useAppStore} from "stores/app";
import {useUserStore} from "stores/user";
import {useRoute, useRouter} from "vue-router";
import TheBreadcrumbs from "layouts/breadcrumbs/TheBreadcrumbs.vue";

const appStore = useAppStore();
const userStore = useUserStore();
const router = useRouter();
const route = useRoute();

const logout = () => {
  userStore.handleLogout();
  router.push(`/login?redirect=${route.fullPath}`);
};
</script>

<style lang="scss" scoped>

</style>
