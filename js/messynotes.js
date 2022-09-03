function togglePrivate() {
  var toggle = document.getElementById("privateSlider");
  var elements = document.getElementsByClassName("private");

  if (toggle.checked) {
    display = "initial";
  } else {
    display = "none";
  }

  for (let i = 0; i < elements.length; i++) {
    elements[i].style.display = display;
  }
}

function attachSlider() {
  var toggle = document.getElementById("privateSlider");
  if (toggle == null) {
    return
  }
  toggle.addEventListener(
    "change",
    togglePrivate
  );
  var body = document.getElementsByTagName("body")[0];
  body.addEventListener(
    "keyup", event => {
      if (event.key == "P") {
        toggle.checked = !toggle.checked;
        togglePrivate();
      }
    }
  );
}

attachSlider()
