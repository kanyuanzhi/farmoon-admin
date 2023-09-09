<template>
  <div class="q-pa-sm q-gutter-sm">
    <q-breadcrumbs class="text-white" active-color="white">
      <q-breadcrumbs-el v-for="item in breadcrumbs" :key="item.path"
                        :label="item.meta.title" :icon="item.meta.icon" :to="item.path"/>
    </q-breadcrumbs>
  </div>
</template>

<script setup>
import {onMounted, ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()

const isDashboard = (route) => {
  const name = route && route.name
  if (!name) {
    return false
  }
  return name.trim().toLocaleLowerCase() === "home".toLocaleLowerCase()
}

const getBreadcrumbs = (route) => {
  const matched = route.matched.filter(item => item.meta && item.meta.title)
  const first = matched[0]

  if (!isDashboard(first)) {
    matched.unshift({path: '/home', name: "home", meta: {title: '主页', icon: "home"}})
  }

  return matched
}

const breadcrumbs = ref(getBreadcrumbs(route))

watch(route, () => {
  breadcrumbs.value = getBreadcrumbs(route)
})


</script>

<style lang="scss" scoped>
</style>
