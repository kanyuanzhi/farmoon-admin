<template>
  <q-drawer
    v-model="appStore.sidebarOpened"
    bordered
  >
    <TheSidebarHeader/>
    <q-list class="rounded-borders">
      <q-scroll-area style="height: calc(100vh - 50px)">
        <SidebarItem
          v-for="route in permissionStore.routes"
          :key="route.path"
          :item="route"
          :level="0"
          :basePath="route.path"
        />
      </q-scroll-area>
    </q-list>
  </q-drawer>
</template>

<script setup>

import SidebarItem from "layouts/sidebar/SidebarItem.vue";
import {usePermissionStore} from "stores/permission";
import {useAppStore} from "stores/app";
import TheSidebarHeader from "layouts/sidebar/TheSidebarHeader.vue";

const appStore = useAppStore();
const permissionStore = usePermissionStore();
// console.log(permissionStore.routes);

</script>

<style lang="scss" scoped>

</style>
