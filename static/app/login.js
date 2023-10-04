// login.js

document.getElementById("login-form").addEventListener("submit", function(event) {
    const usernameInput = document.getElementById("username");
    const passwordInput = document.getElementById("password");

    // Validation for username (minimum 4 characters)
    if (usernameInput.value.length < 4) {
        alert("Username must be at least 4 characters long.");
        event.preventDefault();
        return;
    }

    // Validation for password (minimum 6 characters)
    if (passwordInput.value.length < 6) {
        alert("Password must be at least 6 characters long.");
        event.preventDefault();
        return;
    }
});

