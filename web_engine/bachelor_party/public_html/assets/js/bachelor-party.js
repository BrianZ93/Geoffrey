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

// AWS
// Listener for react wrapper to send guests
window.receivedData = {}; // Initialize a global variable to store the received data

window.addEventListener("message", (event) => {
  if (event.origin !== window.location.origin) {
    return; // Ignore messages from different origins for security reasons
  }
  window.receivedData = event.data; // Store the received data in the global variable
  // console.log("Data received from parent:", window.receivedData);

  // Log all emails for debugging
  const guests = window.receivedData.items || [];
  const emails = guests.map((guest) => guest.email);
  // console.info("Emails in guest list:", emails);

  // Update email validation with the received guests data
  document.getElementById("email").addEventListener("blur", validateEmail);
});

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

function validateEmail() {
  const emailInput = document.getElementById("email").value;
  const messageElem = document.getElementById("message");
  const guests = window.receivedData.items || [];
  const emails = guests.map((guest) => getSelection.email);

  console.log("Checking email:", emailInput);
  console.log("Guests data:", guests);
  guests.forEach((guest) => console.info("Comparing with:", guest.email));

  if (!validateEmailFormat(emailInput)) {
    messageElem.textContent = "Invalid email format.";
    messageElem.className = "error";
    disableSubmitButton();
    return;
  } else {
    messageElem.textContent = "Email is valid.";
    messageElem.className = "success";
    enableSubmitButton();
  }

  // const guestExists = guests.some(
  //   (guest) => guest.email.toLowerCase() === emailInput.toLowerCase()
  // );
  // if (guestExists) {
  //   messageElem.textContent = "Email is valid and found in the guest list.";
  //   messageElem.className = "success";
  //   enableSubmitButton();
  // } else {
  //   messageElem.textContent = "Email not found in the guest list.";
  //   messageElem.className = "error";
  //   disableSubmitButton();
  // }
}

function validateEmailFormat(email) {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return re.test(String(email).toLowerCase());
}

document
  .getElementById("rsvpForm")
  .addEventListener("submit", function (event) {
    const messageElem = document.getElementById("message");
    if (messageElem.className !== "success") {
      event.preventDefault();
      alert("Please provide a valid email address found in the guest list.");
    }
  });

// eslint-disable-next-line no-unused-vars
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

// Lambda ARN
// arn:aws:lambda:us-east-2:061841606853:function:UpdateRSVP

// Submit RSVP
document
  .getElementById("rsvpForm")
  .addEventListener("submit", async function (event) {
    event.preventDefault(); // Prevent the form from submitting the traditional way

    const email = document.getElementById("email").value;
    const note = document.getElementById("note").value;

    const eventDetail = {
      email: email,
      note: note,
      rsvp: rsvp,
    };

    try {
      const response = await fetch(
        `https://op9d8hqeqf.execute-api.us-east-2.amazonaws.com/dev?email=${encodeURIComponent(
          email
        )}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (response.ok) {
        const guest = await response.json();
        console.log("Guest retrieved successfully:", guest);
        alert("Guest retrieved successfully!");
      } else if (response.status === 404) {
        alert("Email not found.");
      } else {
        const errorText = await response.text(); // Get error details
        console.error("Failed to retrieve guest:", errorText);
        alert("Failed to retrieve guest.");
      }
    } catch (error) {
      console.error("Error retrieving guest:", error);
      alert("An error occurred while retrieving guest.");
    }
  });

// // Function to handle the form submission
// async function handleRSVPUpdate(event) {
//   event.preventDefault(); // Prevent the form from submitting the traditional way

//   // Get the email, event_id, and attending values from the input fields
//   const email = document.getElementById("email").value;
//   const event_id = document.getElementById("event_id").value;
//   const attending = document.getElementById("attending").checked; // Assuming it's a checkbox

//   // Make sure email and event_id are not empty
//   if (!email || !event_id) {
//     alert("Please enter an email address and event ID.");
//     return;
//   }

//   try {
//     const response = await fetch(
//       "https://j71yii2ym0.execute-api.us-east-2.amazonaws.com/dev/updateRSVP",
//       {
//         method: "PUT",
//         headers: {
//           "Content-Type": "application/json",
//         },
//         body: JSON.stringify({
//           email: email,
//           event_id: event_id,
//           attending: attending,
//         }),
//       }
//     );

//     if (response.ok) {
//       const result = await response.json();
//       console.log("RSVP updated successfully:", result);
//       alert("RSVP updated successfully!");
//     } else if (response.status === 404) {
//       alert("Email or event not found.");
//     } else {
//       const errorText = await response.text(); // Get error details
//       console.error("Failed to update RSVP:", errorText);
//       alert("Failed to update RSVP.");
//     }
//   } catch (error) {
//     console.error("Error updating RSVP:", error);
//     alert("An error occurred while updating RSVP.");
//   }
// }

// // Add event listener to the form
// document
//   .getElementById("rsvpUpdateForm")
//   .addEventListener("submit", handleRSVPUpdate);
