function ToggleLogout() {
  document.getElementById("logout-dropdown").classList.toggle("hidden");
}
async function logout() {
  const response = await fetch("/logout", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const myJson = await response.json();
  console.log(myJson);
  if (myJson.message == "logged out") {
    window.location.replace("/login");
  }
}

function ToggleChangeService() {
  document.getElementById("modal-box").classList.toggle("hidden");
}
