<template>
  <q-page class="flex flex-center">
    <div id="map"/>
  </q-page>
</template>

<script setup>

import { Mapbox } from "@antv/l7-maps";
import {
  RasterLayer, Scene, PointLayer, PolygonLayer, LineLayer,
} from "@antv/l7";
import { Choropleth } from "@antv/l7plot";

import { onMounted } from "vue";
// 同样你也可以初始化一个 Mapbox 地图

onMounted(() => {
  const scene = new Scene({
    id: "map",
    map: new Mapbox({
      // style: {
      //   "version": 8,
      //   "name": "blank",
      //   "sprite": "http://localhost:1234/api/sprites/streets/sprite",
      //   // "glyphs": "https://gw.alipayobjects.com/os/antvdemo/assets/mapbox/glyphs/{fontstack}/{range}.pbf",
      //   "sources": {
      //     openmaptiles: {
      //       "type": "vector",
      //       "url": "http://localhost:1234/api/tilesets/osm-2020-02-10-v3.11_asia_china/tilejson"
      //     }
      //   },
      //   "layers": [
      //     {
      //       id: "openmaptiles",
      //       type: "background",
      //       paint: {
      //         "background-color": "white",
      //       },
      //     },
      //   ],
      // },
      // [fill, line, symbol, circle, heatmap, fill-extrusion, raster, hillshade, background]
      style: "http://localhost:1234/api/styles/osm",
      // center: [103.83735604457024, 1.360253881403068],
      center: [104.13689, 33.96498, 14],
      // center: [108.83735604457024, 30],
      pitch: 0,
      zoom: 3.210275860702593,
      // rotation: 19.313180925794313,
    }),
    // map: new Mapbox({
    //   pitch: 0,
    //   style: "blank",
    //   center: [3.438, 40.16797],
    //   zoom: 0.51329,
    // }),
  });

  // scene.on("loaded", () => {
    // const baseLayer = new RasterLayer({ zIndex: 1 }).source(
    //     "http://localhost:1234/api/tilesets/osm-2020-02-10-v3.11_asia_china/{z}/{x}/{y}.pbf",
    //     // "https://map.geoq.cn/arcgis/rest/services/ChinaOnlineStreetPurplishBlue/MapServer/tile/{z}/{y}/{x}",
    //     {
    //       parser: {
    //         type: "rasterTile",
    //         tileSize: 256,
    //       },
    //     },
    // );
    // scene.addLayer(mask);
    // scene.addLayer(baseLayer);
  // });

  scene.on("loaded", () => {
    fetch("http://localhost:1234/api/assets/province-sm.geojson").then(res => {
      return res.json();
    }).then(list => {
      console.log(list);
      const polygonLayer = new PolygonLayer(
          { name: "fill" },
      ).source(list).
          shape("fill").
          scale("density", {
            type: "quantize",
          }).
          style({
            opacity: 0.5,
          }).active(true);
      const lineLayer = new LineLayer().source(list).shape("line").color("rgb(246,171,5)").size(0.4).style({
        opacity: 1.0,
      });
      const pointLayer = new PointLayer({}).source(list).shape("name", "text").size(12).color("#084081").style({
        textAnchor: "center", // 文本相对锚点的位置 center|left|right|top|bottom|top-left
        textOffset: [0, 0], // 文本相对锚点的偏移量 [水平, 垂直]
        spacing: 1, // 字符间距
        padding: [1, 1], // 文本包围盒 padding [水平，垂直]，影响碰撞检测结果，避免相邻文本靠的太近
        stroke: "#ffffff", // 描边颜色
        strokeWidth: 1, // 描边宽度
        strokeOpacity: 1.0,
        // rotation: 60, // 常量旋转
        // rotation: { // 字段映射旋转
        //   field: "t",
        //   value: [30, 270],
        // },

      });
      scene.addLayer(polygonLayer);
      scene.addLayer(lineLayer);
      // scene.addLayer(pointLayer);
    });
  });
  // fetch(
  //     "https://gw.alipayobjects.com/os/alisis/geo-data-v0.1.1/administrative-data/area-list.json",
  // ).then((response) => response.json()).then((list) => {
  //   const data = list.filter(({ level, parent }) => level === "city" && parent === 330000).
  //       map((item) => ({ ...item, value: Math.random() * 5000 }));
  //   console.log(data)
  //
  //   const choropleth = new Choropleth({
  //     source: {
  //       data,
  //       joinBy: {
  //         sourceField: "adcode",
  //         geoField: "adcode",
  //       },
  //     },
  //     viewLevel: {
  //       level: "province",
  //       adcode: 330000,
  //     },
  //     autoFit: true,
  //     color: {
  //       field: "value",
  //       value: ["#B8E1FF", "#7DAAFF", "#3D76DD", "#0047A5", "#001D70"],
  //       scale: { type: "quantize" },
  //     },
  //     style: {
  //       opacity: 1,
  //       stroke: "#ccc",
  //       lineWidth: 0.6,
  //       lineOpacity: 1,
  //     },
  //     label: {
  //       visible: true,
  //       field: "name",
  //       style: {
  //         fill: "#000",
  //         opacity: 0.8,
  //         fontSize: 10,
  //         stroke: "#fff",
  //         strokeWidth: 1.5,
  //         textAllowOverlap: false,
  //         padding: [5, 5],
  //       },
  //     },
  //     state: {
  //       active: { stroke: "black", lineWidth: 1 },
  //     },
  //     tooltip: {
  //       items: ["name", "adcode", "value"],
  //     },
  //     zoom: {
  //       position: "bottomright",
  //     },
  //     legend: {
  //       position: "bottomleft",
  //     },
  //   });
  //
  //   choropleth.addToScene(scene);
  // });
});


</script>

<style scoped lang="scss">
#map {
  width: 100%;
  height: 100%;
  position: absolute;
}
</style>
