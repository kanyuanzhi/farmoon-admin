<template>
  <q-btn unelevated :color="color" :size="size" style="width: 100px;padding-top: 13px">
    <div class="column items-center no-wrap">
      <q-icon :name="icon"/>
      <div class="text-center">
        {{ label }}
      </div>
    </div>
  </q-btn>
</template>

<script setup>
const props = defineProps(["color", "label", "size", "icon"]);
</script>

<style lang="scss" scoped>

</style>
