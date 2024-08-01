function ToggleView(event, id) {
  document.querySelectorAll(".dropdown-edit-view").forEach(function (element) {
    if (element.id !== id) {
      element.classList.add("hidden");
    }
  });
  document.querySelectorAll(".image-edit-view").forEach(function (element) {
    element.classList.remove("bg-[#EEF4FF]");
  });
  document.getElementById(id).classList.toggle("hidden");
  var imgElement = event.currentTarget;
  imgElement.classList.toggle("bg-[#EEF4FF]");
}

function getQueryParams() {
  let params = {};
  let queryParams = new URLSearchParams(window.location.search);

  queryParams.forEach((value, key) => {
    params[key] = value;
  });

  return params;
}
function ViewSchedule(ElementId) {
  param = getQueryParams();
  const params = new URLSearchParams({
    itr: param["itr"],
    schema: param["schema"],
    elementID: ElementId,
  });
  const url = `/viewSchedule?${params.toString()}`;
  window.location.href = url;
}
function EditSchedule(ElementId) {
  param = getQueryParams();
  const params = new URLSearchParams({
    itr: param["itr"],
    schema: param["schema"],
    elementID: ElementId,
  });
  const url = `/editSchedule?${params.toString()}`;
  window.location.href = url;
}

let elementsPerPage = 5;
let currentPage = 1;
const parentDivLarge = document.getElementById("table-content-large");
const parentDivSmall = document.getElementById("table-content-small");
const totalElements = parentDivLarge.children.length;
let totalPages = Math.ceil(totalElements / elementsPerPage);
function showPage() {
  const start = (currentPage - 1) * elementsPerPage;
  const end = start + elementsPerPage;
  document.getElementById("fromToElements").textContent =
    start + 1 + "-" + Math.min(end, totalElements) + " of " + totalElements;
  document.getElementById("pageTotalPage").textContent =
    currentPage + "/" + totalPages;
  const childrenSmall = parentDivSmall.children;
  const children = parentDivLarge.children;
  for (let i = 0; i < children.length; i++) {
    children[i].classList.add("hidden");
    childrenSmall[i].classList.add("hidden");
  }
  for (let i = start; i < end && i < children.length; i++) {
    children[i].classList.remove("hidden");
    childrenSmall[i].classList.remove("hidden");
  }
}
window.onload = function () {
  showPage();
};
function nextPage() {
  if (currentPage < totalPages) {
    currentPage++;
  }
  showPage();
}
function previousPage() {
  if (currentPage > 1) {
    currentPage--;
  }
  showPage();
}
function ElementsPerPageChange(value) {
  elementsPerPage = Number(value);
  totalPages = Math.ceil(totalElements / elementsPerPage);
  showPage();
}
function JumpToPage(value) {
  if (1 <= value && value <= totalPages) {
    currentPage = value;
    showPage();
  }
}
