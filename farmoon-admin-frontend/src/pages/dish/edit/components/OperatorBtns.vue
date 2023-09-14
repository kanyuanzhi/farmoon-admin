<template>
  <q-card bordered flat class="my-card" style="">
    <q-card-section class="text-center q-py-sm bg-teal-6 text-white">
      <div class="text-subtitle1">操作</div>
    </q-card-section>
    <q-card-section>
      <div class="column justify-between" style="height: 320px">
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <OperatorBtn label="食材" size="lg" color="green" icon="fa-solid fa-wheat-awn"
                         @click="theIngredientDialog.show()"/>
            <OperatorBtn label="调料" size="lg" color="teal" icon="mdi-shaker"
                         @click="theSeasoningDialog.show()"/>
          </div>
        </div>
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <OperatorBtn label="火力" size="lg" color="red-7" icon="local_fire_department"
                         @click="theFireDialog.show()"/>
            <OperatorBtn label="翻炒" size="lg" color="brown-5" icon="mdi-pot-mix"
                         @click="theStirFryDialog.show()"/>
          </div>
        </div>
        <div class="col">
          <div class="row justify-around" style="padding-top: 10px">
            <OperatorBtn label="纯净水" size="lg" color="blue" icon="water_drop"
                         @click="theWaterDialog.show()"/>
            <OperatorBtn label="食用油" size="lg" color="orange" icon="fa-solid fa-bottle-droplet"
                         @click="theOilDialog.show()"/>
          </div>
        </div>
      </div>

    </q-card-section>
    <q-separator inset/>
    <q-card-actions class="justify-around" style="padding-top: 15px">
      <q-btn-group spread unelevated class="full-width">
        <q-btn color="teal-6" label="保存" @click="theSaveDialog.show()"/>
        <q-separator vertical/>
        <q-btn color="teal-6" label="恢复" @click="dishStore.resetEditingDish()"/>
        <q-separator vertical/>
        <q-btn color="teal-6" label="清空" @click="dishStore.newEditingDish()"/>
        <q-separator vertical/>
        <q-btn color="red-6" label="删除" @click="theDeleteDialog.show()"/>
      </q-btn-group>
    </q-card-actions>
    <TheIngredientDialog ref="theIngredientDialog" @submit="onSubmit"/>
    <TheSeasoningDialog ref="theSeasoningDialog" @submit="onSubmit"/>
    <TheFireDialog ref="theFireDialog" @submit="onSubmit"/>
    <TheStirFryDialog ref="theStirFryDialog" @submit="onSubmit"/>
    <TheWaterDialog ref="theWaterDialog" @submit="onSubmit"/>
    <TheOilDialog ref="theOilDialog" @submit="onSubmit"/>

<!--    <TheSaveDialog ref="theSaveDialog"/>-->
<!--    <TheDeleteDialog ref="theDeleteDialog"/>-->
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
// import TheSaveDialog from "pages/dish/edit/components/dialogs/TheSaveDialog.vue";
import OperatorBtn from "pages/dish/edit/components/dialogs/OperatorBtn.vue";
import {useDishStore} from "stores/dish";
// import TheDeleteDialog from "pages/dish/edit/components/dialogs/TheDeleteDialog.vue";
import {newStirFryStep} from "pages/dish/edit/components/dialogs/newStep";

const dishStore = useDishStore();

const theIngredientDialog = ref(null);
const theSeasoningDialog = ref(null);
const theFireDialog = ref(null);
const theStirFryDialog = ref(null);
const theWaterDialog = ref(null);
const theOilDialog = ref(null);

const theSaveDialog = ref(null);
const theDeleteDialog = ref(null);
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
.my-card {
  height: calc(100vh - 50px - 32px - 52px);
}
</style>
