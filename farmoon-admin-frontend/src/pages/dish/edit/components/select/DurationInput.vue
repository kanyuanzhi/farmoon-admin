<template>
  <q-item>
    <q-item-section avatar>时间控制</q-item-section>
    <q-item-section>
      <div class="row q-col-gutter-md">
        <div class="col">
          <q-input
              dense
              filled
              v-model.number="min"
              stack-label
              :disable="disable"
              mask="###"
          >
            <template v-slot:append>
              <span class="text-body2">分</span>
            </template>
          </q-input>
        </div>
        <div class="col">
          <q-input
              dense
              filled
              v-model.number="sec"
              stack-label
              :disable="disable"
              mask="##"
          >
            <template v-slot:append>
              <span class="text-body2">秒</span>
            </template>
          </q-input>
        </div>
      </div>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { ref, watch } from "vue";
import { floor } from "lodash";

const props = defineProps(["duration", "disable"]);
const emits = defineEmits(["update"]);

const min = ref(floor(props.duration / 60));
const sec = ref(props.duration % 60);

watch(min, (v) => {
  emits("update", parseInt(v * 60 + sec.value));
});

watch(sec, (newValue, oldValue) => {
  if (newValue > 60) {
    sec.value = 60;
  }
  emits("update",  parseInt(min.value * 60 + sec.value));
});
</script>

<style lang="scss" scoped>

</style>
