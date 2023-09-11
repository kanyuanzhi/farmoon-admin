<template>
  <q-drawer
    v-model="appStore.sidebarOpened"
    bordered
  >
    <SidebarHeader/>
    <q-list class="rounded-borders">
      <SidebarItem
        v-for="route in permissionStore.routes"
        :key="route.path"
        :item="route"
        :level="0"
        :basePath="route.path"
      />
    </q-list>
  </q-drawer>
</template>

<script setup>

import SidebarItem from "layouts/sidebar/SidebarItem.vue";
import {usePermissionStore} from "stores/permission";
import {useAppStore} from "stores/app";
import SidebarHeader from "layouts/sidebar/SidebarHeader.vue";

const appStore = useAppStore();
const permissionStore = usePermissionStore();
// console.log(permissionStore.routes);

</script>

<style lang="scss" scoped>

</style>
