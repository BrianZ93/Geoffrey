    // Listener for react wrapper to send guests
    window.receivedData = {}; // Initialize a global variable to store the received data
    let rsvp = false

    window.addEventListener("message", (event) => {
      if (event.origin !== window.location.origin) {
        return; // Ignore messages from different origins for security reasons
      }
      window.receivedData = event.data; // Store the received data in the global variable
      console.log("Data received from parent:", window.receivedData);

      // Log all emails for debugging
      const guests = window.receivedData.items || [];
      const emails = guests.map(guest => guest.email);
      console.debug("Emails in guest list:", emails);

      // Update email validation with the received guests data
      document.getElementById('email').addEventListener('blur', validateEmail);
    });

    function enableSubmitButton() {
      const submitBtn = document.getElementById('submitBtn');
      submitBtn.disabled = false;
      submitBtn.classList.remove('disabled');
      submitBtn.removeAttribute('title');
    }

    function disableSubmitButton() {
      const submitBtn = document.getElementById('submitBtn');
      submitBtn.disabled = true;
      submitBtn.classList.add('disabled');
      submitBtn.setAttribute('title', 'Must enter email address to submit RSVP');
    }

    function validateEmail() {
      const emailInput = document.getElementById('email').value;
      const messageElem = document.getElementById('message');
      const guests = window.receivedData.items || [];
      const emails = guests.map(guest => getSelection.email)

      // console.log("Checking email:", emailInput);
      // console.log("Guests data:", guests);
      guests.forEach(guest => console.debug("Comparing with:", guest.email));

      if (!validateEmailFormat(emailInput)) {
        messageElem.textContent = 'Invalid email format.';
        messageElem.className = 'error';
        return;
      }

      const guestExists = guests.some(guest => guest.email.toLowerCase() === emailInput.toLowerCase());
      if (guestExists) {
        messageElem.textContent = 'Email is valid and found in the guest list.';
        messageElem.className = 'success';
        enableSubmitButton();
      } else {
        messageElem.textContent = 'Email not found in the guest list.';
        messageElem.className = 'error';
        disableSubmitButton();
      }
    }

    function validateEmailFormat(email) {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return re.test(String(email).toLowerCase());
    }

    document.getElementById('rsvpForm').addEventListener('submit', function(event) {
      const messageElem = document.getElementById('message');
      if (messageElem.className !== 'success') {
        event.preventDefault();
        alert('Please provide a valid email address found in the guest list.');
      }
    });

    // eslint-disable-next-line no-unused-vars
    function selectTab(tab) {
      document.getElementById("attending-tab").classList.remove("active");
      document.getElementById("regrets-tab").classList.remove("active");
    
      if (tab === "attending") {
        document.getElementById("attending-tab").classList.add("active");
        rsvp = true
      } else {
        document.getElementById("regrets-tab").classList.add("active");
        rsvp = false
      }
    }

    // Submit RSVP
    document.getElementById('rsvpForm').addEventListener('submit', function(event) {
      event.preventDefault(); // Prevent the form from submitting the traditional way

      const email = document.getElementById('email').value;
      const note = document.getElementById('note').value;

      const eventDetail = {
        email: email,
        note: note,
        rsvp: rsvp
      };

      // Dispatch custom event
      const customEvent = new CustomEvent('rsvpSubmitted', { detail: eventDetail });
      window.parent.dispatchEvent(customEvent);
    });