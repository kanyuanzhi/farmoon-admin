<template>
  <q-dialog v-model="shown" transition-show="scale" transition-hide="scale">
    <q-uploader
      label="上传菜品图片（100KB以内）"
      :url=url
      style="max-width: 300px"
      color="teal-6"
      accept=".jpg,.png"
      :max-file-size="100*1024"
      :form-fields="[{name:'id', value:id}]"
      field-name="file"
      @uploaded="onUploaded"
      @rejected="onRejected"
      auto-upload
    />
  </q-dialog>
</template>

<script setup>
import {ref} from "vue";
import {Notify} from "quasar";

const emit = defineEmits(["updateSuccess"]);

const url = process.env.API + "private/dish/update-image"
const uuid = ref("");
const id = ref("");
const from = ref(null)
const shown = ref(false);

const show = (dishId, fromObj) => {
  shown.value = true;
  id.value = String(dishId);
  from.value = fromObj
};

const onUploaded = (info) => {
  const xhr = info.xhr;
  const {data, message} = JSON.parse(xhr.response);
  emit("updateSuccess", data, from.value);
  Notify.create(message);
  shown.value = false
};

const onRejected = (info) => {
  Notify.create({
    message: "上传失败：" + info[0].failedPropValidation,
    type: "negative"
  });
};

defineExpose({
  show
});
</script>

<style lang="scss" scoped>

</style>
