<template>
  <q-page class="flex flex-center">
    <q-card square class="q-pa-md login-card">
      <q-card-section>
        <div class="text-h6">登录</div>
      </q-card-section>
      <q-card-section>
        <q-form ref="validateForm" @submit="onSubmit" @reset="onReset" class="q-gutter-md">
          <q-input
            v-model="username"
            name="username"
            label="用户名"
            :error="$v.username.$error"
            :error-message="$v.username.required.$message"
            @blur="$v.username.$touch()"
            dense
          />
          <q-input
            v-model="password"
            name="password"
            :type="showPassword ? 'text' : 'password'"
            label="密码"
            :error="$v.password.$error"
            :error-message="$v.password.required.$message"
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
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="前往注册" to="/register"/>
        <q-btn flat label="登录" type="submit" @click="onSubmit"/>
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup>
import {ref, watch} from "vue";
import {required, helpers} from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";
import {Notify} from "quasar";
import {useUserStore} from "stores/user";
import {useRoute, useRouter} from "vue-router";

const userStore = useUserStore();

const username = ref("");
const password = ref("");

const showPassword = ref(false);

const rules = {
  username: {
    required: helpers.withMessage("用户名不能为空！", required)
  },
  password: {
    required: helpers.withMessage("密码不能为空！", required)
  },
};

const $v = useVuelidate(rules, {
  username,
  password,
});

const router = useRouter();
const route = useRoute();

const redirect = ref(null);
const otherQuery = ref({});

const getOtherQuery = (query) => {
  return Object.keys(query)
    .reduce((acc, cur) => {
      if (cur !== "redirect") {
        acc[cur] = query[cur];
      }
      return acc;
    }, {});
};

watch(() => route,
  () => {
    const query = route.query;
    if (query) {
      redirect.value = query.redirect;
      otherQuery.value = getOtherQuery(query);
    }
  },
  {immediate: true});
const onSubmit = async () => {
  const valid = await $v.value.$validate();
  if (valid) {
    const ok = await userStore.handleLogin({
      username: username.value,
      password: password.value
    });
    if (ok) {
      Notify.create("登录成功！");
      await router.push({
        path: redirect.value || "/",
        query: otherQuery.value
      });

    }
  } else {
    Notify.create({
      message: "请检查输入项！",
      type: "warning"
    });
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
