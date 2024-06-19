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

// Adding jquery code to close the navbar menu on nav-link click
$(document).ready(function () {
  $(".nav-link").on("click", function () {
    $(".navbar-collapse").collapse("hide");
  });
});

// EVENT ID FROM AWS
const event_id = "76125999-17f0-4f22-b47e-ef4ac16fb1cc";
let rsvp = true;

// Event listener for email input field
document.getElementById("email").addEventListener("blur", validateEmail);

function validateEmailFormat(email) {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return re.test(String(email).toLowerCase());
}

async function retrieveGuest(email, event_id) {
  const spinner = document.getElementById("spinner");
  spinner.style.display = "block"; // Show spinner
  const formContainer = document.querySelector(".form-container");
  formContainer.classList.add("loading"); // Add loading class
  try {
    const response = await fetch(
      "https://op9d8hqeqf.execute-api.us-east-2.amazonaws.com/Production",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          event_id: event_id,
        }),
      }
    );

    const responseBody = await response.json();

    if (response.ok && responseBody.statusCode === 200) {
      console.log("Guest retrieved successfully:", responseBody);
      return true;
    } else if (responseBody.statusCode === 404) {
      console.log("Guest not found for the specified event and email.");
    } else {
      const errorText = JSON.stringify(responseBody); // Get error details
      console.log("Failed to retrieve guest:", errorText);
    }
  } catch (error) {
    console.error("Error retrieving guest:", error);
  } finally {
    spinner.style.display = "none"; // Hide spinner
    formContainer.classList.remove("loading"); // Remove loading class
  }
  return false;
}

function enableSubmitButton() {
  const submitBtn = document.getElementById("submitBtn");
  submitBtn.disabled = false;
  submitBtn.classList.remove("disabled");
  submitBtn.removeAttribute("title");
}

function disableSubmitButton() {
  const submitBtn = document.getElementById("submitBtn");
  submitBtn.disabled = true;
  submitBtn.classList.add("disabled");
  submitBtn.setAttribute("title", "Must enter email address to submit RSVP");
}

async function validateEmail() {
  const emailInput = document.getElementById("email").value;
  const messageElem = document.getElementById("message");

  if (!validateEmailFormat(emailInput)) {
    messageElem.textContent = "Invalid email format.";
    messageElem.className = "error";
    disableSubmitButton();
    return;
  } else {
    await retrieveGuest(emailInput, event_id).then((response) => {
      // console.log(response);
      if (response) {
        messageElem.textContent = "Invitation Found! You may now RSVP";
        messageElem.className = "success";
        enableSubmitButton();
      } else {
        messageElem.textContent = "Invitation not found.";
        messageElem.className = "error";
        disableSubmitButton();
      }
    });
  }
}

function selectTab(tab) {
  document.getElementById("attending-tab").classList.remove("active");
  document.getElementById("regrets-tab").classList.remove("active");

  if (tab === "attending") {
    document.getElementById("attending-tab").classList.add("active");
    rsvp = true;
  } else {
    document.getElementById("regrets-tab").classList.add("active");
    rsvp = false;
  }
}

async function handleRSVPUpdate(event) {
  event.preventDefault(); // Prevent the form from submitting the traditional way
  const spinner = document.getElementById("spinner");
  spinner.style.display = "block"; // Show spinner
  const formContainer = document.querySelector(".form-container");
  formContainer.classList.add("loading"); // Add loading class

  // Get the email, attending, and note values from the input fields
  const email = document.getElementById("email").value;
  const note = document.getElementById("note").value;

  // Make sure email and event_id are not empty
  if (!email || !event_id) {
    alert("Please enter an email address and event ID.");
    return;
  }

  try {
    const response = await fetch(
      "https://j71yii2ym0.execute-api.us-east-2.amazonaws.com/Production",
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          event_id: event_id,
          attending: rsvp,
          note: note,
        }),
      }
    );

    if (response.ok) {
      const result = await response.json();
      console.log("RSVP updated successfully:", result);
    } else if (response.status === 404) {
      console.warn("Email or event not found.");
    } else {
      const errorText = await response.text(); // Get error details
      console.error("Failed to update RSVP:", errorText);
    }
  } catch (error) {
    console.error("Error updating RSVP:", error);
  } finally {
    spinner.style.display = "none"; // Hide spinner
    formContainer.classList.remove("loading"); // Remove loading class
    const messageElem = document.getElementById("message");
    messageElem.textContent = "";
    document.getElementById("email").value = "";
    disableSubmitButton();
    alert("RSVP Submitted!");
  }
}
// Add event listener to the form
document
  .getElementById("rsvpForm")
  .addEventListener("submit", handleRSVPUpdate);
