window.onload = function() {
  let myliffId = "1655870418-BMAVnZrz";
  liff
    .init({
      liffId: myliffId
    })
    .then(() => {
      initializeLiff();
    })
};

function initializeLiff() {
  liff.getProfile().then((profile) => {
    let displayName = profile.displayName;
    document.getElementById("sender").value = displayName.trim();
    let accessToken = liff.getAccessToken();
    console.log(accessToken);
    document.getElementById("accessToken").value = accessToken;
  }).catch(function(error) {
    window.alert(error);
  })
}
