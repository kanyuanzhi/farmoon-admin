import {getAPI} from "src/api";

export async function getCuisineInfo() {
  const cuisineMap = {}
  const cuisineOptions = []
  try {
    const {data} = await getAPI("private/cuisine/list")
    const cuisines = data.cuisines
    cuisines.forEach(cuisine => {
      cuisineOptions.push({label: cuisine.name, value: cuisine.id})
      cuisineMap[cuisine.id] = cuisine.name
    })
  } catch (e) {
    console.log(e)
  }
  return {options: cuisineOptions, map: cuisineMap}
}
