<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn flat dense round icon="menu" aria-label="Menu" @click="appStore.toggleSidebar"/>
      <TheBreadcrumbs/>

      <q-toolbar-title>
      </q-toolbar-title>
      <q-btn flat round style="min-width: 0;min-height: 0">
        <q-avatar size="md">
          <img :src="userStore.getAvatar()" alt=""/>
        </q-avatar>
        <q-menu auto-close style="font-size: 13px">
          <q-list style="width: 150px;">
            <q-item>
              <q-item-section>{{ userStore.getUsername() }}</q-item-section>
            </q-item>
            <q-separator/>
            <q-item clickable to="/user/profile">
              <q-item-section side>
                <q-icon size="xs" tag="span" name="o_account_circle"/>
              </q-item-section>
              <q-item-section class="text-grey-7">我的资料</q-item-section>
            </q-item>
            <q-separator/>
            <q-item clickable @click="logout">
              <q-item-section side>
                <q-icon size="xs" name="logout"/>
              </q-item-section>
              <q-item-section class="text-grey-7">退出登录</q-item-section>
            </q-item>
          </q-list>
        </q-menu>
      </q-btn>

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
