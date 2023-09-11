<template>
  <q-page class="flex flex-center">
    <q-card square class="q-pa-md login-card">
      <q-card-section>
        <div class="text-h6">注册</div>
      </q-card-section>
      <q-card-section>
        <q-form ref="validateForm" @submit="onSubmit" @reset="onReset" class="q-gutter-md">
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
            :type="showPassword ? 'text' : 'password'"
            label="密码"
            :error="$v.password.$error"
            :error-message="$v.password.required.$invalid? $v.password.required.$message : $v.password.minLength.$message"
            @blur="$v.password.$touch()"
            dense
          >
            <template v-slot:after>
              <q-icon
                :name="showPassword ? 'visibility' : 'visibility_off'"
                class="cursor-pointer"
                @click="showPassword = !showPassword"
              />
            </template>
          </q-input>
          <q-input
            v-model="repeatPassword"
            name="repeatPassword"
            :type="showRepeatPassword ? 'text' : 'password'"
            label="重复密码"
            :error="$v.repeatPassword.$error"
            :error-message="$v.repeatPassword.required.$invalid? $v.repeatPassword.required.$message : $v.repeatPassword.sameAsPassword.$message"
            @blur="$v.repeatPassword.$touch()"
            dense
          >
            <template v-slot:after>
              <q-icon
                :name="showRepeatPassword ? 'visibility' : 'visibility_off'"
                class="cursor-pointer"
                @click="showRepeatPassword = !showRepeatPassword"
              />
            </template>
          </q-input>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="前往登录" to="/login"/>
        <q-btn flat label="注册" type="submit" @click="onSubmit"/>
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref } from "vue";
import { required, helpers, sameAs, minLength } from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import { Notify } from "quasar";
import { postAPI } from "src/api";
import { useRouter } from "vue-router";

const username = ref("");
const password = ref("");
const repeatPassword = ref("");

const showPassword = ref(false);
const showRepeatPassword = ref(false);

const rules = {
  username: {
    required: helpers.withMessage("用户名不能为空！", required),
    minLength: helpers.withMessage("用户名长度不能小于6！", minLength(6))
  },
  password: {
    required: helpers.withMessage("密码不能为空！", required),
    minLength: helpers.withMessage("密码长度不能小于6！", minLength(6))
  },
  repeatPassword: {
    required: helpers.withMessage("密码不能为空！", required),
    sameAsPassword: helpers.withMessage("两次密码不一致！", sameAs(password)),
  },
};

const $v = useVuelidate(rules, {
  username,
  password,
  repeatPassword
});

const router = useRouter();
const onSubmit = async () => {
  try {
    const valid = await $v.value.$validate();
    if (valid) {
      const res = await postAPI("/public/register", {
        username: username.value,
        password: password.value,
        repeatPassword: repeatPassword.value
      });
      Notify.create({
        message: "注册成功，请登录",
        type: "positive",
      });
      await router.push("/login");
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
.login-card {
  width: 400px;
}
</style>
