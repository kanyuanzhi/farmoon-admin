<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-3">
        <q-card>
          <q-card-section class="text-center bg-cyan-2">
            <q-avatar size="8em" class="cursor-pointer">
              <img :src="userStore.getAvatar()" alt="" @click="photoUploader.show(userStore.getId(), null)"/>
            </q-avatar>
          </q-card-section>
          <q-card-section v-if="!isEditing">
            <q-btn class="full-width" label="编辑" unelevated color="grey-5" size="sm" @click="isEditing=true"/>
          </q-card-section>

          <q-card-section>
            <q-list padding>
              <q-item>
                <q-item-section side>用户名</q-item-section>
                <q-item-section>
                  <span>{{ userStore.getUsername() }}</span>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>昵称</q-item-section>
                <q-item-section>
                  <span v-if="!isEditing">{{ userStore.getNickname() }}</span>
                  <q-input v-else v-model="nicknameEdited" dense/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>真实姓名</q-item-section>
                <q-item-section>
                  <span v-if="!isEditing">{{ userStore.getRealName() }}</span>
                  <q-input v-else v-model="realNameEdited" dense/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>性别</q-item-section>
                <q-item-section>
                  <span v-if="!isEditing">{{ genderMap[userStore.getGender()] }}</span>
                  <q-select
                    v-else
                    dense
                    options-dense
                    :options="genderOptions"
                    v-model="genderEdited"
                  />
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>手机号</q-item-section>
                <q-item-section>
                  <span v-if="!isEditing">{{ userStore.getMobile() }}</span>
                  <q-input v-else v-model="mobileEdited" dense/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>邮箱</q-item-section>
                <q-item-section>
                  <span v-if="!isEditing">{{ userStore.getEmail() }}</span>
                  <q-input v-else v-model="emailEdited" dense/>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section side>权限</q-item-section>
                <q-item-section>
                  <div>
                    <q-badge v-for="role in userStore.getRoles().sort()" :key="role" class="q-ml-xs"
                             :color="roleColorMap[role]">
                      {{ role }}
                    </q-badge>
                  </div>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
          <q-card-actions v-if="isEditing" align="right">
            <q-btn label="保存" unelevated color="positive" size="sm" @click="onSave"/>
            <q-btn label="取消" unelevated color="grey-5" size="sm" @click="onCancel"/>
          </q-card-actions>
        </q-card>
      </div>
      <div class="col-8"></div>
    </div>
    <PhotoUploader ref="photoUploader" @update-success="onUpdateAvatarSuccess"/>
  </q-page>
</template>

<script setup>
import {useUserStore} from "stores/user";
import {genderMap, genderOptions, roleColorMap} from "pages/user/user";
import {ref} from "vue";
import {putAPI} from "src/api";
import {Notify} from "quasar";
import PhotoUploader from "pages/user/profile/PhotoUploader.vue";

const userStore = useUserStore();

const isEditing = ref(false)
const nicknameEdited = ref(userStore.getNickname())
const realNameEdited = ref(userStore.getRealName())
const genderEdited = ref({
  label: genderMap[userStore.getGender()],
  value: userStore.getGender()
})
const mobileEdited = ref(userStore.getMobile())
const emailEdited = ref(userStore.getEmail())

const photoUploader = ref(null)

const onSave = async () => {
  try {
    const newData = {
      id: userStore.getId(),
      nickname: nicknameEdited.value,
      realName: realNameEdited.value,
      gender: genderEdited.value.value,
      mobile: mobileEdited.value,
      email: emailEdited.value,
    }
    const {message} = await putAPI("private/user/update", newData)
    isEditing.value = false

    userStore.setNickname(nicknameEdited.value)
    userStore.setRealName(realNameEdited.value)
    userStore.setGender(genderEdited.value)
    userStore.setMobile(mobileEdited.value)
    userStore.setEmail(emailEdited.value)
    Notify.create({
      message: message,
      type: "positive"
    })
  } catch (e) {
    console.log(e)
  }
}

const onCancel = () => {
  isEditing.value = false
  nicknameEdited.value = userStore.getNickname()
  realNameEdited.value = userStore.getRealName()
  genderEdited.value = {
    label: genderMap[userStore.getGender()],
    value: userStore.getGender()
  }
  mobileEdited.value = userStore.getMobile()
  emailEdited.value = userStore.getEmail()
}

const onUpdateAvatarSuccess = (userId, avatarUrl, fromObj) => {
  userStore.setAvatar(avatarUrl)
}
</script>

<style lang="scss" scoped>
</style>
