<template>
  <q-page class="q-pa-md">
    <div class="row q-col-gutter-md">
      <div class="col-9">
        <StepList/>
      </div>
      <div class="col-3">
        <OperatorBtns/>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import StepList from "pages/dish/edit/components/StepList.vue";
import OperatorBtns from "pages/dish/edit/components/OperatorBtns.vue";
</script>

<style lang="scss" scoped>

</style>
