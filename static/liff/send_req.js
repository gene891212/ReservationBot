window.onload = function () {
  let userURL = "http://localhost:7000/api/users";
  fetch(userURL)
    .then((response) => {
      return response.json()
    })
    .then((users) => {
      console.log(users);
      users.forEach(element => {
        let userSelect = document.getElementById("reciver");
        let userOption = document.createElement("option");
        userOption.setAttribute("value", element.displayName);
        userOption.innerText = element.displayName;
        userSelect.appendChild(userOption);
      });
    })
    .catch((err) => {
      console.log(err);
    })
};