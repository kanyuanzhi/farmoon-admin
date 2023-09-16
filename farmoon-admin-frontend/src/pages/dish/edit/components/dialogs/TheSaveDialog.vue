<template>
  <div>
    <q-dialog v-model="shown" @hide="onHide">
      <q-card style="width: 400px;" class="q-mt-md">
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
                  v-model="cuisine"
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
import { ref } from 'vue'
import { Notify } from 'quasar'
import { cloneDeep } from 'lodash/lang'
import { useDishStore } from 'stores/dish'
import { getAPI, postAPI, putAPI } from 'src/api'

const dishStore = useDishStore()

const shown = ref(false)

const newName = ref('')
const cuisineOptions = ref([])
const cuisineMap = ref({})
const cuisine = ref({})

const show = async () => {
  shown.value = true
  newName.value = dishStore.editingDish.name
  cuisineOptions.value = []
  const { data } = await getAPI('private/cuisine/list')
  const cuisines = data.cuisines
  cuisines.forEach(cuisine => {
    cuisineOptions.value.push({
      label: cuisine.name,
      value: cuisine.id,
    })
    cuisineMap.value[cuisine.id] = cuisine.name
  })
  cuisine.value = {
    label: cuisineMap.value[dishStore.editingDish.cuisine],
    value: dishStore.editingDish.cuisine,
  }
}

const onSubmit = async (flag) => {
  if (newName.value.trim() === '') {
    Notify.create('请输入菜品名称')
    return
  }
  const newDish = {
    id: dishStore.editingDish.id,
    name: newName.value,
    cuisine: cuisine.value.value,
    steps: dishStore.editingDish.steps,
    uuid: dishStore.editingDish.uuid,
  }
  if (dishStore.editingDish.uuid === '' || flag === 'create') {
    try {
      const { data, message } = await postAPI('private/dish/add', newDish)
      dishStore.editingDish.id = data.dish.id
      dishStore.editingDish.name = data.dish.name
      dishStore.editingDish.cuisine = data.dish.cuisine
      dishStore.editingDish.uuid = data.dish.uuid
      dishStore.originEditingDish = cloneDeep(dishStore.editingDish)
      Notify.create(message)
    } catch (e) {
      Notify.create({
        message: e.toString(),
        type: 'negative',
      })
    }
  } else {
    const { data, message } = await putAPI('private/dish/update-with-steps', newDish)
    try {
      dishStore.editingDish.name = data.dish.name
      dishStore.editingDish.cuisine = data.dish.cuisine
      dishStore.originEditingDish.name = data.dish.name
      dishStore.originEditingDish.cuisine = data.dish.cuisine
      Notify.create(message)
    } catch (e) {
      Notify.create({
        message: e.toString(),
        type: 'negative',
      })
    }
  }
  shown.value = false
}

const onHide = () => {
  // newName.value = "";
  // newCuisine.value = {
  //   label: cuisineMap[0],
  //   value: 0
  // };
}

defineExpose({
  show,
})
</script>

<style lang="scss" scoped></style>
