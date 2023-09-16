<template>
  <q-card bordered flat class="my-card" style="width: 100%">
    <q-card-section class="text-center q-py-sm">
      <div class="text-subtitle2">
        步骤列表{{ dishStore.editingDish.name? " - " + dishStore.editingDish.name : ""}}
      </div>
    </q-card-section>
    <q-separator/>
    <q-card-section>
      <q-list class="scrollable-container" ref="stepList" bordered separator>
        <q-slide-item v-for="item in newSteps" :key="item.key" class="text-black" @right="onDeleteCouple(item)"
                      right-color="red">
          <template v-slot:right>
            <div class="flex">
              <q-icon name="delete"/>
            </div>
          </template>
          <q-item class="row q-px-sm">
            <div class="row col-6">
              <q-item-section side>
                <q-avatar rounded :color="instructionTypeToColor[item.coupleSteps[0].instructionType]"
                          text-color="white" :icon="instructionTypeToIcon[item.coupleSteps[0].instructionType]"
                          size="md"/>
              </q-item-section>
              <q-item-section @click="onStepListItemClick(item.coupleSteps[0], item.indexes[0])">
                <q-item-label>{{ item.indexes[0] + 1 + ". " + item.coupleSteps[0].instructionName }}</q-item-label>
              </q-item-section>

            </div>
            <div class="row col-5" v-if="item.coupleSteps.length === 2">
              <q-item-section side class="q-pl-lg">
                <q-avatar rounded :color="instructionTypeToColor[item.coupleSteps[1].instructionType]"
                          text-color="white" :icon="instructionTypeToIcon[item.coupleSteps[1].instructionType]"
                          size="md"/>
              </q-item-section>
              <q-item-section @click="onStepListItemClick(item.coupleSteps[1], item.indexes[1])">
                <q-item-label>{{ item.indexes[1] + 1 + ". " + item.coupleSteps[1].instructionName }}</q-item-label>
              </q-item-section>
              <q-item-section side class="no-padding">
                <q-icon color="brown-5" name="delete_forever" @click="onDelete(item.indexes[1])"/>
              </q-item-section>
            </div>
            <div class="row col-5" v-else>
              <q-item-section side class="q-pl-lg">
                <q-avatar rounded color="brown-5" text-color="white" icon="mdi-pot-mix" size="md"/>
              </q-item-section>
              <q-item-section class="text-center">
                <q-btn flat icon="add" color="brown-5" class="no-padding" size="12px"
                       @click="theStirFryDialog.show(item.indexes[0])"/>
              </q-item-section>
            </div>
            <div class="flex col-1 justify-end">
              <q-item-section side class="no-padding">
                <q-icon color="teal-6" name="drag_indicator" class="drag-item"/>
              </q-item-section>
            </div>
          </q-item>
        </q-slide-item>
      </q-list>
    </q-card-section>
    <TheIngredientDialog ref="theIngredientDialog" @update="onUpdate"/>
    <TheSeasoningDialog ref="theSeasoningDialog" @update="onUpdate"/>
    <TheFireDialog ref="theFireDialog" @update="onUpdate"/>
    <TheStirFryDialog ref="theStirFryDialog" @update="onUpdate" @submit="onSubmit"/>
    <TheWaterDialog ref="theWaterDialog" @update="onUpdate"/>
    <TheOilDialog ref="theOilDialog" @update="onUpdate"/>
  </q-card>
</template>

<script setup>

import {onMounted, onUnmounted, ref, watch} from "vue";
import Sortable from "sortablejs";
import {useDishStore} from "stores/dish";
import TheIngredientDialog from "pages/dish/edit/components/dialogs/TheIngredientDialog.vue";
import TheFireDialog from "pages/dish/edit/components/dialogs/TheFireDialog.vue";
import TheStirFryDialog from "pages/dish/edit/components/dialogs/TheStirFryDialog.vue";
import TheWaterDialog from "pages/dish/edit/components/dialogs/TheWaterDialog.vue";
import TheOilDialog from "pages/dish/edit/components/dialogs/TheOilDialog.vue";
import TheSeasoningDialog from "pages/dish/edit/components/dialogs/TheSeasoningDialog.vue";
import {random, range, sampleSize} from "lodash";

const dishStore = useDishStore();

const generateNewSteps = () => {
  const newSteps = [];
  for (let i = 0, len = dishStore.editingDish.steps.length; i < len;) {
    const item = [];
    const indexes = [];
    if (dishStore.editingDish.steps[i].instructionType !== "stir_fry") {
      item.push(dishStore.editingDish.steps[i]);
    } else {
      if (i === 0) {
        item.push(dishStore.editingDish.steps[i]);
      } else if (dishStore.editingDish.steps[i - 1].instructionType === "stir_fry") {
        item.push(dishStore.editingDish.steps[i]);
      }
    }
    indexes.push(i);

    if (i + 1 < len && dishStore.editingDish.steps[i + 1].instructionType === "stir_fry") {
      item.push(dishStore.editingDish.steps[i + 1]);
      indexes.push(i + 1);
      i += 2;
    } else {
      i += 1;
    }
    newSteps.push({
      key: 0,
      indexes: indexes,
      coupleSteps: item
    });
  }
  // 生成一组不重复的整数赋给key
  const keys = range(100, 10000);
  const candidateKeys = sampleSize(keys, newSteps.length);
  for (let j = 0, len = newSteps.length; j < len; j++) {
    newSteps[j].key = candidateKeys[j];
  }
  return newSteps;
};

const newSteps = ref(generateNewSteps());

watch(
  () => dishStore.editingDishChangedFlag,
  () => {
    newSteps.value = generateNewSteps();
  },
);

// watch(
//   () => dishStore.editingDish.steps,
//   () => {
//     newSteps.value = generateNewSteps();
//   },
//   { deep: true }
// );

const stepList = ref(null);

onMounted(() => {
  const sortable = new Sortable(stepList.value.$el, {
    onEnd: onSortEnd,
    // dragClass: "sortable-drag",
    // chosenClass: "sortable-chosen",
    animation: 150,  // ms, animation speed moving items when sorting, `0` — without animation
    easing: "cubic-bezier(1, 0, 0, 1)", // Easing for animation. Defaults to null. See https://easings.net/ for examples.
    handle: ".drag-item",  // Specifies which items inside the element should be draggable
  });
  onUnmounted(() => {
    sortable.destroy();
  });
});

const onSortEnd = (event) => {
  const draggedIndexes = newSteps.value[event.oldIndex].indexes;
  const draggedCoupleSteps = newSteps.value[event.oldIndex].coupleSteps;

  const insertedIndexes = newSteps.value[event.newIndex].indexes;
  const insertedCoupleSteps = newSteps.value[event.newIndex].coupleSteps;

  dishStore.editingDish.steps.splice(draggedIndexes[0], draggedIndexes.length === 1 ? 1 : 2);
  dishStore.editingDish.steps.splice(insertedIndexes[0], 0, ...draggedCoupleSteps);
  dishStore.shiftEditingDishChangedFlag();
};

// 前端叫step，后端叫instruction
const instructionTypeToColor = {
  "ingredient": "green",
  "water": "blue",
  "oil": "orange",
  "heat": "red-7",
  "stir_fry": "brown-5",
  "seasoning": "teal",
};

const instructionTypeToIcon = {
  "ingredient": "fa-solid fa-wheat-awn",
  "water": "water_drop",
  "oil": "fa-solid fa-bottle-droplet",
  "heat": "local_fire_department",
  "stir_fry": "mdi-pot-mix",
  "seasoning": "mdi-shaker",
};

const theIngredientDialog = ref(null);
const theSeasoningDialog = ref(null);
const theFireDialog = ref(null);
const theStirFryDialog = ref(null);
const theWaterDialog = ref(null);
const theOilDialog = ref(null);

const onStepListItemClick = (step, index) => {
  switch (step.instructionType) {
    case "ingredient":
      theIngredientDialog.value.updateDialogShow(step, index);
      break;
    case "seasoning":
      theSeasoningDialog.value.updateDialogShow(step, index);
      break;
    case "heat":
      theFireDialog.value.updateDialogShow(step, index);
      break;
    case "stir_fry":
      theStirFryDialog.value.updateDialogShow(step, index);
      break;
    case "water":
      theWaterDialog.value.updateDialogShow(step, index);
      break;
    case "oil":
      theOilDialog.value.updateDialogShow(step, index);
      break;
    default:
      return;
  }
};

const onUpdate = (step, index) => {
  dishStore.editingDish.steps[index] = step;
  dishStore.shiftEditingDishChangedFlag();
};

const onSubmit = (val, index) => {
  if (index === -1) {
    dishStore.editingDish.steps.push(val);
  } else {
    dishStore.editingDish.steps.splice(index + 1, 0, val);
  }
  dishStore.shiftEditingDishChangedFlag();
};

const onDelete = (index) => {
  dishStore.editingDish.steps.splice(index, 1);
  dishStore.shiftEditingDishChangedFlag();
};

const onDeleteCouple = (item) => {
  dishStore.editingDish.steps.splice(item.indexes[0], item.indexes.length === 1 ? 1 : 2);
  dishStore.shiftEditingDishChangedFlag();
};
</script>

<style lang="scss" scoped>
//.my-card {
//  height: calc(100vh - 50px - 32px);
//}
.scrollable-container {
  height: calc(100vh - 50px - 32px - 32px - 44px - 2px);
  overflow-y: auto;
}

.sortable-drag {
  background: #1976d2;
}

.sortable-chosen {
  background: rgba(38, 166, 154, 0.5);
}

::-webkit-scrollbar {
  width: 3px;
  height: 0;
  background-color: transparent;
  //background-color: #009688;
}

/* 修改滑块样式 */
::-webkit-scrollbar-thumb {
  background-color: #009688;
  border-radius: 5px;
}

/* 修改滑轨样式 */
::-webkit-scrollbar-track {
  background-color: #f2f2f2;
}

</style>
