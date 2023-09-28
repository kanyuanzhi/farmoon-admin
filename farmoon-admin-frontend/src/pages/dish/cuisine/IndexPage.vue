<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-8 offset-2">
        <div style="padding: 12px 16px;height: 64px">
          <q-btn label="添加菜系" color="primary" @click="addCuisine"/>
        </div>
        <q-markup-table flat>
          <thead>
          <tr>
            <th class="text-center" style="width: 20%">#</th>
            <th class="text-center" style="width: 30%">名称</th>
            <th class="text-center" style="width: 20%">不可删除</th>
            <th class="text-center" style="width: 30%">操作</th>
          </tr>
          </thead>
          <tbody ref="tableBody">
          <tr v-for="(cuisine, index) in cuisines" :key="cuisine.id">
            <td class="text-center">{{ index + 1 }}</td>
            <td class="text-center">
              {{ cuisine.name }}
              <q-popup-edit class="" v-model="cuisine.name" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                            @update:model-value="onUpdateName(cuisine)">
                <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
              </q-popup-edit>
            </td>
            <td class="text-center">
              <q-toggle
                  dense
                  v-model="cuisine.unDeletable"
                  color="primary"
                  @update:model-value="onUpdateUnDeletable(cuisine)"
              />
            </td>
            <td class="text-center">
              <q-btn dense round flat icon="mdi-delete" size="sm" color="grey-6"
                     :disable="cuisine.unDeletable"
                     @click="deleteCuisine(cuisine.id, index)">
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">删除</q-tooltip>
              </q-btn>
              <q-btn dense class="drag-item q-ml-sm" round flat icon="drag_indicator" size="sm" color="primary">
                <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">拖拽排序</q-tooltip>
              </q-btn>
            </td>
          </tr>
          </tbody>
        </q-markup-table>
      </div>
    </div>

  </q-page>
</template>

<script setup>
import { onBeforeMount, onMounted, onUnmounted, ref } from "vue";
import { deleteAPI, getAPI, postAPI, putAPI } from "src/api";
import { Dialog, Notify } from "quasar";
import Sortable from "sortablejs";

const cuisines = ref([]);

onBeforeMount(async () => {
  const { data } = await getAPI("/private/cuisine/list");
  cuisines.value = data.cuisines;
});

const tableBody = ref(null);
onMounted(async () => {
  const sortable = new Sortable(tableBody.value, {
    onEnd: onSortEnd,
    dragClass: "sortable-drag",
    chosenClass: "sortable-chosen",
    animation: 150,  // ms, animation speed moving items when sorting, `0` — without animation
    easing: "cubic-bezier(1, 0, 0, 1)", // Easing for animation. Defaults to null. See https://easings.net/ for examples.
    handle: ".drag-item",  // Specifies which items inside the element should be draggable
  });
  onUnmounted(() => {
    sortable.destroy();
  });
});

const onSortEnd = async (event) => {
  try {
    const draggedItem = cuisines.value[event.oldIndex];
    cuisines.value.splice(event.oldIndex, 1);
    cuisines.value.splice(event.newIndex, 0, draggedItem);
    cuisines.value.forEach((cuisine, index) => {
      cuisine.sort = index + 1;
    });
    const { message } = await putAPI("/private/cuisine/update-sorts", { cuisines: cuisines.value });
    Notify.create({
      message: message,
      type: "positive",
    });
  } catch (e) {
    console.log(e.toString());
  }
};

const onUpdateName = async (cuisine) => {
  try {
    const { message } = await putAPI("/private/cuisine/update-name", { id: cuisine.id, name: cuisine.name });
    Notify.create({
      message: message,
      type: "positive",
    });
  } catch (e) {
    console.log(e.toString());
  }
};

const onUpdateUnDeletable = async (cuisine) => {
  try {
    const { message } = await putAPI("/private/cuisine/update-unDeletable", {
      id: cuisine.id,
      unDeletable: cuisine.unDeletable,
    });
    Notify.create({
      message: message,
      type: "positive",
    });
  } catch (e) {
    cuisine.unDeletable = !cuisine.unDeletable;
    console.log(e.toString());
  }
};

const deleteCuisine = async (id, index) => {
  Dialog.create({
    message: "确认删除？",
    ok: "确认",
    cancel: "取消",
    focus: "none",
  }).onOk(async () => {
    try {
      const { message } = await deleteAPI("/private/cuisine/delete", { id: id });
      cuisines.value.splice(index, 1);
      Notify.create({
        message: message,
        type: "positive",
      });
    } catch (e) {
      console.log(e.toString());
    }
  });
};

const addCuisine = async () => {
  Dialog.create({
    message: "请输入菜系名称",
    prompt: {
      model: "",
      type: "text", // optional
    },
    ok: "确定",
    cancel: "取消",
    persistent: true,
  }).onOk(async (name) => {
    try {
      const { data, message } = await postAPI("private/cuisine/add", { name: name });
      Notify.create({
        message: message,
        type: "positive",
      });
      cuisines.value.unshift(data.cuisine);
    } catch (e) {
      console.log(e.toString());
    }
  }).onCancel(() => {
  }).onDismiss(() => {
  });
};
</script>

<style lang="scss" scoped>
@import "src/pages/dish/index.scss";
</style>
