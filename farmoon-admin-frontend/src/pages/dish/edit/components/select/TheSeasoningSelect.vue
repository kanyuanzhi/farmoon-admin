<template>
  <q-item dense>
    <q-item-section avatar>种类</q-item-section>
    <q-item-section>
      <q-select
        dense
        options-dense
        filled
        v-model="selectedSeasoning"
        :options="seasoningOptions"
        options-cover
        stack-label
      >
<!--        <template v-slot:selected>-->
<!--          {{ seasoning.pumpNumber === 0 ? "" : (seasoning.pumpNumber + "号泵：" + seasoning.label) }}-->
<!--        </template>-->
      </q-select>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps(["seasoning", "seasoningOptions"]);
const emits = defineEmits(["update"]);

console.log(props.seasoning)

const selectedSeasoning = ref(props.seasoning);

watch(selectedSeasoning, (value) => {
  emits("update", value);
});
</script>

<style lang="scss" scoped>

</style>
