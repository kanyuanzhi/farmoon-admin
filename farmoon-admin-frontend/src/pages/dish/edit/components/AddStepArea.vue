<template>
  <q-card bordered flat class="my-card" style="">
    <q-card-section class="text-center q-py-sm">
      <div class="text-subtitle2">添加步骤</div>
    </q-card-section>
    <q-separator/>
    <q-card-section>
      <div class="column justify-between" style="height: 280px">
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <AddStepButton label="食材" size="md" color="green" icon="fa-solid fa-wheat-awn"
                         @click="theIngredientDialog.show()"/>
            <AddStepButton label="调料" size="md" color="teal" icon="mdi-shaker"
                         @click="theSeasoningDialog.show()"/>
          </div>
        </div>
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <AddStepButton label="火力" size="md" color="red-7" icon="local_fire_department"
                         @click="theFireDialog.show()"/>
            <AddStepButton label="翻炒" size="md" color="brown-5" icon="mdi-pot-mix"
                         @click="theStirFryDialog.show()"/>
          </div>
        </div>
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <AddStepButton label="纯净水" size="md" color="blue" icon="water_drop"
                         @click="theWaterDialog.show()"/>
            <AddStepButton label="食用油" size="md" color="orange" icon="fa-solid fa-bottle-droplet"
                         @click="theOilDialog.show()"/>
          </div>
        </div>
      </div>
    </q-card-section>

    <TheIngredientDialog ref="theIngredientDialog" @submit="onSubmit"/>
    <TheSeasoningDialog ref="theSeasoningDialog" @submit="onSubmit"/>
    <TheFireDialog ref="theFireDialog" @submit="onSubmit"/>
    <TheStirFryDialog ref="theStirFryDialog" @submit="onSubmit"/>
    <TheWaterDialog ref="theWaterDialog" @submit="onSubmit"/>
    <TheOilDialog ref="theOilDialog" @submit="onSubmit"/>

  </q-card>
</template>

<script setup>
import TheOilDialog from "pages/dish/edit/components/dialogs/TheOilDialog.vue";
import TheIngredientDialog from "pages/dish/edit/components/dialogs/TheIngredientDialog.vue";
import TheStirFryDialog from "pages/dish/edit/components/dialogs/TheStirFryDialog.vue";
import TheSeasoningDialog from "pages/dish/edit/components/dialogs/TheSeasoningDialog.vue";
import TheFireDialog from "pages/dish/edit/components/dialogs/TheFireDialog.vue";
import TheWaterDialog from "pages/dish/edit/components/dialogs/TheWaterDialog.vue";
import {ref} from "vue";
import AddStepButton from "pages/dish/edit/components/dialogs/AddStepButton.vue";
import {useDishStore} from "stores/dish";
import {newStirFryStep} from "pages/dish/edit/components/dialogs/newStep";

const dishStore = useDishStore();

const theIngredientDialog = ref(null);
const theSeasoningDialog = ref(null);
const theFireDialog = ref(null);
const theStirFryDialog = ref(null);
const theWaterDialog = ref(null);
const theOilDialog = ref(null);

const onSubmit = (val, index) => {
  if (index === -1) {
    dishStore.editingDish.steps.push(val);
  } else {
    dishStore.editingDish.steps.splice(index + 1, 0, val);
  }
  if (val.instructionType !== "stir_fry") {
    dishStore.editingDish.steps.push(newStirFryStep(dishStore.lastStirFryGear, 0));
  }
  dishStore.shiftEditingDishChangedFlag()
};
</script>

<style lang="scss" scoped>
//.my-card {
//  height: calc(100vh - 50px - 32px - 52px);
//}
</style>
