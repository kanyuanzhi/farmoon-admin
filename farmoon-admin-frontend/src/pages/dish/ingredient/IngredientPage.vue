<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-8 offset-2">
        <q-table
            flat
            :rows="rows"
            :columns="columns"
            row-key="id"
            :loading="loading"
            v-model:pagination="pagination"
            @request="onRequest"
            :pagination-label="getPaginationString"
            rows-per-page-label="每页显示："
            :filter="filter">
          <template v-slot:top="scope">
            <q-btn label="添加食材" color="primary" @click="addRow"/>
            <q-space/>
            <q-select
                dense
                options-dense
                clearable
                multiple
                v-model="typeFilter"
                :options="typeOptions"
                label="筛选类别"
                style="width: 300px"
                :disable="!enableTypeFilter"
                @update:model-value="onRequest(scope)"
            >
              <template v-slot:before>
                <q-toggle dense v-model="enableTypeFilter" @update:model-value="onRequest(scope)"/>
              </template>
            </q-select>
          </template>
          <template v-slot:body="props">
            <q-tr :props="props">
              <q-td key="index" :props="props">
                {{ props.row.index }}
              </q-td>
              <q-td key="name" :props="props">
                {{ props.row.name }}
                <q-popup-edit v-model="props.row.name" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                              @update:model-value="updateRow(props.row)">
                  <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
                </q-popup-edit>
              </q-td>
              <q-td key="type" :props="props">
                {{ typeMap[props.row.type] }}
                <q-popup-edit v-model="props.row.type" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                              @update:model-value="updateRow(props.row)">
                  <q-option-group
                      :options="typeOptions"
                      v-model="scope.value"
                  />
                </q-popup-edit>
              </q-td>
              <q-td key="operate" :props="props">
                <q-btn class="q-ml-sm" dense round flat icon="publish" size="sm" color="primary"/>
                <q-btn class="q-ml-sm" dense round flat icon="delete" size="sm" color="grey-6"
                       @click="deleteRow(props.row)"/>
              </q-td>
            </q-tr>
          </template>
        </q-table>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { deleteAPI, getAPI, postAPI, putAPI } from 'src/api'
import { getIngredientTypeInfo } from 'pages/dish'
import { Dialog, Notify } from 'quasar'
import { remove } from 'lodash'

const enableTypeFilter = ref(false)
const typeFilter = ref([])

const filter = ref('')
const loading = ref(false)

const typeMap = ref({})
const typeOptions = ref([])

const defaultPageIndex = 1
const defaultPageSize = 10
const startRow = ref(1)
const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: defaultPageIndex,
  rowsPerPage: defaultPageSize,
  rowsNumber: 10,
})
const getPaginationString = (firstRowIndex, endRowIndex, totalRowsNumber) => {
  return `${firstRowIndex}-${endRowIndex}（${totalRowsNumber}）`
}

const rows = ref([])

const columns = [
  { name: 'index', align: 'center', label: '#', field: 'index', headerStyle: 'width: 10%' },
  { name: 'name', requires: true, align: 'center', label: '名称', field: 'name', headerStyle: 'width: 30%' },
  { name: 'type', align: 'center', label: '类别', field: 'type', headerStyle: 'width: 30%' },
  { name: 'operate', align: 'center', label: '操作', headerStyle: 'width: 30%' },
]

onMounted(async () => {
  loading.value = true
  const { options, map } = await getIngredientTypeInfo()
  console.log(options, map)
  typeOptions.value = options
  typeMap.value = map
  await getRowsNumber()
  await getRows(defaultPageIndex, defaultPageSize)
  loading.value = false
})

const onRequest = async (props) => {
  loading.value = true
  await getRowsNumber()
  const { page, rowsPerPage, sortBy, descending } = props.pagination
  startRow.value = (page - 1) * rowsPerPage + 1
  const pageSize = rowsPerPage === 0 ? pagination.value.rowsNumber : rowsPerPage
  await getRows(page, pageSize)
  pagination.value.page = page
  pagination.value.rowsPerPage = rowsPerPage
  pagination.value.sortBy = sortBy
  pagination.value.descending = descending
  loading.value = false
}
const getRowsNumber = async () => {
  const { data } = await getAPI('/private/ingredient/count',
      {
        enableTypeFilter: enableTypeFilter.value,
        typeFilter: typeFilter.value?.map((val) => val.value).join(','),
      })
  if (data.count !== undefined) {
    pagination.value.rowsNumber = data.count
  } else {
    pagination.value.rowsNumber = data.data.count
  }
}

const getRows = async (pageIndex, pageSize) => {
  const { data } = await getAPI('/private/ingredient/list', {
    pageIndex: pageIndex,
    pageSize: pageSize,
    enableTypeFilter: enableTypeFilter.value,
    typeFilter: typeFilter.value?.map((val) => val.value).join(','),
  })
  if (data.ingredients !== undefined) {
    rows.value.splice(0, rows.value.length, ...data.ingredients)
  } else {
    rows.value.splice(0, rows.value.length, ...data.data.ingredients)
  }
  rows.value.forEach((item, index) => {
    item.index = index + startRow.value
    if (item.roles === null) {
      item.roles = []
    }
  })
}

const addRow = async () => {
  Dialog.create({
    message: '请输入食材名称',
    prompt: {
      model: '',
      type: 'text', // optional
    },
    ok: '确定',
    cancel: '取消',
    persistent: true,
  }).onOk(async (name) => {
    try {
      const { data, message } = await postAPI('private/ingredient/add', { name: name })
      Notify.create(message)
      rows.value.unshift(data.ingredient)
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
      type: row.type,
    }
    const { message } = await putAPI('private/ingredient/update', newData)
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
      const { message } = await deleteAPI('/private/ingredient/delete', { id: id })
      remove(rows.value, (row) => row.id === id)
      Notify.create(message)
    } catch (e) {
      console.log(e.toString())
    }
  }).onCancel(() => {

  }).onDismiss(() => {
  })
}
</script>

<style scoped lang="scss">
:deep(.q-table__top) {
  height: 64px;
}
</style>
