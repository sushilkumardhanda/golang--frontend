const userAction = async () => {
  userName = document.getElementById("email").value;
  passWord = document.getElementById("password").value;
  myBody = { username: userName, password: passWord };

  const response = await fetch("/login", {
    method: "POST",
    body: JSON.stringify(myBody),
    headers: {
      "Content-Type": "application/json",
    },
  });
  const myJson = await response.json();
  if (myJson.message == "already logged in") {
    console.log(myJson);
    let modalBox = document.getElementById("modal-box");
    modalBox.classList.remove("hidden");
    let okayButton = document.getElementById("okay");
    okayButton.addEventListener("click", () => {
      modalBox.classList.add("hidden");
      loginConfirm();
    });

    let cancelButton = document.getElementById("cancel");
    cancelButton.addEventListener("click", () => {
      modalBox.classList.add("hidden");
    });
  } else if (myJson.message) {
    window.location.replace("/welcome");
  } else {
    // let errorMsg = document.getElementById("error-msg");
    // errorMsg.innerText = myJson.error;
    // console.log(myJson);
  }
};
const loginConfirm = () => {
  fetch("/loginConfirm", {
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
      if (data.message == "logged in") {
        console.log(data);
        window.location.replace("/welcome");
      }
    })
    .catch((error) => {
      console.error(
        "There has been a problem with your fetch operation:",
        error
      );
    });
};

let loginButton = document.getElementById("login");
loginButton.addEventListener("click", () => {
  userAction();
});

function TogglePassword() {
  var x = document.getElementById("password");
  document.getElementById("hide").classList.toggle("hidden");
  document.getElementById("unhide").classList.toggle("hidden");
  if (x.type === "password") {
    x.type = "text";
  } else {
    x.type = "password";
  }
}
