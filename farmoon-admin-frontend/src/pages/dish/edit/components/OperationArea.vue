<template>
  <q-card bordered flat class="my-card" style="">
    <q-card-section class="text-center q-py-sm">
      <div class="text-subtitle2">操作</div>
    </q-card-section>
    <q-separator/>
    <q-card-section>
      <q-btn-group spread unelevated class="full-width">
        <q-btn color="teal-6" label="保存" @click="theSaveDialog.show()"/>
        <q-separator vertical/>
        <q-btn color="teal-6" label="恢复" @click="dishStore.resetEditingDish()"/>
        <q-separator vertical/>
        <q-btn color="teal-6" label="清空" @click="dishStore.newEditingDish()"/>
        <q-separator vertical/>
        <q-btn color="teal-6" label="删除" @click="theDeleteDialog.show()"/>
      </q-btn-group>
    </q-card-section>
    <TheSaveDialog ref="theSaveDialog"/>
    <!--    <TheDeleteDialog ref="theDeleteDialog"/>-->
  </q-card>
</template>

<script setup>
import { ref } from 'vue'
import { useDishStore } from 'stores/dish'
// import TheDeleteDialog from "pages/dish/edit/components/dialogs/TheDeleteDialog.vue";
import TheSaveDialog from 'pages/dish/edit/components/dialogs/TheSaveDialog.vue'

const dishStore = useDishStore()

const theSaveDialog = ref(null)
const theDeleteDialog = ref(null)
</script>

<style lang="scss" scoped>
//.my-card {
//  height: calc(100vh - 50px - 32px - 52px);
//}
</style>
