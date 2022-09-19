// export const languageNames = new Intl.DisplayNames(["hu"], {
//   type: "language",
// });

export function capitalizeString(s) {
  if (typeof s !== "string") return "";
  return s.charAt(0).toUpperCase() + s.slice(1);
}

export function formatBytesToString(bytes, decimals = 2) {
  if (bytes === 0) return "0 Bytes";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
}

export function includesAny(arr, values) {
  return values.some((v) => arr.includes(v));
}

export function debounce(func, wait = 300, immediate = true) {
  var timeout;

  return function () {
    var context = this,
      args = arguments;

    var callNow = immediate && !timeout;

    clearTimeout(timeout);

    timeout = setTimeout(function () {
      timeout = null;

      if (!immediate) {
        func.apply(context, args);
      }
    }, wait);

    if (callNow) func.apply(context, args);
  };
}

export function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export function clone(obj) {
  return JSON.parse(JSON.stringify(obj));
}
