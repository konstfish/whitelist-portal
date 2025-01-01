fetch('https://api.ipify.org')
    .then(response => response.text())
    .then(data => document.getElementById('address').value = data);