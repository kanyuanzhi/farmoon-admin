<template>
  <q-card style="width: 400px">
    <!--    <q-card-section>-->
    <!--      <div class="text-subtitle1">添加用户</div>-->
    <!--    </q-card-section>-->
    <q-card-section>
      <q-form ref="validateForm" @submit="onSubmit" @reset="onReset" class="q-gutter-md">
        <div class="text-subtitle2 text-center">*必填项</div>
        <q-separator class="q-mt-sm"/>
        <q-input
          v-model="username"
          name="username"
          label="用户名"
          :error="$v.username.$error"
          :error-message="$v.username.required.$invalid? $v.username.required.$message : $v.username.minLength.$message"
          @blur="$v.username.$touch()"
          dense
        />
        <q-input
          v-model="password"
          name="password"
          label="密码"
          :error="$v.password.$error"
          :error-message="$v.password.required.$invalid? $v.password.required.$message : $v.password.minLength.$message"
          @blur="$v.password.$touch()"
          dense
        />
        <div class="text-subtitle2 text-center">选填项</div>
        <q-separator class="q-mt-sm"/>
        <q-input v-model="nickname" name="nickname" label="昵称" dense/>
        <q-input v-model="realName" name="realName" label="真实姓名" dense/>
        <q-select
          v-model="gender"
          name="gender"
          :options="genderOptions"
          label="性别"
          dense
        />
        <q-input v-model="mobile" name="mobile" label="手机号" dense/>
        <q-input v-model="email" name="email" label="邮箱" dense/>
        <q-select
          v-model="roles"
          name="roles"
          multiple
          :options="roleOptions"
          label="权限"
          dense
        />
      </q-form>
    </q-card-section>
    <q-card-actions align="right">
      <q-btn flat label="取消" v-close-popup/>
      <q-btn flat label="确认" @click="onSubmit"/>
    </q-card-actions>
  </q-card>
</template>

<script setup>
import {ref} from "vue";
import {helpers, minLength, required, sameAs} from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import {useRouter} from "vue-router";
import {postAPI} from "src/api";
import {Notify} from "quasar";
import {genderOptions, genderMap, roleOptions} from "pages/user/user";

const emit = defineEmits(["addSuccess"])

const username = ref("")
const password = ref("")
const nickname = ref("")
const realName = ref("")
const gender = ref({label: "未知", value: "N"})
const mobile = ref("")
const email = ref("")
const roles = ref([])

const rules = {
  username: {
    required: helpers.withMessage("用户名不能为空！", required),
    minLength: helpers.withMessage("用户名长度不能小于6！", minLength(6))
  },
  password: {
    required: helpers.withMessage("密码不能为空！", required),
    minLength: helpers.withMessage("密码长度不能小于6！", minLength(6))
  },
};

const $v = useVuelidate(rules, {
  username,
  password,
});

const router = useRouter();
const onSubmit = async () => {
  try {
    const valid = await $v.value.$validate();
    if (valid) {
      const {data, message} = await postAPI("/private/user/add", {
        username: username.value,
        password: password.value,
        nickname: nickname.value,
        realName: realName.value,
        gender: gender.value.value,
        mobile: mobile.value,
        email: email.value,
        roles: roles.value.map(role => role.value),
      });
      Notify.create({
        message: message,
        type: "positive",
      });
      emit("addSuccess", data)
    } else {
      Notify.create({
        message: "请检查输入项！",
        type: "warning",
      });
    }
  } catch (e) {
    console.log(e.toString());
  }
};

const onReset = () => {

};
</script>

<style lang="scss" scoped>

</style>
