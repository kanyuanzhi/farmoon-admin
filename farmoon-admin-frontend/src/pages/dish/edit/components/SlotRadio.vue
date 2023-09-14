<template>
  <q-item dense>
    <q-item-section avatar>{{ label }}</q-item-section>
    <q-item-section>
      <div class="q-gutter-lg">
        <q-radio
          v-for="s in slotOptions"
          :key="s"
          v-model="selectedSlot"
          dense
          checked-icon="task_alt"
          unchecked-icon="panorama_fish_eye"
          :val="s"
          :label="s"
        />
      </div>
    </q-item-section>
  </q-item>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";

const props = defineProps(["label", "slotCount", "slotNumber"]);
const emits = defineEmits(["update"]);

const selectedSlot = ref(props.slotNumber);

onMounted(() => {
  selectedSlot.value = props.slotNumber;
});

const slotOptions = [];
for (let i = 1; i < props.slotCount + 1; i++) {
  slotOptions.push(String(i));
}

watch(selectedSlot, (value) => {
  emits("update", value);
});

</script>

<style lang="scss" scoped>

</style>
