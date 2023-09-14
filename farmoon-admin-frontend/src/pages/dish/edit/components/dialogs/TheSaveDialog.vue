<template>
  <div>
    <q-dialog v-model="shown" @hide="onHide" position="top">
      <q-card style="width: 400px;margin-top: 50px" class="q-mt-md">
        <q-card-section class="bg-teal-6 text-white q-py-sm">
          <div class="text-h6">保存菜品</div>
        </q-card-section>
        <q-card-section>
          <q-item>
            <q-item-section avatar>名称</q-item-section>
            <q-item-section>
              <q-input
                v-model="newName"
                filled
                dense
                autofocus
                @blur="onInputBlur($event, 'newName')"
                @focus="onInputFocus($event, 'newName')"
              >
              </q-input>
            </q-item-section>
          </q-item>
          <q-item>
            <q-item-section avatar>菜系</q-item-section>
            <q-item-section>
              <q-select
                dense
                filled
                v-model="newCuisine"
                :options="cuisineOptions"
                options-cover
                stack-label
              >
              </q-select>
            </q-item-section>
          </q-item>
          <q-item dense>
            <q-item-section avatar></q-item-section>
            <q-item-section>
            <span class="text-grey-7" style="font-size: 12px">
              <span class="text-red">*</span>
              保存操作将会重置该菜品已有的三种自定义口味
            </span>
            </q-item-section>
          </q-item>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn v-close-popup flat color="">取消</q-btn>
          <q-btn unelevated size="md" color="teal-6" @click="onSubmit('save')">保存</q-btn>
          <q-btn unelevated color="teal-6" @click="onSubmit('create')">新建</q-btn>
        </q-card-actions>

      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Notify } from "quasar";
import { createDish, updateDish } from "src/api/dish";
import { getCuisines } from "src/api/cuisine";
import { UseAppStore } from "stores/appStore";
import { cloneDeep } from "lodash/lang";

const useAppStore = UseAppStore();

const shown = ref(false);

const newName = ref("");
const cuisineOptions = ref([]);
const cuisineMap = ref({});
const newCuisine = ref({});

const show = async () => {
  shown.value = true;
  newName.value = useAppStore.editingDish.name;
  cuisineOptions.value = [];
  const { data } = await getCuisines();
  const cuisines = data.data;
  cuisines.forEach(cuisine => {
    cuisineOptions.value.push({
      label: cuisine.name,
      value: cuisine.id,
    });
    cuisineMap.value[cuisine.id] = cuisine.name;
  });
  newCuisine.value = {
    label: cuisineMap.value[useAppStore.editingDish.cuisine],
    value: useAppStore.editingDish.cuisine,
  };
};

const inputNameToPara = {
  newName,
};

const customKeyboard = ref(null);
const onInputFocus = (e, inputName) => {
  customKeyboard.value.setInputName(inputName);
  customKeyboard.value.setInput(e.target.value, inputName);
};

const onInputBlur = (e, inputName) => {
};

const setDefaultName = (defaultName) => {
  name.value = defaultName;
};

const onChange = (input, inputName) => {
  inputNameToPara[inputName].value = input;
};
const onSubmit = async (flag) => {
  if (newName.value.trim() === "") {
    Notify.create("请输入菜品名称");
    return;
  }
  const newDish = {
    name: newName.value,
    cuisine: newCuisine.value.value,
    steps: useAppStore.editingDish.steps,
    uuid: useAppStore.editingDish.uuid
  };
  if (useAppStore.editingDish.uuid === "" || flag === "create") {
    const { data } = await createDish(newDish);
    if (data.message === "success") {
      useAppStore.editingDish.name = newName.value;
      useAppStore.editingDish.cuisine = newCuisine.value.value;
      useAppStore.editingDish.uuid = data.data;
      useAppStore.originEditingDish = cloneDeep(useAppStore.editingDish);
      Notify.create(flag === "create" ? "新建成功" : "保存成功");
    } else {
      Notify.create(flag === "create" ? "新建失败" : "保存失败");
    }
  } else {
    const { data } = await updateDish(newDish);
    if (data.message === "success") {
      useAppStore.editingDish.name = newName.value;
      useAppStore.editingDish.cuisine = newCuisine.value.value;
      useAppStore.originEditingDish.name = newName.value;
      useAppStore.originEditingDish.cuisine = newCuisine.value.value;
      Notify.create("保存成功");
    } else {
      Notify.create("保存失败");
    }
  }
  shown.value = false;
};

const onHide = () => {
  // newName.value = "";
  // newCuisine.value = {
  //   label: cuisineMap[0],
  //   value: 0
  // };
};

defineExpose({
  show,
});
</script>

<style lang="scss" scoped></style>
