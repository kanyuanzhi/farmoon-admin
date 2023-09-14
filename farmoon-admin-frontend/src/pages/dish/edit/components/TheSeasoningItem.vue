<template>
  <div class="row">
    <TheSeasoningSelect class="col" ref="theSeasoningSelect" :seasoning="seasoning"
                        :seasoning-options="seasoningOptions" @update="onSeasoningUpdate"/>
    <NumberSelect class="col" label="分量" unit="克" :number="seasoning.weight" :min="minWeight" :max="maxWeight"
                  :step="weightStep" @update="onWeightUpdate"/>
    <q-btn size="xs" flat icon="clear" text-color="grey-8" @click="emits('delete')"></q-btn>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import TheSeasoningSelect from "pages/dish/edit/components/select/TheSeasoningSelect.vue";
import NumberSelect from "pages/dish/edit/components/select/NumberSelect.vue";

const props = defineProps(["seasoning", "seasoningOptions", "index"]);
const emits = defineEmits(["delete", "seasoning-select", "weight-select"]);

onMounted(() => {
  generateWeightSelection(props.seasoning.pumpNumber);
});

const minWeight = ref(1);
const maxWeight = ref(10);
const weightStep = ref(1);

const generateWeightSelection = (pumpNumber) => {
  if ([1, 2, 3, 4, 5, 6].indexOf(pumpNumber) > -1) {
    minWeight.value = 1;
    maxWeight.value = 100;
    weightStep.value = 1;
  } else if ([7, 8].indexOf(pumpNumber) > -1) {
    minWeight.value = 5;
    maxWeight.value = 150;
    weightStep.value = 5;
  } else {
    minWeight.value = 1;
    maxWeight.value = 20;
    weightStep.value = 1;
  }
};

const onSeasoningUpdate = (val) => {
  generateWeightSelection(val.pumpNumber);
  emits("seasoning-select", val, props.index)
};

const onWeightUpdate = (val) => {
  emits("weight-select", val, props.index)
};
</script>

<style lang="scss" scoped></style>
