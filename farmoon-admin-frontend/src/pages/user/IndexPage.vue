<template>
  <q-page class="q-pa-md">
    <q-table
      ref="table"
      :rows="rows"
      :columns="columns"
      row-key="id"
      flat
      :loading="loading"
      v-model:pagination="pagination"
      :selection="isDeleteInBatch? 'multiple' : 'none'"
      v-model:selected="selectedRows"
      :selected-rows-label="getSelectedString"
      @request="onRequest"
      :pagination-label="getPaginationString"
      rows-per-page-label="每页显示："
      :filter="filter"
    >
      <template v-slot:header-selection="scope">
        <q-checkbox dense v-model="scope.selected"/>
      </template>
      <template v-slot:top="scope">
        <q-btn label="添加用户" color="primary" @click="addUserDialogShown=true"/>
        <q-btn v-if="!isDeleteInBatch" class="q-ml-sm" label="批量删除用户" color="grey-6"
               @click.prevent="isDeleteInBatch=true"/>
        <template v-else>
          <q-btn outline class="q-ml-sm" label="取消删除" color="grey-6" @click="deleteCancel"/>
          <q-btn outline class="q-ml-sm" label="确认删除" color="negative"
                 @click="deleteRows(selectedRows)"/>
        </template>
        <q-space/>

        <q-toggle v-model="enableGenderFilter" @update:model-value="onRequest(scope)"></q-toggle>
        <q-select
          dense
          v-model="genderFilter"
          :options="genderOptions"
          label="筛选性别"
          style="width: 90px"
          :disable="!enableGenderFilter"
          @update:model-value="onRequest(scope)"
        />
        <q-toggle v-model="enableRolesFilter" @update:model-value="onRequest(scope)"></q-toggle>
        <q-select
          dense
          v-model="rolesFilter"
          multiple
          :options="roleOptions"
          label="筛选权限"
          style="width: 160px"
          :disable="!enableRolesFilter"
          @update:model-value="onRequest(scope)"
        />
        <q-input dense class="q-ml-md" debounce="300" v-model="filter" placeholder="模糊搜索">
          <template v-slot:append>
            <q-icon name="search"/>
          </template>
        </q-input>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-if="isDeleteInBatch" key="select">
            <q-checkbox dense v-model="props.selected"/>
          </q-td>
          <q-td key="index" :props="props">
            {{ props.row.index }}
          </q-td>
          <q-td key="avatar" :props="props">
            <q-avatar size="md" class="cursor-pointer">
              <img :src="props.row.avatar" alt="" @click="photoUploader.show(props.row.id, props.row)"/>
            </q-avatar>
          </q-td>
          <q-td key="username" :props="props">
            {{ props.row.username }}
          </q-td>
          <q-td key="nickname" :props="props">
            {{ props.row.nickname }}
            <q-popup-edit v-model="props.row.nickname" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
            </q-popup-edit>
          </q-td>
          <q-td key="realName" :props="props">
            {{ props.row.realName }}
            <q-popup-edit v-model="props.row.realName" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
            </q-popup-edit>
          </q-td>
          <q-td key="gender" :props="props">
            {{ genderMap[props.row.gender] }}
            <q-popup-edit v-model="props.row.gender" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-option-group
                :options="genderOptions"
                type="radio"
                v-model="scope.value"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="mobile" :props="props">
            {{ props.row.mobile }}
            <q-popup-edit v-model="props.row.mobile" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
            </q-popup-edit>
          </q-td>
          <q-td key="email" :props="props">
            {{ props.row.email }}
            <q-popup-edit v-model="props.row.email" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
            </q-popup-edit>
          </q-td>
          <q-td key="roles" :props="props">
            <q-badge v-for="role in props.row.roles.sort()" :key="role" class="q-ml-xs" :color="roleColorMap[role]">
              {{ role }}
            </q-badge>
            <q-popup-edit v-model="props.row.roles" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRoles(props.row)">
              <q-option-group
                :options="roleOptions"
                type="checkbox"
                v-model="scope.value"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="operate" :props="props">
            <q-btn dense round flat icon="key" size="md" color="primary" @click="updatePassword(props.row)"/>
            <q-btn class="q-ml-sm" dense round flat icon="mdi-delete" size="md" color="grey-6"
                   @click="deleteRows([props.row])"/>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-dialog v-model="addUserDialogShown" persistent>
      <AddUserDialog @add-success="onAddUserSuccess"/>
    </q-dialog>
    <PhotoUploader ref="photoUploader" @update-success="onUpdateAvatarSuccess"/>
  </q-page>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {deleteAPI, getAPI, putAPI} from "src/api";
import {Dialog, Notify} from "quasar";
import {remove} from "lodash";
import AddUserDialog from "pages/user/components/AddUserDialog.vue";
import PhotoUploader from "pages/user/components/PhotoUploader.vue";
import {genderOptions, genderMap, roleOptions, roleColorMap} from "pages/user/index";

const enableGenderFilter = ref(false)
const genderFilter = ref(null)

const enableRolesFilter = ref(false)
const rolesFilter = ref([])

const filter = ref("")
const isDeleteInBatch = ref(false)

const loading = ref(false)

const table = ref(null)
const rows = ref([])

const columns = [
  {name: 'index', align: 'center', label: '#', field: 'index', headerStyle: "width: 5%"},
  {name: 'avatar', align: 'center', label: '头像', field: 'avatar', headerStyle: "width: 8%"},
  {name: 'username', required: true, align: 'center', label: '用户名', field: 'username', headerStyle: "width: 10%"},
  {name: 'nickname', align: 'center', label: '昵称', field: 'nickname', headerStyle: "width: 10%"},
  {name: 'realName', align: 'center', label: '真实姓名', field: 'realName', headerStyle: "width: 10%"},
  {name: 'gender', align: 'center', label: '性别', field: 'gender', headerStyle: "width: 5%"},
  {name: 'mobile', align: 'center', label: '手机号', field: 'mobile', headerStyle: "width: 10%"},
  {name: 'email', align: 'center', label: '邮箱', field: 'email', headerStyle: "width: 15%"},
  {name: 'roles', align: 'center', label: '权限', field: 'roles', headerStyle: "width: 15%"},
  {name: 'operate', align: 'center', label: '操作', headerStyle: "width: 12%"},
]

const selectedRows = ref([])
const getSelectedString = () => {
  return selectedRows.value.length === 0 ? '' : `已选中${selectedRows.value.length}条记录`
}

const defaultPageIndex = 1
const defaultPageSize = 10
const startRow = ref(1)

const pagination = ref({
  sortBy: 'desc',
  descending: false,
  page: defaultPageIndex,
  rowsPerPage: defaultPageSize,
  rowsNumber: 10
})

const getPaginationString = (firstRowIndex, endRowIndex, totalRowsNumber) => {
  return `${firstRowIndex}-${endRowIndex}（${totalRowsNumber}）`
}

const addUserDialogShown = ref(false)
const photoUploader = ref(null)

onMounted(async () => {
  loading.value = true
  await getRowsNumber()
  await getRows(defaultPageIndex, defaultPageSize)
  loading.value = false
})

const onRequest = async (props) => {
  loading.value = true
  await getRowsNumber()
  const {page, rowsPerPage, sortBy, descending} = props.pagination
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
  const {data} = await getAPI("/private/user/count", {
      filter: filter.value.trim(),
      enableGenderFilter: enableGenderFilter.value,
      genderFilter: genderFilter.value === null ? "" : genderFilter.value.value,
      enableRolesFilter: enableRolesFilter.value,
      rolesFilter: rolesFilter.value.map((val) => val.value).join(",")
    })
  if (data.count !== undefined) {
    pagination.value.rowsNumber = data.count
  } else {
    pagination.value.rowsNumber = data.data.count
  }
}

const getRows = async (pageIndex, pageSize) => {
  const {data} = await getAPI("/private/user/list", {
    pageIndex: pageIndex,
    pageSize: pageSize,
    filter: filter.value.trim(),
    enableGenderFilter: enableGenderFilter.value,
    genderFilter: genderFilter.value === null ? "" : genderFilter.value.value,
    enableRolesFilter: enableRolesFilter.value,
    rolesFilter: rolesFilter.value.map((val) => val.value).join(",")
  })
  if (data.users !== undefined) {
    rows.value.splice(0, rows.value.length, ...data.users)
  } else {
    rows.value.splice(0, rows.value.length, ...data.data.users)
  }
  rows.value.forEach((item, index) => {
    item.index = index + startRow.value
    if (item.roles === null) {
      item.roles = []
    }
  })
}

const updateRow = async (row) => {
  try {
    const newData = {
      id: row.id,
      nickname: row.nickname,
      realName: row.realName,
      gender: row.gender,
      mobile: row.mobile,
      email: row.email,
    }
    const {message} = await putAPI("private/user/update", newData)
    Notify.create(message)
  } catch (e) {
    console.log(e.toString())
  }
}

const updateRoles = async (row) => {
  try {
    const {message} = await putAPI("private/user/update-roles", {id: row.id, roles: row.roles})
    Notify.create(message)
  } catch (e) {
    console.log(e.toString())
  }
}

const updatePassword = async (row) => {
  Dialog.create({
    message: "请输入新密码：",
    prompt: {
      model: "",
      type: "text",
    },
    ok: "确定",
    cancel: "取消"
  }).onOk(async (data) => {
    console.log(data)
    if (data.trim().length < 6) {
      Notify.create({
        message: "密码长度不能小于6",
        type: "warning"
      })
      return
    }
    try {
      const {message} = await putAPI("private/user/update-password", {id: row.id, password: data})
      Notify.create(message)
    } catch (e) {
      console.log(e.toString())
    }
  }).onDismiss(() => {
  })
}

const deleteRows = async (selectedRows) => {
  if (selectedRows.length === 0) {
    Notify.create({
      message: "请选择要删除的用户",
      type: "warning"
    })
    return
  }
  Dialog.create({
    message: "确认删除？",
    ok: "确认",
    cancel: "取消",
    focus: "none"
  }).onOk(async () => {
    try {
      const ids = selectedRows.map((row) => row.id)
      const {message} = await deleteAPI("/private/user/delete", {ids: ids})
      remove(rows.value, (row) => ids.includes(row.id))
      Notify.create(message)
    } catch (e) {
      console.log(e.toString())
    }
  })
}

const deleteCancel = () => {
  isDeleteInBatch.value = false
  selectedRows.value = []
}

const onAddUserSuccess = (user) => {
  addUserDialogShown.value = false
  onRequest(table.value)
}

const onUpdateAvatarSuccess = (userId, avatarUrl, fromObj) => {
  fromObj.avatar = avatarUrl
}
</script>

<style lang="scss" scoped>

</style>
