function openTab(event, tabName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName("tabcontent");
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }
  tablinks = document.getElementsByClassName("tablinks");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(
      "border-2 shadow-xl",
      "cursor-pointer"
    );
    var imgElement = tablinks[i].querySelector("img");
    if (imgElement) {
      imgElement.classList.add("hidden");
    }
  }
  document.getElementById(tabName).style.display = "block";
  var tabElement = event.currentTarget;
  tabElement.className = tabElement.className.replace(
    "cursor-pointer",
    "border-2 shadow-xl"
  );
  var imgElement = tabElement.querySelector("img");
  if (imgElement) {
    imgElement.classList.remove("hidden");
  }
}
let selectedSchema = "";
let selectedITR = "";
let selectedService = "";
function schemaButton(event) {
  var schemaElement = event.currentTarget;
  selectedSchema = schemaElement.textContent.trim();
  console.log(selectedSchema);
  let schemaButtons = document.getElementsByClassName("button-schema");
  for (i = 0; i < schemaButtons.length; i++) {
    schemaButtons[i].className = schemaButtons[i].className.replace(
      "text-white bg-[#487DE7]",
      "bg-white"
    );
  }

  schemaElement.className = schemaElement.className.replace(
    "bg-white",
    "text-white bg-[#487DE7]"
  );
}

function itrButton(event) {
  var itrElement = event.currentTarget;
  selectedITR = itrElement.textContent.trim();
  console.log(selectedITR);
  getSchemas(selectedITR);
  let itrButtons = document.getElementsByClassName("button-itr");
  for (i = 0; i < itrButtons.length; i++) {
    itrButtons[i].className = itrButtons[i].className.replace(
      "text-white bg-[#487DE7]",
      "bg-white"
    );
  }

  itrElement.className = itrElement.className.replace(
    "bg-white",
    "text-white bg-[#487DE7]"
  );
}

function serviceButton(event) {
  var serviceElement = event.currentTarget;
  selectedService = serviceElement.textContent.trim();
  console.log(selectedService);
  let serviceButtons = document.getElementsByClassName("button-service");
  for (i = 0; i < serviceButtons.length; i++) {
    serviceButtons[i].className = serviceButtons[i].className.replace(
      "text-white bg-[#487DE7]",
      "bg-white"
    );
  }

  serviceElement.className = serviceElement.className.replace(
    "bg-white",
    "text-white bg-[#487DE7]"
  );
}

async function getSchemas(itr) {
  itr = itr.replace(/\s+/g, "");
  const data = { ITR: itr };
  const response = await fetch("/schemaList", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  const res = await response.json();
  console.log(res);
  displaySchemas(res.schemaList);
}
function displaySchemas(list) {
  const container = document.getElementById("schema-container");
  container.innerHTML = "";
  selectedSchema = "";
  // Iterate over the array and create div elements
  list.forEach((item) => {
    const div = document.createElement("div");
    div.className =
      "button-schema border rounded p-2 bg-white text-center text-[10px] md:text-[16px]";
    div.textContent = item;
    div.onclick = schemaButton; // Assign the click event handler
    container.appendChild(div); // Append the div to the container
  });
}

function nextButtonClicked() {
  selectedITR = selectedITR.replace(/\s+/g, "");
  const params = new URLSearchParams({
    itr: selectedITR,
    schema: selectedSchema,
  });
  const url = `/schedule?${params.toString()}`;
  window.location.href = url;
}
