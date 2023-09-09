<template>
  <q-layout view="lHh LpR lFf">
    <TheHeader/>

    <TheSidebar/>

    <q-page-container>
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import TheSidebar from "layouts/sidebar/TheSidebar.vue";
import TheHeader from "layouts/header/TheHeader.vue";

</script>
