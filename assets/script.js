(function () {
  document.body.addEventListener('htmx:afterSwap', function (event) {
    var newTitle = event.detail.elt.getAttribute('data-new-title');
    if (newTitle) {
      document.title = newTitle;
    }
  });
})()

