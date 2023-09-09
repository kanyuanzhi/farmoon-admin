<template>
  <div v-if="!item.hidden">
    <template
      v-if="hasOneShowingChild(item.children, item) && (!onlyOneChild.children||onlyOneChild.noShowingChildren) && !item.alwaysShow">
      <q-expansion-item
        v-if="onlyOneChild.meta"
        expand-separator
        :icon="onlyOneChild.meta.icon"
        :label="onlyOneChild.meta.title"
        :header-inset-level="level"
        :to="resolvePath(onlyOneChild.path)"
        hide-expand-icon
      >
      </q-expansion-item>
    </template>
    <template v-else>
      <q-expansion-item
        v-if="item.meta"
        expand-separator
        :icon="item.meta.icon"
        :label="item.meta.title"
        :header-inset-level="level"
      >
        <SidebarItem
          v-for="item in item.children"
          :key="item.path"
          :item="item"
          :level="level + 0.5"
          :basePath="resolvePath(item.path)"/>
      </q-expansion-item>
    </template>
  </div>
</template>

<script setup>
import path from "path-browserify";
import {ref} from "vue";

const props = defineProps(["item", "basePath", "level"])

const onlyOneChild = ref(null)
const hasOneShowingChild = (children = [], parent) => {
  const showingChildren = children.filter((child) => {
    if (child.hidden) {
      return false
    } else {
      onlyOneChild.value = child
      return true
    }
  })

  if (showingChildren.length === 1) {
    return true
  }

  if (showingChildren.length === 0) {
    onlyOneChild.value = {...parent, path: '', noShowingChildren: true}
    return true
  }

  return false
}

const resolvePath = (routePath) => {
  if (isExternal(routePath)) {
    return routePath
  }
  if (isExternal(props.basePath)) {
    return props.basePath
  }
  return path.resolve(props.basePath, routePath)
}

const isExternal = (path) => {
  return /^(https?:|mailto:|tel:)/.test(path)
}
</script>

<style lang="scss" scoped>
:deep(.q-item__section--avatar){
  min-width: 40px;
}
</style>
