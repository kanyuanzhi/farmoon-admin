<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-8 offset-2">
        <div style="padding: 12px 16px;height: 64px">
          <q-btn label="添加食材形状" color="primary" @click="addRow"/>
        </div>
        <q-markup-table flat>
          <thead>
          <tr>
            <th class="text-center" style="width: 20%">#</th>
            <th class="text-center" style="width: 50%">名称</th>
            <th class="text-center" style="width: 30%">操作</th>
          </tr>
          </thead>
          <tbody ref="tableBody">
          <tr v-for="(row, index) in rows" :key="row.id">
            <td class="text-center">{{ index + 1 }}</td>
            <td class="text-center">
              {{ row.name }}
              <q-popup-edit class="" v-model="row.name" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                            @update:model-value="updateRow(row)">
                <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
              </q-popup-edit>
            </td>
            <td class="text-center">
              <q-btn dense round flat icon="mdi-delete" size="sm" color="grey-6"
                     :disable="row.unDeletable"
                     @click="deleteRow(row)"/>
              <q-btn dense class="drag-item q-ml-sm" round flat icon="drag_indicator" size="sm" color="primary"/>
            </td>
          </tr>
          </tbody>
        </q-markup-table>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { deleteAPI, getAPI, postAPI, putAPI } from 'src/api'
import { Dialog, Notify } from 'quasar'
import { random, remove } from 'lodash'
import Sortable from 'sortablejs'
import { cloneDeep } from 'lodash/lang'

const loading = ref(false)

const defaultPageIndex = 1
const defaultPageSize = 10

const rows = ref([])

const tableBody = ref(null)
onMounted(async () => {
  const sortable = new Sortable(tableBody.value, {
    onEnd: onSortEnd,
    dragClass: 'sortable-drag',
    chosenClass: 'sortable-chosen',
    animation: 150,  // ms, animation speed moving items when sorting, `0` — without animation
    easing: 'cubic-bezier(1, 0, 0, 1)', // Easing for animation. Defaults to null. See https://easings.net/ for examples.
    handle: '.drag-item',  // Specifies which items inside the element should be draggable
  })
  onUnmounted(() => {
    sortable.destroy()
  })
  loading.value = true
  await getRowsNumber()
  await getRows(defaultPageIndex, defaultPageSize)
  loading.value = false
})

const getRowsNumber = async () => {
  const { data } = await getAPI('/private/ingredient-shape/count')
}

const getRows = async (pageIndex, pageSize) => {
  const { data } = await getAPI('/private/ingredient-shape/list', {
    pageIndex: pageIndex,
    pageSize: pageSize,
  })
  if (data.ingredientShapes !== undefined) {
    rows.value.splice(0, rows.value.length, ...data.ingredientShapes)
  } else {
    rows.value.splice(0, rows.value.length, ...data.data.ingredientShapes)
  }
  rows.value.forEach((item, index) => {
    item.index = index + 1
  })
}

const addRow = async () => {
  Dialog.create({
    message: '请输入食材形状',
    prompt: {
      model: '',
      type: 'text', // optional
    },
    ok: '确定',
    cancel: '取消',
    persistent: true,
  }).onOk(async (name) => {
    try {
      const { data, message } = await postAPI('private/ingredient-shape/add', { name: name })
      Notify.create(message)
      rows.value.unshift(data.ingredientShape)
    } catch (e) {
      console.log(e.toString())
    }
  }).onCancel(() => {
  }).onDismiss(() => {
  })
}

const updateRow = async (row) => {
  try {
    const newData = {
      id: row.id,
      name: row.name,
    }
    const { message } = await putAPI('private/ingredient-shape/update', newData)
    Notify.create(message)
  } catch (e) {
    console.log(e.toString())
  }
}

const deleteRow = async (row) => {
  const { id } = row
  console.log(id)
  Dialog.create({
    message: '确认删除？',
    ok: '确认',
    cancel: '取消',
    focus: 'none',
  }).onOk(async () => {
    try {
      const { message } = await deleteAPI('/private/ingredient-shape/delete', { id: id })
      remove(rows.value, (row) => row.id === id)
      Notify.create(message)
    } catch (e) {
      console.log(e.toString())
    }
  }).onCancel(() => {

  }).onDismiss(() => {
  })
}

const onSortEnd = async (event) => {
  try {
    const draggedItem = rows.value[event.oldIndex]
    rows.value.splice(event.oldIndex, 1)
    rows.value.splice(event.newIndex, 0, draggedItem)
    rows.value.forEach((row, index) => {
      row.sort = index + 1
    })
    const { message } = await putAPI('/private/ingredient-type/update-sorts', { ingredientShapes: rows.value })
    Notify.create({
      message: message,
      type: 'positive',
    })
  } catch (e) {
    console.log(e.toString())
  }
}
</script>

<style scoped lang="scss">
@import "src/pages/dish/index.scss";

</style>
