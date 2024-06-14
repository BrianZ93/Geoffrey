// public/messageListener.js
window.receivedData = {}; // Initialize a global variable to store the received data

window.addEventListener("message", (event) => {
  if (event.origin !== window.location.origin) {
    return; // Ignore messages from different origins for security reasons
  }
  window.receivedData = event.data; // Store the received data in the global variable
  console.log("Data received from parent:", window.receivedData);
  // You can also trigger other actions here, like updating the DOM
});
