/*!
=========================================================
* Creative Studio Landing page
=========================================================

* Copyright: 2019 DevCRUD (https://devcrud.com)
* Licensed: (https://devcrud.com/licenses)
* Coded by www.devcrud.com

=========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
*/

// smooth scroll
$(document).ready(function () {
  $(".navbar .nav-link").on("click", function (event) {
    if (this.hash !== "") {
      event.preventDefault();

      var hash = this.hash;

      $("html, body").animate(
        {
          scrollTop: $(hash).offset().top,
        },
        700,
        function () {
          window.location.hash = hash;
        }
      );
    }
  });
});

function selectTab(tab) {
  document.getElementById("attending-tab").classList.remove("active");
  document.getElementById("regrets-tab").classList.remove("active");

  if (tab === "attending") {
    document.getElementById("attending-tab").classList.add("active");
  } else {
    document.getElementById("regrets-tab").classList.add("active");
  }
}
