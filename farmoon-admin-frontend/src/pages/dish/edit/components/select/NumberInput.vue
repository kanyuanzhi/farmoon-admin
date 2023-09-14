<template>
  <q-item>
    <q-item-section avatar>{{ label }}</q-item-section>
    <q-item-section>
      <q-input
        v-model.number="inputNumber"
        filled
        dense
        :disable="disable"
        mask="###########"
      >
        <template v-slot:append>
          <span class="text-body2">{{ unit }}</span>
        </template>
      </q-input>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps(["number", "unit", "label", "disable"]);
const emits = defineEmits(["update"]);

const inputNumber = ref(props.number);

watch(inputNumber, (value) => {
  emits("update", value);
});
</script>

<style lang="scss" scoped>

</style>
