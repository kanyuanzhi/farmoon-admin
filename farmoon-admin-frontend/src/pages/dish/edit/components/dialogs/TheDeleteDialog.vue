<template>
  <div>
    <q-dialog v-model="shown" persistent @hide="onHide" position="standard">
      <q-card style="width: 400px;" class="q-px-sm q-mt-md">
        <q-card-section>
          <q-avatar size="30px" icon="signal_wifi_off" color="red-6" text-color="white"/>
          <span class="q-ml-sm text-subtitle1">确认删除当前菜品？</span>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn v-close-popup flat color="">取消</q-btn>
          <q-btn unelevated color="teal-6" @click="onSubmit">确认</q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Notify } from "quasar";
import { createDish, deleteDish, updateDish } from "src/api/dish";
import { getCuisines } from "src/api/cuisine";
import { UseAppStore } from "stores/appStore";
import { cloneDeep } from "lodash/lang";

const useAppStore = UseAppStore();

const shown = ref(false);

const show = async () => {
  shown.value = true;
};

const onSubmit = async () => {
  const { data } = await deleteDish(useAppStore.editingDish.uuid);
  if (data.message === "success") {
    useAppStore.newEditingDish()
    Notify.create("删除成功");
  } else {
    Notify.create("删除失败");
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
