getAllUsers();

function getAllUsers() {
  let userURL = "/api/users";
  fetch(userURL)
    .then((response) => {
      return response.json()
    })
    .then((users) => {
      users.forEach(element => {
        let userSelect = document.getElementById("reciver");
        let userOption = document.createElement("option");
        userOption.setAttribute("value", element.displayName);
        userOption.innerText = element.displayName;
        userSelect.appendChild(userOption);
      });
    })
    .catch((err) => {
      alert(err);
    })
}
