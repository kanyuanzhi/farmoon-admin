<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-8 offset-2">
        <div style="padding: 12px 16px;height: 64px">
          <q-btn label="添加调料" color="primary" @click="addSeasoning"/>
        </div>
        <q-markup-table flat>
          <thead>
          <tr>
            <th class="text-center" style="width: 20%">#</th>
            <th class="text-center" style="width: 20%">名称</th>
            <th class="text-center" style="width: 20%">泵号</th>
            <th class="text-center" style="width: 20%">加料分量与加料时长比例</th>
            <th class="text-center" style="width: 20%">操作</th>
          </tr>
          </thead>
          <tbody ref="tableBody">
          <tr v-for="(seasoning, index) in seasonings" :key="seasoning.id">
            <td class="text-center">{{ index + 1 }}</td>
            <td class="text-center">
              {{ seasoning.name }}
              <q-popup-edit class="" v-model="seasoning.name" buttons label-set="确认" label-cancel="取消"
                            v-slot="scope" @update:model-value="onUpdate(seasoning)">
                <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
              </q-popup-edit>
            </td>
            <td class="text-center">
              {{ seasoning.pump }}
              <q-popup-edit class="" v-model="seasoning.pump" buttons label-set="确认" label-cancel="取消"
                            v-slot="scope" @update:model-value="onUpdate(seasoning)">
                <q-input v-model.number="scope.value" dense autofocus counter mask="##" @keyup.enter="scope.set"/>
              </q-popup-edit>
            </td>
            <td class="text-center">
              {{ seasoning.ratio }}
              <q-popup-edit class="" v-model="seasoning.ratio" buttons label-set="确认" label-cancel="取消"
                            v-slot="scope" @update:model-value="onUpdate(seasoning)">
                <q-input v-model.number="scope.value" dense autofocus counter mask="####" @keyup.enter="scope.set"/>
              </q-popup-edit>
            </td>
            <td class="text-center">
              <q-btn dense round flat icon="mdi-delete" size="sm" color="grey-6"
                     :disable="seasoning.unDeletable" @click="deleteSeasoning(seasoning.id, index)"/>
            </td>
          </tr>
          </tbody>
        </q-markup-table>
      </div>
    </div>

  </q-page>
</template>

<script setup>
import { onBeforeMount, onMounted, onUnmounted, ref } from 'vue'
import { deleteAPI, getAPI, postAPI, putAPI } from 'src/api'
import { Dialog, Notify } from 'quasar'
import Sortable from 'sortablejs'

const seasonings = ref([])

onBeforeMount(async () => {
  const { data } = await getAPI('/private/seasoning/list')
  seasonings.value = data.seasonings
})

const tableBody = ref(null)
onMounted(async () => {

})

const onUpdate = async (seasoning) => {
  try {
    const { message } = await putAPI('/private/seasoning/update', {
      id: seasoning.id,
      name: seasoning.name,
      pump: seasoning.pump,
      ratio: seasoning.ratio,
    })
    Notify.create({
      message: message,
      type: 'positive',
    })
  } catch (e) {
    console.log(e.toString())
  }
}

const deleteSeasoning = async (id, index) => {
  Dialog.create({
    message: '确认删除？',
    ok: '确认',
    cancel: '取消',
    focus: 'none',
  }).onOk(async () => {
    try {
      const { message } = await deleteAPI('/private/seasoning/delete', { id: id })
      seasonings.value.splice(index, 1)
      Notify.create({
        message: message,
        type: 'positive',
      })
    } catch (e) {
      console.log(e.toString())
    }
  })
}

const addSeasoning = async () => {
  Dialog.create({
    message: '请输入调料名称',
    prompt: {
      model: '',
      type: 'text', // optional
    },
    ok: '确定',
    cancel: '取消',
    persistent: true,
  }).onOk(async (name) => {
    try {
      const { data, message } = await postAPI('private/seasoning/add', { name: name })
      Notify.create({
        message: message,
        type: 'positive',
      })
      seasonings.value.unshift(data.seasoning)
    } catch (e) {
      console.log(e.toString())
    }
  }).onCancel(() => {
  }).onDismiss(() => {
  })
}
</script>

<style lang="scss" scoped>
@import "src/pages/dish/index.scss";

</style>
