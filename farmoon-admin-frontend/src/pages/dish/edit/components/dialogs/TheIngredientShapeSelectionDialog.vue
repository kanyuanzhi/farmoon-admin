<template>
  <q-dialog v-model="shown" @hide="onHide">
    <q-card class="column" style="width: 300px">
      <q-card-section>
        <div class="text-subtitle1">选择食材形状</div>
      </q-card-section>
      <q-separator/>
      <q-card-section>
        <q-list separator padding>
          <q-scroll-area style="height: 400px; max-width: 300px;">
            <q-item dense v-for="shape in shapes" :key="shape.id" clickable v-ripple @click="onSelect($event, shape.name)">
              <q-item-section>{{ shape.name }}</q-item-section>
            </q-item>
          </q-scroll-area>
        </q-list>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref } from 'vue'
import { getAPI } from 'src/api'

const emit = defineEmits(['select'])

const shown = ref(false)

const shapes = ref([])

const show = async () => {
  shown.value = true
  try {
    const shapeData = await getAPI('private/ingredient-shape/list')
    shapes.value = shapeData.data.ingredientShapes
  } catch (e) {
    console.log(e.toString())
  }
}

const onSelect = (e, val) => {
  emit('select', val)
  shown.value = false
}

const onHide = () => {
  shapes.value = []
}

defineExpose({
  show,
})
</script>

<style lang="scss" scoped>

</style>
