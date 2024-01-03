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
        :selection="isDeleteInBatch || isExportInBatch? 'multiple' : 'none'"
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
        <template v-if="!isDeleteInBatch">
          <q-btn v-if="!isExportInBatch" class="q-ml-sm" label="批量导出菜品" color="primary"
                 @click.prevent="isExportInBatch=true"/>
          <template v-else>
            <q-btn outline class="q-ml-sm" label="取消导出" color="grey-6" @click="exportCancel"/>
            <q-btn outline class="q-ml-sm" label="确认导出" color="positive"
                   @click="exportSteps(selectedRows)"/>
          </template>
        </template>
        <template v-if="!isExportInBatch">
          <q-btn v-if="!isDeleteInBatch" class="q-ml-sm" label="批量删除菜品" color="grey-6"
                 @click.prevent="isDeleteInBatch=true"/>
          <template v-else>
            <q-btn outline class="q-ml-sm" label="取消删除" color="grey-6" @click="deleteCancel"/>
            <q-btn outline class="q-ml-sm" label="确认删除" color="negative"
                   @click="deleteRows(selectedRows)"/>
          </template>
        </template>
        <q-space/>
        <q-toggle v-model="enableOwnerFilter" @update:model-value="onRequest(scope)"></q-toggle>
        <q-select
            dense
            options-dense
            clearable
            v-model="ownerFilter"
            :options="ownerOptions"
            label="筛选用户"
            style="width: 300px"
            :disable="!enableOwnerFilter"
            @update:model-value="onRequest(scope)"
        />
        <q-input dense class="q-ml-md" debounce="300" v-model="filter" placeholder="模糊搜索(名称或用户)">
          <template v-slot:append>
            <q-icon name="search"/>
          </template>
        </q-input>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-if="isDeleteInBatch || isExportInBatch" key="select">
            <q-checkbox dense v-model="props.selected"/>
          </q-td>
          <q-td key="index" :props="props">
            {{ props.row.index }}
          </q-td>
          <q-td key="uuid" :props="props">
            {{ props.row.uuid }}
          </q-td>
          <q-td key="image" :props="props">
            <q-img
                class="cursor-pointer"
                :src="props.row.image"
                width="80px"
                height="45px"
                fit="fill"
                alt="修改"
                @click="dishImageUploader.show(props.row.id, props.row)"
            />
          </q-td>
          <q-td key="name" :props="props">
            {{ props.row.name }}
            <q-popup-edit v-model="props.row.name" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-input v-model="scope.value" dense autofocus counter @keyup.enter="scope.set"/>
            </q-popup-edit>
          </q-td>
          <q-td key="cuisine" :props="props">
            {{ cuisineMap[props.row.cuisine] }}
            <q-popup-edit v-model="props.row.cuisine" buttons label-set="确认" label-cancel="取消" v-slot="scope"
                          @update:model-value="updateRow(props.row)">
              <q-option-group
                  :options="ownerOptions"
                  v-model="scope.value"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="owner" :props="props">
            {{ props.row.owner }}
          </q-td>
          <q-td key="operate" :props="props">
            <q-btn dense round flat icon="add" size="sm" color="teal-6" @click="addToOfficials(props.row)">
              <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">添加至官方菜品</q-tooltip>
            </q-btn>
            <q-btn class="q-ml-sm" dense round flat icon="qr_code" size="sm" color="teal-6"
                   @click="exportQrCode(props.row)">
              <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">导出二维码</q-tooltip>
            </q-btn>
            <q-btn class="q-ml-sm" dense round flat icon="text_snippet" size="sm" color="teal-6"
                   @click="exportSteps([props.row])">
              <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">导出步骤</q-tooltip>
            </q-btn>
            <q-btn class="q-ml-sm" dense round flat icon="delete" size="sm" color="grey-6"
                   @click="deleteRows([props.row])">
              <q-tooltip anchor="top middle" self="bottom middle" :offset="[10, 10]">删除</q-tooltip>
            </q-btn>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <DishImageUploader ref="dishImageUploader" @update-success="onUpdateImageSuccess"/>
  </q-page>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { deleteAPI, getAPI, postAPI, putAPI } from "src/api";
import { Dialog, Notify } from "quasar";
import { remove } from "lodash";
import DishImageUploader from "pages/dish/components/DishImageUploader.vue";
import { getCuisineInfo, getOwners } from "pages/dish/index";
import { useDishStore } from "stores/dish";
import { useRouter } from "vue-router";
import { saveAs } from "file-saver";

const ownerOptions = ref([]);
const cuisineMap = ref({});

const enableOwnerFilter = ref(false);
const ownerFilter = ref("");

const filter = ref("");
const isDeleteInBatch = ref(false);
const isExportInBatch = ref(false);

const loading = ref(false);

const table = ref(null);
const rows = ref([]);

const columns = [
  { name: "index", align: "center", label: "#", field: "index", headerStyle: "width: 5%" },
  { name: "uuid", align: "center", label: "uuid", field: "uuid", headerStyle: "width: 20%" },
  { name: "image", align: "center", label: "图片", field: "image", headerStyle: "width: 10%" },
  { name: "name", required: true, align: "center", label: "名称", field: "name", headerStyle: "width: 20%" },
  { name: "cuisine", align: "center", label: "菜系", field: "cuisine", headerStyle: "width: 10%" },
  { name: "owner", align: "center", label: "用户", field: "owner", headerStyle: "width: 20%" },
  { name: "operate", align: "center", label: "操作", headerStyle: "width: 15%" },
];

const selectedRows = ref([]);
const getSelectedString = () => {
  return selectedRows.value.length === 0 ? "" : `已选中${selectedRows.value.length}条记录`;
};

const defaultPageIndex = 1;
const defaultPageSize = 10;
const startRow = ref(1);

const pagination = ref({
  sortBy: "desc",
  descending: false,
  page: defaultPageIndex,
  rowsPerPage: defaultPageSize,
  rowsNumber: 10,
});

const getPaginationString = (firstRowIndex, endRowIndex, totalRowsNumber) => {
  return `${firstRowIndex}-${endRowIndex}（${totalRowsNumber}）`;
};

const dishImageUploader = ref(null);

onMounted(async () => {
  loading.value = true;
  const { map } = await getCuisineInfo();
  cuisineMap.value = map;
  const { options } = await getOwners();
  ownerOptions.value = options;
  await getRowsNumber();
  await getRows(defaultPageIndex, defaultPageSize);
  loading.value = false;
});

const onRequest = async (props) => {
  loading.value = true;
  await getRowsNumber();
  const { page, rowsPerPage, sortBy, descending } = props.pagination;
  startRow.value = (page - 1) * rowsPerPage + 1;
  const pageSize = rowsPerPage === 0 ? pagination.value.rowsNumber : rowsPerPage;
  await getRows(page, pageSize);
  pagination.value.page = page;
  pagination.value.rowsPerPage = rowsPerPage;
  pagination.value.sortBy = sortBy;
  pagination.value.descending = descending;
  loading.value = false;
};

const getRowsNumber = async () => {
  const { data } = await getAPI("/private/dish/count",
      {
        filter: filter.value.trim(),
        enableOwnerFilter: enableOwnerFilter.value,
        ownerFilter: ownerFilter.value?.value,
        isOfficial: false,
      });
  if (data.count !== undefined) {
    pagination.value.rowsNumber = data.count;
  } else {
    pagination.value.rowsNumber = data.data.count;
  }
};

const getRows = async (pageIndex, pageSize) => {
  const { data } = await getAPI("/private/dish/list", {
    pageIndex: pageIndex,
    pageSize: pageSize,
    filter: filter.value.trim(),
    enableOwnerFilter: enableOwnerFilter.value,
    ownerFilter: ownerFilter.value?.value,
    isOfficial: false,
  });
  if (data.dishes !== undefined) {
    rows.value.splice(0, rows.value.length, ...data.dishes);
  } else {
    rows.value.splice(0, rows.value.length, ...data.data.dishes);
  }
  rows.value.forEach((item, index) => {
    item.index = index + startRow.value;
    if (item.roles === null) {
      item.roles = [];
    }
  });
};

const updateRow = async (row) => {
  try {
    const newData = {
      id: row.id,
      name: row.name,
      cuisine: row.cuisine,
    };
    const { message } = await putAPI("private/dish/update", newData);
    Notify.create(message);
  } catch (e) {
    console.log(e.toString());
  }
};

const deleteRows = async (selectedRows) => {
  if (selectedRows.length === 0) {
    Notify.create({
      message: "请选择要删除的菜品",
      type: "warning",
    });
    return;
  }
  Dialog.create({
    message: "确认删除？",
    ok: "确认",
    cancel: "取消",
    focus: "none",
  }).onOk(async () => {
    try {
      const ids = selectedRows.map((row) => row.id);
      const { message } = await deleteAPI("/private/dish/delete-personals", { ids: ids });
      await onRequest(table.value);
      Notify.create(message);
      deleteCancel()
    } catch (e) {
      console.log(e.toString());
    }
  });
};

const addToOfficials = async (row) => {
  try {
    const { message } = await postAPI("private/dish/add-to-officials", { id: row.id });
    Notify.create(message);
    // await onRequest(table.value);
    // rows.value.push(data.dish);
  } catch (e) {
    console.log(e.toString());
  }
};

const toppingRow = async (row) => {
  try {
    const { message, data } = await putAPI("private/dish/topping", { id: row.id });
    Notify.create(message);
    await onRequest(table.value);
  } catch (e) {
    console.log(e.toString());
  }
};

const exportCancel = () => {
  isExportInBatch.value = false;
  selectedRows.value = [];
};

const exportSteps = async (selectedRows) => {
  try {
    if (selectedRows.length === 0) {
      Notify.create({
        message: "请选择要导出步骤的菜品",
        type: "warning",
      });
      return;
    }
    if (selectedRows.length === 1) {
      const ids = selectedRows.map((row) => row.id);
      const { message, data } = await postAPI("/private/dish/export-steps", { ids: ids });
      let stepsStr = "";
      let stepsTxt = "";
      data.dishes.forEach((dish) => {
        dish.steps.forEach((step, index) => {
          stepsStr += "<span class='text-subtitle2'>" + String(index + 1) + ". " + step.instructionName + "</span><br>";
          stepsTxt += String(index + 1) + ". " + step.instructionName + "\n";
        });
      });
      Dialog.create({
        title: data.dishes[0].name,
        message: stepsStr,
        ok: "导出",
        cancel: "取消",
        focus: "none",
        html: true,
      }).onOk(async () => {
        try {
          let strData = new Blob([stepsTxt], { type: "text/plain;charset=utf-8" });
          await saveAs(strData, data.dishes[0].name + ".txt");
        } catch (e) {
          console.log(e.toString());
        }
      });
    } else {
      const ids = selectedRows.map((row) => row.id);
      const { message, data } = await postAPI("/private/dish/export-steps", { ids: ids });
      Notify.create(message);
      data.dishes.forEach((dish) => {
        let stepsTxt = "";
        dish.steps.forEach((step, index) => {
          stepsTxt += String(index + 1) + ". " + step.instructionName + "\n";
        });
        let strData = new Blob([stepsTxt], { type: "text/plain;charset=utf-8" });
        saveAs(strData, dish.name + ".txt");
      });
    }
    exportCancel()
  } catch (e) {
    console.log(e.toString());
  }
};

const exportQrCode = async (row) => {
  try {
    const { message, data } = await getAPI("/private/dish/get-qr-code", { id: row.id });
    const binaryData = atob(data.qrCode);
    const arrayBuffer = new ArrayBuffer(binaryData.length);
    const uint8Array = new Uint8Array(arrayBuffer);
    for (let i = 0; i < binaryData.length; i++) {
      uint8Array[i] = binaryData.charCodeAt(i);
    }
    const blob = new Blob([uint8Array], { type: "image/png" });
    saveAs(blob, row.name + ".png");
    Notify.create(message);
  } catch (e) {
    console.log(e.toString());
  }
};

const deleteCancel = () => {
  isDeleteInBatch.value = false;
  selectedRows.value = [];
};

const onUpdateImageSuccess = (imageUrl, fromObj) => {
  fromObj.image = imageUrl;
};

const dishStore = useDishStore();
const router = useRouter();
const editDish = (row) => {
  dishStore.setEditingDish(row);
  router.push("/dish/edit");
};
</script>

<style lang="scss" scoped>

</style>
