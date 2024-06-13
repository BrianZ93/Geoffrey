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

// Replace these with your actual AWS credentials
const awsAccessKeyId = "YOUR_ACCESS_KEY_ID";
const awsSecretAccessKey = "YOUR_SECRET_ACCESS_KEY";
const awsRegion = "YOUR_AWS_REGION"; // e.g. 'us-west-2'

// Configure AWS SDK with access key and secret key
AWS.config.update({
  region: awsRegion,
  credentials: new AWS.Credentials(awsAccessKeyId, awsSecretAccessKey),
});

// Initialize the DynamoDB Document Client
const docClient = new AWS.DynamoDB.DocumentClient();

// Function to fetch events from DynamoDB
async function fetchEvents() {
  const params = {
    TableName: "Events",
  };

  try {
    const data = await docClient.scan(params).promise();
    return data.Items;
  } catch (error) {
    console.error("Error fetching events:", error);
    return [];
  }
}

// Function to load events and log them to the console
async function loadEvents() {
  const events = await fetchEvents();
  console.log("Fetched Events:", events);
}
