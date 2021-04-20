window.onload = function () {
  let myliffId = "1655873565-oW53qqal";
  liff
    .init({
      liffId: myliffId
    })
    .then(() => {
      initializeLiff();
    })
};

function initializeLiff() {
  registerButtonHandlers();
  if (liff.isLoggedIn()) {
    document.getElementById('liffLoginButton').disabled = true;
  } else {
    document.getElementById('liffLogoutButton').disabled = true;
  }
  getProfile();
}

function getProfile() {
  liff.getProfile()
    .then((profile) => {
      let displayName = profile.displayName;
      let accessToken = liff.getAccessToken();
      document.getElementById("sender").value = displayName.trim();
      document.getElementById("accessToken").value = accessToken;
    })
    .catch(function (error) {
      window.alert(error);
    })
}

function registerButtonHandlers() {
  document.getElementById('liffLoginButton').addEventListener('click', function () {
    console.log('click');
    if (!liff.isLoggedIn()) {
      // set `redirectUri` to redirecst the user to a URL other than the front page of your LIFF app.
      liff.login();
    }
  });
}
