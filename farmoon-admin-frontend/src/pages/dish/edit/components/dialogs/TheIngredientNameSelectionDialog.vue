<template>
  <q-dialog v-model="shown">
    <q-card class="column" style="width: 400px">
      <q-card-section>
        <div class="text-subtitle1">选择食材</div>
      </q-card-section>
      <q-separator/>
      <q-card-section class="row">
        <q-list class="col" separator dense>
          <q-item-section>
            <q-item-label caption>肉类</q-item-label>
          </q-item-section>
          <q-scroll-area style="height: 400px; max-width: 300px;">
            <q-item  v-for="opt in meatOptions" :key="opt.key" clickable v-ripple @click="onSelect($event, opt.name)">
              <q-item-section>{{ opt.name }}</q-item-section>
            </q-item>
          </q-scroll-area>
        </q-list>
        <q-separator vertical/>
        <q-list class="col q-pl-md" separator>
          <q-item-section side top>
            <q-item-label caption>菜类</q-item-label>
          </q-item-section>
          <q-scroll-area style="height: 400px; max-width: 300px;">
            <q-item dense v-for="opt in vegetableOptions" :key="opt.key" clickable v-ripple @click="onSelect($event, opt.name)">
              <q-item-section>{{ opt.name }}</q-item-section>
            </q-item>
          </q-scroll-area>
        </q-list>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref } from "vue";

const emit = defineEmits(["select"]);

const shown = ref(false);
const meatOptions = ref([
  {
    key: "1",
    name: "猪肉"
  },
  {
    key: "2",
    name: "牛肉"
  },
  {
    key: "3",
    name: "羊肉"
  },
  {
    key: "4",
    name: "鸡肉"
  },
  {
    key: "5",
    name: "鸭肉"
  },
  {
    key: "6",
    name: "鱼肉"
  },
  {
    key: "6",
    name: "火腿"
  },
]);
const vegetableOptions = ref([
  {
    key: "1",
    name: "胡萝卜"
  },
  {
    key: "2",
    name: "西兰花"
  },
  {
    key: "3",
    name: "洋葱"
  },
  {
    key: "4",
    name: "土豆"
  },
  {
    key: "5",
    name: "番茄"
  },
  {
    key: "6",
    name: "黄瓜"
  },
  {
    key: "7",
    name: "菠菜"
  },
  {
    key: "8",
    name: "生菜"
  },
  {
    key: "9",
    name: "青豆"
  },
  {
    key: "10",
    name: "青椒"
  },
  {
    key: "11",
    name: "茄子"
  },
  {
    key: "12",
    name: "花菜"
  },
  {
    key: "13",
    name: "豌豆"
  },
  {
    key: "14",
    name: "南瓜"
  },
  {
    key: "15",
    name: "菜花"
  }
])

const show = () => {
  shown.value = true;
  // getIngredients()
  //   .then(res => {
  //     options.value = res.data;
  //   });
};

const onSelect = (e, val) => {
  emit("select", val);
  shown.value = false;
};

defineExpose({
  show
});
</script>

<style lang="scss" scoped>

</style>
