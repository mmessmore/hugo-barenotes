

function togglePrivate() {
  console.log("toggle");
  var slider = document.getElementById("privateSlider");
  var elements = document.getElementsByClassName("private");

  if (elements.length == 0) {
    return;
  }

  // assume everything is the same
  var display = "none";
  var checked = false;
  console.log(`checked: ${slider.checked}`);
  if (slider.checked == false) {
    console.log("hidden");
    display = "block";
    checked = true;
  }

  for (let i = 0; i < elements.length; i++) {
    elements[i].style.display = display;
  }
  slider.checked = checked;
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
        togglePrivate();
      }
    });
}

attachSlider()
