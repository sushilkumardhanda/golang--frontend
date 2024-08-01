var initialContentFormula = document.getElementById(
  "code-editor-formula"
).textContent;
var editor = CodeMirror.fromTextArea(
  document.getElementById("code-editor-formula"),
  {
    lineNumbers: true,
    mode: "text/x-perl",
    theme: "3024-day",
    readOnly: true,
  }
);
editor.setSize(null, 95);
editor.setValue(initialContentFormula);

function editButtonClickFormula() {
  document.getElementById("edit-formula").classList.toggle("hidden");
  document.getElementById("saveCancel-formula").classList.toggle("hidden");
  editor.setOption("readOnly", false);
}

function saveButtonClickFormula() {
  let content = editor.getValue();
  // Save the content (e.g., send it to a server or save to local storage)
  console.log("Saved content:", content);
  initialContentFormula = content; // Update initial content
  editor.setOption("readOnly", true); // make CodeMirror editor read-only
  document.getElementById("saveCancel-formula").classList.toggle("hidden");
  document.getElementById("edit-formula").classList.toggle("hidden");
}

function cancelButtonClickFormula() {
  document.getElementById("saveCancel-formula").classList.toggle("hidden");
  document.getElementById("edit-formula").classList.toggle("hidden");
  editor.setValue(initialContentFormula);
  editor.setOption("readOnly", true);
}

var initialContentChange =
  document.getElementById("code-editor-change").textContent;
var editorChange = CodeMirror.fromTextArea(
  document.getElementById("code-editor-change"),
  {
    lineNumbers: true,
    mode: "text/x-perl",
    theme: "3024-day",
    readOnly: true,
  }
);
editorChange.setSize(null, 95);
editorChange.setValue(initialContentChange);

function editButtonClickChange() {
  document.getElementById("edit-change").classList.toggle("hidden");
  document.getElementById("saveCancel-change").classList.toggle("hidden");
  editorChange.setOption("readOnly", false);
}

function saveButtonClickChange() {
  let content = editorChange.getValue();
  // Save the content (e.g., send it to a server or save to local storage)
  console.log("Saved content:", content);
  initialContentChange = content; // Update initial content
  editorChange.setOption("readOnly", true); // make CodeMirror editor read-only
  document.getElementById("saveCancel-change").classList.toggle("hidden");
  document.getElementById("edit-change").classList.toggle("hidden");
}

function cancelButtonClickChange() {
  document.getElementById("saveCancel-change").classList.toggle("hidden");
  document.getElementById("edit-change").classList.toggle("hidden");
  editorChange.setValue(initialContentChange);
  editorChange.setOption("readOnly", true);
}

var initialContentSave =
  document.getElementById("code-editor-save").textContent;
var editorSave = CodeMirror.fromTextArea(
  document.getElementById("code-editor-save"),
  {
    lineNumbers: true,
    mode: "text/x-perl",
    theme: "3024-day",
    readOnly: true,
  }
);

editorSave.setSize(null, 95);

editorSave.setValue(initialContentSave);

function editButtonClickSave() {
  document.getElementById("edit-save").classList.toggle("hidden");
  document.getElementById("saveCancel-save").classList.toggle("hidden");
  editorSave.setOption("readOnly", false);
}

function saveButtonClickSave() {
  let content = editorSave.getValue();
  // Save the content (e.g., send it to a server or save to local storage)
  console.log("Saved content:", content);
  initialContentSave = content; // Update initial content
  editorSave.setOption("readOnly", true); // make CodeMirror editor read-only
  document.getElementById("saveCancel-save").classList.toggle("hidden");
  document.getElementById("edit-save").classList.toggle("hidden");
}

function cancelButtonClickSave() {
  document.getElementById("saveCancel-save").classList.toggle("hidden");
  document.getElementById("edit-save").classList.toggle("hidden");
  editorSave.setValue(initialContentSave);
  editorSave.setOption("readOnly", true);
}

function getQueryParams() {
  let params = {};
  let queryParams = new URLSearchParams(window.location.search);

  queryParams.forEach((value, key) => {
    params[key] = value;
  });

  return params;
}
function AdvanceSetting() {
  param = getQueryParams();
  const params = new URLSearchParams({
    itr: param["itr"],
    schema: param["schema"],
    elementID: param["elementID"],
  });
  const url = `/advanceSetting?${params.toString()}`;
  window.location.href = url;
}
