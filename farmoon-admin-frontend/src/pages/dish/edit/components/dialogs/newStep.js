import { cloneDeep } from "lodash/lang";
import { Notify } from "quasar";

function newStep(type, name) {
  const step = {
    key: Date.now(),
    instructionType: type,
    instructionName: name,
  };
  return step;
}

export function newIngredientStep(name, shape, weight, slotNumber) {
  const stepName = "添加" + name + (shape === "" ? "" : "（" + shape + "）") +
    weight + "克，使用" + slotNumber + "号菜盒";
  const step =
    newStep("ingredient", stepName);
  step["name"] = name;
  step["shape"] = shape;
  step["weight"] = weight;
  step["slotNumber"] = slotNumber;
  return step;
}

export function newSeasoningStep(seasonings) {
  const newSeasonings = [];
  const stepNames = [];
  for (let i = 0, len = seasonings.length; i < len; i++) {
    if (seasonings[i].label === "") continue;
    newSeasonings.push(cloneDeep(seasonings[i]));
    stepNames.push(seasonings[i].label + seasonings[i].weight + "克");
  }
  if (newSeasonings.length === 0) {
    Notify.create("至少添加1种调料");
    return null;
  }
  const stepName = "添加" + stepNames.join("，");
  const step =
    newStep("seasoning", stepName);
  step["seasonings"] = newSeasonings;
  return step;
}

export function newHeatStep(temperature, judgeType, targetTemperature, duration) {
  let judgeStr = "";
  switch (judgeType) {
    case 1:
      judgeStr = "持续监测锅底温度至" + targetTemperature + "℃";
      break;
    case 2:
      // judgeStr = "持续监测红外温度至" + targetTemperature + "℃";
      judgeStr =
        "持续监测温度至" + targetTemperature + "℃";
      break;
    case 3:
      judgeStr = "持续" + duration + "秒";
      break;
    case 4:
      judgeStr = "无温度监测";
      break;
    default:
      Notify.create("温度控制方式错误");
      return;
  }
  const stepName = "加热" + temperature + "℃，" + judgeStr;
  const step =
    newStep("heat", stepName);
  step["temperature"] = temperature;
  step["judgeType"] = judgeType;
  step["targetTemperature"] = targetTemperature;
  step["duration"] = duration;
  return step;
}

export function newStirFryStep(gear, duration) {
  const stepName = "翻炒" + gear + "档，持续" + duration + "秒";
  const step =
    newStep("stir_fry", stepName);
  step["gear"] = gear;
  step["duration"] = duration;
  return step;
}

export function newWaterStep(weight) {
  const stepName = "添加纯净水" + weight + "克";
  const step =
    newStep("water", stepName);
  step["weight"] = weight;
  step["pumpNumber"] = 6;
  return step;
}

export function newOilStep(weight,) {
  const stepName = "添加食用油" + weight + "克";
  const step =
    newStep("oil", stepName);
  step["weight"] = weight;
  step["pumpNumber"] = 1;
  return step;
}
