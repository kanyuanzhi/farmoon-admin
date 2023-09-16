import { defineStore } from 'pinia'
import { constantRoutes, asyncRoutes } from 'src/router/routes'
import { cloneDeep } from 'lodash/lang'

export const useDishStore = defineStore('dish', {
  state: () => ({
    editingDish: {
      id: 0,
      steps: [],
      name: '',
      cuisine: 1,
      uuid: '',
    },
    editingDishChangedFlag: true,
    originEditingDish: {
      id: 0,
      steps: [],
      name: '',
      cuisine: 1,
      uuid: '',
    },
    lastStirFryGear: 0,

  }),
  getters: {},
  actions: {
    setEditingDish (dish) {
      this.editingDish = cloneDeep(dish)
      this.originEditingDish = cloneDeep(dish)
      this.shiftEditingDishChangedFlag()
    },
    resetEditingDish () {
      this.editingDish = cloneDeep(this.originEditingDish)
      this.shiftEditingDishChangedFlag()
    },
    clearEditingDish () {
      this.editingDish = {
        id: 0,
        steps: [],
        name: '',
        cuisine: 1,
        uuid: '',
      }
      this.originEditingDish = {
        id: 0,
        steps: [],
        name: '',
        cuisine: 1,
        uuid: '',
      }
      this.shiftEditingDishChangedFlag()
    },
    setLastStirFryGear (gear) {
      this.lastStirFryGear = gear
    },
    shiftEditingDishChangedFlag () {
      this.editingDishChangedFlag = !this.editingDishChangedFlag
    },
  },
})
