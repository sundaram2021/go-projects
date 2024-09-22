document.getElementById('apiRequestForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const url = this.querySelector('input[type="text"]').value;
    const method = this.querySelector('select').value;
    const body = this.querySelector('textarea').value;
    const headersObj = {};
    document.querySelectorAll('.header-pair').forEach(header => {
        const key = header.querySelector('input[placeholder="Header Key"]').value;
        const value = header.querySelector('input[placeholder="Header Value"]').value;
        if (key && value) {
            headersObj[key] = value;
        }
    });

    const requestData = JSON.stringify({
        url: url,
        method: method,
        body: body,
        headers: headersObj
    });

    const responseText = document.getElementById('responseText');

    fetch('/request', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: requestData
    })
    .then(response => response.json())
    .then(data => {
        let formattedResponse = `Status: ${data.status}\n`;
        formattedResponse += Object.entries(data.headers).map(([key, value]) => `${key}: ${value.join("; ")}`).join("\n") + "\n\n";
        formattedResponse += data.body;
        responseText.textContent = formattedResponse;
    })
    .catch(error => {
        responseText.textContent = 'Error: ' + error.message;
    });
});


function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:\s*)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function(match) {
        let cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}
