<template>
  <my-upload field="img"
             @crop-success="cropSuccess"
             @crop-upload-success="cropUploadSuccess"
             @crop-upload-fail="cropUploadFail"
             @src-file-set="srcFileSet"
             v-model="shown"
             :width="300"
             :height="300"
             :url="uploadUrl"
             :params="params"
             img-format="png"></my-upload>
</template>

<script setup>
import myUpload from 'vue-image-crop-upload';
import {ref} from "vue";
import {Notify} from "quasar";

const emit = defineEmits(["updateSuccess"])

const shown = ref(false)
const id = ref(0)
const from = ref(null)

const params = ref({
  id: id.value
})

const uploadUrl = process.env.API + "private/user/update-avatar"
let avatarUrl = ""

const srcFileSet = (fileName, fileType, fileSize) => {
}
const cropSuccess = (imgDataUrl, field) => {
  avatarUrl = imgDataUrl
}

const cropUploadSuccess = (jsonData, field) => {
  const {code, message} = jsonData
  if (code === 1) {
    Notify.create(message)
    emit("updateSuccess", id.value, avatarUrl, from.value)
  } else {
    Notify.create({
      message: message,
      type: "negative"
    })
  }
}

const cropUploadFail = (status, field) => {
  Notify.create({
    message: status,
    type: "negative"
  })
}

const show = (userId, fromObj) => {
  shown.value = true
  id.value = userId
  from.value = fromObj
  params.value.id = userId
  console.log(userId, fromObj)
}

defineExpose({
  show
})
</script>

<style lang="scss" scoped>

</style>
