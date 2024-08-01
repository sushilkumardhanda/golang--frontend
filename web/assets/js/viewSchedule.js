function ToggleView(id) {
  document.getElementById(id).classList.toggle("hidden");
}

function getQueryParams() {
  let params = {};
  let queryParams = new URLSearchParams(window.location.search);

  queryParams.forEach((value, key) => {
    params[key] = value;
  });

  return params;
}
function EditSchedule(ElementId, event) {
  event.stopPropagation();
  param = getQueryParams();
  const params = new URLSearchParams({
    itr: param["itr"],
    schema: param["schema"],
    elementID: ElementId,
  });
  const url = `/editSchedule?${params.toString()}`;
  window.location.href = url;
}
