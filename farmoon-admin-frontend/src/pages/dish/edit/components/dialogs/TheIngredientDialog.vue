<template>
  <div>
    <q-dialog v-model="shown" @hide="onHide">
      <q-card style="width: 400px" class="q-mt-md">
        <q-card-section class="bg-teal-6 text-white q-py-sm">
          <div class="text-h6">添加食材</div>
        </q-card-section>
        <q-card-section>
          <q-item>
            <q-item-section avatar>食材</q-item-section>
            <q-item-section>
              <q-input
                v-model="name"
                filled
                dense
                @blur="onInputBlur($event, 'name')"
                @focus="onInputFocus($event, 'name')"
              >
                <template v-slot:after>
                  <q-btn round dense flat icon="toc" @click="theIngredientNameSelectionDialog.show()"/>
                </template>
              </q-input>
            </q-item-section>
          </q-item>

          <q-item>
            <q-item-section avatar>形状</q-item-section>
            <q-item-section>
              <q-input
                v-model="shape"
                filled
                dense
                @blur="onInputBlur($event, 'shape')"
                @focus="onInputFocus($event, 'shape')"
              >
                <template v-slot:after>
                  <q-btn round dense flat icon="toc" @click="theIngredientShapeSelectionDialog.show()"/>
                </template>
              </q-input>
            </q-item-section>
          </q-item>

          <!--          <NumberSelect ref="numberSelect" label="分量" unit="克" :min="0" :max="500" :step="5"-->
          <!--                        :number="weight" @update="(v)=>weight=v"/>-->
          <NumberInput ref="numberInput" label="分量" unit="克" :number="weight" @update="(v)=>weight=v"/>

          <SlotRadio ref="slotRadio" :slotNumber="slotNumber" label="菜盒" :slot-count="4" @update="(v)=>slotNumber=v"/>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn v-close-popup flat color="teal-6">取消</q-btn>
          <q-btn color="teal-6" unelevated @click="onSubmit">提交</q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
    <TheIngredientNameSelectionDialog ref="theIngredientNameSelectionDialog" @select="(val)=>name=val"/>
    <TheIngredientShapeSelectionDialog ref="theIngredientShapeSelectionDialog" @select="(val)=>shape=val"/>
  </div>
</template>

<script setup>
import {ref} from "vue";
import TheIngredientNameSelectionDialog from "pages/dish/edit/components/dialogs/TheIngredientNameSelectionDialog.vue";
import TheIngredientShapeSelectionDialog
  from "pages/dish/edit/components/dialogs/TheIngredientShapeSelectionDialog.vue";
import SlotRadio from "pages/dish/edit/components/SlotRadio.vue";
import NumberSelect from "pages/dish/edit/components/select/NumberSelect.vue";
import {Notify} from "quasar";
import {newIngredientStep} from "pages/dish/edit/components/dialogs/newStep";
import NumberInput from "pages/dish/edit/components/select/NumberInput.vue";

const emits = defineEmits(["update", "submit"]);

const shown = ref(false);

const name = ref("");
const shape = ref("");
const weight = ref(0);
const slotNumber = ref("1");

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
  name.value = step.name;
  shape.value = step.shape;
  weight.value = step.weight;
  slotNumber.value = step.slotNumber;
};

const inputNameToPara = {
  name,
  shape,
};

const customKeyboard = ref(null);
const onInputFocus = (e, inputName) => {
  customKeyboard.value.setInputName(inputName);
  customKeyboard.value.setInput(e.target.value, inputName);
};

const onInputBlur = (e, inputName) => {
};

const onChange = (input, inputName) => {
  inputNameToPara[inputName].value = input;
};

const theIngredientNameSelectionDialog = ref(null);
const theIngredientShapeSelectionDialog = ref(null);

const onSubmit = () => {
  if (name.value.trim() === "") {
    Notify.create("请添加食材");
    return;
  }
  try {
    const newStep = newIngredientStep(name.value.trim(), shape.value.trim(), weight.value, slotNumber.value);
    emits(isUpdate ? "update" : "submit", newStep, stepIndex);
  } catch (e) {
    return;
  }
  isUpdate = false;
  stepIndex = 0;
  shown.value = false;
};

const onHide = () => {
  name.value = "";
  shape.value = "";
  weight.value = 0;
  slotNumber.value = "1";
};

defineExpose({
  show,
  updateDialogShow,
});
</script>

<style lang="scss" scoped></style>
