<template>
  <q-page class="q-pa-md">
    <div class="row q-col-gutter-md">
      <div class="col-9">
        <StepList/>
      </div>
      <div class="col-3 column">
        <div class="col-2">
          <OperationArea/>
        </div>
        <div class="col-6">
          <AddStepArea/>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import StepList from "pages/dish/edit/components/StepList.vue";
import AddStepArea from "pages/dish/edit/components/AddStepArea.vue";
import OperationArea from "pages/dish/edit/components/OperationArea.vue";
</script>

<style lang="scss" scoped>

</style>
