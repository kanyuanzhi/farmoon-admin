import { api } from "src/boot/axios";
import { Notify } from "quasar";

export function getAPI(url, params) {
  return api({
    url: url,
    method: "get",
    params: params
  });
}

export function postAPI(url, params) {
  return api({
    url: url,
    method: "post",
    data: params,
  });
}

export function postBlobAPI(url, params) {
  return api({
    url: url,
    method: "post",
    data: params,
    responseType: "blob"
  });
}

export function putAPI(url, params) {
  return api({
    url: url,
    method: "put",
    data: params
  });
}

export function deleteAPI(url, params) {
  return api({
    url: url,
    method: "delete",
    data: params
  });
}

export async function downloadAPI(url, params, fileName) {
  const data = await postBlobAPI(url, params);
  if (!data || data.size === 0) {
    Notify.create({
      type: "negative",
      message: "download error!",
    });
    return;
  }
  if (typeof window.navigator.msSaveBlob !== "undefined") {
    window.navigator.msSaveBlob(new Blob([data]), fileName);
  } else {
    let urlHref = window.URL.createObjectURL(new Blob([data]));
    let link = document.createElement("a");
    link.style.display = "none";
    link.href = urlHref;
    link.setAttribute("download", fileName);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(urlHref);
  }
}
