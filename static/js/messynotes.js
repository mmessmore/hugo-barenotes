function togglePrivate(invert=true) {
  console.log("toggle");
  var toggle = document.getElementById("privateSlider");
  var elements = document.getElementsByClassName("private");

  if (elements.length == 0) {
    return;
  }

  // assume everything is the same
  var display = "none";
  var checked = toggle.checked;
  if (invert) {
    checked = ! checked;
  }
  console.log(toggle);
  console.log(`toggle.checked: ${toggle.checked}`);
  console.log(`checked: ${toggle.checked}`);
  if (toggle.checked == true) {
    console.log("hidden");
    display = "block";
    checked = false;
  }

  for (let i = 0; i < elements.length; i++) {
    elements[i].style.display = display;
  }


  toggle.checked = checked;
}

function attachSlider() {
  var toggle = document.getElementById("privateSlider");
  if (toggle == null) {
    return
  }
  toggle.checked = false;
  toggle.addEventListener(
    "change",
    togglePrivate
  );
  var body = document.getElementsByTagName("body")[0];
  body.addEventListener(
    "keyup", event => {
      if (event.key == "P") {
        togglePrivate(false);
      }
    });
}

attachSlider()
