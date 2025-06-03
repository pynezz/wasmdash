// app.js
// Some JS for the go-templ app

document.addEventListener("DOMContentLoaded", function() {
    // Get the form element
    const form = document.getElementById("form");

    // Add an event listener for the form submission
    form.addEventListener("submit", function(event) {
        // Prevent the default form submission
        event.preventDefault();

        // Get the input value
        const input = document.getElementById("input").value;

        // Log the input value to the console
        console.log("Form submitted with input:", input);

        // Optionally, you can send this data to a server or process it further
    });
});
