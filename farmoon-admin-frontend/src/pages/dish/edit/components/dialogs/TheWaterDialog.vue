<template>
  <div>
    <q-dialog v-model="shown" @hide="onHide">
      <q-card style="width: 400px" class="q-mt-md">
        <q-card-section class="bg-teal-6 text-white q-py-sm">
          <div class="text-h6">添加纯净水</div>
        </q-card-section>
        <q-card-section>
<!--          <NumberSelect ref="numberSelect" label="分量" unit="克" :min="5" :max="150" :step="5"-->
<!--                        :number="weight" @update="(v)=>weight=v"/>-->
          <NumberInput ref="numberInput" label="分量" unit="克" :number="weight" @update="(v)=>weight=v"/>

        </q-card-section>
        <q-card-actions align="right">
          <q-btn v-close-popup flat color="teal-6">取消</q-btn>
          <q-btn color="teal-6" unelevated @click="onSubmit">提交</q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { ref } from "vue";
import NumberSelect from "pages/dish/edit/components/select/NumberSelect.vue";
import { newWaterStep } from "pages/dish/edit/components/dialogs/newStep";
import NumberInput from "pages/dish/edit/components/select/NumberInput.vue";

const emits = defineEmits(["update", "submit"]);

const shown = ref(false);

const weight = ref(0);

let isUpdate = false;
let stepIndex = 0;

const show = (index = -1) => {
  shown.value = true;
  stepIndex = index;
};

const updateDialogShow = (step, index) => {
  shown.value = true;
  isUpdate = true;
  stepIndex = index;
  weight.value = step.weight;
};
const onSubmit = () => {
  try {
    const newStep = newWaterStep(weight.value);
    emits(isUpdate ? "update" : "submit", newStep, stepIndex);
  } catch (e) {
    return;
  }
  isUpdate = false;
  stepIndex = 0;
  shown.value = false;
};

const onHide = () => {
  weight.value = 0;
};

defineExpose({
  show,
  updateDialogShow
});
</script>

<style lang="scss" scoped></style>
