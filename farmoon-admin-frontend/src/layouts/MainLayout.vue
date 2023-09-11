<template>
  <q-layout view="lHh LpR lFf">
    <!--    <q-layout view="hHh lpR fFf">-->
    <TheHeader/>

    <TheSidebar/>

    <q-page-container>
      <q-scroll-area style="height: calc(100vh - 50px)">
        <router-view/>
      </q-scroll-area>
<!--      <q-page-scroller class="scroll" position="bottom-right" :scroll-offset="150" :offset="[18, 18]">-->
<!--        <q-btn fab icon="keyboard_arrow_up" color="accent"/>-->
<!--      </q-page-scroller>-->
    </q-page-container>
  </q-layout>
</template>

<script setup>
import TheSidebar from "layouts/sidebar/TheSidebar.vue";
import TheHeader from "layouts/header/TheHeader.vue";

</script>

<style lang="scss" scoped>

</style>
