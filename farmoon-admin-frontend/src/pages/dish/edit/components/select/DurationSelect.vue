<template>
  <q-item>
    <q-item-section avatar>时间控制</q-item-section>
    <q-item-section>
      <div class="row q-col-gutter-md">
        <div class="col">
          <q-select
            dense
            options-dense
            filled
            v-model="min"
            :options="minOptions"
            options-cover
            stack-label
            :disable="disable"
          >
            <template v-slot:append>
              <span class="text-body2">分</span>
            </template>
          </q-select>
        </div>
        <div class="col">
          <q-select
            dense
            options-dense
            filled
            v-model="sec"
            :options="secOptions"
            options-cover
            stack-label
            :disable="disable"
          >
            <template v-slot:append>
              <span class="text-body2">秒</span>
            </template>
          </q-select>
        </div>
      </div>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { ref, watch } from "vue";
import { useQuasar } from "quasar";
import { floor } from "lodash";

const $q = useQuasar();

const props = defineProps(["duration", "disable"]);
const emits = defineEmits(["update"]);

const min = ref({
  label: floor(props.duration / 60) < 10 ? "0" + floor(props.duration / 60) : String(floor(props.duration / 60)),
  value: floor(props.duration / 60)
});
const sec = ref({
  label: props.duration % 60 < 10 ? "0" + props.duration % 60 : String(props.duration % 60),
  value: props.duration % 60
});

const minOptions = [];
const secOptions = [];
for (let i = 0; i < 61; i++) {
  minOptions.push({
    label: i < 10 ? "0" + i : String(i),
    value: i
  });
  if (i !== 60) {
    secOptions.push({
      label: i < 10 ? "0" + i : String(i),
      value: i
    });
  }
}

watch(min, (v) => {
  emits("update", v.value * 60 + sec.value.value);
});

watch(sec, (v) => {
  emits("update", min.value.value * 60 + v.value);
});
</script>

<style lang="scss" scoped>

</style>
