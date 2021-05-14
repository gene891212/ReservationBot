window.onload = function () {
  let myliffId = '1655874416-zG2KbgK3';
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
  processDate();
};

function fillZero(num) {
  return num.toString().padStart(2, '0')
}

function processDate() {
  let date = new Date;
  let dateFormat = `${date.getFullYear()}-${fillZero(date.getUTCMonth())}-${fillZero(date.getDate())}`;
  let timeFormat = `${date.getHours()}:${fillZero(date.getMinutes() + 1)}`
  document.getElementById('date').setAttribute('min', dateFormat);
  document.getElementById('date').setAttribute('value', dateFormat);
  document.getElementById('time').setAttribute('value', timeFormat);
}

function initializeLiff() {
  displayIsInClientInfo();
  registerButtonHandlers();
  if (liff.isLoggedIn()) {
    let accessToken = liff.getAccessToken();
    document.getElementById('accessToken').value = accessToken;
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

  document.getElementById('reciver').addEventListener('change', function () {
    let reciverName = document.getElementById('reciver').value;
    let url = `/api/user/${reciverName}`;
    fetch(url)
      .then((response) => {
        return response.json()
      })
      .then((user) => {
        console.log(user);
        document.getElementById('reciverImg').setAttribute('src', user.pictureUrl);
      })
      .catch((err) => {
        alert(err);
      })
  });
}
