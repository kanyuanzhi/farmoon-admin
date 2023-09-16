<template>
  <q-dialog v-model="shown" @hide="onHide">
    <q-card class="column" style="width: 90%">
      <q-card-section>
        <div class="text-subtitle1">选择食材</div>
      </q-card-section>
      <q-separator/>
      <q-card-section class="row q-gutter-md">
        <div v-for="type in ingredientTypes" :key="type.id" class="col">
          <div class="q-pl-md text-weight-bold">{{ type.name }}</div>
          <q-list dense>
            <q-scroll-area style="height: 400px; max-width: 300px;">
              <q-item v-for="(ingredient, index) in ingredients[type.id]" :key="ingredient.id" dense clickable v-ripple
                      @click="onSelect($event, ingredient.name)">
                <q-item-section>{{ index + 1 + '. ' + ingredient.name }}</q-item-section>
              </q-item>
            </q-scroll-area>
          </q-list>
        </div>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref } from 'vue'
import { getAPI } from 'src/api'

const emit = defineEmits(['select'])

const shown = ref(false)

const ingredients = ref({}) // 食材
const ingredientTypes = ref([]) // 食材类型

const show = async () => {
  shown.value = true
  try {
    const ingredientTypeData = await getAPI('private/ingredient-type/list')
    ingredientTypes.value = ingredientTypeData.data.ingredientTypes
    const ingredientData = await getAPI('private/ingredient/list')
    ingredientData.data.ingredients.forEach((ingredient) => {
      if (!ingredients.value[ingredient.type]) {
        ingredients.value[ingredient.type] = []
      }
      ingredients.value[ingredient.type].push(ingredient)
    })
  } catch (e) {
    console.log(e.toString())
  }
}

const onSelect = (e, val) => {
  emit('select', val)
  shown.value = false
}

const onHide = () => {
  ingredients.value = {}
  ingredientTypes.value = []
}

defineExpose({
  show,
})
</script>

<style lang="scss" scoped>
@media (min-width: 1000px) {
  .q-dialog__inner--minimized > div {
    max-width: 1000px;
  }
}
</style>
