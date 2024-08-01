function ToggleSearchFilter() {
  document.getElementById("search-filter").classList.toggle("hidden");
}
let selectedfilter = "";
function filterButton(event) {
  var filterElement = event.currentTarget;
  selectedfilter = filterElement.textContent.trim();
  let filterButtons = document.getElementsByClassName("filter-buttons");
  for (i = 0; i < filterButtons.length; i++) {
    filterButtons[i].className = filterButtons[i].className.replace(
      "text-white bg-[#4A7EE7]",
      "bg-white"
    );
  }

  filterElement.className = filterElement.className.replace(
    "bg-white",
    "text-white bg-[#4A7EE7]"
  );
}
function ApplyFilter() {
  document.getElementById("filter-div").innerHTML = selectedfilter;
  document.getElementById("filter").classList.toggle("hidden");
  document.getElementById("search-filter").classList.toggle("hidden");
}
function ClearFilter() {
  selectedfilter = "";
  let filterButtons = document.getElementsByClassName("filter-buttons");
  for (i = 0; i < filterButtons.length; i++) {
    filterButtons[i].className = filterButtons[i].className.replace(
      "text-white bg-[#4A7EE7]",
      "bg-white"
    );
  }
  document.getElementById("filter").classList.toggle("hidden");
}
