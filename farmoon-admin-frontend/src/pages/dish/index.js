import { getAPI } from 'src/api'

export async function getCuisineInfo () {
  const cuisineMap = {}
  const cuisineOptions = []
  try {
    const { data } = await getAPI('private/cuisine/list')
    const cuisines = data.cuisines
    cuisines.forEach(cuisine => {
      cuisineOptions.push({ label: cuisine.name, value: cuisine.id })
      cuisineMap[cuisine.id] = cuisine.name
    })
  } catch (e) {
    console.log(e)
  }
  return { options: cuisineOptions, map: cuisineMap }
}

export async function getIngredientTypeInfo () {
  const ingredientTypeMap = {}
  const ingredientTypeOptions = []
  try {
    const { data } = await getAPI('private/ingredient-type/list')
    const ingredientTypes = data.ingredientTypes
    ingredientTypes.forEach(ingredientType => {
      ingredientTypeOptions.push(
        { label: ingredientType.name, value: ingredientType.id })
      ingredientTypeMap[ingredientType.id] = ingredientType.name
    })
  } catch (e) {
    console.log(e)
  }
  return { options: ingredientTypeOptions, map: ingredientTypeMap }
}
