window.onload = function () {
  let myliffId = "1655874416-zG2KbgK3";
  liff
    .init({
      liffId: myliffId
    })
    .then(() => {
      initializeLiff();
    })
    .catch(() => {
      window.alert(err);
    })
    getAllUsers();
};

function getAllUsers() {
  let userURL = "/api/users";
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
}

function initializeLiff() {
  displayIsInClientInfo();
  registerButtonHandlers();
  if (liff.isLoggedIn()) {
    let accessToken = liff.getAccessToken();
    document.getElementById("accessToken").value = accessToken;
    document.getElementById('liffLoginButton').disabled = true;
  } else {
    document.getElementById('liffLogoutButton').disabled = true;
  }
}

function displayIsInClientInfo() {
  if (liff.isInClient()) {
    document.getElementById('liffLoginButton').setAttribute('hidden', 'hidden');
    document.getElementById('liffLogoutButton').setAttribute('hidden', 'hidden');
  }
}

function registerButtonHandlers() {
  // Login button
  document.getElementById('liffLoginButton').addEventListener('click', function () {
    console.log('click');
    if (!liff.isLoggedIn()) {
      // set `redirectUri` to redirecst the user to a URL other than the front page of your LIFF app.
      liff.login();
    } else {
      document.getElementById('liffLogoutButton').disabled = true;
    }
  });
  // Logout button
  document.getElementById('liffLogoutButton').addEventListener('click', function () {
    if (liff.isLoggedIn()) {
      liff.logout();
      window.location.reload();
    }
  });
}
