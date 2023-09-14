<template>
  <q-item>
    <q-item-section avatar>{{ label }}</q-item-section>
    <q-item-section>
      <q-slider
        dense
        v-model="selectedGear"
        :color="color"
        marker-labels
        markers
        :min="gearMin"
        :max="gearMax"
      />
    </q-item-section>
  </q-item>
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps(["label", "gearMin", "gearMax", "color", "gear"]);
const emits = defineEmits(["update"])

const selectedGear = ref(props.gear);

watch(selectedGear,(v)=>{
  emits("update", v)
})
</script>

<style lang="scss" scoped>

</style>
