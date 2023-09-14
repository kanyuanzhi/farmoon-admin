<template>
  <q-dialog v-model="shown" @hide="onHide">
    <q-card style="width: 600px" class="q-mt-md">
      <q-card-section class="bg-teal-6 text-white q-py-sm">
        <div class="text-h6">添加翻炒</div>
      </q-card-section>
      <q-card-section>
        <GearSlider ref="gearSlider" :gear="gear" label="档位" color="amber-10" :gear-min="0" :gear-max="5"
                    @update="(v)=>{gear=v}"/>

        <DurationSelect ref="durationSelect" :duration="duration" @update="(v)=>{duration=v}"/>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn v-close-popup flat color="teal-6">取消</q-btn>
        <q-btn color="teal-6" unelevated @click="onSubmit">提交</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref } from "vue";
import GearSlider from "pages/dish/edit/components/select/GearSlider.vue";
import DurationSelect from "pages/dish/edit/components/select/DurationSelect.vue";
import { newStirFryStep } from "pages/dish/edit/components/dialogs/newStep";
import { useDishStore } from "stores/dish";

const dishStore = useDishStore();

const emits = defineEmits(["update", "submit"]);

const shown = ref(false);

const gear = ref(0);
const duration = ref(0);

const gearSlider = ref(null);
const timeSelect = ref(null);

let isUpdate = false;
let stepIndex = 0;

const show = (index = -1) => {
  shown.value = true;
  stepIndex = index;
  gear.value = dishStore.lastStirFryGear;
};

const updateDialogShow = (step, index) => {
  shown.value = true;
  isUpdate = true;
  stepIndex = index;
  gear.value = step.gear;
  duration.value = step.duration;
};

const onSubmit = () => {
  dishStore.setLastStirFryGear(gear.value);
  try {
    const newStep = newStirFryStep(gear.value, duration.value);
    emits(isUpdate ? "update" : "submit", newStep, stepIndex);
  } catch (e) {
    return;
  }
  isUpdate = false;
  stepIndex = 0;
  shown.value = false;
};

const onHide = () => {
  gear.value = 0;
  duration.value = 0;
};

defineExpose({
  show,
  updateDialogShow,
});
</script>

<style lang="scss" scoped></style>
