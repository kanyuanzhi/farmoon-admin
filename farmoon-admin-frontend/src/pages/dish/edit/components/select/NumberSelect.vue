<template>
  <q-item>
    <q-item-section avatar>{{ label }}</q-item-section>
    <q-item-section>
      <q-select
        dense
        options-dense
        filled
        v-model="selectedNumber"
        :options="numberOptions"
        options-cover
        stack-label
        :disable="disable"
      >
        <template v-slot:append>
          <span class="text-body2">{{ unit }}</span>
        </template>
      </q-select>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";

const props = defineProps(["min", "max", "step", "number", "unit", "label", "disable"]);
const emits = defineEmits(["update"]);

const selectedNumber = ref(props.number);

const numberOptions = ref([]);
onMounted(() => {
  for (let i = props.min; i < props.max + 1; i += props.step) {
    numberOptions.value.push(i);
  }
});

watch(
  () => props.max,
  () => {
    numberOptions.value = [];
    for (let i = props.min; i < props.max + 1; i += props.step) {
      numberOptions.value.push(i);
    }
  });

watch(selectedNumber, (value) => {
  emits("update", value);
});
</script>

<style lang="scss" scoped>

</style>
