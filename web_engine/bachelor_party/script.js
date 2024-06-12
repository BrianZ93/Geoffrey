window.addEventListener("scroll", () => {
  const sections = document.querySelectorAll(".section");
  const navItems = document.querySelectorAll(".nav-item");
  let currentSection = "";

  sections.forEach((section) => {
    const sectionTop = section.offsetTop;
    const sectionHeight = section.clientHeight;
    if (pageYOffset >= sectionTop - sectionHeight / 3) {
      currentSection = section.getAttribute("id");
    }
  });

  navItems.forEach((item) => {
    item.classList.remove("active");
    if (item.getAttribute("href").substring(1) === currentSection) {
      item.classList.add("active");
    }
  });
});

function openRSVP() {
  document.getElementById("rsvp-modal").style.display = "block";
}

function closeRSVP() {
  document.getElementById("rsvp-modal").style.display = "none";
}
