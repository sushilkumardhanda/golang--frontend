let initialEnumCondition =
  document.getElementById("enum-condition").textContent;
let enumCondition = CodeMirror.fromTextArea(
  document.getElementById("enum-condition"),
  {
    lineNumbers: true,
    mode: "text/x-perl",
    theme: "3024-day",
    readOnly: false,
  }
);

enumCondition.setValue(initialEnumCondition);
enumCondition.setSize(null, 95);

function getQueryParams() {
  let params = {};
  let queryParams = new URLSearchParams(window.location.search);

  queryParams.forEach((value, key) => {
    params[key] = value;
  });

  return params;
}

function EditSchedule() {
  param = getQueryParams();
  const params = new URLSearchParams({
    itr: param["itr"],
    schema: param["schema"],
    elementID: param["elementID"],
  });
  const url = `/editSchedule?${params.toString()}`;
  window.location.href = url;
}
