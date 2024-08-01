const logout = () => {
  fetch("/logout", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok " + response.statusText);
      }
      return response.json();
    })
    .then((data) => {
      if (data.message == "logged out") {
        console.log(data);
        window.location.replace("http://localhost:8000/login");
      }
    })
    .catch((error) => {
      console.error(
        "There has been a problem with your fetch operation:",
        error
      );
    });
};

let cancelButton = document.getElementById("logout");
cancelButton.addEventListener("click", () => {
  logout();
});
