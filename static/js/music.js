function insertAfter(referenceNode, newNode) {
  referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
}

window.onload = (event) => {
  musicElements = document.querySelectorAll("div.abc-music");
  musicElements.forEach(el => {
    code = el.textContent;
    el.style.display = "none";
    uuid = crypto.randomUUID();
    page = document.createElement("div");
    page.id = uuid;
    insertAfter(el, page);
    ABCJS.renderAbc(
      uuid,
      code,
      { add_classes: true }
    );
  });
};
